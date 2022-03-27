package firewall

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/wafService"
)

type FirewallController struct {
	wafService wafService.WafService
	wafRouteService wafService.WafRouteService
	wafAccessLogService wafService.WafAccessLogService
}

func FirewallControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/firewall")
		aus      = FirewallController{wafService: wafService.WafService{}}
	)
	//query platform data
	adinMenu.Get("/getWaf", hero.Handler(aus.getWaf))

	adinMenu.Post("/saveWaf", hero.Handler(aus.saveWaf))

	adinMenu.Post("/updateWaf", hero.Handler(aus.updateWaf))

	adinMenu.Post("/deleteWaf", hero.Handler(aus.deleteWaf))

	adinMenu.Post("/startWaf", hero.Handler(aus.startWaf))

	adinMenu.Post("/stopWaf", hero.Handler(aus.stopWaf))

	adinMenu.Post("/refreshWafConfig", hero.Handler(aus.refreshWafConfig))

	//query platform data
	adinMenu.Get("/getWafRoute", hero.Handler(aus.getWafRoute))

	adinMenu.Post("/saveWafRoute", hero.Handler(aus.saveWafRoute))

	adinMenu.Post("/updateWafRoute", hero.Handler(aus.updateWafRoute))

	adinMenu.Post("/deleteWafRoute", hero.Handler(aus.deleteWafRoute))

	//query platform data
	adinMenu.Get("/getWafAccessLog", hero.Handler(aus.getWafAccessLog))

	adinMenu.Post("/saveWafAccessLog", hero.Handler(aus.saveWafAccessLog))

	//query platform data
	adinMenu.Get("/getWafAccessLogMap", hero.Handler(aus.getWafAccessLogMap))

}

/**
query waf
*/
func (aus *FirewallController) getWaf(ctx iris.Context) {
	reslut := aus.wafService.GetWafs(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWaf(ctx iris.Context) {
	reslut := aus.wafService.SaveWafs(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWaf(ctx iris.Context) {
	reslut := aus.wafService.UpdateWafs(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWaf(ctx iris.Context) {
	reslut := aus.wafService.DeleteWafs(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) startWaf(ctx iris.Context) {
	reslut := aus.wafService.StartWaff(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) stopWaf(ctx iris.Context) {
	reslut := aus.wafService.StopWaff(ctx)

	ctx.JSON(reslut)
}
/**
delete waf
*/
func (aus *FirewallController) refreshWafConfig(ctx iris.Context) {
	reslut := aus.wafService.RefreshWafConfig(ctx)

	ctx.JSON(reslut)
}



/**
query waf
*/
func (aus *FirewallController) getWafRoute(ctx iris.Context) {
	reslut := aus.wafRouteService.GetWafRoutes(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafRoute(ctx iris.Context) {
	reslut := aus.wafRouteService.SaveWafRoutes(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafRoute(ctx iris.Context) {
	reslut := aus.wafRouteService.UpdateWafRoutes(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafRoute(ctx iris.Context) {
	reslut := aus.wafRouteService.DeleteWafRoutes(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafAccessLog(ctx iris.Context) {
	reslut := aus.wafAccessLogService.GetWafAccessLogs(ctx)

	ctx.JSON(reslut)
}

/**
query waf
*/
func (aus *FirewallController) getWafAccessLogMap(ctx iris.Context) {
	reslut := aus.wafAccessLogService.GetWafAccessLogMap(ctx)

	ctx.JSON(reslut)
}


/**
save waf
*/
func (aus *FirewallController) saveWafAccessLog(ctx iris.Context) {
	reslut := aus.wafAccessLogService.SaveWafAccessLogs(ctx)

	ctx.JSON(reslut)
}