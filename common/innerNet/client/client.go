package client

import "github.com/zihao-boy/zihao/entity/dto/innerNet"

func StartClient(innerNetClientDto *innerNet.InnerNetClientDto) error {
	var tcpClient *TcpClient
	var err error
	if tcpClient, err = NewTcpClient(innerNetClientDto); err != nil {
		return err;
	}
	err = tcpClient.Start()

	if err != nil{
		return err
	}

	return nil

}
