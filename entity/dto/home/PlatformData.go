package home

//platform data
type PlatformDataDto struct {
	HostCount   int64 `json:"hostCount"`   //主机数
	CpuCount    string `json:"cpuCount"`    //cpu 核数
	MemCount    string `json:"memCount"`    //内存数
	DiskCount   string `json:"diskCount"`   //磁盘数
	AppCount    int64 `json:"appCount"`    //应用数
	DockerCount int64 `json:"dockerCount"` //容器数
}
