package ruleAdapt

import (
	"errors"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"strings"
)

type AccurateRuleAdapt struct {
}

func (accurate *AccurateRuleAdapt) validate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (nextRule bool, err error) {

	if rule.Accurate.Action == waf.Waf_ip_black_white_type_W {
		nextRule, err = accurate.whiteValidate(w, r, log, dto, rule)
	} else {
		nextRule, err = accurate.blackValidate(w, r, log, dto, rule)
	}

	if err != nil {
		log.State = waf.State_custom_custom
		log.Message = "精准防护"
	}

	return nextRule, err
}

// white accurate
func (accurate *AccurateRuleAdapt) whiteValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool, error) {

	var (
		includeFlag bool
		headerKey   string
		headerValue string
	)

	// has in
	includeFlag = accurate.hasMatch(rule, log, headerKey, headerValue, r)

	if rule.Accurate.Include == "Y" && includeFlag {
		return true, nil
	}

	if rule.Accurate.Include == "N" && !includeFlag {
		return true, nil
	}
	return false, errors.New("该资源无法访问")
}

// black ip
func (accurate *AccurateRuleAdapt) blackValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool, error) {
	var (
		includeFlag bool
		headerKey   string
		headerValue string
	)

	// has in
	includeFlag = accurate.hasMatch(rule, log, headerKey, headerValue, r)

	if rule.Accurate.Include == "Y" && includeFlag {
		return false, errors.New("该资源无法访问")
	}

	if rule.Accurate.Include == "N" && !includeFlag {
		return false, errors.New("该资源无法访问")
	}
	return true, nil
}

func (accurate *AccurateRuleAdapt) hasMatch(rule *waf.WafRuleDataDto, log *waf.WafAccessLogDto, headerKey string, headerValue string, r *http.Request) bool {
	var includeFlag bool
	if rule.Accurate.TypeCd == waf.Waf_accurate_type_url {
		includeFlag = strings.Contains(log.Url, rule.Accurate.IncludeValue)
	} else {
		includeFlag = true
		if strings.Contains(rule.Accurate.IncludeValue, "=") {
			headerKey = strings.Trim(strings.Split(rule.Accurate.IncludeValue, "=")[0], " ")
			headerValue = strings.Trim(strings.Split(rule.Accurate.IncludeValue, "=")[1], " ")
		} else {
			headerKey = rule.Accurate.IncludeValue
			headerValue = ""
		}
		tHeaderValue := r.Header.Get(headerKey)
		if utils.IsEmpty(tHeaderValue) {
			includeFlag = false
		}
		if !utils.IsEmpty(headerValue) && tHeaderValue != tHeaderValue {
			includeFlag = false
		}
	}
	return includeFlag
}
