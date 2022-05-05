package firewallService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/firewallRuleDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/entity/dto/firewall"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type FirewallRuleService struct {
	firewallDao             firewallRuleDao.FirewallRuleDao
}

// get db link
// all db by this user
func (firewallService *FirewallRuleService) GetFirewallRuleAll(FirewallRuleDto firewall.FirewallRuleDto) ([]*firewall.FirewallRuleDto, error) {
	var (
		err          error
		FirewallRuleDtos []*firewall.FirewallRuleDto
	)

	FirewallRuleDtos, err = firewallService.firewallDao.GetFirewallRules(FirewallRuleDto)
	if err != nil {
		return nil, err
	}

	return FirewallRuleDtos, nil

}

/**
查询 系统信息
*/
func (firewallService *FirewallRuleService) GetFirewallRules(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		firewallDto  = firewall.FirewallRuleDto{}
		firewallDtos []*firewall.FirewallRuleDto
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
	firewallDto.GroupId = ctx.URLParam("groupId")
	firewallDto.Inout = ctx.URLParam("inout")

	total, err = firewallService.firewallDao.GetFirewallRuleCount(firewallDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	firewallDtos, err = firewallService.firewallDao.GetFirewallRules(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}



	return result.SuccessData(firewallDtos, total, row)

}

/**
保存 系统信息
*/
func (firewallService *FirewallRuleService) SaveFirewallRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}
	firewallDto.RuleId = seq.Generator()
	//FirewallRuleDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = firewallService.firewallDao.SaveFirewallRule(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	go shell.ExecFirewallRule()

	return result.SuccessData(firewallDto)

}

/**
修改 系统信息
*/
func (firewallService *FirewallRuleService) UpdateFirewallRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	//firewallDto.Id = ctx.FormValue("id")

	//firewallDto.Name = ctx.FormValue("name")

	err = firewallService.firewallDao.UpdateFirewallRule(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	go shell.ExecFirewallRule()

	return result.SuccessData(firewallDto)

}

/**
删除 系统信息
*/
func (firewallService *FirewallRuleService) DeleteFirewallRules(ctx iris.Context) result.ResultDto {
	var (
		err    error
		firewallDto firewall.FirewallRuleDto
	)
	if err = ctx.ReadJSON(&firewallDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = firewallService.firewallDao.DeleteFirewallRule(firewallDto)
	if err != nil {
		return result.Error(err.Error())
	}

	go shell.ExecFirewallRule()

	return result.SuccessData(firewallDto)

}

func (firewallService *FirewallRuleService) GetFirewallRulesByHost(ctx iris.Context) interface{} {

	// query rule
	hostFirewallGroupDto :=firewall.HostFirewallGroupDto{
		HostId: ctx.URLParam("hostId"),
	}
	firewallRuleDtos,_ :=firewallService.firewallDao.GetFirewallRulesByHost(hostFirewallGroupDto)

	if firewallRuleDtos == nil{
		firewallRuleDtos = []*firewall.FirewallRuleDto{}
	}

	return result.SuccessData(firewallRuleDtos)

}

