package dnsMap

const (
	Type_A = "A"
)
type DnsMapDto struct {
	Host string `json:"host"`
	Type string `json:"type"`
	Value string `json:"value"`
}
