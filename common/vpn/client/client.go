package client

import "github.com/zihao-boy/zihao/common/vpn/config"

func StartClient(cfg *config.Config) error {
	var tcpClient *TcpClient
	var err error
	if tcpClient, err = NewTcpClient(cfg); err != nil {
		return err;
	}
	err = tcpClient.Start()

	if err != nil{
		return err
	}

	return nil

}
