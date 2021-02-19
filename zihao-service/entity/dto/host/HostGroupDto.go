package host

/**
主机组实体
 */
type HostGroupDto struct {
	GroupId string `json:"groupId" sql:"-"`
	Name string `json:"name" `
	Description string `json:"description" `
	TenantId string `json:"tenantId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
