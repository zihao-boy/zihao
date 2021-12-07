package k8sContainerScheduling

import (
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
)
//k8s 调度器
func Scheduling(hosts []*host.HostDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {
	return result.Success(),nil
}


//default  stop
// base on mem
//add by wuxw 2021-12-07
func StopContainer(containerDto *appService.AppServiceContainerDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {
	return result.Success(),nil
}