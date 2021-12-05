package appService

import "github.com/zihao-boy/zihao/entity/dto"

type AppServiceHostsDto struct {
	dto.PageDto
	HostsId       string `json:"hostsId" sql:"-"`
	AsId     string `json:"asId" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Hostname   string `json:"hostname" sql:"-"`
	Ip     string `json:"ip" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
