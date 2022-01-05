package service

import (
	"github.com/zihao-boy/zihao/common/containerScheduling"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	assetsDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppServiceService struct {
	appServiceDao dao.AppServiceDao
	hostDao assetsDao.HostDao
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
	appServiceDto.AsId = ctx.URLParam("asId")
	appServiceDto.AsGroupId = ctx.URLParam("asGroupId")
	appServiceDto.AsType = ctx.URLParam("asType")
	appServiceDto.AsName = ctx.URLParam("asName")

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
			appServiceService.appServiceDao.SaveAppServicePort(*appServicePort)
		}
	}

	if len(appServiceDto.AppServiceHosts) > 0{
		for _,appServiceHost := range appServiceDto.AppServiceHosts{
			appServiceHost.HostsId = seq.Generator()
			appServiceHost.TenantId = user.TenantId
			appServiceHost.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceHosts(*appServiceHost)
		}
	}

	if len(appServiceDto.AppServiceDirs) > 0{
		for _,appServiceDir := range appServiceDto.AppServiceDirs{
			appServiceDir.DirId = seq.Generator()
			appServiceDir.TenantId = user.TenantId
			appServiceDir.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceDir(*appServiceDir)
		}
	}

	if len(appServiceDto.AppServiceVars) > 0{
		for _,appServiceVar := range appServiceDto.AppServiceVars{
			appServiceVar.AvId = seq.Generator()
			appServiceVar.TenantId = user.TenantId
			appServiceVar.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceVar(*appServiceVar)
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

func (appServiceService *AppServiceService) CopyAppServices(ctx iris.Context) result.ResultDto {

	var (
		err           error
		appServiceDto appService.AppServiceDto
		newAppServiceDto *appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	oldAppServiceDto := appService.AppServiceDto{
		AsId: appServiceDto.AsId,
	}

	appServiceDtos, err := appServiceService.appServiceDao.GetAppServices(oldAppServiceDto)

	if len(appServiceDtos) < 1{
		return result.Error("应用不存在")
	}

	newAppServiceDto = appServiceDtos[0]

	newAppServiceDto.AsId = seq.Generator()
	newAppServiceDto.AsName = appServiceDto.AsName
	newAppServiceDto.AsDesc = appServiceDto.AsDesc
	newAppServiceDto.ImagesId = appServiceDto.ImagesId
	newAppServiceDto.State = appService.STATE_STOP

	err = appServiceService.appServiceDao.SaveAppService(*newAppServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	portDto := appService.AppServicePortDto{
		AsId: appServiceDto.AsId,
	}
	portDtos,_ := appServiceService.appServiceDao.GetAppServicePort(portDto)

	if len(portDtos) > 0{
		for _,appServicePort := range portDtos{
			appServicePort.PortId = seq.Generator()
			appServicePort.TenantId = newAppServiceDto.TenantId
			appServicePort.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServicePort(*appServicePort)
		}
	}

	hostsDto := appService.AppServiceHostsDto{
		AsId: appServiceDto.AsId,
	}
	hostsDtos,_ := appServiceService.appServiceDao.GetAppServiceHosts(hostsDto)

	if len(hostsDtos) > 0{
		for _,appServiceHost := range hostsDtos{
			appServiceHost.HostsId = seq.Generator()
			appServiceHost.TenantId = newAppServiceDto.TenantId
			appServiceHost.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceHosts(*appServiceHost)
		}
	}
	dirDto := appService.AppServiceDirDto{
		AsId: appServiceDto.AsId,
	}
	dirDtos,_ := appServiceService.appServiceDao.GetAppServiceDir(dirDto)

	if len(dirDtos) > 0{
		for _,appServiceDir := range dirDtos{
			appServiceDir.DirId = seq.Generator()
			appServiceDir.TenantId = newAppServiceDto.TenantId
			appServiceDir.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceDir(*appServiceDir)
		}
	}

	varDto := appService.AppServiceVarDto{
		AsId: appServiceDto.AsId,
	}
	varDtos,_ := appServiceService.appServiceDao.GetAppServiceVars(varDto)
	if len(varDtos) > 0{
		for _,appServiceVar := range varDtos{
			appServiceVar.AvId = seq.Generator()
			appServiceVar.TenantId = newAppServiceDto.TenantId
			appServiceVar.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceVar(*appServiceVar)
		}
	}

	return result.SuccessData(newAppServiceDto)
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
	appServiceVarDto.AsId = ctx.URLParam("asId")
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
	appServiceHostsDto.AsId = ctx.URLParam("asId")
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
	appServiceDirDto.AsId = ctx.URLParam("asId")
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
	appServicePortDto.AsId = ctx.URLParam("asId")

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
	appServiceContainerDto.AsId = ctx.URLParam("asId")
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

//应用启用方法

func (appServiceService *AppServiceService) StartAppService(ctx iris.Context) interface{} {
	var (
		err           error
		appServiceDto appService.AppServiceDto
		hosts []*host.HostDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	//appServiceDto.State = appService.STATE_STOP //避免 状态不一致导致的坑
	appServiceDtos,_ := appServiceService.appServiceDao.GetAppServices(appServiceDto)

	if len(appServiceDtos) <1{
		return result.Error("应用不存在")
	}

	//修改应用为启动中
	//appServiceDto.State = appService.STATE_DOING
	//appServiceService.appServiceDao.UpdateAppService(appServiceDto)

	tmpAppServiceDto := appServiceDtos[0]

	if tmpAppServiceDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST{
		hostDto :=host.HostDto{
			HostId: tmpAppServiceDto.AsDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	}else{
		hostDto :=host.HostDto{
			GroupId: tmpAppServiceDto.AsDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	}

	param,err:=containerScheduling.ContainerScheduling(hosts,tmpAppServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return param
}

func (appServiceService *AppServiceService) StopAppService(ctx iris.Context) interface{} {
	var (
		appServiceDto appService.AppServiceDto
	)

	if err := ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	//appServiceDto.State = appService.STATE_ONLINE //避免 状态不一致导致的坑
	appServiceDtos,_ := appServiceService.appServiceDao.GetAppServices(appServiceDto)

	if len(appServiceDtos) <1{
		return result.Error("应用不存在")
	}

	//修改应用为启动中
	//appServiceDto.State = appService.STATE_DOING
	//appServiceService.appServiceDao.UpdateAppService(appServiceDto)

	tmpAppServiceDto := appServiceDtos[0]

	var param, err = containerScheduling.StopContainer(tmpAppServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return param
}

// get faster deploy app service log
func (appServiceService *AppServiceService) GetFasterDeploys(ctx iris.Context) result.ResultDto {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		fasterDeployDto  = appService.FasterDeployDto{}
		fasterDeployDtos []*appService.FasterDeployDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	fasterDeployDto.Row = row * page

	fasterDeployDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	fasterDeployDto.TenantId = user.TenantId
	fasterDeployDto.AsGroupId = ctx.URLParam("asGroupId")

	total, err = appServiceService.appServiceDao.GetFasterDeployCount(fasterDeployDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	fasterDeployDtos, err = appServiceService.appServiceDao.GetFasterDeploys(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(fasterDeployDtos, total, row)
}

func (appServiceService *AppServiceService) SaveFasterDeploys(ctx iris.Context) result.ResultDto {

	var (
		err           error
		fasterDeployDto appService.FasterDeployDto
	)

	if err = ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	fasterDeployDto.TenantId = user.TenantId

	err = appServiceService.appServiceDao.SaveFasterDeploy(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(fasterDeployDto)
}

func (appServiceService *AppServiceService) UpdateFasterDeploys(ctx iris.Context) result.ResultDto {
	var (
		err           error
		fasterDeployDto appService.FasterDeployDto
	)

	if err = ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.UpdateFasterDeploy(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(fasterDeployDto)

}

func (appServiceService *AppServiceService) DeleteFasterDeploys(ctx iris.Context) result.ResultDto {
	var (
		err           error
		fasterDeployDto appService.FasterDeployDto
	)

	if err = ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appServiceService.appServiceDao.DeleteFasterDeploy(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(fasterDeployDto)

}


