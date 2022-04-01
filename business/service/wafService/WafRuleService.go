package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafRuleService struct {
	wafDao             wafDao.WafRuleDao
	wafHostnameCertDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafRuleService) GetWafRuleAll(WafRuleDto waf.WafRuleDto) ([]*waf.WafRuleDto, error) {
	var (
		err          error
		WafRuleDtos []*waf.WafRuleDto
	)

	WafRuleDtos, err = wafService.wafDao.GetWafRules(WafRuleDto)
	if err != nil {
		return nil, err
	}

	return WafRuleDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafRuleService) GetWafRules(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafRuleDto{}
		wafDtos []*waf.WafRuleDto
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

	total, err = wafService.wafDao.GetWafRuleCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafRules(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafRuleService) SaveWafRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.RuleId = seq.Generator()
	//WafRuleDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafRule(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafRuleService) UpdateWafRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafRule(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafRuleService) DeleteWafRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafRule(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
