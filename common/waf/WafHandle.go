package waf

import (
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net"
	"net/http"
	"strings"
)

func RootHandle(w http.ResponseWriter, r *http.Request)  {
	var wafHandleRoute WafHandleRoute

	accessLog := analysisRequest(r)

	url,err := wafHandleRoute.GetWafRouteUrl(accessLog,r)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	req,err := http.NewRequest(r.Method,url,r.Body)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	accessLog.RequestLength = string(r.ContentLength)
	accessLog.ResponseCode = req.Response.Status
	accessLog.ResponseLength = string(req.Response.ContentLength)
	req.Write(w)
}


// analysis
func analysisRequest(r *http.Request) waf.WafAccessLogDto{
	accessLog := waf.WafAccessLogDto{
		XRealIp:ClientIP(r),
		Scheme:"",
		ResponseCode:"200",
		Method:r.Method,
		HttpHost:r.Host,
		UpstreamAddr:"",
		Url:r.Host+r.URL.String(),
		RequestLength:"0",
		ResponseLength:"0",
		State:"",
		Message:"",
	}
	return accessLog
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

