package waf

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/ip"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net"
	"net/http"
	"os"
)

var wafServer WafServer
// waf server
// author wuxw
type WafServer struct {
	httpListener   net.Listener
	httpListeners  net.Listener
	wafHandleRoute WafHandleRoute
}

// start waf
func (waf *WafServer) StartWaf(wafDataDto waf.SlaveWafDataDto) error {
	waf.InitWafConfig(wafDataDto)

	//init ip sets
	go waf.initIps()

	gateMux := http.NewServeMux()
	gateMux.HandleFunc("/", RootHandle)
	ctxGateMux := AddContextHandler(gateMux)
	// start http port server
	go waf.startHttpServer(wafDataDto.Waf.Port, ctxGateMux)
	go waf.startHttpsServer(wafDataDto.Waf.HttpsPort, ctxGateMux)
	return nil
}

func (waf *WafServer) InitWafConfig(wafDataDto waf.SlaveWafDataDto) error {
	var tmpWafCerts []WafCert
	wafData.wafDto = wafDataDto.Waf
	//load route
	wafData.routes = wafDataDto.Routes

	// rule
	wafData.rules = wafDataDto.Rules

	if wafDataDto.Certs != nil && len(wafDataDto.Certs) > 0 {
		for _, cert := range wafDataDto.Certs {
			tlsCert, err := tls.X509KeyPair([]byte(cert.CertContent), []byte(cert.PrivKeyContent))
			if err != nil {
				fmt.Println("RPCSelectCertificates X509KeyPair", err)
				continue
			}
			tmpWafCert := WafCert{
				Hostname: cert.Hostname,
				TlsCert:  &tlsCert,
			}
			tmpWafCerts = append(tmpWafCerts, tmpWafCert)
		}
	}
	wafData.wafCerts = tmpWafCerts

	ip.Service_IP = ip.GetServerIp(wafDataDto.ServerIpUrl)

	return nil
}

// stop waf
func (waf *WafServer) StopWaf() (err error) {
	if wafServer.httpListener != nil {
		err = wafServer.httpListener.Close()
	}
	if   wafServer.httpListeners != nil{
		err = wafServer.httpListeners.Close()
	}
	return err
}

func (waf *WafServer) startHttpServer(httpPort string, ctxGateMux http.Handler) {
	var err error
	wafServer.httpListener, err= net.Listen("tcp", ":"+httpPort)
	if err != nil {
		msg := "Port " + httpPort + " is occupied."
		fmt.Println(msg, err)
		return ;
	}
	fmt.Println("Listen HTTP ", httpPort)
	// err = http.Serve(listen, ctxGateMux)
	err = http.Serve(wafServer.httpListener, ctxGateMux)
	if err != nil {
		fmt.Println("http.Serve error", err)
		return
	}
	defer wafServer.httpListener.Close()

}

func (waf *WafServer) startHttpsServer(httpsPort string, ctxGateMux http.Handler) {
	var err error
	tlsconfig := &tls.Config{
		GetCertificate: func(helloInfo *tls.ClientHelloInfo) (*tls.Certificate, error) {
			cert, err := GetCertificateByDomain(helloInfo)
			return cert, err
		},
		NextProtos: []string{"h2", "http/1.1"},
		MaxVersion: tls.VersionTLS13,
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
//			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
//			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		},
	}

	wafServer.httpListeners, err = tls.Listen("tcp", ":"+httpsPort, tlsconfig)
	if err != nil {
		msg := "Port " + httpsPort + " is occupied."
		fmt.Println(msg, err)
		os.Exit(1)
	}
	//err = http.Serve(listen, ctxGateMux)
	err = http.Serve(wafServer.httpListeners, ctxGateMux)
	if err != nil {
		fmt.Println("http.Serve error", err)
		return
	}
	defer wafServer.httpListeners.Close()
}

func (waf *WafServer) initIps() {

	mastIp, isExist := config.Prop.Property("mastIp")
	if !isExist {
		mastIp = "127.0.0.1:7000"
	}
	url := "http://" + mastIp + "/app/firewall/loadIps"
	data := map[string]interface{}{}
	resp, err := httpReq.SendRequest(url, data, nil,"GET")
	if err != nil {
		fmt.Println(err)
		return
	}

	ip.IPData.InitIPDataByByte(resp)

}

// GetCertificateByDomain ...
func GetCertificateByDomain(helloInfo *tls.ClientHelloInfo) (*tls.Certificate, error) {
	domain := helloInfo.ServerName

	for _, wafCert := range wafData.wafCerts {
		if domain == wafCert.Hostname {
			return wafCert.TlsCert, nil
		}
	}
	return nil, errors.New("Unknown Host: " + domain)
}
