package vpn

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/vpnService"
)

type VpnController struct {
	vpnService vpnService.VpnService
	vpnUserService vpnService.VpnUserService



}

func VpnControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/vpn")
		aus      = VpnController{vpnService: vpnService.VpnService{}}
	)
	//query platform data
	adinMenu.Get("/getVpn", hero.Handler(aus.getVpn))

	adinMenu.Post("/saveVpn", hero.Handler(aus.saveVpn))

	adinMenu.Post("/updateVpn", hero.Handler(aus.updateVpn))

	adinMenu.Post("/deleteVpn", hero.Handler(aus.deleteVpn))



	adinMenu.Post("/startVpn", hero.Handler(aus.startVpn))

	adinMenu.Post("/stopVpn", hero.Handler(aus.stopVpn))


	adinMenu.Post("/refreshVpnConfig", hero.Handler(aus.refreshVpnConfig))

	//query platform data
	adinMenu.Get("/getVpnUser", hero.Handler(aus.getVpnUser))

	adinMenu.Post("/saveVpnUser", hero.Handler(aus.saveVpnUser))

	adinMenu.Post("/updateVpnUser", hero.Handler(aus.updateVpnUser))

	adinMenu.Post("/deleteVpnUser", hero.Handler(aus.deleteVpnUser))
}

/**
query vpn
*/
func (aus *VpnController) getVpn(ctx iris.Context) {
	reslut := aus.vpnService.GetVpns(ctx)

	ctx.JSON(reslut)
}

/**
save vpn
*/
func (aus *VpnController) saveVpn(ctx iris.Context) {
	reslut := aus.vpnService.SaveVpns(ctx)

	ctx.JSON(reslut)
}

/**
update vpn
*/
func (aus *VpnController) updateVpn(ctx iris.Context) {
	reslut := aus.vpnService.UpdateVpns(ctx)

	ctx.JSON(reslut)
}

/**
delete vpn
*/
func (aus *VpnController) deleteVpn(ctx iris.Context) {
	reslut := aus.vpnService.DeleteVpns(ctx)

	ctx.JSON(reslut)
}

/**
delete vpn
*/
func (aus *VpnController) startVpn(ctx iris.Context) {
	reslut := aus.vpnService.StartVpnf(ctx)

	ctx.JSON(reslut)
}

/**
delete vpn
*/
func (aus *VpnController) stopVpn(ctx iris.Context) {
	reslut := aus.vpnService.StopVpnf(ctx)

	ctx.JSON(reslut)
}


/**
delete vpn
*/
func (aus *VpnController) refreshVpnConfig(ctx iris.Context) {
	reslut := aus.vpnService.RefreshVpnConfig(ctx)

	ctx.JSON(reslut)
}





/**
query vpn
*/
func (aus *VpnController) getVpnUser(ctx iris.Context) {
	reslut := aus.vpnUserService.GetVpnUsers(ctx)

	ctx.JSON(reslut)
}

/**
save vpn
*/
func (aus *VpnController) saveVpnUser(ctx iris.Context) {
	reslut := aus.vpnUserService.SaveVpnUsers(ctx)

	ctx.JSON(reslut)
}

/**
update vpn
*/
func (aus *VpnController) updateVpnUser(ctx iris.Context) {
	reslut := aus.vpnUserService.UpdateVpnUsers(ctx)

	ctx.JSON(reslut)
}

/**
delete vpn
*/
func (aus *VpnController) deleteVpnUser(ctx iris.Context) {
	reslut := aus.vpnUserService.DeleteVpnUsers(ctx)

	ctx.JSON(reslut)
}
