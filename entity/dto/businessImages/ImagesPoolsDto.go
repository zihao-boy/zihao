package businessImages

import "github.com/zihao-boy/zihao/entity/dto"

type ImagesPoolsDto struct {
	dto.PageDto
	AppId string `json:"appId" sql:"-"`
	Name  string `json:"name"`
	Version string `json:"version"`
	Compose string `json:"compose"`
	ZihaoAppImagesDtos []ZihaoAppImagesDtos `json:"zihaoAppImagesDtos"`
}

