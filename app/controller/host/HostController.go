package host

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/assets/service"
)

type HostController struct {
	hostService service.HostService
}

func HostControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/host")
		aus      = HostController{hostService: service.HostService{}}
	)
	adinMenu.Get("/getHostGroup", hero.Handler(aus.getHostGroup))

	adinMenu.Post("/saveHostGroup", hero.Handler(aus.saveHostGroup))

	adinMenu.Post("/updateHostGroup", hero.Handler(aus.updateHostGroup))

	adinMenu.Post("/deleteHostGroup", hero.Handler(aus.deleteHostGroup))

	adinMenu.Get("/getHosts", hero.Handler(aus.getHosts))

	adinMenu.Post("/saveHost", hero.Handler(aus.saveHost))

	adinMenu.Post("/updateHost", hero.Handler(aus.updateHost))

	adinMenu.Post("/deleteHost", hero.Handler(aus.deleteHost))

	adinMenu.Get("/getHostToken", hero.Handler(aus.getHostToken))

	adinMenu.Get("/getContainers", hero.Handler(aus.getContainers))

	adinMenu.Get("/getHostResource", hero.Handler(aus.getHostResource))

	adinMenu.Get("/getHostPort", hero.Handler(aus.getHostPort))

	adinMenu.Get("/getHostDisk", hero.Handler(aus.getHostDisk))

	adinMenu.Post("/controlHost", hero.Handler(aus.controlHost))

	adinMenu.Post("/slaveHealth", hero.Handler(aus.slaveHealth))

	adinMenu.Get("/listFiles", hero.Handler(aus.listFiles))

	adinMenu.Post("/removeFile", hero.Handler(aus.removeFile))

	adinMenu.Post("/newFile", hero.Handler(aus.newFile))

	adinMenu.Post("/renameFile", hero.Handler(aus.renameFile))


	adinMenu.Get("/listFileContext", hero.Handler(aus.listFileContext))


	adinMenu.Post("/editFile", hero.Handler(aus.editFile))

	adinMenu.Post("/uploadFile", hero.Handler(aus.uploadFile))

	adinMenu.Post("/downloadFile", hero.Handler(aus.downloadFile))


}

/**
查询 主机组
*/
func (aus *HostController) getHostGroup(ctx iris.Context) {
	reslut := aus.hostService.GetHostGroups(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *HostController) saveHostGroup(ctx iris.Context) {
	reslut := aus.hostService.SaveHostGroups(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *HostController) updateHostGroup(ctx iris.Context) {
	reslut := aus.hostService.UpdateHostGroups(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *HostController) deleteHostGroup(ctx iris.Context) {
	reslut := aus.hostService.DeleteHostGroups(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *HostController) getHosts(ctx iris.Context) {
	reslut := aus.hostService.GetHosts(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *HostController) saveHost(ctx iris.Context) {
	reslut := aus.hostService.SaveHost(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *HostController) updateHost(ctx iris.Context) {
	reslut := aus.hostService.UpdateHost(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *HostController) deleteHost(ctx iris.Context) {
	reslut := aus.hostService.DeleteHost(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *HostController) getHostToken(ctx iris.Context) {
	reslut := aus.hostService.GetHostToken(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *HostController) getContainers(ctx iris.Context) {
	reslut := aus.hostService.GetContainers(ctx)

	ctx.JSON(reslut)
}

/**
查询主机资源
*/
func (aus *HostController) getHostResource(ctx iris.Context) {
	reslut := aus.hostService.GetHostResource(ctx)

	ctx.JSON(reslut)
}

/**
查询主机监听端口
*/
func (aus *HostController) getHostPort(ctx iris.Context) {
	reslut := aus.hostService.GetHostPort(ctx)

	ctx.JSON(reslut)
}

/**
查询主机监听端口
*/
func (aus *HostController) getHostDisk(ctx iris.Context) {
	reslut := aus.hostService.GetHostDisk(ctx)

	ctx.JSON(reslut)
}

/**
控制主机
**/
func (aus *HostController) controlHost(ctx iris.Context) {
	reslut := aus.hostService.ControlHost(ctx)

	ctx.JSON(reslut)
}

/**
控制主机
**/
func (aus *HostController) slaveHealth(ctx iris.Context) {
	reslut := aus.hostService.SlaveHealth(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *HostController) listFiles(ctx iris.Context) {
	reslut := aus.hostService.ListFiles(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *HostController) removeFile(ctx iris.Context) {
	reslut := aus.hostService.RemoveFile(ctx)

	ctx.JSON(reslut)
}


/**
查询文件
**/
func (aus *HostController) newFile(ctx iris.Context) {
	reslut := aus.hostService.NewFile(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *HostController) renameFile(ctx iris.Context) {
	reslut := aus.hostService.RenameFile(ctx)

	ctx.JSON(reslut)
}


/**
查询文件
**/
func (aus *HostController) listFileContext(ctx iris.Context) {
	reslut := aus.hostService.ListFileContext(ctx)

	ctx.JSON(reslut)
}


/**
查询文件
**/
func (aus *HostController) editFile(ctx iris.Context) {
	reslut := aus.hostService.EditFile(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *HostController) uploadFile(ctx iris.Context) {
	reslut := aus.hostService.UploadFile(ctx)

	ctx.JSON(reslut)
}

//下载文件
func (aus *HostController) downloadFile(ctx iris.Context) {
	aus.hostService.DownloadFile(ctx)
}




