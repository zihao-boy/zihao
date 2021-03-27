package privilege

type PrivilegeGroupDto struct {
	PgId string `json:"pgId" sql:"-"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string	`json:"statusCd" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
}
