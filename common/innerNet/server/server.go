package server

import (
	encrypt2 "github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
)
// start server

var (
	loginManager  *LoginManager

	tcpServer *TcpServer
)

func StartServer(innerNetDataDto innerNet.SlaveInnerNetDataDto) (err error) {

	if innerNetDataDto.Users != nil && len(innerNetDataDto.Users)>0{
		for _,user:= range innerNetDataDto.Users{
			user.Token = encrypt2.Md5(user.Username+user.Password)
		}
	}
	for _,privilege:= range innerNetDataDto.Privileges{
		privilege.Token = encrypt2.Md5(privilege.SrcUserName+privilege.SrcPassword)
	}
	UserPrivileges = innerNetDataDto.Privileges

	loginManager, err = NewLoginManager(innerNetDataDto)

	if err != nil {
		return err
	}

	tcpServer, err = NewTcpServer(innerNetDataDto, loginManager)
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

func InitInnerNetConfig(innerNetDataDto innerNet.SlaveInnerNetDataDto) error {
	if innerNetDataDto.Users != nil && len(innerNetDataDto.Users)>0{
		for _,user:= range innerNetDataDto.Users{
			user.Token = encrypt2.Md5(user.Username+user.Password)
		}
	}
	for _, user := range innerNetDataDto.Users {
		loginManager.Tokens[user.Token] = *user
	}
	for _,privilege:= range innerNetDataDto.Privileges{
		privilege.Token = encrypt2.Md5(privilege.SrcUserName+privilege.SrcPassword)
	}
	UserPrivileges = innerNetDataDto.Privileges
	return nil
}
