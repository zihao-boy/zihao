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

	if len(appServiceDto.AppServicePorts) > 0{
		for _,appServicePort := range appServiceDto.AppServicePorts{
			appServicePort.PortId = seq.Generator()
			appServicePort.TenantId = user.TenantId
			appServicePort.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServicePort(appServicePort)
		}
	}

	if len(appServiceDto.AppServiceHosts) > 0{
		for _,appServiceHost := range appServiceDto.AppServiceHosts{
			appServiceHost.HostsId = seq.Generator()
			appServiceHost.TenantId = user.TenantId
			appServiceHost.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceHosts(appServiceHost)
		}
	}

	if len(appServiceDto.AppServiceDirs) > 0{
		for _,appServiceDir := range appServiceDto.AppServiceDirs{
			appServiceDir.DirId = seq.Generator()
			appServiceDir.TenantId = user.TenantId
			appServiceDir.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceDir(appServiceDir)
		}
	}

	if len(appServiceDto.AppServiceVars) > 0{
		for _,appServiceVar := range appServiceDto.AppServiceVars{
			appServiceVar.AvId = seq.Generator()
			appServiceVar.TenantId = user.TenantId
			appServiceVar.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceVar(appServiceVar)
		}
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

func (appServiceService *AppServiceService) GetAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServiceHostsDto  = appService.AppServiceHostsDto{}
		appServiceHostsDtos []*appService.AppServiceHostsDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServiceHostsDto.Row = row * page

	appServiceHostsDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceHostsDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServiceHostsCount(appServiceHostsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServiceHostsDtos, err = appServiceService.appServiceDao.GetAppServiceHosts(appServiceHostsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceHostsDtos, total, row)
}

func (appServiceService *AppServiceService) SaveAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceHostsDto appService.AppServiceHostsDto
	)

	if err = ctx.ReadJSON(&appServiceHostsDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceHostsDto.TenantId = user.TenantId
	appServiceHostsDto.HostsId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppServiceHosts(appServiceHostsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceHostsDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceHostsDto appService.AppServiceHostsDto
	)

	if err = ctx.ReadJSON(&appServiceHostsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppServiceHost(appServiceHostsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceHostsDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceHostsDto appService.AppServiceHostsDto
	)

	if err = ctx.ReadJSON(&appServiceHostsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppServiceHosts(appServiceHostsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceHostsDto)
}

func (appServiceService *AppServiceService) GetAppServiceDir(ctx iris.Context) interface{} {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServiceDirDto  = appService.AppServiceDirDto{}
		appServiceDirDtos []*appService.AppServiceDirDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServiceDirDto.Row = row * page

	appServiceDirDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDirDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServiceDirCount(appServiceDirDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServiceDirDtos, err = appServiceService.appServiceDao.GetAppServiceDir(appServiceDirDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDirDtos, total, row)
}

func (appServiceService *AppServiceService) SaveAppServiceDir(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceDirDto appService.AppServiceDirDto
	)

	if err = ctx.ReadJSON(&appServiceDirDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDirDto.TenantId = user.TenantId
	appServiceDirDto.DirId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppServiceDir(appServiceDirDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDirDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceDir(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceDirDto appService.AppServiceDirDto
	)

	if err = ctx.ReadJSON(&appServiceDirDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppServiceDir(appServiceDirDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDirDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceDir(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceDirDto appService.AppServiceDirDto
	)

	if err = ctx.ReadJSON(&appServiceDirDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppServiceDir(appServiceDirDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceDirDto)
}

func (appServiceService *AppServiceService) GetAppServicePort(ctx iris.Context) interface{} {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServicePortDto  = appService.AppServicePortDto{}
		appServicePortDtos []*appService.AppServicePortDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServicePortDto.Row = row * page

	appServicePortDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServicePortDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServicePortCount(appServicePortDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServicePortDtos, err = appServiceService.appServiceDao.GetAppServicePort(appServicePortDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServicePortDtos, total, row)
}

func (appServiceService *AppServiceService) SaveAppServicePort(ctx iris.Context) interface{} {
	var (
		err           error
		appServicePortDto appService.AppServicePortDto
	)

	if err = ctx.ReadJSON(&appServicePortDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServicePortDto.TenantId = user.TenantId
	appServicePortDto.PortId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppServicePort(appServicePortDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServicePortDto)
}

func (appServiceService *AppServiceService) UpdateAppServicePort(ctx iris.Context) interface{} {
	var (
		err           error
		appServicePortDto appService.AppServicePortDto
	)

	if err = ctx.ReadJSON(&appServicePortDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppServicePort(appServicePortDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServicePortDto)
}

func (appServiceService *AppServiceService) DeleteAppServicePort(ctx iris.Context) interface{} {
	var (
		err           error
		appServicePortDto appService.AppServicePortDto
	)

	if err = ctx.ReadJSON(&appServicePortDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppServicePort(appServicePortDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServicePortDto)
}

func (appServiceService *AppServiceService) GetAppServiceContainer(ctx iris.Context) interface{} {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		appServiceContainerDto  = appService.AppServiceContainerDto{}
		appServiceContainerDtos []*appService.AppServiceContainerDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appServiceContainerDto.Row = row * page

	appServiceContainerDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceContainerDto.TenantId = user.TenantId

	total, err = appServiceService.appServiceDao.GetAppServiceContainerCount(appServiceContainerDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appServiceContainerDtos, err = appServiceService.appServiceDao.GetAppServiceContainer(appServiceContainerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceContainerDtos, total, row)
}

func (appServiceService *AppServiceService) SaveAppServiceContainer(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceContainerDto appService.AppServiceContainerDto
	)

	if err = ctx.ReadJSON(&appServiceContainerDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceContainerDto.TenantId = user.TenantId
	appServiceContainerDto.ContainerId = seq.Generator()

	err = appServiceService.appServiceDao.SaveAppServiceContainer(appServiceContainerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceContainerDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceContainer(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceContainerDto appService.AppServiceContainerDto
	)

	if err = ctx.ReadJSON(&appServiceContainerDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateAppServiceContainer(appServiceContainerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceContainerDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceContainer(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceContainerDto appService.AppServiceContainerDto
	)

	if err = ctx.ReadJSON(&appServiceContainerDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteAppServiceContainer(appServiceContainerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appServiceContainerDto)
}
