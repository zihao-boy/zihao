package vpn

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type VpnHostsDto struct {
	dto.PageDto
	VpnHostId     string    `json:"vpnHostId" sql:"-" `
	VpnId     string    `json:"vpnId" sql:"-" `
	HostId     string    `json:"hostId" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

