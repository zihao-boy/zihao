package resources

import "github.com/zihao-boy/zihao/entity/dto"

//ftp struct
type ResourcesFtpDto struct {
	dto.PageDto

	FtpId      string `json:"ftpId" sql:"-"`
	Name       string `json:"name" `
	Ip         string `json:"ip" `
	Port       string `json:"port" `
	Username   string `json:"username" `
	Passwd     string `json:"passwd" `
	Path       string `json:"path" `
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
}
// oss struct

type ResourcesOssDto struct {
	dto.PageDto

	OssId      string `json:"ossId" sql:"-"`
	Name       string `json:"name" `
	OssType       string `json:"ossType" sql:"-" `
	Bucket         string `json:"bucket" `
	AccessKeySecret       string `json:"accessKeySecret" sql:"-"`
	AccessKeyId   string `json:"accessKeyId" sql:"-"`
	Endpoint     string `json:"endpoint" `
	Path       string `json:"path" `
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
}

// oss struct
type ResourcesDbDto struct {
	dto.PageDto
	DbId      string `json:"dbId" sql:"-"`
	Name       string `json:"name" `
	Ip         string `json:"ip" `
	Port       string `json:"port" `
	Username   string `json:"username" `
	Password     string `json:"password" `
	DbName       string `json:"dbName" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
}

// back up struct
type ResourcesBackUpDto struct {
	dto.PageDto
	Id      string `json:"id" `
	Name       string `json:"name" `
	ExecTime         string `json:"execTime" sql:"-"`
	TypeCd       string `json:"typeCd" sql:"-"`
	SrcId   string `json:"srcId" sql:"-"`
	SrcObject     string `json:"srcObject" sql:"-" `
	TargetId       string `json:"targetId" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
}



