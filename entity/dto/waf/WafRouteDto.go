package waf

import "time"

type WafRouteDto struct {
	RouteId     string    `json:"routeId" sql:"-" `
	WafId     string    `json:"wafId" sql:"-" `
	Hostname     string    `json:"hostname" `
	Ip     string    `json:"ip" `
	Port     string    `json:"port" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

