package appVersion

import "github.com/zihao-boy/zihao/entity/dto"

type AppVersionDto struct {
	dto.PageDto
	AvId       string `json:"avId" sql:"-"`
	Name       string `json:"name" sql:"-"`
	Remark     string `json:"remark" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
