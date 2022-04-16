package innerNet

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type InnerNetPrivilegeDto struct {
	dto.PageDto
	PId     string    `json:"pId" sql:"-" `
	SrcUserId   string    `json:"srcUserId" sql:"-"  `
	TargetUserId   string    `json:"targetUserId" sql:"-"  `
	TargetPort string `json:"targetPort" sql:"-" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	Token string `json:"token"`
}
