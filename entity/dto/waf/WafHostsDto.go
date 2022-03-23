package waf

import "time"

type WafHostsDto struct {
	WafHostId     string    `json:"wafHostId" sql:"-" `
	WafId     string    `json:"wafId" sql:"-" `
	HostId     string    `json:"hostId" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

