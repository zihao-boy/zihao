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
	wafRuleGroupService wafService.WafRuleGroupService
	wafRuleService wafService.WafRuleService
	wafIpBlackWhiteService wafService.WafIpBlackWhiteService
	wafAreaService wafService.WafAreaService
	wafCCService wafService.WafCCService
	wafAccurateService wafService.WafAccurateService



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

	adinMenu.Get("/loadIps", hero.Handler(aus.loadIps))

	//query platform data
	adinMenu.Get("/getWafAccessLogMap", hero.Handler(aus.getWafAccessLogMap))

	//query platform data
	adinMenu.Get("/getWafAccessLogTop5", hero.Handler(aus.getWafAccessLogTop5))

	adinMenu.Get("/getWafAccessLogIntercept", hero.Handler(aus.getWafAccessLogIntercept))



	adinMenu.Get("/getWafRuleGroup", hero.Handler(aus.getWafRuleGroup))

	adinMenu.Post("/saveWafRuleGroup", hero.Handler(aus.saveWafRuleGroup))

	adinMenu.Post("/updateWafRuleGroup", hero.Handler(aus.updateWafRuleGroup))
	adinMenu.Post("/startWafRuleGroup", hero.Handler(aus.startWafRuleGroup))

	adinMenu.Post("/deleteWafRuleGroup", hero.Handler(aus.deleteWafRuleGroup))

	adinMenu.Get("/getWafRule", hero.Handler(aus.getWafRule))

	adinMenu.Post("/saveWafRule", hero.Handler(aus.saveWafRule))

	adinMenu.Post("/updateWafRule", hero.Handler(aus.updateWafRule))

	adinMenu.Post("/deleteWafRule", hero.Handler(aus.deleteWafRule))

	adinMenu.Get("/getWafIpBlackWhite", hero.Handler(aus.getWafIpBlackWhite))

	adinMenu.Post("/saveWafIpBlackWhite", hero.Handler(aus.saveWafIpBlackWhite))

	adinMenu.Post("/updateWafIpBlackWhite", hero.Handler(aus.updateWafIpBlackWhite))

	adinMenu.Post("/deleteWafIpBlackWhite", hero.Handler(aus.deleteWafIpBlackWhite))

	adinMenu.Get("/getWafArea", hero.Handler(aus.getWafArea))

	adinMenu.Post("/saveWafArea", hero.Handler(aus.saveWafArea))

	adinMenu.Post("/updateWafArea", hero.Handler(aus.updateWafArea))

	adinMenu.Post("/deleteWafArea", hero.Handler(aus.deleteWafArea))

	adinMenu.Get("/getWafCC", hero.Handler(aus.getWafCC))

	adinMenu.Post("/saveWafCC", hero.Handler(aus.saveWafCC))

	adinMenu.Post("/updateWafCC", hero.Handler(aus.updateWafCC))

	adinMenu.Post("/deleteWafCC", hero.Handler(aus.deleteWafCC))

	adinMenu.Get("/getWafAccurate", hero.Handler(aus.getWafAccurate))

	adinMenu.Post("/saveWafAccurate", hero.Handler(aus.saveWafAccurate))

	adinMenu.Post("/updateWafAccurate", hero.Handler(aus.updateWafAccurate))

	adinMenu.Post("/deleteWafAccurate", hero.Handler(aus.deleteWafAccurate))
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
query waf
*/
func (aus *FirewallController) getWafAccessLogTop5(ctx iris.Context) {
	reslut := aus.wafAccessLogService.GetWafAccessLogTop5(ctx)

	ctx.JSON(reslut)
}

/**
query waf
*/
func (aus *FirewallController) getWafAccessLogIntercept(ctx iris.Context) {
	reslut := aus.wafAccessLogService.GetWafAccessLogIntercept(ctx)

	ctx.JSON(reslut)
}


/**
save waf
*/
func (aus *FirewallController) saveWafAccessLog(ctx iris.Context) {
	reslut := aus.wafAccessLogService.SaveWafAccessLogs(ctx)

	ctx.JSON(reslut)
}

func (aus *FirewallController) loadIps(ctx iris.Context)  {
	aus.wafAccessLogService.LoadIps(ctx)
}

/**
query waf
*/
func (aus *FirewallController) getWafRuleGroup(ctx iris.Context) {
	reslut := aus.wafRuleGroupService.GetWafRuleGroups(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafRuleGroup(ctx iris.Context) {
	reslut := aus.wafRuleGroupService.SaveWafRuleGroups(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafRuleGroup(ctx iris.Context) {
	reslut := aus.wafRuleGroupService.UpdateWafRuleGroups(ctx)

	ctx.JSON(reslut)
}


/**
delete waf
*/
func (aus *FirewallController) startWafRuleGroup(ctx iris.Context) {
	reslut := aus.wafRuleGroupService.StartWafRuleGroup(ctx)

	ctx.JSON(reslut)
}


/**
delete waf
*/
func (aus *FirewallController) deleteWafRuleGroup(ctx iris.Context) {
	reslut := aus.wafRuleGroupService.DeleteWafRuleGroups(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafRule(ctx iris.Context) {
	reslut := aus.wafRuleService.GetWafRules(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafRule(ctx iris.Context) {
	reslut := aus.wafRuleService.SaveWafRules(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafRule(ctx iris.Context) {
	reslut := aus.wafRuleService.UpdateWafRules(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafRule(ctx iris.Context) {
	reslut := aus.wafRuleService.DeleteWafRules(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafIpBlackWhite(ctx iris.Context) {
	reslut := aus.wafIpBlackWhiteService.GetWafIpBlackWhites(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafIpBlackWhite(ctx iris.Context) {
	reslut := aus.wafIpBlackWhiteService.SaveWafIpBlackWhites(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafIpBlackWhite(ctx iris.Context) {
	reslut := aus.wafIpBlackWhiteService.UpdateWafIpBlackWhites(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafIpBlackWhite(ctx iris.Context) {
	reslut := aus.wafIpBlackWhiteService.DeleteWafIpBlackWhites(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafArea(ctx iris.Context) {
	reslut := aus.wafAreaService.GetWafAreas(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafArea(ctx iris.Context) {
	reslut := aus.wafAreaService.SaveWafAreas(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafArea(ctx iris.Context) {
	reslut := aus.wafAreaService.UpdateWafAreas(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafArea(ctx iris.Context) {
	reslut := aus.wafAreaService.DeleteWafAreas(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafCC(ctx iris.Context) {
	reslut := aus.wafCCService.GetWafCCs(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafCC(ctx iris.Context) {
	reslut := aus.wafCCService.SaveWafCCs(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafCC(ctx iris.Context) {
	reslut := aus.wafCCService.UpdateWafCCs(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafCC(ctx iris.Context) {
	reslut := aus.wafCCService.DeleteWafCCs(ctx)

	ctx.JSON(reslut)
}


/**
query waf
*/
func (aus *FirewallController) getWafAccurate(ctx iris.Context) {
	reslut := aus.wafAccurateService.GetWafAccurates(ctx)

	ctx.JSON(reslut)
}

/**
save waf
*/
func (aus *FirewallController) saveWafAccurate(ctx iris.Context) {
	reslut := aus.wafAccurateService.SaveWafAccurates(ctx)

	ctx.JSON(reslut)
}

/**
update waf
*/
func (aus *FirewallController) updateWafAccurate(ctx iris.Context) {
	reslut := aus.wafAccurateService.UpdateWafAccurates(ctx)

	ctx.JSON(reslut)
}

/**
delete waf
*/
func (aus *FirewallController) deleteWafAccurate(ctx iris.Context) {
	reslut := aus.wafAccurateService.DeleteWafAccurates(ctx)

	ctx.JSON(reslut)
}