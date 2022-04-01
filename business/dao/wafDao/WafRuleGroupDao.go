package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafRuleGroup_count string = `
	select count(1) total
from waf_rule_group t
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
	query_wafRuleGroup string = `
select t.*
from waf_rule_group t
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

	insert_wafRuleGroup string = `
	insert into waf_rule_group(group_id, group_name, state)
VALUES(#GroupId#,#GroupName#,#State#)
`

	update_wafRuleGroup string = `
	update waf_rule_group set
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
	delete_wafRuleGroup string = `
	update waf_rule_group  set
                          status_cd = '1'
                          where status_cd = '0'
		$if GroupId != '' then
		and group_id = #GroupId#
		$endif
	`
)

type WafRuleGroupDao struct {
}

/**
查询用户
*/
func (*WafRuleGroupDao) GetWafRuleGroupCount(wafRuleGroupDto waf.WafRuleGroupDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafRuleGroup_count, objectConvert.Struct2Map(wafRuleGroupDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafRuleGroupDao) GetWafRuleGroups(wafRuleGroupDto waf.WafRuleGroupDto) ([]*waf.WafRuleGroupDto, error) {
	var wafRuleGroupDtos []*waf.WafRuleGroupDto
	sqlTemplate.SelectList(query_wafRuleGroup, objectConvert.Struct2Map(wafRuleGroupDto), func(db *gorm.DB) {
		db.Scan(&wafRuleGroupDtos)
	}, false)

	return wafRuleGroupDtos, nil
}

/**
保存服务sql
*/
func (*WafRuleGroupDao) SaveWafRuleGroup(wafRuleGroupDto waf.WafRuleGroupDto) error {
	return sqlTemplate.Insert(insert_wafRuleGroup, objectConvert.Struct2Map(wafRuleGroupDto), false)
}

/**
修改服务sql
*/
func (*WafRuleGroupDao) UpdateWafRuleGroup(wafRuleGroupDto waf.WafRuleGroupDto) error {
	return sqlTemplate.Update(update_wafRuleGroup, objectConvert.Struct2Map(wafRuleGroupDto), false)
}

/**
删除服务sql
*/
func (*WafRuleGroupDao) DeleteWafRuleGroup(wafRuleGroupDto waf.WafRuleGroupDto) error {
	return sqlTemplate.Delete(delete_wafRuleGroup, objectConvert.Struct2Map(wafRuleGroupDto), false)
}
