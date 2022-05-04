package firewall

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)


type HostFirewallGroupDto struct {
	dto.PageDto
	HfgId string `json:"hfgId" sql:"-"`
	GroupId    string    `json:"groupId" sql:"-" `
	GroupName string `json:"groupName" sql:"-"`
	HostId  string    `json:"hostId" sql:"-" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

