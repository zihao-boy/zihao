package innerNet

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const InnerNet_state_stop = "2002"
const InnerNet_state_start = "1001"

type InnerNetDto struct {
	dto.PageDto
	InnerNetId      string         `json:"innerNetId" sql:"-" `
	InnerNetPort    string         `json:"innerNetPort" sql:"-" `
	Tun        string         `json:"tun" `
	TunName    string         `json:"tunName" sql:"-"`
	Dns        string         `json:"dns"`
	Protocol   string         `json:"protocol"`
	InnerNetHosts   []*InnerNetHostsDto `json:"innerNetHosts"`
	HostIds    string         `json:"hostIds"`
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
	State      string         `json:"state"`
}

type SlaveInnerNetDataDto struct {
	ServerIpUrl string `json:"serverIpUrl"`
	InnerNet         InnerNetDto `json:"innerNet"`
	Users      []*InnerNetUserDto
	Privileges []*InnerNetPrivilegeDto
}

type InnerNetClientDto struct {
	ServerAddr string `json:"serverAddr"`
	TunName      string `json:"tunName"`
	Username     string `json:"username"`
	Password     string `json:"password"`

}
