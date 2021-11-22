package appService

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

const (
	//状态  10012 停止 10013 启动 10011 升级
	STATE_STOP string = "10012"
	STATE_START string ="10013"
	STATE_UPGRADE string ="10011"

	//服务类型，001 数据库 002 缓存 003 计算应用
	AS_TYPE_SQL string = "001"
	AS_TYPE_CACHE string = "002"
	AS_TYPE_SERVICE string = "003"

)

type AppServiceDto struct{
	dto.PageDto
	AsId string `json:"asId" sql:"-"`
	AsName string `json:"asName" sql:"-"`
	AsType string `json:"asType" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	AsDesc string `json:"asDesc" sql:"-"`
	State string `json:"state" `
	AsCount string `json:"asCount" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
	StateName string `json:"stateName" sql:"-"`
	AsTypeName string `json:"asTypeName" sql:"-"`

}