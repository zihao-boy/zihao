package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type WafDto struct {
	dto.PageDto
	WafId     string    `json:"wafId" sql:"-" `
	WafName     string    `json:"wafName" sql:"-" `
	Port     string    `json:"port" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	State      string    `json:"state"`
}

