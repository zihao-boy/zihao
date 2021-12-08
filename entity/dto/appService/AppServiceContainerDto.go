package appService

import "github.com/zihao-boy/zihao/entity/dto"

type AppServiceContainerDto struct {
	dto.PageDto
	ContainerId       string `json:"containerId" sql:"-"`
	AsId     string `json:"asId" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	HostId   string `json:"hostId" sql:"-"`
	State     string `json:"state" sql:"-"`
	Message     string `json:"message" sql:"-"`
	UpdateTime     string `json:"updateTime" sql:"-"`
	DockerContainerId     string `json:"dockerContainerId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	Ip string `json:"ip"`
	Hostname string `json:"hostname"`
}
