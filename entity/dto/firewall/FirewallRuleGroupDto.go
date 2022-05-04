package firewall

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const (
	Firewall_Rule_Group_State_T = "T" //start
	Firewall_Rule_Group_State_F = "F" //stop
)

type FirewallRuleGroupDto struct {
	dto.PageDto
	GroupId    string    `json:"groupId" sql:"-" `
	GroupName  string    `json:"groupName" sql:"-" `
	State      string    `json:"state" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

