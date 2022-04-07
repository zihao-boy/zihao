package vpn

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type VpnUserDto struct {
	dto.PageDto
	UserId     string    `json:"userId" sql:"-" `
	Username   string    `json:"username"  `
	Password   string    `json:"password" `
	Tel        string    `json:"tel"`
	Ip         string    `json:"ip"`
	LoginTime  time.Time `json:"loginTime" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	Token string `json:"token"`
}
