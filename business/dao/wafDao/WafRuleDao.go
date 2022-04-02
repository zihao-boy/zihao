package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafRule_count string = `
	select count(1) total
from waf_rule t
					where t.status_cd = '0'
					$if RuleId != '' then
					and t.rule_id = #RuleId#
					$endif
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if RuleName != '' then
					and t.rule_name = #RuleName#
					$endif
					$if Scope != '' then
					and t.scope = #Scope#
					$endif
					$if ObjType != '' then
					and t.obj_type = #ObjType#
					$endif
					$if ObjId != '' then
					and t.obj_id = #ObjId#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif

	`
	query_wafRule string = `
select t.*
from waf_rule t
					where t.status_cd = '0'
					$if RuleId != '' then
					and t.rule_id = #RuleId#
					$endif
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if RuleName != '' then
					and t.rule_name = #RuleName#
					$endif
					$if Scope != '' then
					and t.scope = #Scope#
					$endif
					$if ObjType != '' then
					and t.obj_type = #ObjType#
					$endif
					$if ObjId != '' then
					and t.obj_id = #ObjId#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafRule string = `
	insert into waf_rule(rule_id,group_id,rule_name, scope,obj_id,obj_type,seq, state)
VALUES(#RuleId#,#GroupId#,#RuleName#,#Scope#,#ObjId#,#ObjType#,#Seq#,#State#)
`

	update_wafRule string = `
	update waf_rule set
					$if GroupId != '' then
					group_id = #GroupId#,
					$endif
					$if RuleName != '' then
					rule_name = #RuleName#,
					$endif
					$if Scope != '' then
					scope = #Scope#,
					$endif
					$if ObjType != '' then
					obj_type = #ObjType#,
					$endif
					$if ObjId != '' then
					obj_id = #ObjId#,
					$endif
					$if State != '' then
					state = #State#,
					$endif
					$if Seq != '' then
					seq = #Seq#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if RuleId != '' then
					and rule_id = #RuleId#
					$endif
	`
	delete_wafRule string = `
	update waf_rule  set
                          status_cd = '1'
                          where status_cd = '0'
					$if RuleId != '' then
					and rule_id = #RuleId#
					$endif
	`
)

type WafRuleDao struct {
}

/**
查询用户
*/
func (*WafRuleDao) GetWafRuleCount(wafRuleDto waf.WafRuleDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafRule_count, objectConvert.Struct2Map(wafRuleDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafRuleDao) GetWafRules(wafRuleDto waf.WafRuleDto) ([]*waf.WafRuleDto, error) {
	var wafRuleDtos []*waf.WafRuleDto
	sqlTemplate.SelectList(query_wafRule, objectConvert.Struct2Map(wafRuleDto), func(db *gorm.DB) {
		db.Scan(&wafRuleDtos)
	}, false)

	return wafRuleDtos, nil
}

/**
保存服务sql
*/
func (*WafRuleDao) SaveWafRule(wafRuleDto waf.WafRuleDto) error {
	return sqlTemplate.Insert(insert_wafRule, objectConvert.Struct2Map(wafRuleDto), false)
}

/**
修改服务sql
*/
func (*WafRuleDao) UpdateWafRule(wafRuleDto waf.WafRuleDto) error {
	return sqlTemplate.Update(update_wafRule, objectConvert.Struct2Map(wafRuleDto), false)
}

/**
删除服务sql
*/
func (*WafRuleDao) DeleteWafRule(wafRuleDto waf.WafRuleDto) error {
	return sqlTemplate.Delete(delete_wafRule, objectConvert.Struct2Map(wafRuleDto), false)
}
