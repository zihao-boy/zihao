package ruleAdapt

import (
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
)

func Rule(w http.ResponseWriter, r *http.Request,
	accessLog *waf.WafAccessLogDto,
	rules []*waf.WafRuleDataDto,
	tRouteDto *waf.WafRouteDto) error {

	if rules == nil || len(rules) < 1 {
		return nil
	}

	for _, rule := range rules {
		nextRule,err := doRule(w, r, accessLog, tRouteDto, rule)

		if err != nil {
			return err
		}
		// eg white ip
		if !nextRule {
			return nil
		}
	}

	return nil
}

func doRule(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (nextRule bool, err error) {

	var (
		ipRuleAdapt IpRuleAdapt
		areaRuleAdapt AreaRuleAdapt
	)
	nextRule = true

	switch {
	case rule.ObjType == waf.Waf_obj_type_ip:
		nextRule,err = ipRuleAdapt.validate(w, r, log, dto, rule)
	case rule.ObjType == waf.Waf_obj_type_area:
		nextRule,err = areaRuleAdapt.validate(w, r, log, dto, rule)
	default:
		err = nil
		nextRule = true
	}

	return nextRule,err
}
