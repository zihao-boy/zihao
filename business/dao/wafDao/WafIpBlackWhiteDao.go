package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafIpBlackWhite_count string = `
	select count(1) total
from waf_ip_black_white t
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif

	`
	query_wafIpBlackWhite string = `
select t.*,wr.scope,wr.state,wr.seq,wrg.group_name,wrg.group_id
from waf_ip_black_white t
left join waf_rule wr on t.id = wr.obj_id and wr.obj_type = 'ip' and wr.status_cd = '0'
left join waf_rule_group wrg on wr.group_id = wrg.group_id and wrg.status_cd = '0'
					where t.status_cd = '0'
					$if Id != '' then
					and t.id = #Id#
					$endif
					$if TypeCd != '' then
					and t.type_cd = #TypeCd#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafIpBlackWhite string = `
	insert into waf_ip_black_white(id,type_cd,ip)
VALUES(#Id#,#TypeCd#,#Ip#)
`

	update_wafIpBlackWhite string = `
	update waf_ip_black_white set
					$if TypeCd != '' then
					 type_cd = #TypeCd#,
					$endif
					$if Ip != '' then
					 ip = #Ip#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
	delete_wafIpBlackWhite string = `
	update waf_ip_black_white  set
                          status_cd = '1'
                          where status_cd = '0'
					$if Id != '' then
					and id = #Id#
					$endif
	`
)

type WafIpBlackWhiteDao struct {
}

/**
查询用户
*/
func (*WafIpBlackWhiteDao) GetWafIpBlackWhiteCount(wafIpBlackWhiteDto waf.WafIpBlackWhiteDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafIpBlackWhite_count, objectConvert.Struct2Map(wafIpBlackWhiteDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafIpBlackWhiteDao) GetWafIpBlackWhites(wafIpBlackWhiteDto waf.WafIpBlackWhiteDto) ([]*waf.WafIpBlackWhiteDto, error) {
	var wafIpBlackWhiteDtos []*waf.WafIpBlackWhiteDto
	sqlTemplate.SelectList(query_wafIpBlackWhite, objectConvert.Struct2Map(wafIpBlackWhiteDto), func(db *gorm.DB) {
		db.Scan(&wafIpBlackWhiteDtos)
	}, false)

	return wafIpBlackWhiteDtos, nil
}

/**
保存服务sql
*/
func (*WafIpBlackWhiteDao) SaveWafIpBlackWhite(wafIpBlackWhiteDto waf.WafIpBlackWhiteDto) error {
	return sqlTemplate.Insert(insert_wafIpBlackWhite, objectConvert.Struct2Map(wafIpBlackWhiteDto), false)
}

/**
修改服务sql
*/
func (*WafIpBlackWhiteDao) UpdateWafIpBlackWhite(wafIpBlackWhiteDto waf.WafIpBlackWhiteDto) error {
	return sqlTemplate.Update(update_wafIpBlackWhite, objectConvert.Struct2Map(wafIpBlackWhiteDto), false)
}

/**
删除服务sql
*/
func (*WafIpBlackWhiteDao) DeleteWafIpBlackWhite(wafIpBlackWhiteDto waf.WafIpBlackWhiteDto) error {
	return sqlTemplate.Delete(delete_wafIpBlackWhite, objectConvert.Struct2Map(wafIpBlackWhiteDto), false)
}
