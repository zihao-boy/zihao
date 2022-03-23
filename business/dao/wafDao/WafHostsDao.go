package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafHosts_count string = `
	select count(1) total
from waf_hosts t
					where t.status_cd = '0'
					$if WafId != '' then
					and t.waf_id = #WafId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif

	`
	query_wafHosts string = `
select t.*
from waf_hosts t
					where t.status_cd = '0'
					$if WafId != '' then
					and t.waf_id = #WafId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafHosts string = `
	insert into waf_hosts(waf_host_id, waf_id, host_id)
VALUES(#WafHostId#,#WafId#,#HostId#)
`

	update_wafHosts string = `
	update waf_hosts set
		$if HostId != '' then
		host_id = #HostId#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if WafHostId != '' then
		and waf_host_id = #WafHostId#
		$endif
		$if WafId != '' then
		and t.waf_id = #WafId#
		$endif
	`
	delete_wafHosts string = `
	update waf_hosts  set
                          status_cd = '1'
                          where status_cd = '0'
		$if WafHostId != '' then
		and waf_host_id = #WafHostId#
		$endif
		$if WafId != '' then
		and t.waf_id = #WafId#
		$endif
	`
)

type WafHostsDao struct {
}

/**
查询用户
*/
func (*WafHostsDao) GetWafHostsCount(wafHostsDto waf.WafHostsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafHosts_count, objectConvert.Struct2Map(wafHostsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafHostsDao) GetWafHostss(wafHostsDto waf.WafHostsDto) ([]*waf.WafHostsDto, error) {
	var wafHostsDtos []*waf.WafHostsDto
	sqlTemplate.SelectList(query_wafHosts, objectConvert.Struct2Map(wafHostsDto), func(db *gorm.DB) {
		db.Scan(&wafHostsDtos)
	}, false)

	return wafHostsDtos, nil
}

/**
保存服务sql
*/
func (*WafHostsDao) SaveWafHosts(wafHostsDto waf.WafHostsDto) error {
	return sqlTemplate.Insert(insert_wafHosts, objectConvert.Struct2Map(wafHostsDto), false)
}

/**
修改服务sql
*/
func (*WafHostsDao) UpdateWafHosts(wafHostsDto waf.WafHostsDto) error {
	return sqlTemplate.Update(update_wafHosts, objectConvert.Struct2Map(wafHostsDto), false)
}

/**
删除服务sql
*/
func (*WafHostsDao) DeleteWafHosts(wafHostsDto waf.WafHostsDto) error {
	return sqlTemplate.Delete(delete_wafHosts, objectConvert.Struct2Map(wafHostsDto), false)
}
