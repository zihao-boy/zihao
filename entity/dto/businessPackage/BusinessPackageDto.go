package businessPackage

import "github.com/zihao-boy/zihao/entity/dto"

type BusinessPackageDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Varsion      string `json:"varsion" sql:"-"`
	Path         string `json:"path" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
}
