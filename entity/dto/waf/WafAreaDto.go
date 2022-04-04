package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const(
	Waf_area_type_W = "W"
	Waf_area_type_B = "B"
)

type WafAreaDto struct {
	dto.PageDto
	WafRuleDto
	Id    string    `json:"id"  `
	TypeCd    string    `json:"typeCd" sql:"-" `
	AreaName  string    `json:"areaName"  sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
