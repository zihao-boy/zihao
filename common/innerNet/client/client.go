package client

import (
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"time"
)

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

	go func() {
		for{
			time.Sleep(10 * time.Second)
			if tcpClient.HeartbeatTime.After(time.Now()){
				continue
			}
			tcpClient.Recover()
		}

	}()

	return nil

}
