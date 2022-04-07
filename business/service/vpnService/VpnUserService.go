package vpnService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/vpnDao"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"strconv"
)

type VpnUserService struct {
	vpnDao             vpnDao.VpnUserDao
}

// get db link
// all db by this user
func (vpnService *VpnUserService) GetVpnUserAll(VpnUserDto vpn.VpnUserDto) ([]*vpn.VpnUserDto, error) {
	var (
		err          error
		VpnUserDtos []*vpn.VpnUserDto
	)

	VpnUserDtos, err = vpnService.vpnDao.GetVpnUsers(VpnUserDto)
	if err != nil {
		return nil, err
	}

	return VpnUserDtos, nil

}

/**
查询 系统信息
*/
func (vpnService *VpnUserService) GetVpnUsers(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		vpnDto  = vpn.VpnUserDto{}
		vpnDtos []*vpn.VpnUserDto
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

	total, err = vpnService.vpnDao.GetVpnUserCount(vpnDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	vpnDtos, err = vpnService.vpnDao.GetVpnUsers(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDtos, total, row)

}

/**
保存 系统信息
*/
func (vpnService *VpnUserService) SaveVpnUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnUserDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}
	vpnDto.UserId = seq.Generator()
	//VpnUserDto.Path = filepath.Join(curDest, fileHeader.Filename)
	vpnDto.LoginTime = date.GetNowTime()

	err = vpnService.vpnDao.SaveVpnUser(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}

/**
修改 系统信息
*/
func (vpnService *VpnUserService) UpdateVpnUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnUserDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	//vpnDto.Id = ctx.FormValue("id")

	//vpnDto.Name = ctx.FormValue("name")

	err = vpnService.vpnDao.UpdateVpnUser(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(vpnDto)

}

/**
删除 系统信息
*/
func (vpnService *VpnUserService) DeleteVpnUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		vpnDto vpn.VpnUserDto
	)
	if err = ctx.ReadJSON(&vpnDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = vpnService.vpnDao.DeleteVpnUser(vpnDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(vpnDto)

}
