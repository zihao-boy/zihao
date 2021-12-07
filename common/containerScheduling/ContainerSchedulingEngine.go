package containerScheduling

import (
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/containerScheduling/defalutContainerScheduling"
	"github.com/zihao-boy/zihao/common/containerScheduling/k8sContainerScheduling"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
)

const (
	Scheduling_default string = "default"
	Scheduling_k8s     string = "k8s"
)

// 容器调度引擎
// 目前支持 k8s 和 默认 引擎
func ContainerScheduling(hosts []*host.HostDto, appServiceDto *appService.AppServiceDto) (interface{}, error) {

	//调度器
	var (
		scheduling = config.G_AppConfig.ContainerScheduling
		err        error
		resultDto  result.ResultDto
		appServiceDao dao.AppServiceDao
	)

	//查询 目录映射
	dirDto :=appService.AppServiceDirDto{
		AsId: appServiceDto.AsId,
		TenantId: appServiceDto.TenantId,
	}
	dirDtos,_ := appServiceDao.GetAppServiceDir(dirDto)
	appServiceDto.AppServiceDirs = dirDtos

	//查询 hosts
	hostsDto :=appService.AppServiceHostsDto{
		AsId: appServiceDto.AsId,
		TenantId: appServiceDto.TenantId,
	}
	hostsDtos,_ := appServiceDao.GetAppServiceHosts(hostsDto)
	appServiceDto.AppServiceHosts = hostsDtos

	//查询 环境变量
	varDto :=appService.AppServiceVarDto{
		AsId: appServiceDto.AsId,
		TenantId: appServiceDto.TenantId,
	}
	varDtos,_ := appServiceDao.GetAppServiceVars(varDto)
	appServiceDto.AppServiceVars = varDtos

	//查询 环境变量
	portDto :=appService.AppServicePortDto{
		AsId: appServiceDto.AsId,
		TenantId: appServiceDto.TenantId,
	}
	portDtos,_ := appServiceDao.GetAppServicePort(portDto)
	appServiceDto.AppServicePorts = portDtos

	//根据配置触发 调度引擎  目前暂时支持 k8s 和默认调度器
	if Scheduling_k8s == scheduling {
		resultDto , err = defalutContainerScheduling.Scheduling(hosts,appServiceDto)
	}else if(Scheduling_default == scheduling){
		resultDto , err = k8sContainerScheduling.Scheduling(hosts,appServiceDto)
	}

	return resultDto, err
}
