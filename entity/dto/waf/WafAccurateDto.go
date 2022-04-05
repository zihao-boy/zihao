package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)


type WafAccurateDto struct {
	dto.PageDto
	WafRuleDto
	Id    string    `json:"id"  `
	Action    string    `json:"action" `
	TypeCd    string    `json:"typeCd" sql:"-" `
	Include  string    `json:"include" `
	IncludeValue  string    `json:"includeValue"  sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}