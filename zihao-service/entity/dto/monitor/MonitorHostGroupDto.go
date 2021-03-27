package monitor

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type MonitorHostGroupDto struct {
	dto.PageDto
	MhgId string `json:"mhgId" sql:"-"`
	Name string `json:"name"`
	MonCron string `json:"monCron" sql:"-"`
	State string `json:"state" `
	StateName string `json:"stateName" sql:"-"`
	MonDate string `json:"monDate" sql:"-"`
	NoticeType  string `json:"noticeType" sql:"-"`
	NoticeTypeName string `json:"noticeTypeName" sql:"-"`
	CreateTime  string `json:"createTime" sql:"-"`
	StatusCd  string `json:"statusCd" sql:"-"`
	Remark  string `json:"remark"`
	TenantId string `json:"tenantId" sql:"-"`

}
