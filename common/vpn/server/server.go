package server

import (
	"github.com/zihao-boy/zihao/common/vpn/config"
)
// start server

func StartServer(cfg *config.Config) error {

	loginManager, err := NewLoginManager(cfg)

	if err != nil {
		return err
	}

	tcpServer, err := NewTcpServer(cfg, loginManager)
	if err != nil {
		return err
	}

	loginManager.Start()
	tcpServer.Start()

	return nil
}
