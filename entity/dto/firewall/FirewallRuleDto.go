package firewall

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type FirewallRuleDto struct {
	dto.PageDto
	RuleId     string    `json:"ruleId" sql:"-" `
	GroupId    string    `json:"groupId" sql:"-" `
	Inout      string    `json:"inout" `
	AllowLimit string    `json:"allowLimit" sql:"-"`
	DstObj     string    `json:"dstObj" sql:"-"`
	Seq        string    `json:"seq" `
	Protocol   string    `json:"protocol" `
	SrcObj     string    `json:"srcObj" sql:"-"`
	Remark     string    `json:"remark" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
