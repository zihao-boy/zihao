package workflow

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)


type WorkflowDto struct {
	dto.PageDto
	WorkflowId      string         `json:"workflowId" sql:"-" `
	Name       string         `json:"name" `
	Yaml       string         `json:"yaml" `
	TenantId       string         `json:"tenantId" `
	State       string         `json:"state" `
	JobTime       string         `json:"jobTime" `
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
}

type WorkflowStepDto struct {
	dto.PageDto
	StepId      string         `json:"stepId" sql:"-" `
	Step    string         `json:"step" `
	Name       string         `json:"name" `
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
}

type WorkflowStepParamDto struct {
	dto.PageDto
	ParamId      string         `json:"paramId" sql:"-" `
	StepId      string         `json:"stepId" sql:"-" `
	ParamName string `json:"paramName" sql:"-"`
	ParamSpec string `json:"paramSpec" sql:"-"`
	ParamValue string `json:"paramValue" sql:"-"`
	Seq string `json:"seq"`
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
}

