package vpnService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/vpnDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"strconv"
	"strings"
)

type VpnService struct {
	vpnDao      vpnDao.VpnDao
	vpnHostsDao vpnDao.VpnHostsDao
}

// get db link
// all db by this user
func (vpnService *VpnService) GetVpnAll(VpnDto vpn.VpnDto) ([]*vpn.VpnDto, error) {
	var (
		err     error
		VpnDtos []*vpn.VpnDto
	)

	VpnDtos, err = vpnService.vpnDao.GetVpns(VpnDto)
	if err != nil {
		return nil, err
	}

	return VpnDtos, nil

}

/**
查询 系统信息
*/
func (vpnService *VpnService) GetVpns(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		vpnDto  = vpn.VpnDto{}
		vpnDtos []*vpn.VpnDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	vpnDto.Row = row * page

	vpnDto.Page = (page - 1) * row

	total, err = vpnService.vpnDao.GetVpnCount(vpnDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	vpnDtos, err = vpnService.vpnDao.GetVpns(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, vpnDto := range vpnDtos {

		vpnHostsDto := vpn.VpnHostsDto{
			VpnId: vpnDto.VpnId,
		}
		vpnHostsDtos, _ := vpnService.vpnHostsDao.GetVpnHostss(vpnHostsDto)
		vpnDto.VpnHosts = vpnHostsDtos
	}

	return result.SuccessData(vpnDtos, total, row)

}

/**
保存 系统信息
*/
func (vpnService *VpnService) SaveVpns(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}
	vpnDto.VpnId = seq.Generator()
	//VpnDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = vpnService.vpnDao.SaveVpn(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}

/**
修改 系统信息
*/
func (vpnService *VpnService) UpdateVpns(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	//vpnDto.Id = ctx.FormValue("id")

	//vpnDto.Name = ctx.FormValue("name")

	err = vpnService.vpnDao.UpdateVpn(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if utils.IsEmpty(vpnDto.HostIds) {
		return result.SuccessData(vpnDto)
	}

	vpnHostsDto := vpn.VpnHostsDto{
		VpnId: vpnDto.VpnId,
	}
	vpnService.vpnHostsDao.DeleteVpnHosts(vpnHostsDto)

	for _, hostId := range strings.Split(vpnDto.HostIds, ",") {

		vpnHostsDto = vpn.VpnHostsDto{
			VpnId:     vpnDto.VpnId,
			VpnHostId: seq.Generator(),
			HostId:    hostId,
		}
		vpnService.vpnHostsDao.SaveVpnHosts(vpnHostsDto)
	}

	return result.SuccessData(vpnDto)

}

/**
删除 系统信息
*/
func (vpnService *VpnService) DeleteVpns(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = vpnService.vpnDao.DeleteVpn(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}

func (vpnService *VpnService) StartVpnf(ctx iris.Context) interface{} {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start vpn
	tmpVpnDto := vpn.VpnDto{
		VpnId: vpnDto.VpnId,
	}
	vpnDtos, err := vpnService.vpnDao.GetVpns(tmpVpnDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(vpnDtos) < 1 {
		return result.Error("未查询到数据")
	}

	resultDto, _ := shell.ExecStartVpn(vpnService.getVpnConfig(*vpnDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpVpnDto = vpn.VpnDto{
		VpnId: vpnDto.VpnId,
		State: vpn.Vpn_state_start,
	}

	err = vpnService.vpnDao.UpdateVpn(tmpVpnDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(vpnDto)

}

func (vpnService *VpnService) StopVpnf(ctx iris.Context) interface{} {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start vpn
	tmpVpnDto := vpn.VpnDto{
		VpnId: vpnDto.VpnId,
	}
	vpnDtos, err := vpnService.vpnDao.GetVpns(tmpVpnDto)

	if err != nil || len(vpnDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecStopVpn(vpnService.getVpnConfig(*vpnDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpVpnDto = vpn.VpnDto{
		VpnId: vpnDto.VpnId,
		State: vpn.Vpn_state_stop,
	}

	err = vpnService.vpnDao.UpdateVpn(tmpVpnDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(vpnDto)

}

func (vpnService *VpnService) RefreshVpnConfig(ctx iris.Context) interface{} {
	var (
		err    error
		vpnDto vpn.VpnDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}
	// start vpn
	tmpVpnDto := vpn.VpnDto{
		VpnId: vpnDto.VpnId,
	}
	vpnDtos, err := vpnService.vpnDao.GetVpns(tmpVpnDto)

	if err != nil || len(vpnDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecRefreshVpnConfig(vpnService.getVpnConfig(*vpnDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	return result.SuccessData(vpnDto)
}

// get vpn config
func (vpnService *VpnService) getVpnConfig(vpnDto vpn.VpnDto) vpn.SlaveVpnDataDto {
	var (
		vpnUserDao        vpnDao.VpnUserDao
	)

	vpnHostsDto := vpn.VpnHostsDto{
		VpnId: vpnDto.VpnId,
	}
	vpnHostsDtos, _ := vpnService.vpnHostsDao.GetVpnHostss(vpnHostsDto)
	vpnDto.VpnHosts = vpnHostsDtos

	// start vpn
	tmpVpnUserDto := vpn.VpnUserDto{
		
	}
	users, _ := vpnUserDao.GetVpnUsers(tmpVpnUserDto)


	return vpn.SlaveVpnDataDto{
		ServerIpUrl: config.G_AppConfig.ServerIpUrl,
		Vpn:         vpnDto,
		Users:      users,
	}

}
