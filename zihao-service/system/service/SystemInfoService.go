package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/date"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/system"
)

const(
	SYSTEM_NAME string = "梓豪平台"
	VERSION string = "v1.0"
)

type SystemInfoService struct {

}

/**
查询 系统信息
 */
func (*SystemInfoService) Info(context iris.Context) system.SystemDto {
	var systemDto = system.SystemDto{Id: seq.Generator(),Name:SYSTEM_NAME,Version: VERSION,Time: date.GetNowTimeString()}
	return systemDto
}
