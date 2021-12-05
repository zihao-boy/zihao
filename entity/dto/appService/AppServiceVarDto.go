package appService

import "github.com/zihao-boy/zihao/entity/dto"

type AppServiceVarDto struct {
	dto.PageDto
	AvId       string `json:"avId" sql:"-"`
	AsId     string `json:"asId" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	VarSpec   string `json:"varSpec" sql:"-"`
	VarName     string `json:"varName" sql:"-"`
	VarValue      string `json:"varValue" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
