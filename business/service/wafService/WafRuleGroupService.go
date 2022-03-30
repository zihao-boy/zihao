package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafRuleGroupService struct {
	wafDao             wafDao.WafRuleGroupDao
	wafHostnameCertDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafRuleGroupService) GetWafRuleGroupAll(WafRuleGroupDto waf.WafRuleGroupDto) ([]*waf.WafRuleGroupDto, error) {
	var (
		err          error
		WafRuleGroupDtos []*waf.WafRuleGroupDto
	)

	WafRuleGroupDtos, err = wafService.wafDao.GetWafRuleGroups(WafRuleGroupDto)
	if err != nil {
		return nil, err
	}

	return WafRuleGroupDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafRuleGroupService) GetWafRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafRuleGroupDto{}
		wafDtos []*waf.WafRuleGroupDto
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

	total, err = wafService.wafDao.GetWafRuleGroupCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafRuleGroups(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafRuleGroupService) SaveWafRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleGroupDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.GroupId = seq.Generator()
	//WafRuleGroupDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafRuleGroup(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafRuleGroupService) UpdateWafRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleGroupDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafRuleGroup(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafRuleGroupService) DeleteWafRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafRuleGroupDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafRuleGroup(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
