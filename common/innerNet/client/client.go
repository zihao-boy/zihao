package client

import "github.com/zihao-boy/zihao/entity/dto/vpn"

func StartClient(vpnClientDto *vpn.VpnClientDto) error {
	var tcpClient *TcpClient
	var err error
	if tcpClient, err = NewTcpClient(vpnClientDto); err != nil {
		return err;
	}
	err = tcpClient.Start()

	if err != nil{
		return err
	}

	return nil

}
