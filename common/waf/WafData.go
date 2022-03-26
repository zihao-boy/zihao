package waf

import (
	"crypto/tls"
	"github.com/zihao-boy/zihao/entity/dto/waf"
)

var wafData WafData

type WafData struct {
	wafDto   waf.WafDto
	routes   []*waf.WafRouteDto
	wafCerts []WafCert
}

type WafCert struct {
	Hostname string
	TlsCert  *tls.Certificate
}
