package appService

import "github.com/zihao-boy/zihao/entity/dto"

const (
	//状态  10012 停止 10013 启动 10011 升级
	STATE_STOP    string = "10012"
	STATE_START   string = "10013"
	STATE_DOING  string ="10014"
	STATE_UPGRADE string = "10011"

	//服务类型，001 数据库 002 缓存 003 计算应用
	AS_TYPE_SQL     string = "001"
	AS_TYPE_CACHE   string = "002"
	AS_TYPE_SERVICE string = "003"

	AS_DEPLOY_TYPE_GROUP string ="1001" // 多主机
	AS_DEPLOY_TYPE_HOST string ="2002" //单主机
)

type AppServiceDto struct {
	dto.PageDto
	AsId       string `json:"asId" sql:"-"`
	AsName     string `json:"asName" sql:"-"`
	AsType     string `json:"asType" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
	AsDesc     string `json:"asDesc" sql:"-"`
	State      string `json:"state" `
	AsCount    string `json:"asCount" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	StateName  string `json:"stateName" sql:"-"`
	AsTypeName string `json:"asTypeName" sql:"-"`
	AsGroupId  string `json:"asGroupId" sql:"-"`
	AsDeployType string `json:"asDeployType" sql:"-"`
	AsDeployId  string `json:"asDeployId" sql:"-"`
	ImagesId  string `json:"imagesId" sql:"-"`
	ImagesName  string `json:"imagesName" sql:"-"`
	ImagesVersion  string `json:"imagesVersion" sql:"-"`
	AvgName  string `json:"avgName" sql:"-"`
	HostGroupName  string `json:"hostGroupName" sql:"-"`
	HostName  string `json:"hostName" sql:"-"`
	ImagesUrl  string `json:"imagesUrl" sql:"-"`

	AppServicePorts []*AppServicePortDto `json:"appServicePorts"`
	AppServiceHosts []*AppServiceHostsDto `json:"appServiceHosts"`
	AppServiceDirs []*AppServiceDirDto `json:"appServiceDirs"`
	AppServiceVars []*AppServiceVarDto `json:"appServiceVars"`
}
