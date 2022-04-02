package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafAreaService struct {
	wafDao             wafDao.WafAreaDao
	wafRuleDao wafDao.WafRuleDao
	wafHostnameCertDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafAreaService) GetWafAreaAll(WafAreaDto waf.WafAreaDto) ([]*waf.WafAreaDto, error) {
	var (
		err          error
		WafAreaDtos []*waf.WafAreaDto
	)

	WafAreaDtos, err = wafService.wafDao.GetWafAreas(WafAreaDto)
	if err != nil {
		return nil, err
	}

	return WafAreaDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafAreaService) GetWafAreas(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafAreaDto{}
		wafDtos []*waf.WafAreaDto
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

	total, err = wafService.wafDao.GetWafAreaCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafAreas(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafAreaService) SaveWafAreas(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAreaDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.Id = seq.Generator()
	//WafAreaDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafArea(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	wafRuleDto := waf.WafRuleDto{
		RuleId:seq.Generator(),
		GroupId:wafDto.GroupId,
		RuleName:wafDto.AreaName,
		Scope:wafDto.Scope,
		ObjId:wafDto.Id,
		ObjType:waf.Waf_obj_type_area,
		Seq:wafDto.Seq,
		State:wafDto.State,
	}
	err = wafService.wafRuleDao.SaveWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafAreaService) UpdateWafAreas(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAreaDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafArea(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:wafDto.Id,
		ObjType:waf.Waf_obj_type_area,
	}
	qWafRuleDtos , _ := wafService.wafRuleDao.GetWafRules(qWafRuleDto)

	if qWafRuleDtos == nil ||  len(qWafRuleDtos) <1{
		return result.Success()
	}

	wafRuleDto := waf.WafRuleDto{
		RuleId:qWafRuleDtos[0].RuleId,
		GroupId:wafDto.GroupId,
		RuleName:wafDto.AreaName,
		Scope:wafDto.Scope,
		Seq:wafDto.Seq,
		State:wafDto.State,
	}
	err = wafService.wafRuleDao.UpdateWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafAreaService) DeleteWafAreas(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAreaDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafArea(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:wafDto.Id,
		ObjType:waf.Waf_obj_type_area,
	}
	qWafRuleDtos , _ := wafService.wafRuleDao.GetWafRules(qWafRuleDto)

	if qWafRuleDtos == nil ||  len(qWafRuleDtos) <1{
		return result.Success()
	}

	wafRuleDto := waf.WafRuleDto{
		RuleId:qWafRuleDtos[0].RuleId,
	}

	err = wafService.wafRuleDao.DeleteWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
