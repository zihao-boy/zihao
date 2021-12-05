package appService

import "github.com/zihao-boy/zihao/entity/dto"

type AppServicePortDto struct {
	dto.PageDto
	PortId       string `json:"portId" sql:"-"`
	AsId     string `json:"asId" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	SrcPort   string `json:"srcPort" sql:"-"`
	TargetPort     string `json:"targetPort" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
