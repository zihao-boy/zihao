package client

import (
	"fmt"
	"github.com/songgao/water"
	encrypt2 "github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/common/innerNet/encrypt"
	"github.com/zihao-boy/zihao/common/innerNet/header"
	"github.com/zihao-boy/zihao/common/innerNet/iface"
	"github.com/zihao-boy/zihao/common/innerNet/io"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"golang.zx2c4.com/wireguard/tun"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type TcpClient struct {
	ServerAdd         string
	InnerNetClientDto *innerNet.InnerNetClientDto
	TcpConn           *net.TCPConn
	TunConn           *water.Interface
	WinTunConn        tun.Device
	HeartbeatTime     time.Time
}

func NewTcpClient(innerNetClientDto *innerNet.InnerNetClientDto) (*TcpClient, error) {
	saddr, tname, mtu := innerNetClientDto.ServerAddr, innerNetClientDto.TunName, 1500
	addr, err := net.ResolveTCPAddr("", saddr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return nil, err
	}

	tun, err := iface.NewTunServer(tname, mtu)
	if err != nil {
		return nil, err
	}

	return &TcpClient{
		ServerAdd:         saddr,
		InnerNetClientDto: innerNetClientDto,
		TcpConn:           conn,
		TunConn:           tun.TunConn,
		WinTunConn:        tun.WinTunConn,
		HeartbeatTime: time.Now().Add(60 * time.Second),
	}, nil
}

func (tc *TcpClient) writeToServer() {
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.InnerNetClientDto.Username + tc.InnerNetClientDto.Password)))
	//var frame ethernet.Frame
	var (
		n   int
		err error
	)

	for {
		data := make([]byte, 1500)
		//n, err := tc.TunConn.Read(data)
		sysType := runtime.GOOS
		if sysType == "windows" {
			n, err = tc.WinTunConn.Read(data, 0)
		} else {
			n, err = tc.TunConn.Read(data)
		}
		fmt.Println("网卡数据", err, n)
		if err == nil && n > 0 {
			protocol, src, dst, err := header.GetBase(data)
			fmt.Println("网卡数据解析 ", protocol, src, dst, err)

			if err == nil {
				endata, err := encrypt.EncryptAES(data[:n], encryptKey)
				fmt.Println("网卡数据加密 ", endata, err)
				if err == nil {
					io.WritePacket(tc.TcpConn, endata)
					fmt.Println("ToServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, n, src, dst)
				}
			}
		}
	}
}

// sudo ifconfig utun7 192.168.2.4 255.255.255.0 up
// sudo route -n add -net 192.168.2.0 -netmask 255.255.255.0 192.168.2.4
func (tc *TcpClient) readFromServer() error {
	var (
		cmd *exec.Cmd
	)
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.InnerNetClientDto.Username + tc.InnerNetClientDto.Password)))
	for {
		if data, err := io.ReadPacket(tc.TcpConn); err == nil {
			if data, err = encrypt.DecryptAES(data, encryptKey); err == nil {
				if protocol, src, dst, err := header.GetBase(data); err == nil {
					sysType := runtime.GOOS
					if sysType == "windows" {
						tc.WinTunConn.Write(data, 0)
					} else {
						tc.TunConn.Write(data)
					}
					fmt.Println("FromServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, len(data), src, dst)
					continue
				}
				ipData := string(data)
				if strings.HasPrefix(ipData,"ping"){
					tc.HeartbeatTime = time.Now().Add(60*time.Second)
					continue
				}
				if !strings.HasPrefix(ipData, "ip=") {
					continue
				}
				fmt.Println("ipData", ipData)
				ipData = ipData[3:]
				fmt.Println("ipData2", ipData)

				//setting tun

				sysType := runtime.GOOS
				if sysType == "windows" {
					ipName, _ := tc.WinTunConn.Name()
					cmd = exec.Command("cmd", "/C", "netsh interface ip set address name=\""+ipName+"\" source=static addr="+ipData+" mask=255.255.255.0 gateway=none")
					//netsh interface ip set address name="本地连接" source=static addr=192.168.1.6 mask=255.255.255.0 gateway=192.168.0.1 1
					cmd.CombinedOutput()
					ipDatas := strings.Split(ipData, ".")
					ipDatas[3] = "0"
					nIpData := strings.Join(ipDatas, ".")
					cmd = exec.Command("cmd", "/C", "route add "+nIpData+" MASK 255.255.255.0 "+ipData)
					cmd.CombinedOutput()
				} else if sysType == "darwin" {
					shellCmd := "ifconfig " + tc.TunConn.Name() + " " + ipData + " 255.255.255.0 up"
					cmd = exec.Command("bash", "-c", shellCmd)
					fmt.Println(shellCmd)
					cmd.CombinedOutput()
					ipDatas := strings.Split(ipData, ".")
					ipDatas[3] = "0"
					nIpData := strings.Join(ipDatas, ".")
					shellCmd = "route -n add -net " + nIpData + " -netmask 255.255.255.0 " + ipData
					fmt.Println(shellCmd)
					cmd = exec.Command("bash", "-c", shellCmd)
					cmd.CombinedOutput()
				} else {
					shellCmd := "ifconfig " + tc.TunConn.Name() + " " + ipData + " 255.255.255.0 up"
					cmd = exec.Command("bash", "-c", shellCmd)
					fmt.Println(shellCmd)
					cmd.CombinedOutput()
					ipDatas := strings.Split(ipData, ".")
					ipDatas[3] = "0"
					nIpData := strings.Join(ipDatas, ".")
					shellCmd = "route add -net " + nIpData + " netmask 255.255.255.0 gw " + ipData
					fmt.Println(shellCmd)
					cmd = exec.Command("bash", "-c", shellCmd)
					cmd.CombinedOutput()
				}

			}
		}
	}
}

func (tc *TcpClient) login() error {
	//if len(tc.Cfg.Tokens) <= 0 {
	//	return fmt.Errorf("no token provided")
	//}
	data := []byte(encrypt2.Md5(tc.InnerNetClientDto.Username + tc.InnerNetClientDto.Password))
	if _, err := io.WritePacket(tc.TcpConn, data); err != nil {
		return err
	}
	return nil
}

func (tc *TcpClient) Start() error {
	fmt.Println("TcpClient started")
	if err := tc.login(); err != nil {
		return err
	}
	go tc.writeToServer()
	go tc.readFromServer()

	//heartbeat
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.InnerNetClientDto.Username + tc.InnerNetClientDto.Password)))
	endata, _ := encrypt.EncryptAES([]byte("ping"), encryptKey)
	go func() {
		for{
			n,err := io.WritePacket(tc.TcpConn, endata)
			fmt.Println("client heart beat",n,err,endata)
			time.Sleep(5 * time.Second)
		}
	}()
	return nil
}

func (tc *TcpClient) Stop() error {
	fmt.Println("TcpClient stopped")
	tc.TcpConn.Close()
	// 保存原始设备句柄
	sysType := runtime.GOOS
	if sysType == "windows" {
		tc.WinTunConn.Close()
	} else {
		tc.TunConn.Close()
	}

	return nil
}

func (tc *TcpClient)Recover() error {
	tc.TcpConn.Close()
	saddr := tc.InnerNetClientDto.ServerAddr
	addr, err := net.ResolveTCPAddr("", saddr)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return err
	}

	tc.TcpConn = conn

	tc.login()
	return nil

}
