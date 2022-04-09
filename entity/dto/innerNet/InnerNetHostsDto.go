package innerNet

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type InnerNetHostsDto struct {
	dto.PageDto
	InnerNetHostId     string    `json:"innerNetHostId" sql:"-" `
	InnerNetId     string    `json:"innerNetId" sql:"-" `
	HostId     string    `json:"hostId" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}

