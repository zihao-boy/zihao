package server

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/innerNet/header"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
)

type Dhcp struct {
	InnerNetDataDto *innerNet.SlaveInnerNetDataDto
	Ip         uint32
	Mask       uint32
	UsedIps    map[uint32]bool
}

func NewDhcp(vpnDataDto *innerNet.SlaveInnerNetDataDto) *Dhcp {
	ip, mask := header.ParseNet(vpnDataDto.InnerNet.Tun)
	return &Dhcp{
		InnerNetDataDto: vpnDataDto,
		Ip:         header.Str2IP(ip),
		Mask:       header.MaskNumber2Mask(mask),
		UsedIps:    map[uint32]bool{},
	}
}

func (dhcp *Dhcp) ApplyIp() (string, error) {
	for ip := dhcp.Ip + 1; ip < ((dhcp.Ip & dhcp.Mask) ^ (^dhcp.Mask)); ip++ {
		if _, ok := dhcp.UsedIps[ip]; !ok {
			dhcp.UsedIps[ip] = true
			return header.IP2Str(ip), nil
		}
	}
	return "", fmt.Errorf("no enough ip")
}

func (dhcp *Dhcp) ReleaseIp(ip string) {
	delete(dhcp.UsedIps, header.Str2IP(ip))
}
