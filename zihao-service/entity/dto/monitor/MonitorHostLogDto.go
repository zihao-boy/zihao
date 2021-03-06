package monitor

type MonitorHostLogDto struct {
	LogId string `json:"logId" sql:"-"`
	HostId string `json:"hostId" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	CpuRate string `json:"cpuRate" sql:"-"`
	MemRate string `json:"memRate" sql:"-"`
	DiskRate string `json:"diskRate" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`

}
