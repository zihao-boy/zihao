package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type WafRouteDto struct {
	dto.PageDto
	RouteId     string    `json:"routeId" sql:"-" `
	WafId     string    `json:"wafId" sql:"-" `
	Hostname     string    `json:"hostname" `
	Ip     string    `json:"ip" `
	Port     string    `json:"port" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

