package wafService

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
	"strings"
)

type WafService struct {
	wafDao      wafDao.WafDao
	wafHostsDao wafDao.WafHostsDao
}

// get db link
// all db by this user
func (wafService *WafService) GetWafAll(WafDto waf.WafDto) ([]*waf.WafDto, error) {
	var (
		err     error
		WafDtos []*waf.WafDto
	)

	WafDtos, err = wafService.wafDao.GetWafs(WafDto)
	if err != nil {
		return nil, err
	}

	return WafDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafService) GetWafs(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafDto{}
		wafDtos []*waf.WafDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	wafDto.Row = row * page

	wafDto.Page = (page - 1) * row

	total, err = wafService.wafDao.GetWafCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafs(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, wafDto := range wafDtos {

		wafHostsDto := waf.WafHostsDto{
			WafId: wafDto.WafId,
		}
		wafHostsDtos, _ := wafService.wafHostsDao.GetWafHostss(wafHostsDto)
		wafDto.WafHosts = wafHostsDtos
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafService) SaveWafs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.WafId = seq.Generator()
	//WafDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafService) UpdateWafs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if utils.IsEmpty(wafDto.HostIds) {
		return result.SuccessData(wafDto)
	}

	wafHostsDto := waf.WafHostsDto{
		WafId: wafDto.WafId,
	}
	wafService.wafHostsDao.DeleteWafHosts(wafHostsDto)

	for _, hostId := range strings.Split(wafDto.HostIds, ",") {

		wafHostsDto = waf.WafHostsDto{
			WafId:     wafDto.WafId,
			WafHostId: seq.Generator(),
			HostId:    hostId,
		}
		wafService.wafHostsDao.SaveWafHosts(wafHostsDto)
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafService) DeleteWafs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

func (wafService *WafService) StartWaff(ctx iris.Context) interface{} {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start waf
	tmpWafDto := waf.WafDto{
		WafId: wafDto.WafId,
	}
	wafDtos, err := wafService.wafDao.GetWafs(tmpWafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(wafDtos) < 1 {
		return result.Error("未查询到数据")
	}

	resultDto, _ := shell.ExecStartWaf(wafService.getWafConfig(*wafDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpWafDto = waf.WafDto{
		WafId: wafDto.WafId,
		State: waf.Waf_state_start,
	}

	err = wafService.wafDao.UpdateWaf(tmpWafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(wafDto)

}

func (wafService *WafService) StopWaff(ctx iris.Context) interface{} {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start waf
	tmpWafDto := waf.WafDto{
		WafId: wafDto.WafId,
	}
	wafDtos, err := wafService.wafDao.GetWafs(tmpWafDto)

	if err != nil || len(wafDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecStopWaf(wafService.getWafConfig(*wafDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpWafDto = waf.WafDto{
		WafId: wafDto.WafId,
		State: waf.Waf_state_stop,
	}

	err = wafService.wafDao.UpdateWaf(tmpWafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(wafDto)

}



func (wafService *WafService) RefreshWafConfig(ctx iris.Context) interface{} {
	var (
		err    error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	// start waf
	tmpWafDto := waf.WafDto{
		WafId: wafDto.WafId,
	}
	wafDtos, err := wafService.wafDao.GetWafs(tmpWafDto)

	if err != nil || len(wafDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecRefreshWafConfig(wafService.getWafConfig(*wafDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	return result.SuccessData(wafDto)
}

func (wafService *WafService) getWafConfig(wafDto waf.WafDto) waf.SlaveWafDataDto {
	var (
		wafRouteDao        wafDao.WafRouteDao
		wafHostnameCertDao wafDao.WafHostnameCertDao
		wafRuleGroupDao wafDao.WafRuleGroupDao
	)

	wafHostsDto := waf.WafHostsDto{
		WafId: wafDto.WafId,
	}
	wafHostsDtos, _ := wafService.wafHostsDao.GetWafHostss(wafHostsDto)
	wafDto.WafHosts = wafHostsDtos

	// start waf
	tmpWafRouteDto := waf.WafRouteDto{
		WafId: wafDto.WafId,
	}
	routes, _ := wafRouteDao.GetWafRoutes(tmpWafRouteDto)
	tmpWafHostnameCertDto := waf.WafHostnameCertDto{
	}
	certs, _ := wafHostnameCertDao.GetWafHostnameCerts(tmpWafHostnameCertDto)

	// query rule group
	tWafRuleGroupDto := waf.WafRuleGroupDto{
		State: waf.Waf_Rule_State_T,
	}
	tRuleGrops, _ := wafRuleGroupDao.GetWafRuleGroups(tWafRuleGroupDto)
	rules := wafService.getRules(tRuleGrops)

	return waf.SlaveWafDataDto{
		ServerIpUrl: config.G_AppConfig.ServerIpUrl,
		Waf:    wafDto,
		Routes: routes,
		Certs:  certs,
		Rules:rules,
	}


}

func (wafService *WafService) getRules(grops []*waf.WafRuleGroupDto) []*waf.WafRuleDataDto {
	if grops == nil || len(grops) < 1{
		return nil
	}
	var (
		wafRuleDao wafDao.WafRuleDao
		rules []*waf.WafRuleDataDto
		ruleData *waf.WafRuleDataDto
	)



	tWafRuleDto := waf.WafRuleDto{
		GroupId: grops[0].GroupId,
	}
	wafRules ,_ := wafRuleDao.GetWafRules(tWafRuleDto)

	if wafRules == nil || len(wafRules) < 1{
		return nil
	}

	for _,rule:= range wafRules{
		ruleData = &waf.WafRuleDataDto{
		}
		objectConvert.Struct2Struct(rule,ruleData)

		err := wafService.getRuleObject(ruleData)

		if err != nil{
			continue
		}

		rules = append(rules,ruleData)
	}

	return rules
}

func (wafService *WafService) getRuleObject(data *waf.WafRuleDataDto) error {
	var wafIpBlackWhiteDao wafDao.WafIpBlackWhiteDao
	if data.ObjType == waf.Waf_obj_type_ip{
		tWafIp := waf.WafIpBlackWhiteDto{
			Id:data.ObjId,
		}

		wafIps ,_:= wafIpBlackWhiteDao.GetWafIpBlackWhites(tWafIp)

		if wafIps == nil || len(wafIps) < 1{
			return errors.New("未包含ip")
		}
		data.Ip = wafIps[0]
	}

	return nil

}
