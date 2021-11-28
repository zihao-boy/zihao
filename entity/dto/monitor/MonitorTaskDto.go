package monitor

import "github.com/zihao-boy/zihao/entity/dto"

type MonitorTaskDto struct {
	dto.PageDto
	TaskId         string                `json:"taskId" sql:"-"`
	TaskName       string                `json:"taskName" sql:"-"`
	TemplateId     string                `json:"templateId" sql:"-"`
	TaskCron       string                `json:"taskCron" sql:"-"`
	State          string                `json:"state"`
	StatusCd       string                `json:"statusCd" sql:"-"`
	CreateTime     string                `json:"createTime" sql:"-"`
	TenantId       string                `json:"tenantId" sql:"-"`
	HostId         string                `json:"hostId" sql:"-"`
	TemplateName   string                `json:"templateName" sql:"-"`
	ClassBean      string                `json:"classBean" sql:"-"`
	HostName       string                `json:"hostName" sql:"-"`
	Ip             string                `json:"ip" `
	NoticeType     string                `json:"noticeType" sql:"-"`
	NoticeTypeName string                `json:"noticeTypeName" sql:"-"`
	Attr           []*MonitorTaskAttrDto `json:"templateSpecs"`
}
