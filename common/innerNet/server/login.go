package server

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/innerNet/encrypt"
	"github.com/zihao-boy/zihao/common/innerNet/iface"
	"github.com/zihao-boy/zihao/common/innerNet/io"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)
const Mtu = 1500
//todo: add sync.Mutx for Users change
type LoginManager struct {
	//key: clientProtocol:clientIP:clientPort  value: key for AES
	Users    map[string]*User
	Tokens   map[string]innerNet.InnerNetUserDto
	InnerNetDataDto      *innerNet.SlaveInnerNetDataDto
	TunServer  *iface.TunServer
	DhcpServer *Dhcp
	Mutex    sync.Mutex
}

func NewLoginManager(innerNetDataDto innerNet.SlaveInnerNetDataDto) (*LoginManager, error) {
	var (
		cmd *exec.Cmd
	)
	tunServer, err := iface.NewTunServer(innerNetDataDto.InnerNet.TunName, Mtu)
	if err != nil {
		return nil, err
	}

	lm := &LoginManager{
		Users:      map[string]*User{},
		Tokens:     map[string]innerNet.InnerNetUserDto{},
		InnerNetDataDto:       & innerNetDataDto,
		TunServer:  tunServer,
		DhcpServer: NewDhcp(&innerNetDataDto),
	}

	for _, user := range innerNetDataDto.Users {
		lm.Tokens[user.Token] = *user
	}

	ipData,err := lm.DhcpServer.ApplyIp()
	//setting tun

	sysType := runtime.GOOS
	if sysType == "windows" {
		cmd = exec.Command("cmd", "/C", "ipconfig "+tunServer.TunConn.Name()+" "+ipData+" 255.255.255.0 up")
		cmd.CombinedOutput()
		ipDatas := strings.Split(ipData, ".")
		ipDatas[3] = "0"
		nIpData := strings.Join(ipDatas, ".")
		cmd = exec.Command("cmd", "-c", "route -n add -net "+nIpData+" -netmask 255.255.255.0 "+ipData)
		cmd.CombinedOutput()
	} else if(sysType == "darwin") {
		shellCmd := "ifconfig "+tunServer.TunConn.Name()+" "+ipData+" 255.255.255.0 up"
		cmd = exec.Command("bash", "-c", shellCmd)
		fmt.Println(shellCmd)
		cmd.CombinedOutput()
		ipDatas := strings.Split(ipData, ".")
		ipDatas[3] = "0"
		nIpData := strings.Join(ipDatas, ".")
		shellCmd = "route -n add -net "+nIpData+" -netmask 255.255.255.0 "+ipData
		fmt.Println(shellCmd)
		cmd = exec.Command("bash", "-c", shellCmd)
		cmd.CombinedOutput()
	}else{
		shellCmd := "ifconfig "+tunServer.TunConn.Name()+" "+ipData+" 255.255.255.0 up"
		cmd = exec.Command("bash", "-c", shellCmd)
		fmt.Println(shellCmd)
		cmd.CombinedOutput()
		ipDatas := strings.Split(ipData, ".")
		ipDatas[3] = "0"
		nIpData := strings.Join(ipDatas, ".")
		shellCmd = "route add -net "+nIpData+" netmask 255.255.255.0 gw "+ipData
		fmt.Println(shellCmd)
		cmd = exec.Command("bash", "-c", shellCmd)
		cmd.CombinedOutput()
	}


	return lm, nil
}

func (lm *LoginManager) Login(client string, protocol string, token string,conn net.Conn) error {
	defer lm.Mutex.Unlock()
	lm.Mutex.Lock()
	if innerNetUser, ok := lm.Tokens[token]; ok {
		if user, ok := lm.Users[client]; ok {
			user.Close()
		}
		var localTunIp string
		var  err error
		if innerNetUser.Ip == "0.0.0.0"{
			localTunIp, err = lm.DhcpServer.ApplyIp()
		}else{
			localTunIp,err = innerNetUser.Ip,nil
		}
		if err != nil {
			return err
		}

		user := NewUser(client, protocol, localTunIp, token, nil, lm.Logout)
		oldUser,ok := lm.Users[client] // 假如key存在,则name = 李四 ，ok = true,否则，ok = false
		if ok{
			oldUser.Close()
			delete(lm.Users,client)
		}

		lm.Users[client] = user
		encryptKey := encrypt.GetAESKey([]byte(user.Token))

		if endata, err := encrypt.EncryptAES([]byte("ip="+localTunIp), encryptKey); err == nil {
			if user.Protocol == "tcp" {
				_, err = io.WritePacket(conn, endata)
			}
		}

		return nil
	}
	return fmt.Errorf("token not found")
}

func (lm *LoginManager) Logout(client string) {
	defer lm.Mutex.Unlock()
	lm.Mutex.Lock()
	if user, ok := lm.Users[client]; ok {
		lm.DhcpServer.ReleaseIp(user.LocalTunIp)
		delete(lm.Users, client)

	}
}

func (lm *LoginManager) Start() {
	lm.TunServer.Start()
}

func (lm *LoginManager) Stop() {
	lm.TunServer.Stop()
}

func (lm *LoginManager) StartClient(client string, conn net.Conn) {
	if user, ok := lm.Users[client]; ok {
		user.Conn = conn
		user.Start()
		lm.TunServer.StartClient(client, user.ConnToTunChan, user.TunToConnChan,user.LocalTunIp,user.Protocol)
	}
}

func (lm *LoginManager) GetUser(client string) *User {
	if user, ok := lm.Users[client]; ok {
		return user
	}
	return nil
}
