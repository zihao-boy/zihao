package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/assets/mapper/hostMapper"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
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

	mysql.SelectList(hostMapper.QueryHostGroups,objectConvert.Struct2Map(hostGropDto), func(db *gorm.DB) {
		db.Scan(&hostGroupDtos)
	})


	return hostGroupDtos,nil
}

