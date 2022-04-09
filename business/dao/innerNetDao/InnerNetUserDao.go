package innerNetDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"gorm.io/gorm"
)

const (
	query_innerNetUser_count string = `
	select count(1) total
from inner_net_users t
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
	query_innerNetUser string = `
select t.*
from inner_net_users t
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

	insert_innerNetUser string = `
	insert into inner_net_users(user_id,username,password,tel,ip, login_time)
VALUES(#UserId#,#Username#,#Password#,#Tel#,#Ip#,#LoginTime#)
`

	update_innerNetUser string = `
	update inner_net_users set
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
	delete_innerNetUser string = `
	update inner_net_users  set
                          status_cd = '1'
                          where status_cd = '0'
					$if UserId != '' then
					and user_id = #UserId#
					$endif
	`
)

type InnerNetUserDao struct {
}

/**
查询用户
*/
func (*InnerNetUserDao) GetInnerNetUserCount(innerNetUserDto innerNet.InnerNetUserDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_innerNetUser_count, objectConvert.Struct2Map(innerNetUserDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InnerNetUserDao) GetInnerNetUsers(innerNetUserDto innerNet.InnerNetUserDto) ([]*innerNet.InnerNetUserDto, error) {
	var innerNetUserDtos []*innerNet.InnerNetUserDto
	sqlTemplate.SelectList(query_innerNetUser, objectConvert.Struct2Map(innerNetUserDto), func(db *gorm.DB) {
		db.Scan(&innerNetUserDtos)
	}, false)

	return innerNetUserDtos, nil
}

/**
保存服务sql
*/
func (*InnerNetUserDao) SaveInnerNetUser(innerNetUserDto innerNet.InnerNetUserDto) error {
	return sqlTemplate.Insert(insert_innerNetUser, objectConvert.Struct2Map(innerNetUserDto), false)
}

/**
修改服务sql
*/
func (*InnerNetUserDao) UpdateInnerNetUser(innerNetUserDto innerNet.InnerNetUserDto) error {
	return sqlTemplate.Update(update_innerNetUser, objectConvert.Struct2Map(innerNetUserDto), false)
}

/**
删除服务sql
*/
func (*InnerNetUserDao) DeleteInnerNetUser(innerNetUserDto innerNet.InnerNetUserDto) error {
	return sqlTemplate.Delete(delete_innerNetUser, objectConvert.Struct2Map(innerNetUserDto), false)
}
