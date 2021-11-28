package monitor

import "github.com/zihao-boy/zihao/entity/dto"

type MonitorTaskTemplateSpecDto struct {
	dto.PageDto

	SpecId     string `json:"specId" sql:"-"`
	TemplateId string `json:"templateId" sql:"-"`
	SpecCd     string `json:"specCd" sql:"-"`
	SpecName   string `json:"specName" sql:"-"`
	SpecDesc   string `json:"specDesc" sql:"-"`
	IsShow     string `json:"isShow" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
