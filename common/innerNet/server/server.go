package server

import (
	encrypt2 "github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
)
// start server

var (
	loginManager  *LoginManager

	tcpServer *TcpServer
)

func StartServer(vpnDataDto vpn.SlaveVpnDataDto) (err error) {

	if vpnDataDto.Users != nil && len(vpnDataDto.Users)>0{
		for _,user:= range vpnDataDto.Users{
			user.Token = encrypt2.Md5(user.Username+user.Password)
		}
	}

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

func InitVpnConfig(vpnDataDto vpn.SlaveVpnDataDto) error {
	if vpnDataDto.Users != nil && len(vpnDataDto.Users)>0{
		for _,user:= range vpnDataDto.Users{
			user.Token = encrypt2.Md5(user.Username+user.Password)
		}
	}
	for _, user := range vpnDataDto.Users {
		loginManager.Tokens[user.Token] = *user
	}
	return nil
}
