package vpn

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const Vpn_state_stop = "2002"
const Vpn_state_start = "1001"

type VpnDto struct {
	dto.PageDto
	VpnId      string         `json:"vpnId" sql:"-" `
	VpnPort    string         `json:"vpnPort" sql:"-" `
	Tun        string         `json:"tun" `
	TunName    string         `json:"tunName" sql:"-"`
	Dns        string         `json:"dns"`
	Protocol   string         `json:"protocol"`
	VpnHosts   []*VpnHostsDto `json:"vpnHosts"`
	HostIds    string         `json:"hostIds"`
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
	State      string         `json:"state"`
}

type SlaveVpnDataDto struct {
	ServerIpUrl string `json:"serverIpUrl"`
	Vpn         VpnDto `json:"vpn"`
	Users      []*VpnUserDto
}

type VpnClientDto struct {
	ServerAddr string `json:"serverAddr"`
	TunName      string `json:"tunName"`
	Username     string `json:"username"`
	Password     string `json:"password"`

}
