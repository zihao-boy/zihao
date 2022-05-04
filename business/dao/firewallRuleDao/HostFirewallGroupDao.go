package firewallRuleDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/firewall"
	"gorm.io/gorm"
)

const (
	query_hostFirewallGroup_count string = `
	select count(1) total
from host_firewall_group t
					where t.status_cd = '0'
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if HfgId != '' then
					and t.hfg_id = #HfgId#
					$endif

	`
	query_hostFirewallGroup string = `

select t.*,frg.*
from host_firewall_group t
left join firewall_rule_group frg on t.group_id = frg.group_id and frg.status_cd = '0'
					where t.status_cd = '0'
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if HfgId != '' then
					and t.hfg_id = #HfgId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_hostFirewallGroup string = `
	insert into host_firewall_group(group_id, host_id, hfg_id)
VALUES(#GroupId#,#HostId#,#HfgId#)
`

	update_hostFirewallGroup string = `
	update host_firewall_group set
					$if HostId != '' then
					 host_id = #HostId#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if HfgId != '' then
					and hfg_id = #HfgId#
					$endif
	`
	delete_hostFirewallGroup string = `
	update host_firewall_group  set
                          status_cd = '1'
                          where status_cd = '0'
					$if HfgId != '' then
					and hfg_id = #HfgId#
					$endif
	`
)

type HostFirewallGroupDao struct {
}

/**
查询用户
*/
func (*HostFirewallGroupDao) GetHostFirewallGroupCount(hostFirewallGroupDto firewall.HostFirewallGroupDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_hostFirewallGroup_count, objectConvert.Struct2Map(hostFirewallGroupDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*HostFirewallGroupDao) GetHostFirewallGroups(hostFirewallGroupDto firewall.HostFirewallGroupDto) ([]*firewall.HostFirewallGroupDto, error) {
	var hostFirewallGroupDtos []*firewall.HostFirewallGroupDto
	sqlTemplate.SelectList(query_hostFirewallGroup, objectConvert.Struct2Map(hostFirewallGroupDto), func(db *gorm.DB) {
		db.Scan(&hostFirewallGroupDtos)
	}, false)

	return hostFirewallGroupDtos, nil
}

/**
保存服务sql
*/
func (*HostFirewallGroupDao) SaveHostFirewallGroup(hostFirewallGroupDto firewall.HostFirewallGroupDto) error {
	return sqlTemplate.Insert(insert_hostFirewallGroup, objectConvert.Struct2Map(hostFirewallGroupDto), false)
}

/**
修改服务sql
*/
func (*HostFirewallGroupDao) UpdateHostFirewallGroup(hostFirewallGroupDto firewall.HostFirewallGroupDto) error {
	return sqlTemplate.Update(update_hostFirewallGroup, objectConvert.Struct2Map(hostFirewallGroupDto), false)
}

/**
删除服务sql
*/
func (*HostFirewallGroupDao) DeleteHostFirewallGroup(hostFirewallGroupDto firewall.HostFirewallGroupDto) error {
	return sqlTemplate.Delete(delete_hostFirewallGroup, objectConvert.Struct2Map(hostFirewallGroupDto), false)
}
