package dbLink

import "github.com/zihao-boy/zihao/entity/dto"


// db link struct
type DbLinkDto struct {
	dto.PageDto
	Id       string `json:"id"`
	Name       string `json:"name"`
	Ip       string `json:"ip"`
	Port       string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	DbName       string `json:"dbName" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
}

// sql
type DbSqlDto struct {
	DbId string `json:"dbId"`
	Sql string `json:"sql"`
}
