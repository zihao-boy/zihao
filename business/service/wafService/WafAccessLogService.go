package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafAccessLogService struct {
	wafDao wafDao.WafAccessLogDao
}

// get db link
// all db by this user
func (wafService *WafAccessLogService) GetWafAccessLogAll(WafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogDto, error) {
	var (
		err        error
		WafAccessLogDtos []*waf.WafAccessLogDto
	)

	WafAccessLogDtos, err = wafService.wafDao.GetWafAccessLogs(WafAccessLogDto)
	if err != nil {
		return nil, err
	}

	return WafAccessLogDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafAccessLogService) GetWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafAccessLogDto{}
		wafDtos []*waf.WafAccessLogDto
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

	total, err = wafService.wafDao.GetWafAccessLogCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafAccessLogs(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafAccessLogService) SaveWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.RequestId = seq.Generator()
	//WafAccessLogDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafAccessLog(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafAccessLogService) UpdateWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafAccessLog(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafAccessLogService) DeleteWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafAccessLog(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
