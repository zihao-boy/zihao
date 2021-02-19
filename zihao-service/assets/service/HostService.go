package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
	"github.com/zihao-boy/zihao/zihao-service/assets/dao"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
)

type HostService struct {
	hostDao dao.HostDao
}



/**
查询 系统信息
*/
func (hostService *HostService) GetHostGroups(ctx iris.Context)  result.ResultDto {
	var (
		err       error
		hostGroupDto = host.HostGroupDto{}
		hostGroupDtos []*host.HostGroupDto
	)
	hostGroupDtos,err = hostService.hostDao.GetHostGroups(hostGroupDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDtos)

}