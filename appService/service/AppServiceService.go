package service

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	assetsDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/containerScheduling"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/queue/dockerfileQueue"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"github.com/zihao-boy/zihao/entity/dto/composeYaml"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	dao2 "github.com/zihao-boy/zihao/softService/dao"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const maxSize = 1000 * iris.MB // 第二种方法
type AppServiceService struct {
	appServiceDao         dao.AppServiceDao
	hostDao               assetsDao.HostDao
	businessPackageDao    dao2.BusinessPackageDao
	businessDockerfileDao dao2.BusinessDockerfileDao
	businessImagesDao     dao2.BusinessImagesDao
	businessImagesVerDao  dao2.BusinessImagesVerDao
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

	if len(appServiceDto.AppServicePorts) > 0 {
		for _, appServicePort := range appServiceDto.AppServicePorts {
			appServicePort.PortId = seq.Generator()
			appServicePort.TenantId = user.TenantId
			appServicePort.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServicePort(*appServicePort)
		}
	}

	if len(appServiceDto.AppServiceHosts) > 0 {
		for _, appServiceHost := range appServiceDto.AppServiceHosts {
			appServiceHost.HostsId = seq.Generator()
			appServiceHost.TenantId = user.TenantId
			appServiceHost.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceHosts(*appServiceHost)
		}
	}

	if len(appServiceDto.AppServiceDirs) > 0 {
		for _, appServiceDir := range appServiceDto.AppServiceDirs {
			appServiceDir.DirId = seq.Generator()
			appServiceDir.TenantId = user.TenantId
			appServiceDir.AsId = appServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceDir(*appServiceDir)
		}
	}

	if len(appServiceDto.AppServiceVars) > 0 {
		for _, appServiceVar := range appServiceDto.AppServiceVars {
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
		err              error
		appServiceDto    appService.AppServiceDto
		newAppServiceDto *appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	oldAppServiceDto := appService.AppServiceDto{
		AsId: appServiceDto.AsId,
	}

	appServiceDtos, err := appServiceService.appServiceDao.GetAppServices(oldAppServiceDto)

	if len(appServiceDtos) < 1 {
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
	portDtos, _ := appServiceService.appServiceDao.GetAppServicePort(portDto)

	if len(portDtos) > 0 {
		for _, appServicePort := range portDtos {
			appServicePort.PortId = seq.Generator()
			appServicePort.TenantId = newAppServiceDto.TenantId
			appServicePort.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServicePort(*appServicePort)
		}
	}

	hostsDto := appService.AppServiceHostsDto{
		AsId: appServiceDto.AsId,
	}
	hostsDtos, _ := appServiceService.appServiceDao.GetAppServiceHosts(hostsDto)

	if len(hostsDtos) > 0 {
		for _, appServiceHost := range hostsDtos {
			appServiceHost.HostsId = seq.Generator()
			appServiceHost.TenantId = newAppServiceDto.TenantId
			appServiceHost.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceHosts(*appServiceHost)
		}
	}
	dirDto := appService.AppServiceDirDto{
		AsId: appServiceDto.AsId,
	}
	dirDtos, _ := appServiceService.appServiceDao.GetAppServiceDir(dirDto)

	if len(dirDtos) > 0 {
		for _, appServiceDir := range dirDtos {
			appServiceDir.DirId = seq.Generator()
			appServiceDir.TenantId = newAppServiceDto.TenantId
			appServiceDir.AsId = newAppServiceDto.AsId
			appServiceService.appServiceDao.SaveAppServiceDir(*appServiceDir)
		}
	}

	varDto := appService.AppServiceVarDto{
		AsId: appServiceDto.AsId,
	}
	varDtos, _ := appServiceService.appServiceDao.GetAppServiceVars(varDto)
	if len(varDtos) > 0 {
		for _, appServiceVar := range varDtos {
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
		err               error
		page              int64
		row               int64
		total             int64
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
		err              error
		appServiceVarDto appService.AppServiceVarDto
		asIds            []string
	)

	if err = ctx.ReadJSON(&appServiceVarDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	asIds = strings.Split(appServiceVarDto.AsId, ",")
	for _, asId := range asIds {
		appServiceVarDto.AsId = asId
		appServiceVarDto.TenantId = user.TenantId
		appServiceVarDto.AvId = seq.Generator()
		err = appServiceService.appServiceDao.SaveAppServiceVar(appServiceVarDto)
		if err != nil {
			return result.Error(err.Error())
		}
	}

	return result.SuccessData(appServiceVarDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		appServiceVarDto      appService.AppServiceVarDto
		queryAppServiceVarDto appService.AppServiceVarDto

		asIds []string
	)

	if err = ctx.ReadJSON(&appServiceVarDto); err != nil {
		return result.Error("解析入参失败")
	}

	asIds = strings.Split(appServiceVarDto.AsId, ",")
	for _, asId := range asIds {
		queryAppServiceVarDto.AsId = asId
		queryAppServiceVarDto.VarSpec = appServiceVarDto.VarSpec
		appServiceVarDtos, _ := appServiceService.appServiceDao.GetAppServiceVars(queryAppServiceVarDto)
		if appServiceVarDtos == nil || len(appServiceVarDtos) < 1 {
			continue
		}
		for _, appVarDto := range appServiceVarDtos {
			appServiceVarDto.AvId = appVarDto.AvId
			err = appServiceService.appServiceDao.UpdateAppServiceVar(appServiceVarDto)
			if err != nil {
				return result.Error(err.Error())
			}
		}
	}
	return result.SuccessData(appServiceVarDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceVars(ctx iris.Context) result.ResultDto {
	var (
		err              error
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
		err                 error
		page                int64
		row                 int64
		total               int64
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
		err                error
		appServiceHostsDto appService.AppServiceHostsDto

		asIds []string
	)

	if err = ctx.ReadJSON(&appServiceHostsDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	asIds = strings.Split(appServiceHostsDto.AsId, ",")
	for _, asId := range asIds {
		appServiceHostsDto.AsId = asId
		appServiceHostsDto.TenantId = user.TenantId
		appServiceHostsDto.HostsId = seq.Generator()
		err = appServiceService.appServiceDao.SaveAppServiceHosts(appServiceHostsDto)
		if err != nil {
			return result.Error(err.Error())
		}

	}
	return result.SuccessData(appServiceHostsDto)
}

func (appServiceService *AppServiceService) UpdateAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err                    error
		appServiceHostsDto     appService.AppServiceHostsDto
		querAppServiceHostsDto appService.AppServiceHostsDto
		asIds                  []string
	)

	if err = ctx.ReadJSON(&appServiceHostsDto); err != nil {
		return result.Error("解析入参失败")
	}
	asIds = strings.Split(appServiceHostsDto.AsId, ",")
	for _, asId := range asIds {
		querAppServiceHostsDto.AsId = asId
		querAppServiceHostsDto.Hostname = appServiceHostsDto.Hostname
		appServiceHostsDtos, _ := appServiceService.appServiceDao.GetAppServiceHosts(querAppServiceHostsDto)
		if appServiceHostsDtos == nil || len(appServiceHostsDtos) < 1 {
			continue
		}
		for _, appHostDto := range appServiceHostsDtos {
			appServiceHostsDto.HostsId = appHostDto.HostsId
			err = appServiceService.appServiceDao.UpdateAppServiceHost(appServiceHostsDto)
			if err != nil {
				return result.Error(err.Error())
			}
		}

	}
	return result.SuccessData(appServiceHostsDto)
}

func (appServiceService *AppServiceService) DeleteAppServiceHosts(ctx iris.Context) interface{} {
	var (
		err                error
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
		err               error
		page              int64
		row               int64
		total             int64
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
		err              error
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
		err              error
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
		err              error
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
		err                error
		page               int64
		row                int64
		total              int64
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
		err               error
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
		err               error
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
		err               error
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
		err                     error
		page                    int64
		row                     int64
		total                   int64
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
		err                    error
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
		err                    error
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
		err                    error
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
		hosts         []*host.HostDto
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	//appServiceDto.State = appService.STATE_STOP //避免 状态不一致导致的坑
	appServiceDtos, _ := appServiceService.appServiceDao.GetAppServices(appServiceDto)

	if len(appServiceDtos) < 1 {
		return result.Error("应用不存在")
	}

	//修改应用为启动中
	//appServiceDto.State = appService.STATE_DOING
	//appServiceService.appServiceDao.UpdateAppService(appServiceDto)

	tmpAppServiceDto := appServiceDtos[0]

	if tmpAppServiceDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST {
		hostDto := host.HostDto{
			HostId: tmpAppServiceDto.AsDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	} else {
		hostDto := host.HostDto{
			GroupId: tmpAppServiceDto.AsDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	}

	param, err := containerScheduling.ContainerScheduling(hosts, tmpAppServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return param
}

// restart more apps
func (appServiceService *AppServiceService) RestartAppServices(ctx iris.Context) interface{} {

	var (
		err                   error
		appServiceDto         appService.AppServiceDto
		restartAppServicesDto appService.RestartAppServicesDto
	)

	if err = ctx.ReadJSON(&restartAppServicesDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDto.TenantId = user.TenantId

	appServiceDtos, _ := appServiceService.appServiceDao.GetAppServices(appServiceDto)
	if len(appServiceDtos) < 1 {
		return result.Error("应用不存在")
	}

	go appServiceService.doRestartAppServices(appServiceDtos, restartAppServicesDto)

	return result.Success()
}

// do restart app
func (appServiceService *AppServiceService) doRestartAppServices(appServiceDtos []*appService.AppServiceDto, restartAppServicesDto appService.RestartAppServicesDto) {
	for _, tmpAppServiceDto := range appServiceDtos {
		if !hasAppService(tmpAppServiceDto, restartAppServicesDto.AsIds) {
			continue
		}
		var (
			hosts []*host.HostDto
		)
		if tmpAppServiceDto.State != appService.STATE_STOP {
			containerScheduling.StopContainer(tmpAppServiceDto)
		}
		if tmpAppServiceDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST {
			hostDto := host.HostDto{
				HostId: tmpAppServiceDto.AsDeployId,
			}
			hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
		} else {
			hostDto := host.HostDto{
				GroupId: tmpAppServiceDto.AsDeployId,
			}
			hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
		}
		containerScheduling.ContainerScheduling(hosts, tmpAppServiceDto)
	}
}

func (appServiceService *AppServiceService) StopAppService(ctx iris.Context) interface{} {
	var (
		appServiceDto appService.AppServiceDto
	)

	if err := ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	//appServiceDto.State = appService.STATE_ONLINE //避免 状态不一致导致的坑
	appServiceDtos, _ := appServiceService.appServiceDao.GetAppServices(appServiceDto)

	if len(appServiceDtos) < 1 {
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
		err              error
		page             int64
		row              int64
		total            int64
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

// faster deploy
func (appServiceService *AppServiceService) SaveFasterDeploys(ctx iris.Context) result.ResultDto {

	var (
		err                   error
		fasterDeployDto       appService.FasterDeployDto
		businessPackageDto    businessPackage.BusinessPackageDto
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
		appServiceDto         appService.AppServiceDto
		appServicePort        appService.AppServicePortDto
	)

	if err = ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	businessPackageDto.Id = fasterDeployDto.PackageId
	businessPackageDto.TenantId = user.TenantId
	businessPackageDtos, _ := appServiceService.businessPackageDao.GetBusinessPackages(businessPackageDto)
	if businessPackageDtos == nil || len(businessPackageDtos) < 1 {
		return result.Error("项目包不存在")
	}

	fasterDeployDto.TenantId = user.TenantId
	fasterDeployDto.DeployId = seq.Generator()
	businessPackageDto.Id = seq.Generator()
	curDest := filepath.Join("businessPackage", user.TenantId, businessPackageDto.Id)
	dest := filepath.Join(config.WorkSpace, curDest)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}
	dest = filepath.Join(dest, "start_"+fasterDeployDto.AppName+".sh")
	if utils.IsFile(dest) {
		//f, err = os.OpenFile(dest, os.O_RDWR, 0600)
		os.Remove(dest)
	}
	fDest, err := os.Create(dest)

	defer fDest.Close()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = fDest.Write([]byte(fasterDeployDto.ShellPackageId))
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	//businessPackageDto.Path = filepath.Join(curDest, fileHeader.Filename)
	businessPackageDto.Path = filepath.Join(businessPackageDto.Id, "start_"+fasterDeployDto.AppName+".sh")
	businessPackageDto.Varsion = "V" + date.GetNowAString()
	businessPackageDto.Name = "start_" + fasterDeployDto.AppName + ".sh"

	err = appServiceService.businessPackageDao.SaveBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	fasterDeployDto.ShellPackageId = businessPackageDto.Id

	err = appServiceService.appServiceDao.SaveFasterDeploy(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//generator dockerfile

	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	businessDockerfileDto.Id = seq.Generator()
	businessDockerfileDto.Version = "V" + date.GetNowAString()
	businessDockerfileDto.Name = fasterDeployDto.AppName

	var dockerfile = ""
	if fasterDeployDto.DeployType == appService.DeployTypeJava {
		dockerfile += "FROM registry.cn-beijing.aliyuncs.com/sxd/ubuntu-java8:1.0\n"
	} else {
		dockerfile += "FROM centos:centos7\n"
	}
	dockerfile += "MAINTAINER zihao <928255095@qq.com>\n"
	dockerfile += "ADD " + businessPackageDtos[0].Path + " /root \n"
	dockerfile += "ADD " + businessPackageDto.Path + " /root \n"

	dockerfile += "RUN chmod u+x /root/" + businessPackageDto.Name + "\n"

	dockerfile += "CMD [\"/root/" + businessPackageDto.Name + "\"]\n"

	businessDockerfileDto.Dockerfile = dockerfile
	//save log
	logPath := filepath.Join(config.WorkSpace, "businessPackage/"+user.TenantId, businessDockerfileDto.Id+".log")

	businessDockerfileDto.LogPath = logPath

	err = appServiceService.businessDockerfileDao.SaveBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//generator images
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	//queue
	businessDockerfileDto.ImagesId = seq.Generator()
	businessDockerfileDto.VerId = seq.Generator()

	dockerfileQueue.SendData(&businessDockerfileDto)

	// save app
	appServiceDto.TenantId = user.TenantId
	appServiceDto.State = appService.STATE_STOP
	appServiceDto.AsId = seq.Generator()
	appServiceDto.AsName = fasterDeployDto.AppName
	appServiceDto.AsType = appService.AS_TYPE_SERVICE
	appServiceDto.AsDesc = fasterDeployDto.AppName
	appServiceDto.AsCount = "1"
	appServiceDto.AsGroupId = fasterDeployDto.AsGroupId
	appServiceDto.AsDeployType = fasterDeployDto.AsDeployType
	appServiceDto.AsDeployId = fasterDeployDto.AsDeployId
	appServiceDto.ImagesId = businessDockerfileDto.ImagesId
	appServiceDto.VerId = businessDockerfileDto.VerId

	err = appServiceService.appServiceDao.SaveAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if fasterDeployDto.OpenPort != "" {
		appServicePort.PortId = seq.Generator()
		appServicePort.TenantId = user.TenantId
		appServicePort.AsId = appServiceDto.AsId
		appServicePort.SrcPort = fasterDeployDto.OpenPort
		appServicePort.TargetPort = fasterDeployDto.OpenPort
		appServiceService.appServiceDao.SaveAppServicePort(appServicePort)
	}

	return result.SuccessData(fasterDeployDto)
}

func (appServiceService *AppServiceService) UpdateFasterDeploys(ctx iris.Context) result.ResultDto {
	var (
		err             error
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
		err             error
		fasterDeployDto appService.FasterDeployDto
	)

	if err = ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败")
	}

	// query fasterDeployDto
	//
	//tempFasterDeployDto := appService.FasterDeployDto{
	//	DeployId: fasterDeployDto.DeployId,
	//}
	//
	//fasterDeployDtos, _ := appServiceService.appServiceDao.GetFasterDeploys(tempFasterDeployDto)
	//
	//if len(fasterDeployDtos) < 1 {
	//	return result.Error("未找到快速部署")
	//}
	//
	////delete business package
	//tempBusinessPackageDto := businessPackage.BusinessPackageDto{
	//	Id: fasterDeployDtos[0].PackageId,
	//}
	//appServiceService.businessPackageDao.DeleteBusinessPackage(tempBusinessPackageDto)
	////delete business package shell
	//tempBusinessPackageDto = businessPackage.BusinessPackageDto{
	//	Id: fasterDeployDtos[0].ShellPackageId,
	//}
	//appServiceService.businessPackageDao.DeleteBusinessPackage(tempBusinessPackageDto)

	//delete app service

	err = appServiceService.appServiceDao.DeleteFasterDeploy(fasterDeployDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//delete

	return result.SuccessData(fasterDeployDto)

}

// export app config

func (appServiceService *AppServiceService) ExportAppService(ctx iris.Context) {
	var (
		err            error
		appServiceDto  appService.AppServiceDto
		composeYamlDto = composeYaml.ComposeYamlDto{
			Version: "2",
		}
		serviceDto composeYaml.ServiceDto
		services   []interface{}
	)

	//if err = ctx.ReadJSON(&appServiceDto); err != nil {
	//	return
	//}
	asIds := ctx.URLParam("asIds")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appServiceDto.TenantId = user.TenantId

	appServices, err := appServiceService.appServiceDao.GetAppServices(appServiceDto)
	if err != nil {
		return
	}

	for _, appS := range appServices {
		//get service
		if !hasAppService(appS, asIds) {
			continue
		}
		serviceDto = appServiceService.getServiceDto(appS)
		servicesDto := composeYaml.ServicesDto{
			Name:    appS.AsName,
			Service: serviceDto,
		}
		services = append(services, servicesDto.ToMap())
		composeYamlDto.Services = services
	}

	data, _ := yaml.Marshal(composeYamlDto)

	responseWriter := ctx.ResponseWriter()
	responseWriter.Header().Set("Content-Disposition", "attachment; filename=docker-compose.yml")
	responseWriter.Header().Set("Content-Type", "application/octet-stream")
	//responseWriter.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	responseWriter.Write(data)
	responseWriter.Flush()
}

// container images
func hasAppService(appServiceDto *appService.AppServiceDto, asIds string) bool {
	for _, images := range strings.Split(asIds, ",") {
		if images == appServiceDto.AsId {
			return true
		}
	}

	return false
}

func (appServiceService *AppServiceService) getServiceDto(appServiceDto *appService.AppServiceDto) composeYaml.ServiceDto {
	var (
		serviceDto composeYaml.ServiceDto
	)

	serviceDto.Image = appServiceDto.ImagesUrl

	portDto := appService.AppServicePortDto{
		AsId: appServiceDto.AsId,
	}
	portDtos, _ := appServiceService.appServiceDao.GetAppServicePort(portDto)

	if len(portDtos) > 0 {
		for _, port := range portDtos {
			portStr := port.SrcPort + ":" + port.TargetPort
			serviceDto.Ports = append(serviceDto.Ports, portStr)
		}
	}

	hostsDto := appService.AppServiceHostsDto{
		AsId: appServiceDto.AsId,
	}
	hostsDtos, _ := appServiceService.appServiceDao.GetAppServiceHosts(hostsDto)

	if len(hostsDtos) > 0 {
		for _, host := range hostsDtos {
			hostStr := host.Hostname + ":" + host.Ip
			serviceDto.ExtraHosts = append(serviceDto.ExtraHosts, hostStr)
		}
	}
	dirDto := appService.AppServiceDirDto{
		AsId: appServiceDto.AsId,
	}
	dirDtos, _ := appServiceService.appServiceDao.GetAppServiceDir(dirDto)

	if len(dirDtos) > 0 {
		for _, dir := range dirDtos {
			dirStr := dir.SrcDir + ":" + dir.TargetDir
			serviceDto.Volumes = append(serviceDto.Volumes, dirStr)
		}
	}

	varDto := appService.AppServiceVarDto{
		AsId: appServiceDto.AsId,
	}
	varDtos, _ := appServiceService.appServiceDao.GetAppServiceVars(varDto)
	if len(varDtos) > 0 {
		for _, vari := range varDtos {
			variStr := vari.VarSpec + ":" + vari.VarValue
			serviceDto.Environment = append(serviceDto.Environment, variStr)
		}
	}

	return serviceDto
}

// import app service

func (appServiceService *AppServiceService) ImportAppService(ctx iris.Context) result.ResultDto {
	var (
		composeYamlDto = composeYaml.ComposeYamlZiHaoDto{
		}
		serviceName   string
		appServiceDto appService.AppServiceDto
	)
	ctx.SetMaxRequestBodySize(maxSize)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	file, _, _ := ctx.FormFile("uploadFile")
	defer func() {
		file.Close()
	}()
	asType := ctx.FormValue("asType")
	asGroupId := ctx.FormValue("asGroupId")
	asDeployType := ctx.FormValue("asDeployType")
	asDeployId := ctx.FormValue("asDeployId")

	content, _ := ioutil.ReadAll(file)

	yaml.Unmarshal(content, &composeYamlDto)

	if len(composeYamlDto.Services) < 1 {
		return result.Error("yml文件中不包含应用信息")
	}

	for _, service := range composeYamlDto.Services {
		serviceMap := service.(map[string]interface{})

		for key := range serviceMap {
			serviceName = key
		}
		//
		//verId := seq.Generator()
		appServiceDto = appService.AppServiceDto{
			AsId:         seq.Generator(),
			AsName:       serviceName,
			AsType:       asType,
			TenantId:     user.TenantId,
			AsDesc:       serviceName,
			State:        "10012",
			AsCount:      "1",
			AsGroupId:    asGroupId,
			AsDeployType: asDeployType,
			AsDeployId:   asDeployId,
			//ImagesId: imagesId,
			//VerId:verId,
		}
		appServiceService.doImportAppService(appServiceDto, serviceMap[serviceName], user)
	}

	// zihao cmd
	if utils.IsEmpty(composeYamlDto.ZihaoCmd) {
		return result.Success()
	}

	var (
		hosts []*host.HostDto
	)

	if asDeployType == appService.AS_DEPLOY_TYPE_HOST {
		hostDto := host.HostDto{
			HostId: asDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	} else {
		hostDto := host.HostDto{
			GroupId: asDeployId,
		}
		hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
	}
	for _, host := range hosts {
		shell.ExecCommonShell(*host, composeYamlDto.ZihaoCmd)
	}

	return result.Success()
}

func (appServiceService *AppServiceService) doImportAppService(appServiceDto appService.AppServiceDto, info interface{}, user *user.UserDto) result.ResultDto {

	var (
		imagesId   string
		version    string = "V" + date.GetNowAString()
		serviceDto composeYaml.ServiceDto
		verId      string = seq.Generator()
		err        error
	)
	//objectConvert.Map2Struct(info.(map[string]interface{}), &serviceDto)
	data, _ := json.Marshal(info)
	json.Unmarshal(data, &serviceDto)
	//get images
	businessImagesDto := businessImages.BusinessImagesDto{
		Name:     appServiceDto.AsName,
		TenantId: appServiceDto.TenantId,
	}
	images, _ := appServiceService.businessImagesDao.GetBusinessImagess(businessImagesDto)

	if images == nil || len(images) < 1 {
		imagesId = seq.Generator()
		// save images
		businessImagesDto = businessImages.BusinessImagesDto{
			Id:           imagesId,
			Name:         appServiceDto.AsName,
			TenantId:     user.TenantId,
			CreateUserId: user.UserId,
			Version:      version,
			ImagesType:   businessImages.IMAGES_TYPE_IMPORT,
			ImagesFlag:   businessImages.IMAGES_FLAG_CUSTOM,
			TypeUrl:      serviceDto.Image,
		}
		err = appServiceService.businessImagesDao.SaveBusinessImages(businessImagesDto)
		if err != nil {
			return result.Error(err.Error())
		}
	} else {
		imagesId = images[0].Id
		businessImagesDto = businessImages.BusinessImagesDto{
			Id:      imagesId,
			Version: version,
			TypeUrl: serviceDto.Image,
		}
		err = appServiceService.businessImagesDao.UpdateBusinessImages(businessImagesDto)
		if err != nil {
			return result.Error(err.Error())
		}
	}
	appServiceDto.ImagesId = imagesId
	appServiceDto.VerId = verId

	//save images version
	businessImagesVerDto := businessImages.BusinessImagesVerDto{
		Id:       verId,
		ImagesId: imagesId,
		Version:  version,
		TypeUrl:  serviceDto.Image,
		TenantId: user.TenantId,
	}
	err = appServiceService.businessImagesVerDao.SaveBusinessImagesVer(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//save app service
	err = appServiceService.appServiceDao.SaveAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(serviceDto.Ports) > 0 {
		for _, port := range serviceDto.Ports {
			if !strings.Contains(port, ":") {
				continue
			}
			appServicePort := appService.AppServicePortDto{
				PortId:     seq.Generator(),
				AsId:       appServiceDto.AsId,
				TenantId:   appServiceDto.TenantId,
				SrcPort:    strings.Split(port, ":")[0],
				TargetPort: strings.Split(port, ":")[1],
			}
			appServiceService.appServiceDao.SaveAppServicePort(appServicePort)
		}
	}

	if len(serviceDto.ExtraHosts) > 0 {
		for _, host := range serviceDto.ExtraHosts {
			if !strings.Contains(host, ":") {
				continue
			}
			appServiceHosts := appService.AppServiceHostsDto{
				HostsId:  seq.Generator(),
				AsId:     appServiceDto.AsId,
				TenantId: appServiceDto.TenantId,
				Hostname: strings.Split(host, ":")[0],
				Ip:       strings.Split(host, ":")[1],
			}
			appServiceService.appServiceDao.SaveAppServiceHosts(appServiceHosts)
		}
	}

	if len(serviceDto.Volumes) > 0 {
		for _, volume := range serviceDto.Volumes {
			if !strings.Contains(volume, ":") {
				continue
			}
			appServiceDir := appService.AppServiceDirDto{
				DirId:     seq.Generator(),
				AsId:      appServiceDto.AsId,
				TenantId:  appServiceDto.TenantId,
				SrcDir:    strings.Split(volume, ":")[0],
				TargetDir: strings.Split(volume, ":")[1],
			}
			appServiceService.appServiceDao.SaveAppServiceDir(appServiceDir)
		}
	}

	if len(serviceDto.Environment) > 0 {
		for _, env := range serviceDto.Environment {
			if !strings.Contains(env, ":") {
				continue
			}
			appServiceVar := appService.AppServiceVarDto{
				AvId:     seq.Generator(),
				AsId:     appServiceDto.AsId,
				TenantId: appServiceDto.TenantId,
				VarSpec:  strings.SplitN(env, ":", 2)[0],
				VarValue: strings.SplitN(env, ":", 2)[1],
				VarName:  strings.SplitN(env, ":", 2)[0],
			}
			appServiceService.appServiceDao.SaveAppServiceVar(appServiceVar)
		}
	}

	return result.Success()
}

func (appServiceService *AppServiceService) UpgradeAppService(ctx iris.Context) interface{} {

	var (
		err              error
		tmpAppServiceDto appService.AppServiceDto
	)

	if err = ctx.ReadJSON(&tmpAppServiceDto); err != nil {
		return result.Error("解析入参失败")
	}

	appServiceDto := appService.AppServiceDto{
		AsId:  tmpAppServiceDto.AsId,
	}
	appServiceDtos, err := appServiceService.appServiceDao.GetAppServices(appServiceDto)

	if err != nil || len(appServiceDtos) < 1 {
		return result.Error("服务不存在或者未运行状态")
	}
	var hosts []*host.HostDto
	for _, appServiceDto := range appServiceDtos {
		tmpAppServiceDto := appService.AppServiceDto{
			AsId:    appServiceDto.AsId,
			VerId:   tmpAppServiceDto.VerId,
			AsCount: tmpAppServiceDto.AsCount,
		}
		appServiceService.appServiceDao.UpdateAppService(tmpAppServiceDto)

		//stop app service
		containerScheduling.StopContainer(appServiceDto)

		//start app service

		if appServiceDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST {
			hostDto := host.HostDto{
				HostId: appServiceDto.AsDeployId,
			}
			hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
		} else {
			hostDto := host.HostDto{
				GroupId: appServiceDto.AsDeployId,
			}
			hosts, _ = appServiceService.hostDao.GetHosts(hostDto)
		}

		if len(hosts) < 1 {
			return result.Error("主机不存在")
		}
		appServiceDto.VerId = tmpAppServiceDto.VerId
		appServiceDto.AsCount = tmpAppServiceDto.AsCount
		containerScheduling.ContainerScheduling(hosts, appServiceDto)
	}
	return result.Success()
}
