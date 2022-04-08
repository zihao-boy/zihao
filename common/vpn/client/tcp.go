package client

import (
	"fmt"
	"github.com/songgao/water"
	encrypt2 "github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/common/vpn/encrypt"
	"github.com/zihao-boy/zihao/common/vpn/header"
	"github.com/zihao-boy/zihao/common/vpn/iface"
	"github.com/zihao-boy/zihao/common/vpn/io"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"net"
)

type TcpClient struct {
	ServerAdd    string
	VpnClientDto *vpn.VpnClientDto
	TcpConn      *net.TCPConn
	TunConn      *water.Interface
}

func NewTcpClient(vpnClientDto *vpn.VpnClientDto) (*TcpClient, error) {
	saddr, tname, mtu := vpnClientDto.ServerAddr, vpnClientDto.TunName, 1500
	addr, err := net.ResolveTCPAddr("", saddr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return nil, err
	}

	tun, err := iface.NewTunServer(tname, mtu)
	if err != nil {
		return nil, err
	}

	return &TcpClient{
		ServerAdd:    saddr,
		VpnClientDto: vpnClientDto,
		TcpConn:      conn,
		TunConn:      tun.TunConn,
	}, nil
}

func (tc *TcpClient) writeToServer() {
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.VpnClientDto.Username+tc.VpnClientDto.Password)))
	data := make([]byte, 1500*2)
	for {
		if n, err := tc.TunConn.Read(data); err == nil && n > 0 {
			if protocol, src, dst, err := header.GetBase(data); err == nil {
				if endata, err := encrypt.EncryptAES(data[:n], encryptKey); err == nil {
					io.WritePacket(tc.TcpConn, endata)
					fmt.Println("ToServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, n, src, dst)
				}
			}
		}
	}
}

func (tc *TcpClient) readFromServer() error {
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.VpnClientDto.Username+tc.VpnClientDto.Password)))
	for {
		if data, err := io.ReadPacket(tc.TcpConn); err == nil {
			if data, err = encrypt.DecryptAES(data, encryptKey); err == nil {
				if protocol, src, dst, err := header.GetBase(data); err == nil {
					tc.TunConn.Write(data)
					fmt.Println("FromServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, len(data), src, dst)
				}
			}
		}
	}
}

func (tc *TcpClient) login() error {
	//if len(tc.Cfg.Tokens) <= 0 {
	//	return fmt.Errorf("no token provided")
	//}
	data := []byte(encrypt2.Md5(tc.VpnClientDto.Username+tc.VpnClientDto.Password))
	if _, err := io.WritePacket(tc.TcpConn, data); err != nil {
		return err
	}
	return nil
}

func (tc *TcpClient) Start() error {
	fmt.Println("TcpClient started")
	if err := tc.login(); err != nil {
		return err
	}
	go tc.writeToServer()
	go tc.readFromServer()
	return nil
}

func (tc *TcpClient) Stop() error {
	fmt.Println("TcpClient stopped")
	tc.TcpConn.Close()
	tc.TunConn.Close()
	return nil
}
