package installApp

import "github.com/zihao-boy/zihao/entity/dto"

type InstallAppDto struct {
	dto.PageDto
	AppId        string `json:"appId" sql:"-"`
	AppName      string `json:"appName" sql:"-"`
	Version      string `json:"version" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	ExtAppId     string `json:"extAppId" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
}

// install app page param

type InstallAppPageDto struct {
	AsType   string `json:"asType" sql:"-"`
	AsGroupId   string `json:"asGroupId" sql:"-"`
	GroupId   string `json:"groupId" sql:"-"`
	AsDeployType   string `json:"asDeployType" sql:"-"`
	AsDeployId   string `json:"asDeployId" sql:"-"`
	AppId   string `json:"appId" sql:"-"`
}
