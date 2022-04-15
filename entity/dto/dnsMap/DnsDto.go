package dnsMap

type DnsDto struct {
	DnsIp string `json:"dnsIp" sql:"-"`
	Port int `json:"port"`
}

type DnsDataDto struct {
	DnsDto
	Maps []*DnsMapDto
}
