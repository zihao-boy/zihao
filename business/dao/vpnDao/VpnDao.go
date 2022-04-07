package vpnDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"gorm.io/gorm"
)

const (
	query_vpn_count string = `
	select count(1) total
from vpn t
					where t.status_cd = '0'
					$if Tun != '' then
					and t.tun = #Tun#
					$endif
					$if TunName != '' then
					and t.tun_name = #TunName#
					$endif
					$if Dns != '' then
					and t.dns = #Dns#
					$endif
					$if Protocol != '' then
					and t.protocol = #Protocol#
					$endif
					$if VpnPort != '' then
					and t.vpn_port = #VpnPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if VpnId != '' then
					and t.vpn_id = #VpnId#
					$endif

	`
	query_vpn string = `
select t.*
from vpn t
					where t.status_cd = '0'
					$if Tun != '' then
					and t.tun = #Tun#
					$endif
					$if TunName != '' then
					and t.tun_name = #TunName#
					$endif
					$if Dns != '' then
					and t.dns = #Dns#
					$endif
					$if Protocol != '' then
					and t.protocol = #Protocol#
					$endif
					$if VpnPort != '' then
					and t.vpn_port = #VpnPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if VpnId != '' then
					and t.vpn_id = #VpnId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_vpn string = `
	insert into vpn(vpn_id, vpn_port, tun,tun_name,dns,protocol, state)
VALUES(#VpnId#,#VpnPort#,#Tun#,#TunName#,#Dns#,#Protocol#,#State#)
`

	update_vpn string = `
	update vpn set
		$if VpnPort != '' then
		vpn_port = #VpnPort#,
		$endif
		$if Tun != '' then
		tun = #Tun#,
		$endif
		$if TunName != '' then
		tun_name = #TunName#,
		$endif
		$if Dns != '' then
		dns = #Dns#,
		$endif
		$if Protocol != '' then
		protocol = #Protocol#,
		$endif
        $if State != '' then
		state = #State#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if VpnId != '' then
		and vpn_id = #VpnId#
		$endif
	`
	delete_vpn string = `
	update vpn  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if VpnId != '' then
						  and vpn_id = #VpnId#
						  $endif
	`
)

type VpnDao struct {
}

/**
查询用户
*/
func (*VpnDao) GetVpnCount(vpnDto vpn.VpnDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_vpn_count, objectConvert.Struct2Map(vpnDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*VpnDao) GetVpns(vpnDto vpn.VpnDto) ([]*vpn.VpnDto, error) {
	var vpnDtos []*vpn.VpnDto
	sqlTemplate.SelectList(query_vpn, objectConvert.Struct2Map(vpnDto), func(db *gorm.DB) {
		db.Scan(&vpnDtos)
	}, false)

	return vpnDtos, nil
}

/**
保存服务sql
*/
func (*VpnDao) SaveVpn(vpnDto vpn.VpnDto) error {
	return sqlTemplate.Insert(insert_vpn, objectConvert.Struct2Map(vpnDto), false)
}

/**
修改服务sql
*/
func (*VpnDao) UpdateVpn(vpnDto vpn.VpnDto) error {
	return sqlTemplate.Update(update_vpn, objectConvert.Struct2Map(vpnDto), false)
}

/**
删除服务sql
*/
func (*VpnDao) DeleteVpn(vpnDto vpn.VpnDto) error {
	return sqlTemplate.Delete(delete_vpn, objectConvert.Struct2Map(vpnDto), false)
}
