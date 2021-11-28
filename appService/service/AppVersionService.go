package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appVersion"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppVersionService struct {
	appVersionDao dao.AppVersionDao
}

/**
查询 系统信息
*/
func (appVersionService *AppVersionService) GetAppVersionAll(appVersionDto appVersion.AppVersionDto) ([]*appVersion.AppVersionDto, error) {
	var (
		err            error
		appVersionDtos []*appVersion.AppVersionDto
	)

	appVersionDtos, err = appVersionService.appVersionDao.GetAppVersions(appVersionDto)
	if err != nil {
		return nil, err
	}

	return appVersionDtos, nil

}

/**
查询 系统信息
*/
func (appVersionService *AppVersionService) GetAppVersions(ctx iris.Context) result.ResultDto {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appVersionDto  = appVersion.AppVersionDto{}
		appVersionDtos []*appVersion.AppVersionDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionDto.Row = row * page

	appVersionDto.Page = (page - 1) * row

	appVersionDto.AvId = ctx.URLParam("avId")
	appVersionDto.Name = ctx.URLParam("name")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionDto.TenantId = user.TenantId

	total, err = appVersionService.appVersionDao.GetAppVersionCount(appVersionDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionDtos, err = appVersionService.appVersionDao.GetAppVersions(appVersionDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionDtos, total, row)

}

/**
保存 系统信息
*/
func (appVersionService *AppVersionService) SaveAppVersions(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appVersionDto appVersion.AppVersionDto
	)

	if err = ctx.ReadJSON(&appVersionDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionDto.TenantId = user.TenantId
	appVersionDto.AvId = seq.Generator()

	err = appVersionService.appVersionDao.SaveAppVersion(appVersionDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionDto)

}

/**
修改 系统信息
*/
func (appVersionService *AppVersionService) UpdateAppVersions(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appVersionDto appVersion.AppVersionDto
	)

	if err = ctx.ReadJSON(&appVersionDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionService.appVersionDao.UpdateAppVersion(appVersionDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionDto)

}

/**
删除 系统信息
*/
func (appVersionService *AppVersionService) DeleteAppVersions(ctx iris.Context) result.ResultDto {
	var (
		err           error
		appVersionDto appVersion.AppVersionDto
	)

	if err = ctx.ReadJSON(&appVersionDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionService.appVersionDao.DeleteAppVersion(appVersionDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionDto)

}
