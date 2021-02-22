package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
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


/**
保存 系统信息
*/
func (hostService *HostService) SaveHostGroups(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	hostGroupDto.TenantId = user.TenantId
	hostGroupDto.GroupId = seq.Generator()

	err = hostService.hostDao.SaveHostGroup(hostGroupDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}


/**
修改 系统信息
*/
func (hostService *HostService) UpdateHostGroups(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.UpdateHostGroup(hostGroupDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}


/**
删除 系统信息
*/
func (hostService *HostService) DeleteHostGroups(ctx iris.Context)  (result.ResultDto) {
	var (
		err          error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.DeleteHostGroup(hostGroupDto)
	if (err != nil) {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}