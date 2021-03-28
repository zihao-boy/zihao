package appVersionJob

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type AppVersionJobDetailDto struct {
	dto.PageDto
	DetailId string `json:"detailId" sql:"-"`
	JobId string `json:"jobId" sql:"-"`
	LogPath string `json:"logPath" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	State string `json:"state"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
