package vpnDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/vpn"
	"gorm.io/gorm"
)

const (
	query_vpnUser_count string = `
	select count(1) total
from vpn_users t
					where t.status_cd = '0'
					$if UserId != '' then
					and t.user_name = #UserId#
					$endif
					$if Username != '' then
					and t.username = #Username#
					$endif
					$if Password != '' then
					and t.password = #Password#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if Tel != '' then
					and t.tel = #Tel#
					$endif

	`
	query_vpnUser string = `
select t.*
from vpn_users t
					where t.status_cd = '0'
					$if UserId != '' then
					and t.user_name = #UserId#
					$endif
					$if Username != '' then
					and t.username = #Username#
					$endif
					$if Password != '' then
					and t.password = #Password#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if Tel != '' then
					and t.tel = #Tel#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_vpnUser string = `
	insert into vpn_users(user_id,username,password,tel,ip, login_time)
VALUES(#UserId#,#Username#,#Password#,#Tel#,#Ip#,#LoginTime#)
`

	update_vpnUser string = `
	update vpn_users set
					$if Username != '' then
					 username = #Username#,
					$endif
					$if Password != '' then
					 password = #Password#,
					$endif
					$if Ip != '' then
					ip = #Ip#,
					$endif
					$if Tel != '' then
					tel = #Tel#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if UserId != '' then
					and user_id = #UserId#
					$endif
	`
	delete_vpnUser string = `
	update vpn_users  set
                          status_cd = '1'
                          where status_cd = '0'
					$if UserId != '' then
					and user_id = #UserId#
					$endif
	`
)

type VpnUserDao struct {
}

/**
查询用户
*/
func (*VpnUserDao) GetVpnUserCount(vpnUserDto vpn.VpnUserDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_vpnUser_count, objectConvert.Struct2Map(vpnUserDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*VpnUserDao) GetVpnUsers(vpnUserDto vpn.VpnUserDto) ([]*vpn.VpnUserDto, error) {
	var vpnUserDtos []*vpn.VpnUserDto
	sqlTemplate.SelectList(query_vpnUser, objectConvert.Struct2Map(vpnUserDto), func(db *gorm.DB) {
		db.Scan(&vpnUserDtos)
	}, false)

	return vpnUserDtos, nil
}

/**
保存服务sql
*/
func (*VpnUserDao) SaveVpnUser(vpnUserDto vpn.VpnUserDto) error {
	return sqlTemplate.Insert(insert_vpnUser, objectConvert.Struct2Map(vpnUserDto), false)
}

/**
修改服务sql
*/
func (*VpnUserDao) UpdateVpnUser(vpnUserDto vpn.VpnUserDto) error {
	return sqlTemplate.Update(update_vpnUser, objectConvert.Struct2Map(vpnUserDto), false)
}

/**
删除服务sql
*/
func (*VpnUserDao) DeleteVpnUser(vpnUserDto vpn.VpnUserDto) error {
	return sqlTemplate.Delete(delete_vpnUser, objectConvert.Struct2Map(vpnUserDto), false)
}
