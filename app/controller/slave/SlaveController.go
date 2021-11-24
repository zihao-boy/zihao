package slave

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/system/service"
)

type SlaveController struct {
	systenInfoService service.SystemInfoService
	serviceSqlService service.ServiceSqlService
	mappingService    service.MappingService
}

func SlaveControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/slave")
		aus      = SlaveController{systenInfoService: service.SystemInfoService{},
			serviceSqlService: service.ServiceSqlService{},
			mappingService:    service.MappingService{}}
	)
	adinUser.Get("/info", hero.Handler(aus.info))

}

func (aus *SlaveController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}
