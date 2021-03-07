package tenant

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

/**
商户设置
 */
type TenantSettingDto struct {
	dto.PageDto
	SettingId string `json:"settingId" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	SpecCd string `json:"specCd" sql:"-"`
	Value string `json:"value"`
	CreateTime string `json:"CreateTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
	SpecCdName string `json:"specCdName" sql:"-"`
}
