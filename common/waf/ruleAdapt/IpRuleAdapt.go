package ruleAdapt

import (
	"errors"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"strconv"
	"strings"
)

type IpRuleAdapt struct {

}

func (ip *IpRuleAdapt)validate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (nextRule bool,err error) {

	if rule.Ip.TypeCd == waf.Waf_ip_black_white_type_W{
		nextRule,err = ip.whiteValidate(w,r,log,dto,rule)
	}else{
		nextRule,err = ip.blackValidate(w,r,log,dto,rule)
	}

	if err != nil{
		log.State = waf.State_custom_whiteip
		log.Message = "黑白名单"
	}



	return nextRule,err;
}

// white ip
func (ip *IpRuleAdapt) whiteValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool,error) {
	srcIp := log.XRealIp
	srcIps := strings.Split(srcIp,".")

	if len(srcIps) != 4 {
		return false,errors.New("您当前没有权限访问")
	}

	if rule.Ip == nil{
		return true,nil
	}
	ruleIp := rule.Ip.Ip

	ruleIps := strings.Split(ruleIp,".")

	if len(ruleIps) != 4{
		return true,nil
	}

	if srcIps[0] != ruleIps[0] || srcIps[1] != ruleIps[1] || srcIps[2] != ruleIps[2]{
		return false,errors.New("您当前没有权限访问")
	}
	// white ip continue
	if srcIps[3] == ruleIps[3]{
		return true,nil
	}

	if !strings.Contains(ruleIps[3],"-") && srcIps[3] != ruleIps[3]{
		return false,errors.New("您当前没有权限访问")
	}

	 nRuleIps := strings.Split(ruleIps[3],"-")

	minIp4, err := strconv.ParseInt(nRuleIps[0], 10, 64)
	if err != nil{
		return true,nil
	}

	maxIp4, err := strconv.ParseInt(nRuleIps[1], 10, 64)
	if err != nil{
		return true,nil
	}

	ip4, err := strconv.ParseInt(srcIps[3], 10, 64)
	if err != nil{
		return false,errors.New("您当前没有权限访问")
	}

	if ip4 < minIp4 || ip4 > maxIp4{
		return false,errors.New("您当前没有权限访问")
	}

	return true,nil;
}

// black ip
func (ip *IpRuleAdapt) blackValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool,error) {
	srcIp := log.XRealIp
	srcIps := strings.Split(srcIp,".")

	if len(srcIps) != 4 {
		return false,errors.New("您当前没有权限访问")
	}

	ruleIp := rule.Ip.Ip

	ruleIps := strings.Split(ruleIp,".")
	//config error
	if len(ruleIps) != 4{
		return true,nil
	}

	if srcIps[0] != ruleIps[0] || srcIps[1] != ruleIps[1] || srcIps[2] != ruleIps[2]{
		return true,nil
	}

	// white ip continue
	if srcIps[3] == ruleIps[3]{
		return false,errors.New("您当前没有权限访问")
	}


	if !strings.Contains(ruleIps[3],"-") && srcIps[3] != ruleIps[3]{
		return true,nil
	}
	nRuleIps := strings.Split(ruleIps[3],"-")
	//config error
	minIp4, err := strconv.ParseInt(nRuleIps[0], 10, 64)
	if err != nil{
		return true,nil
	}
	//config error
	maxIp4, err := strconv.ParseInt(nRuleIps[1], 10, 64)
	if err != nil{
		return true,nil
	}

	ip4, err := strconv.ParseInt(srcIps[3], 10, 64)
	if err != nil{
		return true,nil
	}

	if ip4 < minIp4 || ip4 > maxIp4{
		return true,nil
	}

	return false,errors.New("您当前没有权限访问")
}