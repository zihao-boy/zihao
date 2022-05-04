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

	adinUser.Post("/uploadFile", hero.Handler(aus.uploadFile))

	adinUser.Post("/downloadFile", hero.Handler(aus.downloadFile))

	adinUser.Post("/downloadDir", hero.Handler(aus.downloadDir))

	adinUser.Post("/execShell", hero.Handler(aus.execShell))


	adinUser.Post("/startWaf", hero.Handler(aus.startWaf))

	adinUser.Post("/stopWaf", hero.Handler(aus.stopWaf))

	adinUser.Post("/refreshWafConfig", hero.Handler(aus.refreshWafConfig))

	adinUser.Post("/startInnerNet", hero.Handler(aus.startInnerNet))

	adinUser.Post("/stopInnerNet", hero.Handler(aus.stopInnerNet))

	adinUser.Post("/refreshInnerNetConfig", hero.Handler(aus.refreshInnerNetConfig))

	adinUser.Post("/startDns", hero.Handler(aus.startDns))

	adinUser.Post("/stopDns", hero.Handler(aus.stopDns))

	adinUser.Post("/refreshDnsConfig", hero.Handler(aus.refreshDnsConfig))

	adinUser.Post("/refreshFirewallRule", hero.Handler(aus.refreshFirewallRule))



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

//开启容器
func (aus *SlaveController) uploadFile(ctx iris.Context) {
	relustDto, _ := aus.systenInfoService.UploadFile(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) downloadFile(ctx iris.Context) {
	 aus.systenInfoService.DownloadFile(ctx)
}

//开启容器
func (aus *SlaveController) downloadDir(ctx iris.Context) {
	aus.systenInfoService.DownloadDir(ctx)
}

//开启容器
func (aus *SlaveController) execShell(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.ExecShell(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) startWaf(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StartWaf(ctx)
	ctx.JSON(relustDto)
}



//停止容器
func (aus *SlaveController) stopWaf(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StopWaf(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) refreshWafConfig(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.RefreshWafConfig(ctx)
	ctx.JSON(relustDto)
}



//开启容器
func (aus *SlaveController) startInnerNet(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StartInnerNet(ctx)
	ctx.JSON(relustDto)
}



//停止容器
func (aus *SlaveController) stopInnerNet(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StopInnerNet(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) refreshInnerNetConfig(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.RefreshInnerNetConfig(ctx)
	ctx.JSON(relustDto)
}




//开启容器
func (aus *SlaveController) startDns(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StartDns(ctx)
	ctx.JSON(relustDto)
}



//停止容器
func (aus *SlaveController) stopDns(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.StopDns(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) refreshDnsConfig(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.RefreshDns(ctx)
	ctx.JSON(relustDto)
}

//开启容器
func (aus *SlaveController) refreshFirewallRule(ctx iris.Context) {
	relustDto,_:=aus.systenInfoService.RefreshFirewallRule(ctx)
	ctx.JSON(relustDto)
}









