package waf

import (
	"errors"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
)

var routes []waf.WafRouteDto

type WafHandleRoute struct {
}

func (route *WafHandleRoute) GetWafRouteUrl(log waf.WafAccessLogDto, r *http.Request) (string, error) {

	if len(routes) < 1 {
		return "", errors.New("请配置路由")
	}

	for _,tRoute := range routes{
		if tRoute.Hostname == log.HttpHost{
			return "http://"+tRoute.Ip+":"+tRoute.Port,nil
		}
	}
	return "",errors.New("未找到【"+log.HttpHost+"】路由配置，请配置路由")
}

func (route *WafHandleRoute) LoadWafRoute() {

}
