package innerNetDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"gorm.io/gorm"
)

const (
	query_innerNetHosts_count string = `
	select count(1) total
from inner_net_hosts t
					where t.status_cd = '0'
					$if InnerNetId != '' then
					and t.inner_net_id = #InnerNetId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif

	`
	query_innerNetHosts string = `
select t.*
from inner_net_hosts t
					where t.status_cd = '0'
					$if InnerNetId != '' then
					and t.inner_net_id = #InnerNetId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_innerNetHosts string = `
	insert into inner_net_hosts(inner_net_host_id, inner_net_id, host_id)
VALUES(#InnerNetHostId#,#InnerNetId#,#HostId#)
`

	update_innerNetHosts string = `
	update inner_net_hosts set
		$if HostId != '' then
		host_id = #HostId#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if InnerNetHostId != '' then
		and inner_net_host_id = #InnerNetHostId#
		$endif
		$if InnerNetId != '' then
		and inner_net_id = #InnerNetId#
		$endif
	`
	delete_innerNetHosts string = `
	update inner_net_hosts  set
                          status_cd = '1'
                          where status_cd = '0'
		$if InnerNetHostId != '' then
		and inner_net_host_id = #InnerNetHostId#
		$endif
		$if InnerNetId != '' then
		and inner_net_id = #InnerNetId#
		$endif
	`
)

type InnerNetHostsDao struct {
}

/**
查询用户
*/
func (*InnerNetHostsDao) GetInnerNetHostsCount(innerNetHostsDto innerNet.InnerNetHostsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_innerNetHosts_count, objectConvert.Struct2Map(innerNetHostsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InnerNetHostsDao) GetInnerNetHostss(innerNetHostsDto innerNet.InnerNetHostsDto) ([]*innerNet.InnerNetHostsDto, error) {
	var innerNetHostsDtos []*innerNet.InnerNetHostsDto
	sqlTemplate.SelectList(query_innerNetHosts, objectConvert.Struct2Map(innerNetHostsDto), func(db *gorm.DB) {
		db.Scan(&innerNetHostsDtos)
	}, false)

	return innerNetHostsDtos, nil
}

/**
保存服务sql
*/
func (*InnerNetHostsDao) SaveInnerNetHosts(innerNetHostsDto innerNet.InnerNetHostsDto) error {
	return sqlTemplate.Insert(insert_innerNetHosts, objectConvert.Struct2Map(innerNetHostsDto), false)
}

/**
修改服务sql
*/
func (*InnerNetHostsDao) UpdateInnerNetHosts(innerNetHostsDto innerNet.InnerNetHostsDto) error {
	return sqlTemplate.Update(update_innerNetHosts, objectConvert.Struct2Map(innerNetHostsDto), false)
}

/**
删除服务sql
*/
func (*InnerNetHostsDao) DeleteInnerNetHosts(innerNetHostsDto innerNet.InnerNetHostsDto) error {
	return sqlTemplate.Delete(delete_innerNetHosts, objectConvert.Struct2Map(innerNetHostsDto), false)
}
