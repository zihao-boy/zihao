package appVersion

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type AppVersionAttrDto struct {
	dto.PageDto
	AttrId string `json:"attrId" sql:"-"`
	AvId string `json:"avId" sql:"-"`
	Version string `json:"version"`
	TenantId string `json:"tenantId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
