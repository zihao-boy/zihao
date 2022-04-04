package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)


type WafCCDto struct {
	dto.PageDto
	WafRuleDto
	Id    string    `json:"id"  `
	Path    string    `json:"path" `
	VisitSec    string    `json:"visitSec" sql:"-" `
	VisitCount  string    `json:"visitCount"  sql:"-"`
	BlockSec  string    `json:"blockSec"  sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}