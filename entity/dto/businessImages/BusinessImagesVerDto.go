package businessImages

import "github.com/zihao-boy/zihao/entity/dto"

type BusinessImagesVerDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	ImagesId         string `json:"imagesId" sql:"-"`
	Version      string `json:"version" sql:"-"`
	TypeUrl   string `json:"typeUrl" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
}