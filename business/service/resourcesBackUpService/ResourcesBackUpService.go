package resourcesBackUpService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesBackUpDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/crontab"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type ResourcesBackUpService struct {
	resourcesBackUpDao resourcesBackUpDao.ResourcesBackUpDao
	hostDao            hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesBackUpService *ResourcesBackUpService) GetResourcesBackUpAll(ResourcesBackUpDto resources.ResourcesBackUpDto) ([]*resources.ResourcesBackUpDto, error) {
	var (
		err                 error
		ResourcesBackUpDtos []*resources.ResourcesBackUpDto
	)

	ResourcesBackUpDtos, err = resourcesBackUpService.resourcesBackUpDao.GetResourcesBackUps(ResourcesBackUpDto)
	if err != nil {
		return nil, err
	}

	return ResourcesBackUpDtos, nil

}

/**
查询 系统信息
*/
func (resourcesBackUpService *ResourcesBackUpService) GetResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		page                int64
		row                 int64
		total               int64
		resourcesBackUpDto  = resources.ResourcesBackUpDto{}
		resourcesBackUpDtos []*resources.ResourcesBackUpDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	resourcesBackUpDto.Row = row * page

	resourcesBackUpDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesBackUpDto.TenantId = user.TenantId

	total, err = resourcesBackUpService.resourcesBackUpDao.GetResourcesBackUpCount(resourcesBackUpDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesBackUpDtos, err = resourcesBackUpService.resourcesBackUpDao.GetResourcesBackUps(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	resourcesBackUpDtos = resourcesBackUpService.freshBackUpObjectName(resourcesBackUpDtos)

	return result.SuccessData(resourcesBackUpDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesBackUpService *ResourcesBackUpService) SaveResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                error
		resourcesBackUpDto resources.ResourcesBackUpDto
	)
	if err = ctx.ReadJSON(&resourcesBackUpDto); err != nil {
		return result.Error("解析入参失败")
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesBackUpDto.TenantId = user.TenantId
	resourcesBackUpDto.Id = seq.Generator()
	resourcesBackUpDto.State = resources.Back_up_state_STOP
	resourcesBackUpDto.BackTime = date.GetNowTimeString()
	//ResourcesBackUpDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesBackUpService.resourcesBackUpDao.SaveResourcesBackUp(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesBackUpDto)

}

/**
修改 系统信息
*/
func (resourcesBackUpService *ResourcesBackUpService) UpdateResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                error
		resourcesBackUpDto resources.ResourcesBackUpDto
	)
	if err = ctx.ReadJSON(&resourcesBackUpDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesBackUpDto.Id = ctx.FormValue("id")

	resourcesBackUpDto.TenantId = user.TenantId
	//resourcesBackUpDto.Name = ctx.FormValue("name")

	err = resourcesBackUpService.resourcesBackUpDao.UpdateResourcesBackUp(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesBackUpDto)

}

/**
删除 系统信息
*/
func (resourcesBackUpService *ResourcesBackUpService) DeleteResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                error
		resourcesBackUpDto resources.ResourcesBackUpDto
	)
	if err = ctx.ReadJSON(&resourcesBackUpDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesBackUpService.resourcesBackUpDao.DeleteResourcesBackUp(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesBackUpDto)

}

func (resourcesBackUpService *ResourcesBackUpService) freshBackUpObjectName(dtos []*resources.ResourcesBackUpDto) []*resources.ResourcesBackUpDto {

	for _,dto := range dtos{
		if dto.TypeCd == resources.Back_up_Type_Cd_db{
			dto.SrcName = dto.SrcDbName
		}else{
			dto.SrcName = dto.SrcHostName
		}

		if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_db{
			dto.TargetName = dto.TargetDbName
		}else if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_oss{
			dto.TargetName = dto.TargetOssName
		}else{
			dto.TargetName = dto.TargetFtpName
		}
	}

	return dtos
}

// 启动资源备份

func (resourcesBackUpService *ResourcesBackUpService) StartResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                error
		resourcesBackUpDto resources.ResourcesBackUpDto
		backUpJob  crontab.BackUpJob
	)
	if err = ctx.ReadJSON(&resourcesBackUpDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesBackUpDto.Id = ctx.FormValue("id")

	resourcesBackUpDto.TenantId = user.TenantId
	resourcesBackUpDto.State = resources.Back_up_state_START

	err = resourcesBackUpService.resourcesBackUpDao.UpdateResourcesBackUp(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	backUpJob.Start()

	return result.Success()
}

// 停止资源备份

func (resourcesBackUpService *ResourcesBackUpService) StopResourcesBackUps(ctx iris.Context) result.ResultDto {
	var (
		err                error
		resourcesBackUpDto resources.ResourcesBackUpDto
		backUpJob  crontab.BackUpJob
	)
	if err = ctx.ReadJSON(&resourcesBackUpDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesBackUpDto.Id = ctx.FormValue("id")

	resourcesBackUpDto.TenantId = user.TenantId
	resourcesBackUpDto.State = resources.Back_up_state_STOP

	err = resourcesBackUpService.resourcesBackUpDao.UpdateResourcesBackUp(resourcesBackUpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	backUpJob.Stop()

	return result.Success()
}
