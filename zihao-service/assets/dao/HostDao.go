package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/assets/mapper/hostMapper"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
)

const(
	insert_service_sql string = "hostDao.SaveHostGroup"
	update_service_sql string = "hostDao.UpdateHostGroup"
	delete_service_sql string = "hostDao.DeleteHostGroup"
)

type HostDao struct {

}

/**
查询用户
*/
func (*HostDao) GetHostGroups(hostGropDto host.HostGroupDto) ([]*host.HostGroupDto,error){
	var (
		hostGroupDtos []*host.HostGroupDto
	)

	sqlTemplate.SelectList(hostMapper.QueryHostGroups,objectConvert.Struct2Map(hostGropDto), func(db *gorm.DB) {
		db.Scan(&hostGroupDtos)
	},false)


	return hostGroupDtos,nil
}


/**
保存服务sql
*/
func (*HostDao) SaveHostGroup(hostGroupDto host.HostGroupDto) error{
	return sqlTemplate.Insert(insert_service_sql,objectConvert.Struct2Map(hostGroupDto),true)
}

/**
修改服务sql
*/
func (*HostDao) UpdateHostGroup(hostGroupDto host.HostGroupDto) error{
	return sqlTemplate.Update(update_service_sql,objectConvert.Struct2Map(hostGroupDto),true)
}

/**
删除服务sql
*/
func (*HostDao) DeleteHostGroup(hostGroupDto host.HostGroupDto) error{
	return sqlTemplate.Delete(delete_service_sql,objectConvert.Struct2Map(hostGroupDto),true)
}


