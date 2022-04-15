package dns

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/dnsService"
)

type DnsController struct {
	dnsService dnsService.DnsService
	dnsUserService dnsService.DnsMapService



}

func DnsControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/dns")
		aus      = DnsController{dnsService: dnsService.DnsService{}}
	)
	//query platform data
	adinMenu.Get("/getDns", hero.Handler(aus.getDns))

	adinMenu.Post("/saveDns", hero.Handler(aus.saveDns))

	adinMenu.Post("/updateDns", hero.Handler(aus.updateDns))

	adinMenu.Post("/deleteDns", hero.Handler(aus.deleteDns))



	adinMenu.Post("/startDns", hero.Handler(aus.startDns))

	adinMenu.Post("/stopDns", hero.Handler(aus.stopDns))


	adinMenu.Post("/refreshDnsConfig", hero.Handler(aus.refreshDnsConfig))

	//query platform data
	adinMenu.Get("/getDnsMap", hero.Handler(aus.getDnsMap))

	adinMenu.Post("/saveDnsMap", hero.Handler(aus.saveDnsMap))

	adinMenu.Post("/updateDnsMap", hero.Handler(aus.updateDnsMap))

	adinMenu.Post("/deleteDnsMap", hero.Handler(aus.deleteDnsMap))
}

/**
query dns
*/
func (aus *DnsController) getDns(ctx iris.Context) {
	reslut := aus.dnsService.GetDnss(ctx)

	ctx.JSON(reslut)
}

/**
save dns
*/
func (aus *DnsController) saveDns(ctx iris.Context) {
	reslut := aus.dnsService.SaveDnss(ctx)

	ctx.JSON(reslut)
}

/**
update dns
*/
func (aus *DnsController) updateDns(ctx iris.Context) {
	reslut := aus.dnsService.UpdateDnss(ctx)

	ctx.JSON(reslut)
}

/**
delete dns
*/
func (aus *DnsController) deleteDns(ctx iris.Context) {
	reslut := aus.dnsService.DeleteDnss(ctx)

	ctx.JSON(reslut)
}

/**
delete dns
*/
func (aus *DnsController) startDns(ctx iris.Context) {
	reslut := aus.dnsService.StartDnsf(ctx)

	ctx.JSON(reslut)
}

/**
delete dns
*/
func (aus *DnsController) stopDns(ctx iris.Context) {
	reslut := aus.dnsService.StopDnsf(ctx)

	ctx.JSON(reslut)
}


/**
delete dns
*/
func (aus *DnsController) refreshDnsConfig(ctx iris.Context) {
	reslut := aus.dnsService.RefreshDnsConfig(ctx)

	ctx.JSON(reslut)
}





/**
query dns
*/
func (aus *DnsController) getDnsMap(ctx iris.Context) {
	reslut := aus.dnsUserService.GetDnsMaps(ctx)

	ctx.JSON(reslut)
}

/**
save dns
*/
func (aus *DnsController) saveDnsMap(ctx iris.Context) {
	reslut := aus.dnsUserService.SaveDnsMaps(ctx)

	ctx.JSON(reslut)
}

/**
update dns
*/
func (aus *DnsController) updateDnsMap(ctx iris.Context) {
	reslut := aus.dnsUserService.UpdateDnsMaps(ctx)

	ctx.JSON(reslut)
}

/**
delete dns
*/
func (aus *DnsController) deleteDnsMap(ctx iris.Context) {
	reslut := aus.dnsUserService.DeleteDnsMaps(ctx)

	ctx.JSON(reslut)
}
