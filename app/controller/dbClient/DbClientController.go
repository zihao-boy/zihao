package dbClient

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/dbLinkService"
)

type DbClientController struct {
	dbLinkService dbLinkService.DbLinkService
}

func DbClientControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/dbClient")
		aus      = DbClientController{
			dbLinkService: dbLinkService.DbLinkService{},
		}
	)
	adinMenu.Get("/getDbLink", hero.Handler(aus.getDbLink))
}

/**
查询 业务包
*/
func (aus *DbClientController) getDbLink(ctx iris.Context) {
	reslut := aus.dbLinkService.GetDbLinks(ctx)
	ctx.JSON(reslut)
}

/**
添加 业务包
*/
func (aus *DbClientController) SaveBusinessPackages(ctx iris.Context) {
	reslut := aus.dbLinkService.SaveDbLinks(ctx)

	ctx.JSON(reslut)
}

/**
修改 业务包
*/
func (aus *DbClientController) UpdateBusinessImagesVer(ctx iris.Context) {
	reslut := aus.dbLinkService.UpdateDbLinks(ctx)

	ctx.JSON(reslut)
}

/**
删除 业务包
*/
func (aus *DbClientController) DeleteBusinessImagesVer(ctx iris.Context) {
	reslut := aus.dbLinkService.DeleteDbLinks(ctx)

	ctx.JSON(reslut)
}
