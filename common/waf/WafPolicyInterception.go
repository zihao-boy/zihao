package waf

import (
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
)

type WafPolicyInterception struct {
	
}

func (policy *WafPolicyInterception)PolicyInterception(w http.ResponseWriter, r *http.Request,tRoute *waf.WafRouteDto) error {


	return nil
}
