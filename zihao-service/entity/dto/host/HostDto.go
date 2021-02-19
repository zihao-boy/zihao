package host

/**
主机 实体
 */
type HostDto struct {
	HostId string `json:"hostId" sql:"-"`
	GroupId string `json:"groupId" sql:"-"`
	Name string `json:"name" `
	Ip string `json:"ip" `
	Username string `json:"username" `
	Passwd string `json:"passwd" `
	Cpu string `json:"cpu" `
	Mem string `json:"mem" `
	Disk string `json:"disk" `
	TenantId string `json:"tenantId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
