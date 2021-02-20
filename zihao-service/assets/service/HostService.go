package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
	"github.com/zihao-boy/zihao/zihao-service/assets/dao"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
)

type HostService struct {
	hostDao dao.HostDao
}



/**
查询 系统信息
*/
func (hostService *HostService) GetHostGroups(ctx iris.Context)  result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err       error
		hostGroupDto = host.HostGroupDto{TenantId: user.TenantId}
		hostGroupDtos []*host.HostGroupDto
	)
	hostGroupDtos,err = hostService.hostDao.GetHostGroups(hostGroupDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDtos)

}