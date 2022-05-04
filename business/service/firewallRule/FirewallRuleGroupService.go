package firewallService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/firewallRuleDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/firewall"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type FirewallRuleGroupService struct {
	firewallDao             firewallRuleDao.FirewallRuleGroupDao
}

// get db link
// all db by this user
func (firewallService *FirewallRuleGroupService) GetFirewallRuleGroupAll(FirewallRuleGroupDto firewall.FirewallRuleGroupDto) ([]*firewall.FirewallRuleGroupDto, error) {
	var (
		err          error
		FirewallRuleGroupDtos []*firewall.FirewallRuleGroupDto
	)

	FirewallRuleGroupDtos, err = firewallService.firewallDao.GetFirewallRuleGroups(FirewallRuleGroupDto)
	if err != nil {
		return nil, err
	}

	return FirewallRuleGroupDtos, nil

}

/**
查询 系统信息
*/
func (firewallService *FirewallRuleGroupService) GetFirewallRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		firewallDto  = firewall.FirewallRuleGroupDto{}
		firewallDtos []*firewall.FirewallRuleGroupDto
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

	total, err = firewallService.firewallDao.GetFirewallRuleGroupCount(firewallDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	firewallDtos, err = firewallService.firewallDao.GetFirewallRuleGroups(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(firewallDtos, total, row)

}

/**
保存 系统信息
*/
func (firewallService *FirewallRuleGroupService) SaveFirewallRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}
	firewallDto.GroupId = seq.Generator()
	firewallDto.State = firewall.Firewall_Rule_Group_State_T
	//FirewallRuleGroupDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = firewallService.firewallDao.SaveFirewallRuleGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(firewallDto)

}

/**
修改 系统信息
*/
func (firewallService *FirewallRuleGroupService) UpdateFirewallRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	//firewallDto.Id = ctx.FormValue("id")

	//firewallDto.Name = ctx.FormValue("name")

	err = firewallService.firewallDao.UpdateFirewallRuleGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(firewallDto)

}

/**
删除 系统信息
*/
func (firewallService *FirewallRuleGroupService) DeleteFirewallRuleGroups(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = firewallService.firewallDao.DeleteFirewallRuleGroup(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(firewallDto)

}

func (firewallService *FirewallRuleGroupService) StartFirewallRuleGroup(ctx iris.Context) interface{} {
	var (
		err    error
		firewallDto firewall.FirewallRuleGroupDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	tmpFirewallRuleGroupDto := firewall.FirewallRuleGroupDto{
		State:firewall.Firewall_Rule_Group_State_F,
	}
	err = firewallService.firewallDao.UpdateFirewallRuleGroup(tmpFirewallRuleGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}
	tmpFirewallRuleGroupDto = firewall.FirewallRuleGroupDto{
		State:firewall.Firewall_Rule_Group_State_T,
		GroupId: firewallDto.GroupId,
	}
	err = firewallService.firewallDao.UpdateFirewallRuleGroup(tmpFirewallRuleGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(firewallDto)
}
