package ruleAdapt

import (
	"errors"
	"github.com/zihao-boy/zihao/common/ip"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"strings"
)

type AreaRuleAdapt struct {
}

func (area *AreaRuleAdapt) validate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (nextRule bool, err error) {

	if rule.Area.TypeCd == waf.Waf_area_type_W {
		nextRule, err = area.whiteValidate(w, r, log, dto, rule)
	} else {
		nextRule, err = area.blackValidate(w, r, log, dto, rule)
	}

	if err != nil {
		log.State = waf.State_custom_geoip
		log.Message = "地理位置"
	}

	return nextRule, err
}

// white area
func (area *AreaRuleAdapt) whiteValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool, error) {
	srcArea := ip.GetOnlyIpAddr(log.XRealIp).Country

	if rule.Area == nil {
		return true, nil
	}
	ruleArea := rule.Area.AreaName

	ruleAreas := strings.Split(ruleArea, ",")

	for _, area := range ruleAreas {
		if strings.Contains(srcArea,area) {
			return false, nil
		}
	}

	return false, errors.New("当前区域没有权限访问")

}

// black area
func (area *AreaRuleAdapt) blackValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool, error) {
	srcArea := ip.GetOnlyIpAddr(log.XRealIp).Country
	if rule.Area == nil {
		return true, nil
	}
	ruleArea := rule.Area.AreaName

	ruleAreas := strings.Split(ruleArea, ",")

	for _, area := range ruleAreas {
		if strings.Contains(srcArea,area) {
			return false, errors.New("当前区域没有权限访问")
		}
	}
	return true, nil
}
