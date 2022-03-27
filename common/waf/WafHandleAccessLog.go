package waf

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/ip"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"strconv"
)

func saveAccessLog(wafAccessLogDto waf.WafAccessLogDto) {


	mastIp, isExist := config.Prop.Property("mastIp")
	if !isExist {
		mastIp = "127.0.0.1:7000"
	}
	url := "http://" + mastIp + "/app/firewall/saveWafAccessLog"

	resp, err := httpReq.SendRequest(url, wafAccessLogDto, nil,"POST")
	if err != nil {
		fmt.Print(err.Error(), url, wafAccessLogDto)
	}
	fmt.Print(resp)

	fmt.Println(wafAccessLogDto)
}

// analysis
func analysisRequest(r *http.Request) waf.WafAccessLogDto {
	wafIp := ip.Service_IP
	slaveId, _ := config.Prop.Property("slaveId")
	accessLog := waf.WafAccessLogDto{
		RequestId:      seq.Generator(),
		WafIp:          wafIp,
		XRealIp:        ClientIP(r),
		Scheme:         "",
		HostId:         slaveId,
		ResponseCode:   "200",
		Method:         r.Method,
		HttpHost:       r.Host,
		UpstreamAddr:   "",
		Url:            r.Host + r.URL.String(),
		RequestLength:  strconv.FormatInt(r.ContentLength, 10),
		ResponseLength: "0",
		State:          waf.State_default,
		Message:        "正常访问",
		CreateTime: date.GetNowTime(),
	}
	return accessLog
}

func refreshAccessLogByRoute(accessLog *waf.WafAccessLogDto, tRoute *waf.WafRouteDto) {
	accessLog.UpstreamAddr = tRoute.Scheme + "://" + tRoute.Ip + ":" + tRoute.Port
	accessLog.Url = tRoute.Scheme + "://" + accessLog.Url
	accessLog.Scheme = tRoute.Scheme
}


