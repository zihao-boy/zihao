package server

import (
	"github.com/zihao-boy/zihao/common/vpn/config"
	"github.com/zihao-boy/zihao/common/vpn/io"
	"net"

)

type TcpServer struct {
	Addr         string
	Cfg          *config.Config
	TcpListener  net.Listener
	LoginManager *LoginManager
}

func NewTcpServer(cfg *config.Config, loginManager *LoginManager) (*TcpServer, error) {
	tcpListener, err := net.Listen("tcp", cfg.ServerAddr)
	if err != nil {
		return nil, err
	}

	return &TcpServer{
		Addr:         cfg.ServerAddr,
		Cfg:          cfg,
		TcpListener:  tcpListener,
		LoginManager: loginManager,
	}, nil
}

func (ts *TcpServer) Start() {
	go func() {
		for {
			if conn, err := ts.TcpListener.Accept(); err == nil {
				go ts.handleRequest(conn)
			}
		}
	}()
}

func (ts *TcpServer) Stop() {
	ts.TcpListener.Close()
}

func (ts *TcpServer) handleRequest(conn net.Conn) {
	client := "tcp:" + conn.RemoteAddr().String()
	if err := ts.login(client, conn); err != nil {
		return
	}
	ts.LoginManager.StartClient(client, conn)
}

func (ts *TcpServer) login(client string, conn net.Conn) error {
	if data, err := io.ReadPacket(conn); err != nil {
		return err

	} else {
		return ts.LoginManager.Login(client, "tcp", string(data))
	}
}
