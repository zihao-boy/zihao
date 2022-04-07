package server

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/vpn/iface"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"net"
	"sync"
)
const Mtu = 1500
//todo: add sync.Mutx for Users change
type LoginManager struct {
	//key: clientProtocol:clientIP:clientPort  value: key for AES
	Users    map[string]*User
	Tokens   map[string]bool
	VpnDataDto      *vpn.SlaveVpnDataDto
	TunServer  *iface.TunServer
	DhcpServer *Dhcp
	Mutex    sync.Mutex
}

func NewLoginManager(vpnDataDto vpn.SlaveVpnDataDto) (*LoginManager, error) {
	tunServer, err := iface.NewTunServer(vpnDataDto.Vpn.TunName, Mtu)
	if err != nil {
		return nil, err
	}

	lm := &LoginManager{
		Users:      map[string]*User{},
		Tokens:     map[string]bool{},
		VpnDataDto:       & vpnDataDto,
		TunServer:  tunServer,
		DhcpServer: NewDhcp(&vpnDataDto),
	}

	for _, user := range vpnDataDto.Users {
		lm.Tokens[user.Token] = true
	}
	return lm, nil
}

func (lm *LoginManager) Login(client string, protocol string, token string) error {
	defer lm.Mutex.Unlock()
	lm.Mutex.Lock()
	if _, ok := lm.Tokens[token]; ok {
		if user, ok := lm.Users[client]; ok {
			user.Close()
		}
		localTunIp, err := lm.DhcpServer.ApplyIp()
		if err != nil {
			return err
		}

		user := NewUser(client, protocol, localTunIp, token, nil, lm.Logout)
		lm.Users[client] = user

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
