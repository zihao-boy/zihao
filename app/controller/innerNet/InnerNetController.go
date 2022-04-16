package innerNet

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/innerNetService"
)

type InnerNetController struct {
	innerNetService innerNetService.InnerNetService
	innerNetUserService innerNetService.InnerNetUserService
	innerNetPrivilegeService innerNetService.InnerNetPrivilegeService
}

func InnerNetControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/innerNet")
		aus      = InnerNetController{innerNetService: innerNetService.InnerNetService{}}
	)
	//query platform data
	adinMenu.Get("/getInnerNet", hero.Handler(aus.getInnerNet))

	adinMenu.Post("/saveInnerNet", hero.Handler(aus.saveInnerNet))

	adinMenu.Post("/updateInnerNet", hero.Handler(aus.updateInnerNet))

	adinMenu.Post("/deleteInnerNet", hero.Handler(aus.deleteInnerNet))



	adinMenu.Post("/startInnerNet", hero.Handler(aus.startInnerNet))

	adinMenu.Post("/stopInnerNet", hero.Handler(aus.stopInnerNet))


	adinMenu.Post("/refreshInnerNetConfig", hero.Handler(aus.refreshInnerNetConfig))

	//query platform data
	adinMenu.Get("/getInnerNetUser", hero.Handler(aus.getInnerNetUser))

	adinMenu.Post("/saveInnerNetUser", hero.Handler(aus.saveInnerNetUser))

	adinMenu.Post("/updateInnerNetUser", hero.Handler(aus.updateInnerNetUser))

	adinMenu.Post("/deleteInnerNetUser", hero.Handler(aus.deleteInnerNetUser))

	//query platform data
	adinMenu.Get("/getInnerNetPrivilege", hero.Handler(aus.getInnerNetPrivilege))

	adinMenu.Post("/saveInnerNetPrivilege", hero.Handler(aus.saveInnerNetPrivilege))

	adinMenu.Post("/updateInnerNetPrivilege", hero.Handler(aus.updateInnerNetPrivilege))

	adinMenu.Post("/deleteInnerNetPrivilege", hero.Handler(aus.deleteInnerNetPrivilege))
}

/**
query innerNet
*/
func (aus *InnerNetController) getInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.GetInnerNets(ctx)

	ctx.JSON(reslut)
}

/**
save innerNet
*/
func (aus *InnerNetController) saveInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.SaveInnerNets(ctx)

	ctx.JSON(reslut)
}

/**
update innerNet
*/
func (aus *InnerNetController) updateInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.UpdateInnerNets(ctx)

	ctx.JSON(reslut)
}

/**
delete innerNet
*/
func (aus *InnerNetController) deleteInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.DeleteInnerNets(ctx)

	ctx.JSON(reslut)
}

/**
delete innerNet
*/
func (aus *InnerNetController) startInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.StartInnerNetf(ctx)

	ctx.JSON(reslut)
}

/**
delete innerNet
*/
func (aus *InnerNetController) stopInnerNet(ctx iris.Context) {
	reslut := aus.innerNetService.StopInnerNetf(ctx)

	ctx.JSON(reslut)
}


/**
delete innerNet
*/
func (aus *InnerNetController) refreshInnerNetConfig(ctx iris.Context) {
	reslut := aus.innerNetService.RefreshInnerNetConfig(ctx)

	ctx.JSON(reslut)
}





/**
query innerNet
*/
func (aus *InnerNetController) getInnerNetUser(ctx iris.Context) {
	reslut := aus.innerNetUserService.GetInnerNetUsers(ctx)

	ctx.JSON(reslut)
}

/**
save innerNet
*/
func (aus *InnerNetController) saveInnerNetUser(ctx iris.Context) {
	reslut := aus.innerNetUserService.SaveInnerNetUsers(ctx)

	ctx.JSON(reslut)
}

/**
update innerNet
*/
func (aus *InnerNetController) updateInnerNetUser(ctx iris.Context) {
	reslut := aus.innerNetUserService.UpdateInnerNetUsers(ctx)

	ctx.JSON(reslut)
}

/**
delete innerNet
*/
func (aus *InnerNetController) deleteInnerNetUser(ctx iris.Context) {
	reslut := aus.innerNetUserService.DeleteInnerNetUsers(ctx)

	ctx.JSON(reslut)
}


/**
query innerNet
*/
func (aus *InnerNetController) getInnerNetPrivilege(ctx iris.Context) {
	reslut := aus.innerNetPrivilegeService.GetInnerNetPrivileges(ctx)

	ctx.JSON(reslut)
}

/**
save innerNet
*/
func (aus *InnerNetController) saveInnerNetPrivilege(ctx iris.Context) {
	reslut := aus.innerNetPrivilegeService.SaveInnerNetPrivileges(ctx)

	ctx.JSON(reslut)
}

/**
update innerNet
*/
func (aus *InnerNetController) updateInnerNetPrivilege(ctx iris.Context) {
	reslut := aus.innerNetPrivilegeService.UpdateInnerNetPrivileges(ctx)

	ctx.JSON(reslut)
}

/**
delete innerNet
*/
func (aus *InnerNetController) deleteInnerNetPrivilege(ctx iris.Context) {
	reslut := aus.innerNetPrivilegeService.DeleteInnerNetPrivileges(ctx)

	ctx.JSON(reslut)
}
