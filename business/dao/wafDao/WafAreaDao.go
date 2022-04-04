package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafArea_count string = `
	select count(1) total
from waf_area t
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if AreaName != '' then
					and t.area_name = #AreaName#
					$endif

	`
	query_wafArea string = `
select t.*,wr.scope,wr.state,wr.seq,wrg.group_name,wrg.group_id
from waf_area t
left join waf_rule wr on t.id = wr.obj_id and wr.obj_type = 'area' and wr.status_cd = '0'
left join waf_rule_group wrg on wr.group_id = wrg.group_id and wrg.status_cd = '0'
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if AreaName != '' then
					and t.area_name = #AreaName#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafArea string = `
	insert into waf_area(id,type_cd,area_name)
VALUES(#Id#,#TypeCd#,#AreaName#)
`

	update_wafArea string = `
	update waf_area set
					$if TypeCd != '' then
					 type_cd = #TypeCd#,
					$endif
					$if AreaName != '' then
					 area_name = #AreaName#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
	delete_wafArea string = `
	update waf_area  set
                          status_cd = '1'
                          where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
)

type WafAreaDao struct {
}

/**
查询用户
*/
func (*WafAreaDao) GetWafAreaCount(wafAreaDto waf.WafAreaDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafArea_count, objectConvert.Struct2Map(wafAreaDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafAreaDao) GetWafAreas(wafAreaDto waf.WafAreaDto) ([]*waf.WafAreaDto, error) {
	var wafAreaDtos []*waf.WafAreaDto
	sqlTemplate.SelectList(query_wafArea, objectConvert.Struct2Map(wafAreaDto), func(db *gorm.DB) {
		db.Scan(&wafAreaDtos)
	}, false)

	return wafAreaDtos, nil
}

/**
保存服务sql
*/
func (*WafAreaDao) SaveWafArea(wafAreaDto waf.WafAreaDto) error {
	return sqlTemplate.Insert(insert_wafArea, objectConvert.Struct2Map(wafAreaDto), false)
}

/**
修改服务sql
*/
func (*WafAreaDao) UpdateWafArea(wafAreaDto waf.WafAreaDto) error {
	return sqlTemplate.Update(update_wafArea, objectConvert.Struct2Map(wafAreaDto), false)
}

/**
删除服务sql
*/
func (*WafAreaDao) DeleteWafArea(wafAreaDto waf.WafAreaDto) error {
	return sqlTemplate.Delete(delete_wafArea, objectConvert.Struct2Map(wafAreaDto), false)
}
