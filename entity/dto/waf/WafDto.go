package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const Waf_state_stop ="2002"
const Waf_state_start ="1001"

type WafDto struct {
	dto.PageDto
	WafId     string    `json:"wafId" sql:"-" `
	WafName     string    `json:"wafName" sql:"-" `
	Port     string    `json:"port" `
	HostIds string `json:"hostIds"`
	WafHosts []*WafHostsDto `json:"wafHosts"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	State      string    `json:"state"`
}

