package appVersionJob

import "github.com/zihao-boy/zihao/entity/dto"

type AppVersionJobImagesDto struct {
	dto.PageDto
	JobImagesId         string `json:"jobImagesId" sql:"-"`
	TenantId            string `json:"tenantId" sql:"-"`
	PackageUrl          string `json:"packageUrl" sql:"-"`
	BusinessPackageName string `json:"businessPackageName" sql:"-"`
	BusinessImagesName  string `json:"businessImagesName" sql:"-"`
	JobId               string `json:"jobId" sql:"-"`
	CreateTime          string `json:"createTime" sql:"-"`
	StatusCd            string `json:"statusCd" sql:"-"`
}
