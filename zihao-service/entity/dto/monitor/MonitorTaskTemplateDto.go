package monitor

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type MonitorTaskTemplateDto struct {
	dto.PageDto
	TemplateId string `json:"templateId" sql:"-"`
	TemplateName string `json:"templateName" sql:"-"`
	TemplateDesc string `json:"templateDesc" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
	ClassBean string `json:"classBean" sql:"-"`
	SpecCd string `json:"specCd" sql:"-"`
	SpecName string `json:"specName" sql:"-"`
	SpecDesc string `json:"specDesc" sql:"-"`
	IsShow string `json:"isShow" sql:"-"`

	MonitorTaskTemplateSpecDto []MonitorTaskTemplateSpecDto `json:"specs"`
}
