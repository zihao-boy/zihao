package firewallService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/firewallRuleDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/firewall"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type HostFirewallGroupService struct {
	firewallDao             firewallRuleDao.HostFirewallGroupDao
}

// get db link
// all db by this user
func (firewallService *HostFirewallGroupService) GetHostFirewallGroupAll(HostFirewallGroupDto firewall.HostFirewallGroupDto) ([]*firewall.HostFirewallGroupDto, error) {
	var (
		err          error
		HostFirewallGroupDtos []*firewall.HostFirewallGroupDto
	)

	HostFirewallGroupDtos, err = firewallService.firewallDao.GetHostFirewallGroups(HostFirewallGroupDto)
	if err != nil {
		return nil, err
	}

	return HostFirewallGroupDtos, nil

}

/**
查询 系统信息
*/
func (firewallService *HostFirewallGroupService) GetHostFirewallGroups(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		firewallDto  = firewall.HostFirewallGroupDto{}
		firewallDtos []*firewall.HostFirewallGroupDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	firewallDto.Row = row * page

	firewallDto.Page = (page - 1) * row

	firewallDto.HostId = ctx.URLParam("hostId")

	total, err = firewallService.firewallDao.GetHostFirewallGroupCount(firewallDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	firewallDtos, err = firewallService.firewallDao.GetHostFirewallGroups(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(firewallDtos, total, row)

}

/**
保存 系统信息
*/
func (firewallService *HostFirewallGroupService) SaveHostFirewallGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.HostFirewallGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}
	firewallDto.HfgId = seq.Generator()
	//HostFirewallGroupDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = firewallService.firewallDao.SaveHostFirewallGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(firewallDto)

}

/**
修改 系统信息
*/
func (firewallService *HostFirewallGroupService) UpdateHostFirewallGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.HostFirewallGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	//firewallDto.Id = ctx.FormValue("id")

	//firewallDto.Name = ctx.FormValue("name")

	err = firewallService.firewallDao.UpdateHostFirewallGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(firewallDto)

}

/**
删除 系统信息
*/
func (firewallService *HostFirewallGroupService) DeleteHostFirewallGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.HostFirewallGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = firewallService.firewallDao.DeleteHostFirewallGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(firewallDto)

}
