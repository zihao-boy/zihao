package waf

import (
	"errors"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
)


type WafHandleRoute struct {
}

func (route *WafHandleRoute) GetWafRoute(r *http.Request) (*waf.WafRouteDto, error) {

	if len(wafData.routes) < 1 {
		return nil, errors.New("请配置路由")
	}

	for _,tRoute := range wafData.routes{
		if tRoute.Hostname == r.Host{
			return tRoute,nil
		}
	}
	return nil,errors.New("未找到【"+r.Host+"】路由配置，请配置路由")
}

func (route *WafHandleRoute) LoadWafRoute() {

}

func (route *WafHandleRoute) StopLoadWafRoute() {

}
