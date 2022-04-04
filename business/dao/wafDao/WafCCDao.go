package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafCC_count string = `
					select count(1) total
					from waf_cc t
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if Path != '' then
					and t.path = #Path#
					$endif
					$if VisitCount != '' then
					and t.visit_count = #VisitCount#
					$endif
					$if VisitSec != '' then
					and t.visit_sec = #VisitSec#
					$endif
					$if BlockSec != '' then
					and t.block_sec = #BlockSec#
					$endif
	`
	query_wafCC string = `
select t.*,wr.scope,wr.state,wr.seq,wrg.group_name,wrg.group_id
from waf_cc t
left join waf_rule wr on t.id = wr.obj_id and wr.obj_type = 'cc' and wr.status_cd = '0'
left join waf_rule_group wrg on wr.group_id = wrg.group_id and wrg.status_cd = '0'
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if Path != '' then
					and t.path = #Path#
					$endif
					$if VisitCount != '' then
					and t.visit_count = #VisitCount#
					$endif
					$if VisitSec != '' then
					and t.visit_sec = #VisitSec#
					$endif
					$if BlockSec != '' then
					and t.block_sec = #BlockSec#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafCC string = `
	insert into waf_cc(id,path,visit_count,visit_sec,block_sec)
VALUES(#Id#,#Path#,#VisitCount#,#VisitSec#,#BlockSec#)
`

	update_wafCC string = `
	update waf_cc set
					$if Path != '' then
					 path = #Path#,
					$endif
					$if VisitCount != '' then
					visit_count = #VisitCount#,
					$endif
					$if VisitSec != '' then
					visit_sec = #VisitSec#,
					$endif
					$if BlockSec != '' then
					block_sec = #BlockSec#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
	delete_wafCC string = `
	update waf_cc  set
                          status_cd = '1'
                          where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
)

type WafCCDao struct {
}

/**
查询用户
*/
func (*WafCCDao) GetWafCCCount(wafCCDto waf.WafCCDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafCC_count, objectConvert.Struct2Map(wafCCDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafCCDao) GetWafCCs(wafCCDto waf.WafCCDto) ([]*waf.WafCCDto, error) {
	var wafCCDtos []*waf.WafCCDto
	sqlTemplate.SelectList(query_wafCC, objectConvert.Struct2Map(wafCCDto), func(db *gorm.DB) {
		db.Scan(&wafCCDtos)
	}, false)

	return wafCCDtos, nil
}

/**
保存服务sql
*/
func (*WafCCDao) SaveWafCC(wafCCDto waf.WafCCDto) error {
	return sqlTemplate.Insert(insert_wafCC, objectConvert.Struct2Map(wafCCDto), false)
}

/**
修改服务sql
*/
func (*WafCCDao) UpdateWafCC(wafCCDto waf.WafCCDto) error {
	return sqlTemplate.Update(update_wafCC, objectConvert.Struct2Map(wafCCDto), false)
}

/**
删除服务sql
*/
func (*WafCCDao) DeleteWafCC(wafCCDto waf.WafCCDto) error {
	return sqlTemplate.Delete(delete_wafCC, objectConvert.Struct2Map(wafCCDto), false)
}
