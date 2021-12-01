package soft

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/softService/service"
)

type SoftController struct {
	businessPackageService service.BusinessPackageService
	businessDockerService service.BusinessDockerfileService
}

func SoftControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/soft")
		aus      = SoftController{
			businessPackageService: service.BusinessPackageService{},
			businessDockerService: service.BusinessDockerfileService{},
		}
	)
	adinMenu.Get("/getBusinessPackages", hero.Handler(aus.GetBusinessPackages))

	adinMenu.Post("/saveBusinessPackages", hero.Handler(aus.SaveBusinessPackages))

	adinMenu.Post("/updateBusinessPackages", hero.Handler(aus.UpdateBusinessPackages))

	adinMenu.Post("/deleteBusinessPackages", hero.Handler(aus.DeleteBusinessPackages))

	adinMenu.Get("/getBusinessDockerfile", hero.Handler(aus.GetBusinessDockerfile))

	adinMenu.Post("/saveBusinessDockerfile", hero.Handler(aus.SaveBusinessDockerfile))

	adinMenu.Post("/updateBusinessDockerfile", hero.Handler(aus.UpdateBusinessDockerfile))

	adinMenu.Post("/deleteBusinessDockerfile", hero.Handler(aus.DeleteBusinessDockerfiles))
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
查询 业务包
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
