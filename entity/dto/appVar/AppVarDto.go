package appVar

import "github.com/zihao-boy/zihao/entity/dto"

type AppVarDto struct {
	dto.PageDto
	AvId       string `json:"avId" sql:"-"`
	AvgId      string `json:"avgId" sql:"-"`
	AvgName    string `json:"avgName" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
	VarName    string `json:"varName" sql:"-"`
	VarType    string `json:"varType" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	VarSpec    string `json:"varSpec" sql:"-"`
}
