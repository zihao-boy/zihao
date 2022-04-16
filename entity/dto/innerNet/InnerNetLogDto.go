package innerNet

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type InnerNetLogDto struct {
	dto.PageDto
	LogId      string    `json:"logId" sql:"-" `
	Username   string    `json:"username"  `
	SrcIp      string    `json:"srcIp" sql:"-" `
	State      string    `json:"state"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	Token      string    `json:"token"`
}
