package workflow

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

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

