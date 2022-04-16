package innerNetDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"gorm.io/gorm"
)

const (
	query_innerNetPrivilege_count string = `
	select count(1) total
from inner_net_privilege t
					where t.status_cd = '0'
					$if PId != '' then
					and t.p_id = #PId#
					$endif
					$if SrcUserId != '' then
					and t.src_user_id = #SrcUserId#
					$endif
					$if TargetUserId != '' then
					and t.target_user_id = #TargetUserId#
					$endif

	`
	query_innerNetPrivilege string = `
select t.*
from inner_net_privilege t
					where t.status_cd = '0'
					$if PId != '' then
					and t.p_id = #PId#
					$endif
					$if SrcUserId != '' then
					and t.src_user_id = #SrcUserId#
					$endif
					$if TargetUserId != '' then
					and t.target_user_id = #TargetUserId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_innerNetPrivilege string = `
	insert into inner_net_privilege(p_id,src_user_id,target_user_id,target_port)
VALUES(#PId#,#SrcUserId#,#TargetUserId#,#TargetPort#)
`

	update_innerNetPrivilege string = `
	update inner_net_privilege set
					$if SrcUserId != '' then
					 src_user_id = #SrcUserId#,
					$endif
					$if TargetUserId != '' then
					target_user_id = #TargetUserId#,
					$endif
					$if TargetPort != '' then
					target_port = #TargetPort#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if PId != '' then
					and p_id = #PId#
					$endif
	`
	delete_innerNetPrivilege string = `
	update inner_net_privilege  set
                          status_cd = '1'
                          where status_cd = '0'
					$if PId != '' then
					and p_id = #PId#
					$endif
	`
)

type InnerNetPrivilegeDao struct {
}

/**
查询用户
*/
func (*InnerNetPrivilegeDao) GetInnerNetPrivilegeCount(innerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_innerNetPrivilege_count, objectConvert.Struct2Map(innerNetPrivilegeDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InnerNetPrivilegeDao) GetInnerNetPrivileges(innerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) ([]*innerNet.InnerNetPrivilegeDto, error) {
	var innerNetPrivilegeDtos []*innerNet.InnerNetPrivilegeDto
	sqlTemplate.SelectList(query_innerNetPrivilege, objectConvert.Struct2Map(innerNetPrivilegeDto), func(db *gorm.DB) {
		db.Scan(&innerNetPrivilegeDtos)
	}, false)

	return innerNetPrivilegeDtos, nil
}

/**
保存服务sql
*/
func (*InnerNetPrivilegeDao) SaveInnerNetPrivilege(innerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) error {
	return sqlTemplate.Insert(insert_innerNetPrivilege, objectConvert.Struct2Map(innerNetPrivilegeDto), false)
}

/**
修改服务sql
*/
func (*InnerNetPrivilegeDao) UpdateInnerNetPrivilege(innerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) error {
	return sqlTemplate.Update(update_innerNetPrivilege, objectConvert.Struct2Map(innerNetPrivilegeDto), false)
}

/**
删除服务sql
*/
func (*InnerNetPrivilegeDao) DeleteInnerNetPrivilege(innerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) error {
	return sqlTemplate.Delete(delete_innerNetPrivilege, objectConvert.Struct2Map(innerNetPrivilegeDto), false)
}
