package tenant

import "github.com/zihao-boy/zihao/entity/dto"

/**
租户实体类
*/
type TenantDto struct {
	dto.PageDto
	TenantType string `json:"tenantType"  sql:"-"`
	TenantName string `json:"tenantName"  sql:"-"`
	TenantId   string `json:"tenantId"  sql:"-"`
	StatusCd   string `json:"statusCd"  sql:"-"`
	State      string `json:"State"  `
	Remark     string `json:"remark"  `
	Phone      string `json:"phone"  `
	PersonName string `json:"personName"  `
	CreateTime string `json:"createTime"  sql:"-"`
	Address    string `json:"address"`
	Username   string `json:"username"`
	Passwd     string `json:"passwd"`
}
