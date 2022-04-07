package vpnDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"gorm.io/gorm"
)

const (
	query_vpnHosts_count string = `
	select count(1) total
from vpn_hosts t
					where t.status_cd = '0'
					$if VpnId != '' then
					and t.vpn_id = #VpnId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif

	`
	query_vpnHosts string = `
select t.*
from vpn_hosts t
					where t.status_cd = '0'
					$if VpnId != '' then
					and t.vpn_id = #VpnId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_vpnHosts string = `
	insert into vpn_hosts(vpn_host_id, vpn_id, host_id)
VALUES(#VpnHostId#,#VpnId#,#HostId#)
`

	update_vpnHosts string = `
	update vpn_hosts set
		$if HostId != '' then
		host_id = #HostId#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if VpnHostId != '' then
		and vpn_host_id = #VpnHostId#
		$endif
		$if VpnId != '' then
		and vpn_id = #VpnId#
		$endif
	`
	delete_vpnHosts string = `
	update vpn_hosts  set
                          status_cd = '1'
                          where status_cd = '0'
		$if VpnHostId != '' then
		and vpn_host_id = #VpnHostId#
		$endif
		$if VpnId != '' then
		and vpn_id = #VpnId#
		$endif
	`
)

type VpnHostsDao struct {
}

/**
查询用户
*/
func (*VpnHostsDao) GetVpnHostsCount(vpnHostsDto vpn.VpnHostsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_vpnHosts_count, objectConvert.Struct2Map(vpnHostsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*VpnHostsDao) GetVpnHostss(vpnHostsDto vpn.VpnHostsDto) ([]*vpn.VpnHostsDto, error) {
	var vpnHostsDtos []*vpn.VpnHostsDto
	sqlTemplate.SelectList(query_vpnHosts, objectConvert.Struct2Map(vpnHostsDto), func(db *gorm.DB) {
		db.Scan(&vpnHostsDtos)
	}, false)

	return vpnHostsDtos, nil
}

/**
保存服务sql
*/
func (*VpnHostsDao) SaveVpnHosts(vpnHostsDto vpn.VpnHostsDto) error {
	return sqlTemplate.Insert(insert_vpnHosts, objectConvert.Struct2Map(vpnHostsDto), false)
}

/**
修改服务sql
*/
func (*VpnHostsDao) UpdateVpnHosts(vpnHostsDto vpn.VpnHostsDto) error {
	return sqlTemplate.Update(update_vpnHosts, objectConvert.Struct2Map(vpnHostsDto), false)
}

/**
删除服务sql
*/
func (*VpnHostsDao) DeleteVpnHosts(vpnHostsDto vpn.VpnHostsDto) error {
	return sqlTemplate.Delete(delete_vpnHosts, objectConvert.Struct2Map(vpnHostsDto), false)
}
