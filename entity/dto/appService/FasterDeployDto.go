package appService

import "github.com/zihao-boy/zihao/entity/dto"

const(
	DeployTypeJava string = "java"
	DeployTypeCommon string = "common"
)

// faster deploy struct

type FasterDeployDto struct {
	dto.PageDto
	DeployId       string `json:"deployId" sql:"-"`
	AppName        string `json:"appName" sql:"-"`
	DeployType     string `json:"deployType" sql:"-"`
	TenantId       string `json:"tenantId" sql:"-"`
	PackageId      string `json:"packageId" sql:"-"`
	ShellPackageId string `json:"shellPackageId" sql:"-"`
	CreateTime     string `json:"createTime" sql:"-"`
	StatusCd       string `json:"statusCd" sql:"-"`
	AsGroupId      string `json:"asGroupId" sql:"-"`
	AsDeployType   string `json:"asDeployType" sql:"-"`
	AsDeployId     string `json:"asDeployId" sql:"-"`
	OpenPort       string `json:"openPort" sql:"-"`
	AvgName       string `json:"avgName" sql:"-"`
	HostGroupName       string `json:"hostGroupName" sql:"-"`
	HostName       string `json:"hostName" sql:"-"`
	PackageName       string `json:"packageName" sql:"-"`
	ShellPackageName       string `json:"shellPackageName" sql:"-"`
}
