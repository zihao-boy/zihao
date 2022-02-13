package resourcesDbService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesDbDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type ResourcesDbService struct {
	resourcesDbDao resourcesDbDao.ResourcesDbDao
	hostDao         hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesDbService *ResourcesDbService) GetResourcesDbAll(ResourcesDbDto resources.ResourcesDbDto) ([]*resources.ResourcesDbDto, error) {
	var (
		err              error
		ResourcesDbDtos []*resources.ResourcesDbDto
	)

	ResourcesDbDtos, err = resourcesDbService.resourcesDbDao.GetResourcesDbs(ResourcesDbDto)
	if err != nil {
		return nil, err
	}

	return ResourcesDbDtos, nil

}

/**
查询 系统信息
*/
func (resourcesDbService *ResourcesDbService) GetResourcesDbs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesDbDto  = resources.ResourcesDbDto{}
		resourcesDbDtos []*resources.ResourcesDbDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	resourcesDbDto.Row = row * page

	resourcesDbDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesDbDto.TenantId = user.TenantId

	total, err = resourcesDbService.resourcesDbDao.GetResourcesDbCount(resourcesDbDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesDbDtos, err = resourcesDbService.resourcesDbDao.GetResourcesDbs(resourcesDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesDbDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesDbService *ResourcesDbService) SaveResourcesDbs(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesDbDto resources.ResourcesDbDto
	)
	if err = ctx.ReadJSON(&resourcesDbDto); err != nil {
		return result.Error("解析入参失败")
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesDbDto.TenantId = user.TenantId
	resourcesDbDto.DbId = seq.Generator()
	//ResourcesDbDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesDbService.resourcesDbDao.SaveResourcesDb(resourcesDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesDbDto)

}

/**
修改 系统信息
*/
func (resourcesDbService *ResourcesDbService) UpdateResourcesDbs(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesDbDto resources.ResourcesDbDto
	)
	if err = ctx.ReadJSON(&resourcesDbDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesDbDto.Id = ctx.FormValue("id")

	resourcesDbDto.TenantId = user.TenantId
	//resourcesDbDto.Name = ctx.FormValue("name")

	err = resourcesDbService.resourcesDbDao.UpdateResourcesDb(resourcesDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesDbDto)

}

/**
删除 系统信息
*/
func (resourcesDbService *ResourcesDbService) DeleteResourcesDbs(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesDbDto resources.ResourcesDbDto
	)
	if err = ctx.ReadJSON(&resourcesDbDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesDbService.resourcesDbDao.DeleteResourcesDb(resourcesDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesDbDto)

}
