package waf

import "time"

type WafDto struct {
	WafId     string    `json:"wafId" sql:"-" `
	WafName     string    `json:"wafName" sql:"-" `
	Port     string    `json:"port" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	State      string    `json:"state"`
}

