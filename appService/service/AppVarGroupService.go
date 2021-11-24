package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appVarGroup"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppVarGroupService struct {
	appVarGroupDao dao.AppVarGroupDao
}

/**
查询 系统信息
*/
func (appVarGroupService *AppVarGroupService) GetAppVarGroupAll(appVarGroupDto appVarGroup.AppVarGroupDto) ([]*appVarGroup.AppVarGroupDto, error) {
	var (
		err             error
		appVarGroupDtos []*appVarGroup.AppVarGroupDto
	)

	appVarGroupDtos, err = appVarGroupService.appVarGroupDao.GetAppVarGroups(appVarGroupDto)
	if err != nil {
		return nil, err
	}

	return appVarGroupDtos, nil

}

/**
查询 系统信息
*/
func (appVarGroupService *AppVarGroupService) GetAppVarGroups(ctx iris.Context) result.ResultDto {
	var (
		err             error
		page            int64
		row             int64
		total           int64
		appVarGroupDto  = appVarGroup.AppVarGroupDto{}
		appVarGroupDtos []*appVarGroup.AppVarGroupDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVarGroupDto.Row = row * page

	appVarGroupDto.Page = (page - 1) * row

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVarGroupDto.TenantId = user.TenantId

	total, err = appVarGroupService.appVarGroupDao.GetAppVarGroupCount(appVarGroupDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVarGroupDtos, err = appVarGroupService.appVarGroupDao.GetAppVarGroups(appVarGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarGroupDtos, total, row)

}

/**
保存 系统信息
*/
func (appVarGroupService *AppVarGroupService) SaveAppVarGroups(ctx iris.Context) result.ResultDto {
	var (
		err            error
		appVarGroupDto appVarGroup.AppVarGroupDto
	)

	if err = ctx.ReadJSON(&appVarGroupDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVarGroupDto.TenantId = user.TenantId
	appVarGroupDto.AvgType = appVarGroup.Avg_Type_public
	appVarGroupDto.AvgId = seq.Generator()

	err = appVarGroupService.appVarGroupDao.SaveAppVarGroup(appVarGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarGroupDto)

}

/**
修改 系统信息
*/
func (appVarGroupService *AppVarGroupService) UpdateAppVarGroups(ctx iris.Context) result.ResultDto {
	var (
		err            error
		appVarGroupDto appVarGroup.AppVarGroupDto
	)

	if err = ctx.ReadJSON(&appVarGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVarGroupService.appVarGroupDao.UpdateAppVarGroup(appVarGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarGroupDto)

}

/**
删除 系统信息
*/
func (appVarGroupService *AppVarGroupService) DeleteAppVarGroups(ctx iris.Context) result.ResultDto {
	var (
		err            error
		appVarGroupDto appVarGroup.AppVarGroupDto
	)

	if err = ctx.ReadJSON(&appVarGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVarGroupService.appVarGroupDao.DeleteAppVarGroup(appVarGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVarGroupDto)

}
