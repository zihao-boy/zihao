package firewall

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/wafService"
)

type FirewallController struct {
	wafService wafService.WafService
	wafRouteService wafService.WafRouteService

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

	//query platform data
	adinMenu.Get("/getWafRoute", hero.Handler(aus.getWafRoute))

	adinMenu.Post("/saveWafRoute", hero.Handler(aus.saveWafRoute))

	adinMenu.Post("/updateWafRoute", hero.Handler(aus.updateWafRoute))

	adinMenu.Post("/deleteWafRoute", hero.Handler(aus.deleteWafRoute))
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