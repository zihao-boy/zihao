package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appVar"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppVarService struct {
	appVarDao dao.AppVarDao
}

/**
查询 系统信息
*/
func (appVarService *AppVarService) GetAppVarAll(appVarDto appVar.AppVarDto) ([]*appVar.AppVarDto, error) {
	var (
		err        error
		appVarDtos []*appVar.AppVarDto
	)

	appVarDtos, err = appVarService.appVarDao.GetAppVars(appVarDto)
	if err != nil {
		return nil, err
	}

	return appVarDtos, nil

}

/**
查询 系统信息
*/
func (appVarService *AppVarService) GetAppVars(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		appVarDto  = appVar.AppVarDto{}
		appVarDtos []*appVar.AppVarDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVarDto.Row = row * page

	appVarDto.Page = (page - 1) * row

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVarDto.TenantId = user.TenantId

	total, err = appVarService.appVarDao.GetAppVarCount(appVarDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVarDtos, err = appVarService.appVarDao.GetAppVars(appVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarDtos, total, row)

}

/**
保存 系统信息
*/
func (appVarService *AppVarService) SaveAppVars(ctx iris.Context) result.ResultDto {
	var (
		err       error
		appVarDto appVar.AppVarDto
	)

	if err = ctx.ReadJSON(&appVarDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVarDto.TenantId = user.TenantId
	appVarDto.AvId = seq.Generator()

	err = appVarService.appVarDao.SaveAppVar(appVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarDto)

}

/**
修改 系统信息
*/
func (appVarService *AppVarService) UpdateAppVars(ctx iris.Context) result.ResultDto {
	var (
		err       error
		appVarDto appVar.AppVarDto
	)

	if err = ctx.ReadJSON(&appVarDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVarService.appVarDao.UpdateAppVar(appVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarDto)

}

/**
删除 系统信息
*/
func (appVarService *AppVarService) DeleteAppVars(ctx iris.Context) result.ResultDto {
	var (
		err       error
		appVarDto appVar.AppVarDto
	)

	if err = ctx.ReadJSON(&appVarDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVarService.appVarDao.DeleteAppVar(appVarDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarDto)

}
