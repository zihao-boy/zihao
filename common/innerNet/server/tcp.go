package server

import (
	"github.com/zihao-boy/zihao/common/innerNet/io"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"net"
)

type TcpServer struct {
	Addr         string
	VpnDataDto   *vpn.SlaveVpnDataDto
	TcpListener  net.Listener
	LoginManager *LoginManager
}

func NewTcpServer(vpnDataDto vpn.SlaveVpnDataDto, loginManager *LoginManager) (*TcpServer, error) {
	tcpListener, err := net.Listen("tcp", ":"+vpnDataDto.Vpn.VpnPort)
	if err != nil {
		return nil, err
	}

	return &TcpServer{
		Addr:         ":" + vpnDataDto.Vpn.VpnPort,
		VpnDataDto:   &vpnDataDto,
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
		return ts.LoginManager.Login(client, "tcp", string(data),conn)
	}
}
