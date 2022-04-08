package server

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/vpn/encrypt"
	"github.com/zihao-boy/zihao/common/vpn/iface"
	"github.com/zihao-boy/zihao/common/vpn/io"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
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
	Tokens   map[string]vpn.VpnUserDto
	VpnDataDto      *vpn.SlaveVpnDataDto
	TunServer  *iface.TunServer
	DhcpServer *Dhcp
	Mutex    sync.Mutex
}

func NewLoginManager(vpnDataDto vpn.SlaveVpnDataDto) (*LoginManager, error) {
	var (
		cmd *exec.Cmd
	)
	tunServer, err := iface.NewTunServer(vpnDataDto.Vpn.TunName, Mtu)
	if err != nil {
		return nil, err
	}

	lm := &LoginManager{
		Users:      map[string]*User{},
		Tokens:     map[string]vpn.VpnUserDto{},
		VpnDataDto:       & vpnDataDto,
		TunServer:  tunServer,
		DhcpServer: NewDhcp(&vpnDataDto),
	}

	for _, user := range vpnDataDto.Users {
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
	} else {
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
	}


	return lm, nil
}

func (lm *LoginManager) Login(client string, protocol string, token string,conn net.Conn) error {
	defer lm.Mutex.Unlock()
	lm.Mutex.Lock()
	if vpnUser, ok := lm.Tokens[token]; ok {
		if user, ok := lm.Users[client]; ok {
			user.Close()
		}
		var localTunIp string
		var  err error
		if vpnUser.Ip == "0.0.0.0"{
			localTunIp, err = lm.DhcpServer.ApplyIp()
		}else{
			localTunIp,err = vpnUser.Ip,nil
		}
		if err != nil {
			return err
		}

		user := NewUser(client, protocol, localTunIp, token, nil, lm.Logout)
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
		lm.TunServer.StartClient(client, user.ConnToTunChan, user.TunToConnChan)
	}
}

func (lm *LoginManager) GetUser(client string) *User {
	if user, ok := lm.Users[client]; ok {
		return user
	}
	return nil
}
