package server

import (
	"github.com/zihao-boy/zihao/entity/dto/vpn"
)
// start server

var (
	loginManager  *LoginManager

	tcpServer *TcpServer
)

func StartServer(vpnDataDto vpn.SlaveVpnDataDto) (err error) {

	loginManager, err = NewLoginManager(vpnDataDto)

	if err != nil {
		return err
	}

	tcpServer, err = NewTcpServer(vpnDataDto, loginManager)
	if err != nil {
		return err
	}

	loginManager.Start()
	tcpServer.Start()

	return nil
}

func StopServer() error {
	if loginManager == nil{
		return nil
	}
	if tcpServer == nil{
		return nil
	}
	loginManager.Stop()
	tcpServer.Stop()
	return nil
}
