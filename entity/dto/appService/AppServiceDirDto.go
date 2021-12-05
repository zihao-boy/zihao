package appService

import "github.com/zihao-boy/zihao/entity/dto"

type AppServiceDirDto struct {
	dto.PageDto
	DirId       string `json:"dirId" sql:"-"`
	AsId     string `json:"asId" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	SrcDir   string `json:"srcDir" sql:"-"`
	TargetDir     string `json:"targetDir" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
