package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/assets/mapper/hostMapper"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
	"gorm.io/gorm"
)

const (
	insert_service_sql string = "hostDao.SaveHostGroup"
	update_service_sql string = "hostDao.UpdateHostGroup"
	delete_service_sql string = "hostDao.DeleteHostGroup"

	get_host_count string = "hostDao.GetHostCount"
	get_hosts      string = "hostDao.GetHosts"
	save_host      string = "hostDao.SaveHost"
	update_host    string = "hostDao.UpdateHost"
	delete_host    string = "hostDao.DeleteHost"
)

type HostDao struct {
}

/**
查询用户
*/
func (*HostDao) GetHostGroupCount(hostGropDto host.HostGroupDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(hostMapper.QueryHostGroupCount, objectConvert.Struct2Map(hostGropDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*HostDao) GetHostGroups(hostGropDto host.HostGroupDto) ([]*host.HostGroupDto, error) {
	var (
		hostGroupDtos []*host.HostGroupDto
	)

	sqlTemplate.SelectList(hostMapper.QueryHostGroups, objectConvert.Struct2Map(hostGropDto), func(db *gorm.DB) {
		db.Scan(&hostGroupDtos)
	}, false)

	return hostGroupDtos, nil
}

/**
保存服务sql
*/
func (*HostDao) SaveHostGroup(hostGroupDto host.HostGroupDto) error {
	return sqlTemplate.Insert(insert_service_sql, objectConvert.Struct2Map(hostGroupDto), true)
}

/**
修改服务sql
*/
func (*HostDao) UpdateHostGroup(hostGroupDto host.HostGroupDto) error {
	return sqlTemplate.Update(update_service_sql, objectConvert.Struct2Map(hostGroupDto), true)
}

/**
删除服务sql
*/
func (*HostDao) DeleteHostGroup(hostGroupDto host.HostGroupDto) error {
	return sqlTemplate.Delete(delete_service_sql, objectConvert.Struct2Map(hostGroupDto), true)
}

/**
查询用户
*/
func (*HostDao) GetHostCount(hostDto host.HostDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(get_host_count, objectConvert.Struct2Map(hostDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, true)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*HostDao) GetHosts(hostDto host.HostDto) ([]*host.HostDto, error) {
	var (
		hostDtos []*host.HostDto
		err      error
	)

	sqlTemplate.SelectList(get_hosts, objectConvert.Struct2Map(hostDto), func(db *gorm.DB) {
		err = db.Scan(&hostDtos).Error
	}, true)

	return hostDtos, err
}

/**
保存主机
*/
func (*HostDao) SaveHost(hostDto host.HostDto) error {
	return sqlTemplate.Insert(save_host, objectConvert.Struct2Map(hostDto), true)
}

/**
修改主机
*/
func (*HostDao) UpdateHost(hostDto host.HostDto) error {
	return sqlTemplate.Update(update_host, objectConvert.Struct2Map(hostDto), true)
}

/**
删除主机
*/
func (*HostDao) DeleteHost(hostDto host.HostDto) error {
	return sqlTemplate.Delete(delete_host, objectConvert.Struct2Map(hostDto), true)
}
