package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafRouteService struct {
	wafDao wafDao.WafRouteDao
}

// get db link
// all db by this user
func (wafService *WafRouteService) GetWafRouteAll(WafRouteDto waf.WafRouteDto) ([]*waf.WafRouteDto, error) {
	var (
		err        error
		WafRouteDtos []*waf.WafRouteDto
	)

	WafRouteDtos, err = wafService.wafDao.GetWafRoutes(WafRouteDto)
	if err != nil {
		return nil, err
	}

	return WafRouteDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafRouteService) GetWafRoutes(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafRouteDto{}
		wafDtos []*waf.WafRouteDto
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

	total, err = wafService.wafDao.GetWafRouteCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafRoutes(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafRouteService) SaveWafRoutes(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafRouteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.RouteId = seq.Generator()
	//WafRouteDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafRoute(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafRouteService) UpdateWafRoutes(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafRouteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafRoute(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafRouteService) DeleteWafRoutes(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafRouteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafRoute(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
