package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafAccurate_count string = `
					select count(1) total
					from waf_accurate t
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if Action != '' then
					and t.action = #Action#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if Include != '' then
					and t.include = #Include#
					$endif
					$if IncludeValue != '' then
					and t.include_value = #IncludeValue#
					$endif
	`
	query_wafAccurate string = `
select t.*,wr.scope,wr.state,wr.seq,wrg.group_name,wrg.group_id
from waf_accurate t
left join waf_rule wr on t.id = wr.obj_id and wr.obj_type = 'accurate' and wr.status_cd = '0'
left join waf_rule_group wrg on wr.group_id = wrg.group_id and wrg.status_cd = '0'
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if Action != '' then
					and t.action = #Action#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if Include != '' then
					and t.include = #Include#
					$endif
					$if IncludeValue != '' then
					and t.include_value = #IncludeValue#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafAccurate string = `
	insert into waf_accurate(id,action,type_cd,include,include_value)
VALUES(#Id#,#Action#,#TypeCd#,#Include#,#IncludeValue#)
`

	update_wafAccurate string = `
	update waf_accurate set
					$if Action != '' then
					action = #Action#,
					$endif
					$if TypeCd != '' then
					type_cd = #TypeCd#,
					$endif
					$if Include != '' then
					include = #Include#,
					$endif
					$if IncludeValue != '' then
					include_value = #IncludeValue#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
	delete_wafAccurate string = `
	update waf_accurate  set
                          status_cd = '1'
                          where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
)

type WafAccurateDao struct {
}

/**
查询用户
*/
func (*WafAccurateDao) GetWafAccurateCount(wafAccurateDto waf.WafAccurateDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafAccurate_count, objectConvert.Struct2Map(wafAccurateDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafAccurateDao) GetWafAccurates(wafAccurateDto waf.WafAccurateDto) ([]*waf.WafAccurateDto, error) {
	var wafAccurateDtos []*waf.WafAccurateDto
	sqlTemplate.SelectList(query_wafAccurate, objectConvert.Struct2Map(wafAccurateDto), func(db *gorm.DB) {
		db.Scan(&wafAccurateDtos)
	}, false)

	return wafAccurateDtos, nil
}

/**
保存服务sql
*/
func (*WafAccurateDao) SaveWafAccurate(wafAccurateDto waf.WafAccurateDto) error {
	return sqlTemplate.Insert(insert_wafAccurate, objectConvert.Struct2Map(wafAccurateDto), false)
}

/**
修改服务sql
*/
func (*WafAccurateDao) UpdateWafAccurate(wafAccurateDto waf.WafAccurateDto) error {
	return sqlTemplate.Update(update_wafAccurate, objectConvert.Struct2Map(wafAccurateDto), false)
}

/**
删除服务sql
*/
func (*WafAccurateDao) DeleteWafAccurate(wafAccurateDto waf.WafAccurateDto) error {
	return sqlTemplate.Delete(delete_wafAccurate, objectConvert.Struct2Map(wafAccurateDto), false)
}
