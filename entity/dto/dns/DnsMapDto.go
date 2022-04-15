package dns

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const (
	Type_A = "A"
)
type DnsMapDto struct {
	dto.PageDto
	DnsMapId string `json:"dnsMapId"`
	Host string `json:"host"`
	Type string `json:"type"`
	Value string `json:"value"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
