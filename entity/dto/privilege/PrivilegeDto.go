package privilege

/**
权限 实体
*/
type PrivilegeDto struct {
	Pid         string `json:"pId"  sql:"-"`
	Name        string `json:"name" `
	Description string `json:"description" `
	CreateTime  string `json:"createTime"  sql:"-"`
	StatusCd    string `json:"statusCd"  sql:"-"`
	Resource    string `json:"resource" `
	Mid         string `json:"mId"  sql:"-"`
}
