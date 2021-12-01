package businessDockerfile

import "github.com/zihao-boy/zihao/entity/dto"

type BusinessDockerfileDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Version      string `json:"version" sql:"-"`
	Dockerfile         string `json:"dockerfile" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Username	string `json:"username" sql:"-"`
}
