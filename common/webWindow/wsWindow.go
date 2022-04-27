package webWindow

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kataras/iris/v12/websocket"
	"github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/webWindow/guac"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"golang.org/x/sync/errgroup"
	"strings"
)

const (
	host_token string = "host_token"
)

type ReqArg struct {
	GuacadAddr    string `form:"guacad_addr"`
	AssetProtocol string `form:"asset_protocol"`
	AssetHost     string `form:"asset_host"`
	AssetPort     string `form:"asset_port"`
	AssetUser     string `form:"asset_user"`
	AssetPassword string `form:"asset_password"`
	ScreenWidth   int    `form:"screen_width"`
	ScreenHeight  int    `form:"screen_height"`
	ScreenDpi     int    `form:"screen_dpi"`
}

type Request struct {
	Operate    string `json:"operate"`
	ZihaoToken string `json:"zihaoToken"`
	HostId     string `json:"hostId"`
	Command    string `json:"command"`
	WinWidth   int    `json:"winWidth"`
	WinHeight  int    `json:"winHeight"`
	Cols       uint32 `json:"cols"`
	Rows       uint32 `json:"rows"`
}

var (
	channels = make(map[string]*guac.SimpleTunnel)
)

func WebSocketConn(req Request, connId string, nsConn *websocket.Conn) {
	var channel *guac.SimpleTunnel
	if !checkUserToken(req.ZihaoToken) {
		return
	}
	loginInfo := getServerInfo(req)
	if loginInfo.GuacadAddr == "" {
		return
	}
	channel, err := windowConnect(loginInfo)
	if err != nil {
		fmt.Println("连接错误",err)
		return
	}
	channels[connId] = channel

	ioCopy(nsConn, channel)
}

func WebSocketHandler(data []byte, connId string, nsConn *websocket.NSConn) {
	var channel *guac.SimpleTunnel
	//done := make(chan bool, 2
	channel = channels[connId]
	if channel == nil {
		return
	}
	writer := channel.AcquireWriter()
	defer channel.ReleaseWriter()
	if _, err := writer.Write(data); nil != err {
		return
	}

	//<-done

}

func windowConnect(arg ReqArg) (s *guac.SimpleTunnel, err error) {
	uid := ""
	fmt.Println("参数：",arg.GuacadAddr, arg.AssetProtocol, arg.AssetHost, arg.AssetPort, arg.AssetUser, arg.AssetPassword, uid, arg.ScreenWidth, arg.ScreenHeight, arg.ScreenDpi)
	pipeTunnel, err := guac.NewGuacamoleTunnel(arg.GuacadAddr, arg.AssetProtocol, arg.AssetHost, arg.AssetPort, arg.AssetUser, arg.AssetPassword, uid, arg.ScreenWidth, arg.ScreenHeight, arg.ScreenDpi)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = pipeTunnel.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return pipeTunnel, nil
}

func ioCopy(nsConn *websocket.Conn, tunnl *guac.SimpleTunnel) {

	reader := tunnl.AcquireReader()
	defer tunnl.ReleaseReader()

	//使用 errgroup 来处理(管理) goroutine for-loop, 防止 for-goroutine zombie
	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		buf := bytes.NewBuffer(make([]byte, 0, guac.MaxGuacMessage*2))

		for {
			ins, err := reader.ReadSome()
			fmt.Println("读取数据")

			if err != nil {
				fmt.Println("读取数据失败",err.Error())
				return err
			}

			if bytes.HasPrefix(ins, guac.InternalOpcodeIns) {
				// messages starting with the InternalDataOpcode are never sent to the websocket
				continue
			}

			if _, err = buf.Write(ins); err != nil {
				fmt.Println("写数据失败",err.Error())

				return err
			}

			// if the buffer has more data in it or we've reached the max buffer size, send the data and reset
			if !reader.Available() || buf.Len() >= guac.MaxGuacMessage {
				mg := websocket.Message{
					Body:     buf.Bytes(),
					IsNative: true,
				}
				nsConn.Write(mg)
				buf.Reset()
			}
		}

	})
	if err := eg.Wait(); err != nil {
		fmt.Println("session-err",err.Error())
	}

}

func checkUserToken(token string) bool {

	cacheToken, err := factory.GetValueAndRemove(host_token)

	if err != nil {
		return false
	}

	if len(cacheToken) < 1 {
		return false
	}

	if token != cacheToken {
		return false
	}
	return true
}

func getServerInfo(req Request) ReqArg {

	var (
		hostDao  dao.HostDao
		hostDtos []*host.HostDto
		err      error
	)

	var hostDto = host.HostDto{
		HostId: req.HostId,
	}

	hostDtos, err = hostDao.GetHosts(hostDto)

	if err != nil || len(hostDtos) < 1 {
		return ReqArg{}
	}

	ips := strings.Split(hostDtos[0].Ip, ":")
	if len(ips) != 2 {
		return ReqArg{}
	}

	return ReqArg{
		GuacadAddr:    config.G_AppConfig.GuacadAddr,
		AssetProtocol: "rdp",
		AssetHost:     ips[0],
		AssetPort:     ips[1],
		AssetUser:     hostDtos[0].Username,
		AssetPassword: hostDtos[0].Passwd,
		ScreenWidth:   req.WinWidth,
		ScreenHeight:  req.WinHeight,
		ScreenDpi:     128,
	}
}

func CloseSsh(connId string) {
	if _, ok := channels[connId]; ok {
		channels[connId].Close()
		delete(channels, connId)
	}
}
