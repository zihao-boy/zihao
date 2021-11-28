package monitor

import "github.com/zihao-boy/zihao/entity/dto"

type MonitorEventDto struct {
	dto.PageDto
	EventId        string `json:"eventId" sql:"-"`
	EventType      string `json:"eventType" sql:"-"`
	EventObjId     string `json:"eventObjId" sql:"-"`
	EventObjName   string `json:"eventObjName" sql:"-"`
	TenantId       string `json:"tenantId" sql:"-"`
	ThresholdValue string `json:"thresholdValue" sql:"-"`
	CurValue       string `json:"curValue" sql:"-"`
	Remark         string `json:"remark"`
	CreateTime     string `json:"createTime" sql:"-"`
	StatusCd       string `json:"statusCd" sql:"-"`
	NoticeType     string `json:"noticeType" sql:"-"`
	State          string `json:"state"`
	StateRemark    string `json:"stateRemark" sql:"-"`
	NoticeTypeName string `json:"noticeTypeName" sql:"-"`
}
