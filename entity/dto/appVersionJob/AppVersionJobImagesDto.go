package appVersionJob

import "github.com/zihao-boy/zihao/entity/dto"

type AppVersionJobImagesDto struct {
	dto.PageDto
	JobImagesId          string `json:"jobImagesId" sql:"-"`
	TenantId             string `json:"tenantId" sql:"-"`
	PackageUrl           string `json:"packageUrl" sql:"-"`
	BusinessPackageId    string `json:"businessPackageId" sql:"-"`
	BusinessDockerfileId string `json:"businessDockerfileId" sql:"-"`
	JobId                string `json:"jobId" sql:"-"`
	CreateTime           string `json:"createTime" sql:"-"`
	StatusCd             string `json:"statusCd" sql:"-"`
	BusinessDockerfileName             string `json:"businessDockerfileName" sql:"-"`
	BusinessPackageName             string `json:"businessPackageName" sql:"-"`
}