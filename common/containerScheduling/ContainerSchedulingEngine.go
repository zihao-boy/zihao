package containerScheduling

import (
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
	)

	//根据配置触发 调度引擎  目前暂时支持 k8s 和默认调度器
	if Scheduling_k8s == scheduling {
		resultDto , err = defalutContainerScheduling.Scheduling(hosts,appServiceDto)
	}else if(Scheduling_default == scheduling){
		resultDto , err = k8sContainerScheduling.Scheduling(hosts,appServiceDto)
	}

	return resultDto, err
}
