package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafHostsService struct {
	wafDao wafDao.WafHostsDao
}

// get db link
// all db by this user
func (wafService *WafHostsService) GetWafHostsAll(WafHostsDto waf.WafHostsDto) ([]*waf.WafHostsDto, error) {
	var (
		err        error
		WafHostsDtos []*waf.WafHostsDto
	)

	WafHostsDtos, err = wafService.wafDao.GetWafHostss(WafHostsDto)
	if err != nil {
		return nil, err
	}

	return WafHostsDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafHostsService) GetWafHostss(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafHostsDto{}
		wafDtos []*waf.WafHostsDto
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

	total, err = wafService.wafDao.GetWafHostsCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafHostss(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafHostsService) SaveWafHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostsDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.WafHostId = seq.Generator()
	//WafHostsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafHosts(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafHostsService) UpdateWafHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostsDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafHosts(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafHostsService) DeleteWafHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostsDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafHosts(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
