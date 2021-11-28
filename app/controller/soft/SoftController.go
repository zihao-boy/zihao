package soft

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/softService/service"
)

type SoftController struct {
	businessPackageService service.BusinessPackageService
}

func SoftControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/soft")
		aus      = SoftController{businessPackageService: service.BusinessPackageService{}}
	)
	adinMenu.Get("/getBusinessPackages", hero.Handler(aus.GetBusinessPackages))

	adinMenu.Post("/saveBusinessPackages", hero.Handler(aus.SaveBusinessPackages))

	adinMenu.Post("/updateBusinessPackages", hero.Handler(aus.UpdateBusinessPackages))

	adinMenu.Post("/deleteBusinessPackages", hero.Handler(aus.DeleteBusinessPackages))
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
