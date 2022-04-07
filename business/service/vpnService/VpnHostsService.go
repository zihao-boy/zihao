package vpnService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/vpnDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"strconv"
)

type VpnHostsService struct {
	vpnDao vpnDao.VpnHostsDao
}

// get db link
// all db by this user
func (vpnService *VpnHostsService) GetVpnHostsAll(VpnHostsDto vpn.VpnHostsDto) ([]*vpn.VpnHostsDto, error) {
	var (
		err        error
		VpnHostsDtos []*vpn.VpnHostsDto
	)

	VpnHostsDtos, err = vpnService.vpnDao.GetVpnHostss(VpnHostsDto)
	if err != nil {
		return nil, err
	}

	return VpnHostsDtos, nil

}

/**
查询 系统信息
*/
func (vpnService *VpnHostsService) GetVpnHostss(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		vpnDto  = vpn.VpnHostsDto{}
		vpnDtos []*vpn.VpnHostsDto
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

	total, err = vpnService.vpnDao.GetVpnHostsCount(vpnDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	vpnDtos, err = vpnService.vpnDao.GetVpnHostss(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDtos, total, row)

}

/**
保存 系统信息
*/
func (vpnService *VpnHostsService) SaveVpnHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		vpnDto vpn.VpnHostsDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}
	vpnDto.VpnHostId = seq.Generator()
	//VpnHostsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = vpnService.vpnDao.SaveVpnHosts(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}

/**
修改 系统信息
*/
func (vpnService *VpnHostsService) UpdateVpnHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		vpnDto vpn.VpnHostsDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	//vpnDto.Id = ctx.FormValue("id")

	//vpnDto.Name = ctx.FormValue("name")

	err = vpnService.vpnDao.UpdateVpnHosts(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}

/**
删除 系统信息
*/
func (vpnService *VpnHostsService) DeleteVpnHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		vpnDto vpn.VpnHostsDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = vpnService.vpnDao.DeleteVpnHosts(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}
