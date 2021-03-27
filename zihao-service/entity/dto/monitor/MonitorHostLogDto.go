package monitor

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type MonitorHostLogDto struct {
	dto.PageDto
	LogId string `json:"logId" sql:"-"`
	HostId string `json:"hostId" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	CpuRate string `json:"cpuRate" sql:"-"`
	MemRate string `json:"memRate" sql:"-"`
	DiskRate string `json:"diskRate" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`

}
