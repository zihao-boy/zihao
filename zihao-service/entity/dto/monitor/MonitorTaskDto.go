package monitor

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type MonitorTaskDto struct {
	dto.PageDto
	TaskId string `json:"taskId" sql:"-"`
	TaskName string `json:"taskName" sql:"-"`
	TemplateId string `json:"templateId" sql:"-"`
	TaskCron string `json:"taskCron" sql:"-"`
	State string `json:"state"`
	StatusCd string `json:"statusCd" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
 	HostId string `json:"hostId" sql:"-"`
}