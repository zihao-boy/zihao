package dns

import (
	dnsMap2 "github.com/zihao-boy/zihao/entity/dto/dns"
	"net"
)
var dnsService DnsServer
func StartDns(dto dnsMap2.DnsDataDto) error  {
	dnsService = DnsServer{
		memo:       addrBag{data: make(map[string][]net.UDPAddr)},
		forwarders: []net.UDPAddr{{IP: net.ParseIP(dto.Dns.DnsIp), Port:53}},
	}

	go dnsService.Listen()
	return nil
}

func StopDns() error {
	return dnsService.conn.Close()
}

func RefreshDns(dto dnsMap2.DnsDataDto) error {
	FreshDnsConfig(dto)
	dnsService.forwarders = []net.UDPAddr{{IP: net.ParseIP(dto.Dns.DnsIp), Port:53}}
	return nil
}