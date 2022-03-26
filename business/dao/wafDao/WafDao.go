package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_waf_count string = `
	select count(1) total
from waf t
					where t.status_cd = '0'
					$if WafName != '' then
					and t.waf_name = #WafName#
					$endif
					$if Port != '' then
					and t.port = #Port#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if WafId != '' then
					and t.id = #WafId#
					$endif

	`
	query_waf string = `
select t.*
from waf t
					where t.status_cd = '0'
					$if WafName != '' then
					and t.waf_name = #WafName#
					$endif
					$if Port != '' then
					and t.port = #Port#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if WafId != '' then
					and t.waf_id = #WafId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_waf string = `
	insert into waf(waf_id, waf_name, port,https_port, state)
VALUES(#WafId#,#WafName#,#Port#,#HttpsPort#,#State#)
`

	update_waf string = `
	update waf set
		$if WafName != '' then
		waf_name = #WafName#,
		$endif
		$if Port != '' then
		port = #Port#,
		$endif
		$if HttpsPort != '' then
		https_port = #HttpsPort#,
		$endif
        $if State != '' then
		state = #State#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if WafId != '' then
		and waf_id = #WafId#
		$endif
	`
	delete_waf string = `
	update waf  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if WafId != '' then
						  and waf_id = #WafId#
						  $endif
	`
)

type WafDao struct {
}

/**
查询用户
*/
func (*WafDao) GetWafCount(wafDto waf.WafDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_waf_count, objectConvert.Struct2Map(wafDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafDao) GetWafs(wafDto waf.WafDto) ([]*waf.WafDto, error) {
	var wafDtos []*waf.WafDto
	sqlTemplate.SelectList(query_waf, objectConvert.Struct2Map(wafDto), func(db *gorm.DB) {
		db.Scan(&wafDtos)
	}, false)

	return wafDtos, nil
}

/**
保存服务sql
*/
func (*WafDao) SaveWaf(wafDto waf.WafDto) error {
	return sqlTemplate.Insert(insert_waf, objectConvert.Struct2Map(wafDto), false)
}

/**
修改服务sql
*/
func (*WafDao) UpdateWaf(wafDto waf.WafDto) error {
	return sqlTemplate.Update(update_waf, objectConvert.Struct2Map(wafDto), false)
}

/**
删除服务sql
*/
func (*WafDao) DeleteWaf(wafDto waf.WafDto) error {
	return sqlTemplate.Delete(delete_waf, objectConvert.Struct2Map(wafDto), false)
}
