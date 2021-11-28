package monitor

//{'cpuRate':2.6,'memTotal':3821760,'memUsed':541840,'diskTotal':1014,'diskUsed':149}
type MonitorCheckHostInfoDto struct {
	CpuRate float64 `json:"cpuRate"`
	MemTotal float64 `json:"memTotal"`
	MemUsed float64 `json:"memUsed"`
	DiskTotal float64 `json:"diskTotal"`
	DiskUsed float64 `json:"diskUsed"`
}
