package dns

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const Dns_state_stop = "2002"
const Dns_state_start = "1001"

type DnsDto struct {
	dto.PageDto
	DnsId      string         `json:"dnsId" sql:"-" `
	DnsIp string `json:"dnsIp" sql:"-"`
	DnsPort string `json:"dnsPort"`
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
	State      string         `json:"state"`
	DnsHosts   []*DnsHostsDto `json:"dnsHosts"`
	HostIds    string         `json:"hostIds"`
}

type DnsDataDto struct {
	Dns         DnsDto `json:"dns"`
	Maps []*DnsMapDto
}
