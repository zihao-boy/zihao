package soft

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/softService/service"
)

type SoftController struct {
	businessPackageService service.BusinessPackageService
	businessDockerService service.BusinessDockerfileService
	businessImagesService service.BusinessImagesService
	businessImagesVerService service.BusinessImagesVerService
}

func SoftControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/soft")
		aus      = SoftController{
			businessPackageService: service.BusinessPackageService{},
			businessDockerService: service.BusinessDockerfileService{},
			businessImagesService: service.BusinessImagesService{},
		}
	)
	adinMenu.Get("/getBusinessPackages", hero.Handler(aus.GetBusinessPackages))

	adinMenu.Post("/saveBusinessPackages", hero.Handler(aus.SaveBusinessPackages))

	adinMenu.Post("/updateBusinessPackages", hero.Handler(aus.UpdateBusinessPackages))

	adinMenu.Post("/deleteBusinessPackages", hero.Handler(aus.DeleteBusinessPackages))

	adinMenu.Get("/getBusinessDockerfile", hero.Handler(aus.GetBusinessDockerfile))

	adinMenu.Post("/saveBusinessDockerfile", hero.Handler(aus.SaveBusinessDockerfile))

	adinMenu.Post("/saveBusinessDockerfileCommon", hero.Handler(aus.saveBusinessDockerfileCommon))

	adinMenu.Post("/updateBusinessDockerfile", hero.Handler(aus.UpdateBusinessDockerfile))

	adinMenu.Post("/deleteBusinessDockerfile", hero.Handler(aus.DeleteBusinessDockerfiles))

	adinMenu.Get("/getBusinessImages", hero.Handler(aus.GetBusinessImages))

	adinMenu.Post("/saveBusinessImages", hero.Handler(aus.SaveBusinessImages))

	adinMenu.Post("/updateBusinessImages", hero.Handler(aus.UpdateBusinessImages))

	adinMenu.Post("/deleteBusinessImages", hero.Handler(aus.DeleteBusinessImages))

	adinMenu.Post("/generatorImages", hero.Handler(aus.GeneratorImages))

	adinMenu.Get("/getImagesPool", hero.Handler(aus.getImagesPool))

	adinMenu.Post("/installImages", hero.Handler(aus.installImages))

	adinMenu.Get("/getBusinessImagesVer", hero.Handler(aus.GetBusinessImagesVer))

	adinMenu.Post("/saveBusinessImagesVer", hero.Handler(aus.SaveBusinessImagesVer))

	adinMenu.Post("/updateBusinessImagesVer", hero.Handler(aus.UpdateBusinessImagesVer))

	adinMenu.Post("/deleteBusinessImagesVer", hero.Handler(aus.DeleteBusinessImagesVer))

	adinMenu.Post("/upload", hero.Handler(aus.upload))

	adinMenu.Get("/listBusinessPackageContext", hero.Handler(aus.listBusinessPackageContext))

	adinMenu.Post("/editBusinessPackageContext", hero.Handler(aus.editBusinessPackageContext))




}

/**
查询 业务包
*/
func (aus *SoftController) GetBusinessPackages(ctx iris.Context) {
	reslut := aus.businessPackageService.GetBusinessPackages(ctx)

	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *SoftController) SaveBusinessPackages(ctx iris.Context) {
	reslut := aus.businessPackageService.SaveBusinessPackages(ctx)

	ctx.JSON(reslut)
}

/**
修改 业务包
*/
func (aus *SoftController) UpdateBusinessPackages(ctx iris.Context) {
	reslut := aus.businessPackageService.UpdateBusinessPackages(ctx)

	ctx.JSON(reslut)
}

/**
删除 业务包
*/
func (aus *SoftController) DeleteBusinessPackages(ctx iris.Context) {
	reslut := aus.businessPackageService.DeleteBusinessPackages(ctx)

	ctx.JSON(reslut)
}


/**
查询 dockerfile
*/
func (aus *SoftController) GetBusinessDockerfile(ctx iris.Context) {
	reslut := aus.businessDockerService.GetBusinessDockerfiles(ctx)

	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *SoftController) SaveBusinessDockerfile(ctx iris.Context) {
	reslut := aus.businessDockerService.SaveBusinessDockerfiles(ctx)

	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *SoftController) saveBusinessDockerfileCommon(ctx iris.Context) {
	reslut := aus.businessDockerService.SaveBusinessDockerfileCommon(ctx)

	ctx.JSON(reslut)
}


/**
修改 业务包
*/
func (aus *SoftController) UpdateBusinessDockerfile(ctx iris.Context) {
	reslut := aus.businessDockerService.UpdateBusinessDockerfiles(ctx)

	ctx.JSON(reslut)
}

/**
删除 业务包
*/
func (aus *SoftController) DeleteBusinessDockerfiles(ctx iris.Context) {
	reslut := aus.businessDockerService.DeleteBusinessDockerfiles(ctx)

	ctx.JSON(reslut)
}


/**
查询 dockerfile
*/
func (aus *SoftController) GetBusinessImages(ctx iris.Context) {
	reslut := aus.businessImagesService.GetBusinessImages(ctx)

	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *SoftController) SaveBusinessImages(ctx iris.Context) {
	reslut := aus.businessImagesService.SaveBusinessImages(ctx)

	ctx.JSON(reslut)
}

/**
修改 业务包
*/
func (aus *SoftController) UpdateBusinessImages(ctx iris.Context) {
	reslut := aus.businessImagesService.UpdateBusinessImages(ctx)

	ctx.JSON(reslut)
}

/**
删除 业务包
*/
func (aus *SoftController) DeleteBusinessImages(ctx iris.Context) {
	reslut := aus.businessImagesService.DeleteBusinessImages(ctx)

	ctx.JSON(reslut)
}

/**
custom generator docker images
base on dockerfile
*/
func (aus *SoftController) GeneratorImages(ctx iris.Context) {
	reslut := aus.businessImagesService.GeneratorImages(ctx)

	ctx.JSON(reslut)
}

/**
查询远程镜像
*/
func (aus *SoftController) getImagesPool(ctx iris.Context) {
	reslut := aus.businessImagesService.GetImagesPool(ctx)

	ctx.JSON(reslut)
}

/**
查询远程镜像
*/
func (aus *SoftController) installImages(ctx iris.Context) {
	reslut := aus.businessImagesService.InstallImages(ctx)

	ctx.JSON(reslut)
}



/**
查询 dockerfile
*/
func (aus *SoftController) GetBusinessImagesVer(ctx iris.Context) {
	reslut := aus.businessImagesVerService.GetBusinessImagesVer(ctx)

	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *SoftController) SaveBusinessImagesVer(ctx iris.Context) {
	reslut := aus.businessImagesVerService.SaveBusinessImagesVer(ctx)

	ctx.JSON(reslut)
}

/**
修改 业务包
*/
func (aus *SoftController) UpdateBusinessImagesVer(ctx iris.Context) {
	reslut := aus.businessImagesVerService.UpdateBusinessImagesVer(ctx)

	ctx.JSON(reslut)
}

/**
删除 业务包
*/
func (aus *SoftController) DeleteBusinessImagesVer(ctx iris.Context) {
	reslut := aus.businessImagesVerService.DeleteBusinessImagesVer(ctx)

	ctx.JSON(reslut)
}

// upload file
func (aus *SoftController) upload(ctx iris.Context) {
	reslut := aus.businessPackageService.Upload(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *SoftController) listBusinessPackageContext(ctx iris.Context) {
	reslut := aus.businessPackageService.ListBusinessPackageContext(ctx)

	ctx.JSON(reslut)
}

/**
查询文件
**/
func (aus *SoftController) editBusinessPackageContext(ctx iris.Context) {
	reslut := aus.businessPackageService.EditBusinessPackageContext(ctx)

	ctx.JSON(reslut)
}





