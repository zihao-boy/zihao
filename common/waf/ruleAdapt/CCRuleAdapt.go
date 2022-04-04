package ruleAdapt

import (
	"errors"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var blockIps []*BlockIp

var ccIps []*CCIp

var expireTime time.Time = date.GetNowTime()

const (
	default_expire_time = "30000s"
)

type BlockIp struct {
	Ip         string
	expireTime time.Time
}

type CCIp struct {
	Path       string
	Ip         string
	Count      int64
	CreateTime time.Time
	VisitSec string
}

type CCRuleAdapt struct {
}

func (cc *CCRuleAdapt) validate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (nextRule bool, err error) {

	defer func() {
		// clear ip data
		cc.clearIpData();
	}()
	nextRule, err = cc.ccValidate(w, r, log, dto, rule)

	if err != nil {
		log.State = waf.State_cc
		log.Message = "CC防护"
	}

	return nextRule, err
}

// black area
func (cc *CCRuleAdapt) ccValidate(w http.ResponseWriter,
	r *http.Request,
	log *waf.WafAccessLogDto,
	dto *waf.WafRouteDto,
	rule *waf.WafRuleDataDto) (bool, error) {
	srcIp := log.XRealIp
	err := cc.getBlockIp(srcIp)
	if err != nil {
		return false, err
	}

	matched, _ := regexp.Match(rule.CC.Path, []byte(log.Url))

	if !matched {
		return true, nil
	}


	ccIp := cc.getCCIp(srcIp,rule.CC.Path,rule.CC.VisitSec)

	visitCount,err := strconv.ParseInt(rule.CC.VisitCount,10,64)

	if err != nil{
		return true,nil
	}

	if ccIp.Count >= visitCount{
		sec,err := time.ParseDuration(rule.CC.BlockSec+"s")
		if err != nil{
			return true,nil
		}
		endTime := date.GetNowTime().Add(sec)
		blockIp := BlockIp{
			Ip:srcIp,
			expireTime: endTime,
		}
		blockIps = append(blockIps,&blockIp)
		return false,errors.New("请休息一会再访问")
	}

	return true, nil
}

// check blockIp
func (cc *CCRuleAdapt) getBlockIp(srcIp string) error {
	if len(blockIps) < 1 {
		return nil
	}

	nowTime := date.GetNowTime()
	for _, tBlockIp := range blockIps {
		if srcIp == tBlockIp.Ip && tBlockIp.expireTime.Before(nowTime) {
			return errors.New("请休息一会再访问")
		}
	}
	return nil
}

func (cc *CCRuleAdapt) getCCIp(ip string,path string,visitSec string) CCIp {
	var ccIp CCIp
	if len(ccIps) < 1 {
		ccIp = CCIp{
			Ip: ip,
			Path:path,
			Count: 1,
			CreateTime: date.GetNowTime(),
			VisitSec: visitSec,
		}
		ccIps = append(ccIps,&ccIp)
		return ccIp
	}

	nowTime := date.GetNowTime()

	for _, tCCIp := range ccIps {
		sec,err := time.ParseDuration(visitSec+"s")
		if err != nil{
			continue
		}
		endTime := tCCIp.CreateTime.Add(sec)
		if ip == tCCIp.Ip && tCCIp.Path == path && nowTime.Before(endTime){
			tCCIp.Count +=1
			return *tCCIp
		}
	}
	ccIp = CCIp{
		Ip: ip,
		Path:path,
		Count: 1,
		CreateTime: date.GetNowTime(),
		VisitSec: visitSec,
	}
	ccIps = append(ccIps,&ccIp)
	return ccIp
}
// clear data
func (cc *CCRuleAdapt) clearIpData() {

	defaultTime,_ :=time.ParseDuration(default_expire_time)

	tExpireTime := expireTime.Add(defaultTime)
	nowTime := date.GetNowTime()
	if nowTime.Before(tExpireTime){
		return
	}
	expireTime = nowTime
	var tBlockIps []*BlockIp
	for _, tBlockIp := range blockIps {
		if nowTime.Before(tBlockIp.expireTime) {
			continue
		}
		tBlockIps = append(tBlockIps,tBlockIp)
	}

	blockIps = tBlockIps

	var tCCIps []*CCIp
	for _, tCCIp := range ccIps {
		sec,err := time.ParseDuration(tCCIp.VisitSec+"s")
		if err != nil{
			continue
		}
		endTime := tCCIp.CreateTime.Add(sec)
		if endTime.Before(nowTime) {
			continue
		}

		tCCIps = append(tCCIps,tCCIp)
	}
	ccIps = tCCIps
}
