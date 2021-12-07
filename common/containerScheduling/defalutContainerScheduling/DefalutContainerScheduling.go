package defalutContainerScheduling

import (
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
)

//default  scheduling
// base on mem
//add by wuxw 2021-12-07
func Scheduling(hosts []*host.HostDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {

	return result.Success(),nil
}