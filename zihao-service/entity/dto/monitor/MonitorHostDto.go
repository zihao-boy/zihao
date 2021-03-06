package monitor

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type MonitorHostDto struct {
	dto.PageDto
	MhId string `json:"mhId" sql:"-"`
	MhgId string `json:"mhgId" sql:"-"`
	HostId string `json:"hostId" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	CpuRate string `json:"cpuRate" sql:"-"`
	MemRate string `json:"memRate" sql:"-"`
	DiskRate string `json:"diskRate" sql:"-"`
	FreeMem string `json:"freeMem" sql:"-"`
	FreeDisk string `json:"freeDisk" sql:"-"`
	MonDisk string `json:"monDisk" sql:"-"`
	MonDate string `json:"monDate" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
	CpuThreshold string `json:"cpuThreshold" sql:"-"`
	MemThreshold string `json:"memThreshold" sql:"-"`
	DiskThreshold string `json:"diskThreshold" sql:"-"`
	Name string `json:"name"`
	Ip string `json:"ip"`
	Passwd string `json:"passwd"`
	Username string `json:"username"`
	NoticeType  string `json:"noticeType" sql:"-"`

}
