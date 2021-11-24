package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/mapping"
	"gorm.io/gorm"
)

const (
	query_mapping_count string = `
		select count(1) total from mapping t
				where t.status_cd = '0'
				$if Domain != '' then
				and t.domain = #Domain#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if Zkeys != '' then
				and t.zkeys=#Zkeys#
				$endif
				$if Value != '' then
				and t.value = #Value#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
    	
	`
	query_mapping string = `
		select * from mapping t
				where t.status_cd = '0'
				$if Domain != '' then
				and t.domain = #Domain#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if Zkeys != '' then
				and t.zkeys=#Zkeys#
				$endif
				$if Value != '' then
				and t.value = #Value#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
				order by t.create_time desc
				$if Page != -1 then
					limit #Page#,#Row#
				$endif
	`

	insert_mapping string = `
	insert into mapping( domain, name, zkeys, value, remark)
values(#Domain#,#Name#,#Zkeys#,#Value#,#Remark#)
`

	update_mapping string = `
	update mapping t set
		$if Domain != '' then
		 t.domain = #Domain#,
		$endif
		$if Name != '' then
		 t.name = #Name#,
		$endif
		$if Zkeys != '' then
		 t.zkeys=#Zkeys#,
		$endif
		$if Value != '' then
		 t.value = #Value#,
		$endif
		$if Remark != '' then
		 t.remark = #Remark#,
		$endif
		t.status_cd = '0'
		where t.status_cd = '0'
		  and t.id = #Id#
	`
	delete_mapping string = `
	update mapping t set
		t.status_cd = '1'
		where t.status_cd = '0'
		  and t.id = #Id#
	`

	query_dict string = `
	select * from t_dict t
		where 1=1
		and t.table_name = #TableName#
		and t.table_columns = #TableColumns#
	`
)

type MappingDao struct {
}

/**
查询用户
*/
func (*MappingDao) GetMappingCount(mappingDto mapping.MappingDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_mapping_count, objectConvert.Struct2Map(mappingDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*MappingDao) GetMappings(mappingDto mapping.MappingDto) ([]*mapping.MappingDto, error) {
	var mappingDtos []*mapping.MappingDto
	sqlTemplate.SelectList(query_mapping, objectConvert.Struct2Map(mappingDto), func(db *gorm.DB) {
		db.Scan(&mappingDtos)
	}, false)

	return mappingDtos, nil
}

/**
查询用户
*/
func (*MappingDao) GetDicts(dictDto mapping.DictDto) ([]*mapping.DictDto, error) {
	var dictDtos []*mapping.DictDto
	sqlTemplate.SelectList(query_dict, objectConvert.Struct2Map(dictDto), func(db *gorm.DB) {
		db.Scan(&dictDtos)
	}, false)

	return dictDtos, nil
}

/**
保存服务sql
*/
func (*MappingDao) SaveMapping(mappingDto mapping.MappingDto) error {
	return sqlTemplate.Insert(insert_mapping, objectConvert.Struct2Map(mappingDto), false)
}

/**
修改服务sql
*/
func (*MappingDao) UpdateMapping(mappingDto mapping.MappingDto) error {
	return sqlTemplate.Update(update_mapping, objectConvert.Struct2Map(mappingDto), false)
}

/**
删除服务sql
*/
func (*MappingDao) DeleteMapping(mappingDto mapping.MappingDto) error {
	return sqlTemplate.Delete(delete_mapping, objectConvert.Struct2Map(mappingDto), false)
}
