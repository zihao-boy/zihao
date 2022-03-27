package waf

import (
	"errors"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"strings"
)


type WafHandleRoute struct {
}

func (route *WafHandleRoute) GetWafRoute(r *http.Request) (*waf.WafRouteDto, error) {

	if len(wafData.routes) < 1 {
		return nil, errors.New("请配置路由")
	}

	for _,tRoute := range wafData.routes{
		if strings.Trim(tRoute.Hostname," ") == r.Host{
			return tRoute,nil
		}
	}

	for _,tRoute := range wafData.routes{
		if strings.Trim(tRoute.Hostname," ") == "*"{
			return tRoute,nil
		}
	}
	//return wafData.routes[0],nil
	return nil,errors.New("梓豪平台（waf）提醒您，当前资源没有权限访问！")
}

func (route *WafHandleRoute) LoadWafRoute() {

}

func (route *WafHandleRoute) StopLoadWafRoute() {

}
