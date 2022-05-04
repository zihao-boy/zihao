package firewall

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type FirewallRuleDto struct {
	dto.PageDto
	RuleId     string    `json:"rule_id" sql:"-" `
	GroupId    string    `json:"group_id" sql:"-" `
	Inout      string    `json:"in_out" sql:"-"`
	AllowLimit string    `json:"allow_limit" sql:"-"`
	DstObj     string    `json:"dst_obj" sql:"-"`
	Seq        string    `json:"seq" `
	Protocol   string    `json:"protocol" `
	SrcObj     string    `json:"src_obj" sql:"-"`
	Remark     string    `json:"remark" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
