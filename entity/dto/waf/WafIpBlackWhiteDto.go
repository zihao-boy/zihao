package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)



type WafIpBlackWhiteDto struct {
	dto.PageDto
	Id    string    `json:"id"  `
	TypeCd    string    `json:"typeCd" sql:"-" `
	Ip  string    `json:"ip"  `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
