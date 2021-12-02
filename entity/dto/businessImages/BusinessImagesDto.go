package businessImages

import "github.com/zihao-boy/zihao/entity/dto"

type BusinessImagesDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Version      string `json:"version" sql:"-"`
	ImagesType   string `json:"imagesType" sql:"-"`
	TypeUrl   string `json:"typeUrl" sql:"-"`
	ImagesFlag   string `json:"imagesFlag" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Username     string `json:"username" sql:"-"`
}