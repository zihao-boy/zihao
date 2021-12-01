package privilege

type PrivilegeUserDto struct {
	PuId          string `json:"puId" sql:"-"`
	Pid           string `json:"pId" sql:"-"`
	PrivilegeFlag string `json:"privilegeFlag" sql:"-"`
	UserId        string `json:"userId" sql:"-"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
	TenantId      string `json:"tenantId" sql:"-"`
}
