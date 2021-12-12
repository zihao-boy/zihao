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

	adinUser.Post("/startContainer", hero.Handler(aus.startContainer))

	adinUser.Post("/stopContainer", hero.Handler(aus.stopContainer))

	adinUser.Post("/listFiles", hero.Handler(aus.listFiles))

	adinUser.Post("/removeFile", hero.Handler(aus.removeFile))

	adinUser.Post("/newFile", hero.Handler(aus.newFile))

	adinUser.Post("/renameFile", hero.Handler(aus.renameFile))

	adinUser.Post("/listFileContext", hero.Handler(aus.listFileContext))


	adinUser.Post("/editFile", hero.Handler(aus.editFile))

}

func (aus *SlaveController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) startContainer(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.StartContainer(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) stopContainer(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.StopContainer(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) listFiles(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.ListFiles(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) removeFile(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.RemoveFile(ctx)
	ctx.JSON(relustDto)
}


//开启容器
func (aus *SlaveController) newFile(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.NewFile(ctx)
	ctx.JSON(relustDto)
}



//开启容器
func (aus *SlaveController) renameFile(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.RenameFile(ctx)
	ctx.JSON(relustDto)
}


//开启容器
func (aus *SlaveController) listFileContext(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.ListFileContext(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) editFile(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.EditFile(ctx)
	ctx.JSON(relustDto)
}





