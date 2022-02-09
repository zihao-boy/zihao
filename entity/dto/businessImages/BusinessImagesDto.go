package businessImages

import "github.com/zihao-boy/zihao/entity/dto"

const (
	IMAGES_TYPE_REMOTE = "1" //远程
	IMAGES_TYPE_DOCKER = "2" //dockerfile生成镜像
	IMAGES_TYPE_IMPORT = "3" // 导入
	IMAGES_FLAG_PUBLIC = "P"
	IMAGES_FLAG_CUSTOM = "C"
)

type BusinessImagesDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Version      string `json:"version" sql:"-"`
	ImagesType   string `json:"imagesType" sql:"-"`
	TypeUrl      string `json:"typeUrl" sql:"-"`
	ImagesFlag   string `json:"imagesFlag" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Username     string `json:"username" sql:"-"`
	BusinessImagesExtDto
}

// images ext
type BusinessImagesExtDto struct {
	dto.PageDto
	Id       string `json:"id" sql:"-"`
	ImagesId string `json:"imagesId" sql:"-"`
	AppId    string `json:"appId" sql:"-"`
	AppName    string `json:"appName" sql:"-"`
	ExtImagesId    string `json:"extImagesId" sql:"-"`
	ExtPublisherId    string `json:"extPublisherId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
}
