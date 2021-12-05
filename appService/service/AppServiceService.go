package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppServiceService struct {
	appServiceDao dao.AppServiceDao
}

/**
查询 系统信息
*/
func (appServiceService *AppServiceService) GetAppServiceAll(appServiceDto appService.AppServiceDto) ([]*appService.AppServiceDto, error) {
	var (
		err            error
		appServiceDtos []*appService.AppServiceDto
	)

	appServiceDtos, err = appServiceService.appServiceDao.GetAppServices(appServiceDto)
	if err != nil {
		return nil, err
	}

	return appServiceDtos, nil

}

/**
查询 系统信息
*/
func (appServiceService *AppServiceService) GetAppServices(ctx iris.Context) result.ResultDto {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServiceDto  = appService.AppServiceDto{}
		appServiceDtos []*appService.AppServiceDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServiceDto.Row = row * page

	appServiceDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServiceCount(appServiceDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServiceDtos, err = appServiceService.appServiceDao.GetAppServices(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDtos, total, row)

}

/**
保存 系统信息
*/
func (appServiceService *AppServiceService) SaveAppServices(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceDto appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDto.TenantId = user.TenantId
	appServiceDto.State = appService.STATE_STOP
	appServiceDto.AsId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDto)

}

/**
修改 系统信息
*/
func (appServiceService *AppServiceService) UpdateAppServices(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceDto appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDto)

}

/**
删除 系统信息
*/
func (appServiceService *AppServiceService) DeleteAppServices(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceDto appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDto)

}

func (appServiceService *AppServiceService) GetAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServiceVarDto  = appService.AppServiceVarDto{}
		appServiceVarDtos []*appService.AppServiceVarDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServiceVarDto.Row = row * page

	appServiceVarDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceVarDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServiceVarCount(appServiceVarDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServiceVarDtos, err = appServiceService.appServiceDao.GetAppServiceVars(appServiceVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceVarDtos, total, row)
}

func (appServiceService *AppServiceService) SaveAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceVarDto appService.AppServiceVarDto
	)

	if err = ctx.ReadJSON(&appServiceVarDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceVarDto.TenantId = user.TenantId
	appServiceVarDto.AvId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppServiceVar(appServiceVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceVarDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceVarDto appService.AppServiceVarDto
	)

	if err = ctx.ReadJSON(&appServiceVarDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppServiceVar(appServiceVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceVarDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appServiceVarDto appService.AppServiceVarDto
	)

	if err = ctx.ReadJSON(&appServiceVarDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppServiceVar(appServiceVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceVarDto)
}
