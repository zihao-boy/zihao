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
	"os/exec"
	"runtime"
	"strings"
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
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.VpnClientDto.Username + tc.VpnClientDto.Password)))
	data := make([]byte, 1500)
	for {
		n, err := tc.TunConn.Read(data)
		if err == nil && n > 0 {
			protocol, src, dst, err := header.GetBase(data)
			if err == nil {
				endata, err := encrypt.EncryptAES(data[:n], encryptKey)
				if err == nil {
					io.WritePacket(tc.TcpConn, endata)
					fmt.Println("ToServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, n, src, dst)
				}
			}
		}
	}
}

// sudo ifconfig utun7 192.168.2.4 255.255.255.0 up
// sudo route -n add -net 192.168.2.0 -netmask 255.255.255.0 192.168.2.4
func (tc *TcpClient) readFromServer() error {
	var (
		cmd *exec.Cmd
	)
	encryptKey := encrypt.GetAESKey([]byte(encrypt2.Md5(tc.VpnClientDto.Username + tc.VpnClientDto.Password)))
	for {
		if data, err := io.ReadPacket(tc.TcpConn); err == nil {
			if data, err = encrypt.DecryptAES(data, encryptKey); err == nil {
				if protocol, src, dst, err := header.GetBase(data); err == nil {
					tc.TunConn.Write(data)
					fmt.Println("FromServer: protocol:%v, len:%v, src:%v, dst:%v", protocol, len(data), src, dst)
					continue
				}
				ipData := string(data)
				if !strings.HasPrefix(ipData, "ip=") {
					continue
				}
				fmt.Println("ipData",ipData)
				ipData = ipData[2:]
				fmt.Println("ipData2",ipData)

				//setting tun

				sysType := runtime.GOOS
				if sysType == "windows" {
					cmd = exec.Command("cmd", "/C", "ipconfig "+tc.TunConn.Name()+" "+ipData+" 255.255.255.0 up")
					cmd.CombinedOutput()
					ipDatas := strings.Split(ipData, ".")
					ipDatas[3] = "0"
					nIpData := strings.Join(ipDatas, ".")
					cmd = exec.Command("cmd", "-c", "route -n add -net "+nIpData+" -netmask 255.255.255.0 "+ipData)
					cmd.CombinedOutput()
				} else {
					cmd = exec.Command("bash", "-c", "ifconfig "+tc.TunConn.Name()+" "+ipData+" 255.255.255.0 up")
					cmd.CombinedOutput()
					ipDatas := strings.Split(ipData, ".")
					ipDatas[3] = "0"
					nIpData := strings.Join(ipDatas, ".")
					cmd = exec.Command("bash", "-c", "route -n add -net "+nIpData+" -netmask 255.255.255.0 "+ipData)
					cmd.CombinedOutput()
				}

			}
		}
	}
}

func (tc *TcpClient) login() error {
	//if len(tc.Cfg.Tokens) <= 0 {
	//	return fmt.Errorf("no token provided")
	//}
	data := []byte(encrypt2.Md5(tc.VpnClientDto.Username + tc.VpnClientDto.Password))
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
