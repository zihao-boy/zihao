package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"gorm.io/gorm"
)

const (
	query_hostAttr_count string = `
		select count(1) total
			from host_attr t
			where t.status_cd = '0'
			$if HostId != '' then
			and t.host_id = #HostId#
			$endif
			$if SpecCd != '' then
			and t.spec_cd = #SpecCd#
			$endif
			$if AttrId != '' then
			and t.attr_id = #AttrId#
			$endif
    	
	`
	query_hostAttr string = `
				select * from host_attr t
					where t.status_cd = '0'
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if SpecCd != '' then
					and t.spec_cd = #SpecCd#
					$endif
					$if AttrId != '' then
					and t.attr_id = #AttrId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_hostAttr string = `
	insert into host_attr(attr_id, host_id, value, spec_cd) 
values (#AttrId#,#HostId#,#Value#,#SpecCd#)
`

	update_hostAttr string = `
	update host_attr  set 
		  value = #Value#
		where
		status_cd = '0'
		$if HostId != '' then
		and host_id = #HostId#
		$endif
		$if SpecCd != '' then
		and spec_cd = #SpecCd#
		$endif
		$if AttrId != '' then
		and attr_id = #AttrId#
		$endif
	`
	delete_hostAttr string = `
	update host_attr set
                          status_cd = '1'
                          where status_cd = '0'
		$if AttrId != '' then
			and attr_id = #AttrId#
		$endif
	`
)

type HostAttrDao struct {
}

/**
查询用户
*/
func (*HostAttrDao) GetHostAttrCount(hostAttrDto host.HostAttrDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_hostAttr_count, objectConvert.Struct2Map(hostAttrDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*HostAttrDao) GetHostAttrs(hostAttrDto host.HostAttrDto) ([]*host.HostAttrDto, error) {
	var hostAttrDtos []*host.HostAttrDto
	sqlTemplate.SelectList(query_hostAttr, objectConvert.Struct2Map(hostAttrDto), func(db *gorm.DB) {
		db.Scan(&hostAttrDtos)
	}, false)

	return hostAttrDtos, nil
}

/**
保存服务sql
*/
func (*HostAttrDao) SaveHostAttr(hostAttrDto host.HostAttrDto) error {
	return sqlTemplate.Insert(insert_hostAttr, objectConvert.Struct2Map(hostAttrDto), false)
}

/**
修改服务sql
*/
func (*HostAttrDao) UpdateHostAttr(hostAttrDto host.HostAttrDto) error {
	return sqlTemplate.Update(update_hostAttr, objectConvert.Struct2Map(hostAttrDto), false)
}

/**
删除服务sql
*/
func (*HostAttrDao) DeleteHostAttr(hostAttrDto host.HostAttrDto) error {
	return sqlTemplate.Delete(delete_hostAttr, objectConvert.Struct2Map(hostAttrDto), false)
}
