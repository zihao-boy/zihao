package firewallRuleDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/firewall"
	"gorm.io/gorm"
)

const (
	query_firewallRuleGroup_count string = `
	select count(1) total
from firewall_rule_group t
					where t.status_cd = '0'
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if GroupName != '' then
					and t.group_name = #GroupName#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif

	`
	query_firewallRuleGroup string = `
select t.*
from firewall_rule_group t
					where t.status_cd = '0'
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if GroupName != '' then
					and t.group_name = #GroupName#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_firewallRuleGroup string = `
	insert into firewall_rule_group(group_id, group_name, state)
VALUES(#GroupId#,#GroupName#,#State#)
`

	update_firewallRuleGroup string = `
	update firewall_rule_group set
					$if GroupName != '' then
					 group_name = #GroupName#,
					$endif
					$if State != '' then
					 state = #State#,
					$endif
		status_cd = '0'
		where status_cd = '0'
		$if GroupId != '' then
		and group_id = #GroupId#
		$endif
	`
	delete_firewallRuleGroup string = `
	update firewall_rule_group  set
                          status_cd = '1'
                          where status_cd = '0'
		$if GroupId != '' then
		and group_id = #GroupId#
		$endif
	`
)

type FirewallRuleGroupDao struct {
}

/**
查询用户
*/
func (*FirewallRuleGroupDao) GetFirewallRuleGroupCount(firewallRuleGroupDto firewall.FirewallRuleGroupDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_firewallRuleGroup_count, objectConvert.Struct2Map(firewallRuleGroupDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*FirewallRuleGroupDao) GetFirewallRuleGroups(firewallRuleGroupDto firewall.FirewallRuleGroupDto) ([]*firewall.FirewallRuleGroupDto, error) {
	var firewallRuleGroupDtos []*firewall.FirewallRuleGroupDto
	sqlTemplate.SelectList(query_firewallRuleGroup, objectConvert.Struct2Map(firewallRuleGroupDto), func(db *gorm.DB) {
		db.Scan(&firewallRuleGroupDtos)
	}, false)

	return firewallRuleGroupDtos, nil
}

/**
保存服务sql
*/
func (*FirewallRuleGroupDao) SaveFirewallRuleGroup(firewallRuleGroupDto firewall.FirewallRuleGroupDto) error {
	return sqlTemplate.Insert(insert_firewallRuleGroup, objectConvert.Struct2Map(firewallRuleGroupDto), false)
}

/**
修改服务sql
*/
func (*FirewallRuleGroupDao) UpdateFirewallRuleGroup(firewallRuleGroupDto firewall.FirewallRuleGroupDto) error {
	return sqlTemplate.Update(update_firewallRuleGroup, objectConvert.Struct2Map(firewallRuleGroupDto), false)
}

/**
删除服务sql
*/
func (*FirewallRuleGroupDao) DeleteFirewallRuleGroup(firewallRuleGroupDto firewall.FirewallRuleGroupDto) error {
	return sqlTemplate.Delete(delete_firewallRuleGroup, objectConvert.Struct2Map(firewallRuleGroupDto), false)
}
