package dns

import "github.com/zihao-boy/zihao/entity/dto"

const (
	Type_A = "A"
)
type DnsMapDto struct {
	dto.PageDto
	DnsMapId string `json:"dnsMapId"`
	Host string `json:"host"`
	Type string `json:"type"`
	Value string `json:"value"`
}
