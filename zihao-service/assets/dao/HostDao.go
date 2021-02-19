package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/assets/mapper/hostMapper"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
)

type HostDao struct {

}

/**
查询用户
*/
func (*HostDao) GetHostGroups(hostGropDto host.HostGroupDto) ([]*host.HostGroupDto,error){
	var hostGroupDtos []*host.HostGroupDto
	db := mysql.G_DB.Raw(hostMapper.QueryHostGroupCount,hostGropDto.GroupId)
	if err:=db.Scan(&hostGroupDtos).Error; err !=nil{
		return nil,err
	}

	return hostGroupDtos,nil
}

