package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type WafHostsDto struct {
	dto.PageDto
	WafHostId     string    `json:"wafHostId" sql:"-" `
	WafId     string    `json:"wafId" sql:"-" `
	HostId     string    `json:"hostId" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

