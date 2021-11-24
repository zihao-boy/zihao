package monitor

import "github.com/zihao-boy/zihao/entity/dto"

type MonitorTaskAttrDto struct {
	dto.PageDto
	TaskId     string `json:"taskId" sql:"-"`
	AttrId     string `json:"attrId" sql:"-"`
	SpecCd     string `json:"specCd" sql:"-"`
	Value      string `json:"value" `
	CreateTime string `json:"createTime" sql:"-"`
	SpecName   string `json:"specName" sql:"-"`
}
