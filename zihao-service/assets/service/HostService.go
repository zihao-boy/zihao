package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/assets/dao"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"strconv"
)

type HostService struct {
	hostDao dao.HostDao
}

const(
	host_token string = "host_token"
)



/**
查询 系统信息
*/
func (hostService *HostService) GetHostGroups(ctx iris.Context)  result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err       error
		page int64
		row int64
		total int64
		hostGroupDto = host.HostGroupDto{TenantId: user.TenantId}
		hostGroupDtos []*host.HostGroupDto
	)

	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	hostGroupDto.Row = row * page

	hostGroupDto.Page = (page -1) * row

	total,err = hostService.hostDao.GetHostGroupCount(hostGroupDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}
	hostGroupDtos,err = hostService.hostDao.GetHostGroups(hostGroupDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDtos,total,row)

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



/**
查询 系统信息
*/
func (hostService *HostService) GetHosts(ctx iris.Context)  result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err       error
		page int64
		row int64
		total int64
		hostDto = host.HostDto{TenantId: user.TenantId}
		hostDtos []*host.HostDto
	)


	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	hostDto.Row = row * page

	hostDto.Page = (page -1) * row

	groupId := ctx.URLParam("groupId")

	hostDto.GroupId=groupId

	total,err = hostService.hostDao.GetHostCount(hostDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}
	hostDtos,err = hostService.hostDao.GetHosts(hostDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	for _,item := range hostDtos{
		item.Passwd=""
	}

	return result.SuccessData(hostDtos,total,row)

}


/**
保存 系统信息
*/
func (hostService *HostService) SaveHost(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	hostDto.TenantId = user.TenantId
	hostDto.HostId = seq.Generator()

	err = hostService.hostDao.SaveHost(hostDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}


/**
修改 系统信息
*/
func (hostService *HostService) UpdateHost(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.UpdateHost(hostDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}


/**
删除 系统信息
*/
func (hostService *HostService) DeleteHost(ctx iris.Context)  (result.ResultDto) {
	var (
		err          error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.DeleteHost(hostDto)
	if (err != nil) {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}

/**
主机生成token
*/
func (hostService *HostService) GetHostToken(ctx iris.Context)  (result.ResultDto) {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId: ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos,err := hostService.hostDao.GetHosts(hostDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1{
		return result.Error("主机不存在")
	}

	var hostToken string = seq.Generator()
	redis.G_Redis.SetValue(host_token,hostToken)

	return result.SuccessData(hostToken)

}