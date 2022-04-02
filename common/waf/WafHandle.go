package waf

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/zihao-boy/zihao/common/waf/ruleAdapt"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// AddContextHandler to add context handler
func AddContextHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// map[GroupPolicyID int64](Value int64)
		ctx := context.WithValue(r.Context(), "groupPolicyHitValue", &sync.Map{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RootHandle(w http.ResponseWriter, r *http.Request) {
	var wafHandleRoute WafHandleRoute
	var wafPolicyInterception WafPolicyInterception

	// analysis access log
	accessLog := analysisRequest(r)
	defer func() {
		go saveAccessLog(accessLog)
	}()

	tRoute, err := wafHandleRoute.GetWafRoute(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	refreshAccessLogByRoute(&accessLog,tRoute)

	// rule
	err = ruleAdapt.Rule(w,r,accessLog,wafData.rules,tRoute)

	if err != nil {
		w.Write([]byte("梓豪平台提醒您，"+err.Error()))
		return
	}


	//&& (app.RedirectHTTPS || app.HSTSEnabled)
	if r.TLS == nil && tRoute.Scheme == waf.Scheme_https {
		RedirectRequest(w, r, "https://"+tRoute.Ip+":"+tRoute.Port+r.URL.Path)
		return
	}

	// firewall start
	err = wafPolicyInterception.PolicyInterception(w, r, tRoute)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// firewall end

	transport := &http.Transport{
		TLSHandshakeTimeout:   30 * time.Second,
		IdleConnTimeout:       30 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := net.Dial("tcp", tRoute.Ip+":"+tRoute.Port)
			return conn, err
		},
		DialTLS: func(network, addr string) (net.Conn, error) {
			conn, err := net.Dial("tcp", tRoute.Ip+":"+tRoute.Port)
			if err != nil {
				return nil, err
			}
			cfg := &tls.Config{
				ServerName:         r.Host,
				NextProtos:         []string{"h2", "http/1.1"},
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			}
			tlsConn := tls.Client(conn, cfg)
			if err := tlsConn.Handshake(); err != nil {
				fmt.Println("tlsConn.Handshake error", err)
			}
			return tlsConn, err //net.Dial("tcp", dest)
		},
	}
	err = http2.ConfigureTransport(transport)
	if err != nil {
		fmt.Println("http2.ConfigureTransport error", err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = tRoute.Scheme
			req.URL.Host = r.Host
		},
		Transport:      transport,
		ModifyResponse: func(response *http.Response) error {
			return rewriteResponse(response,&accessLog)
		},
	}
	//if utils.Debug {
	//	dump, err := httputil.DumpRequest(r, true)
	//	if err != nil {
	//		utils.DebugPrintln("ReverseHandlerFunc DumpRequest", err)
	//	}
	//	fmt.Println(string(dump))
	//}

	proxy.ServeHTTP(w, r)

}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// RedirectRequest for example: redirect 80 to 443
func RedirectRequest(w http.ResponseWriter, r *http.Request, location string) {
	if len(r.URL.RawQuery) > 0 {
		location += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, location, http.StatusPermanentRedirect)
}

func rewriteResponse(resp *http.Response,accessLog *waf.WafAccessLogDto) (err error) {
	r := resp.Request
	accessLog.ResponseCode = strconv.Itoa(resp.StatusCode)
	accessLog.ResponseLength = strconv.FormatInt(resp.ContentLength,10)
	locationStr := resp.Header.Get("Location")
	indexHTTP := strings.Index(locationStr, "http")
	if indexHTTP == 0 {
		locationURL, _ := resp.Location()
		host := locationURL.Hostname()
		port := locationURL.Port()
		if host == r.Host {
			var oldHost, newHost string
			if (port == "") || (port == "80") || (port == "443") {
				oldHost = host
			} else {
				oldHost = host + ":" + port
			}
			var userScheme string
			if resp.Request.TLS != nil {
				userScheme = "https"
				if wafData.wafDto.HttpsPort == ":443" {
					newHost = host
				} else {
					newHost = host + wafData.wafDto.HttpsPort
				}
			} else {
				userScheme = "http"
				if wafData.wafDto.Port == ":80" {
					newHost = host
				} else {
					newHost = host + wafData.wafDto.Port
				}
			}
			newLocation := strings.Replace(locationURL.String(), oldHost, newHost, -1)
			newLocation = strings.Replace(newLocation, locationURL.Scheme, userScheme, 1)
			resp.Header.Set("Location", newLocation)
		}
	}

	// Hide X-Powered-By
	xPoweredBy := resp.Header.Get("X-Powered-By")
	if xPoweredBy != "" {
		resp.Header.Set("X-Powered-By", "ZiHaoWaf")
	}

	// if client http and backend https, remove "; Secure" and replace https by http
	//&& (app.InternalScheme == "https")
	if r.TLS == nil {
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			re := regexp.MustCompile(`;\s*Secure`)
			cookieStr := re.ReplaceAllLiteralString(cookie.Raw, "")
			resp.Header.Set("Set-Cookie", cookieStr)
		}
		origin := resp.Header.Get("Access-Control-Allow-Origin")
		if len(origin) > 0 {
			resp.Header.Set("Access-Control-Allow-Origin", strings.Replace(origin, "https", "http", 1))
		}
		csp := resp.Header.Get("Content-Security-Policy")
		if len(csp) > 0 {
			resp.Header.Set("Content-Security-Policy", strings.Replace(origin, "https", "http", -1))
		}
	}
	//body, err := httputil.DumpResponse(resp, true)
	//fmt.Println("Dump Response:")
	//fmt.Println(string(body))
	return nil
}
