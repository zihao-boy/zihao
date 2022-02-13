package resourcesOssService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesOssDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type ResourcesOssService struct {
	resourcesOssDao resourcesOssDao.ResourcesOssDao
	hostDao         hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesOssService *ResourcesOssService) GetResourcesOssAll(ResourcesOssDto resources.ResourcesOssDto) ([]*resources.ResourcesOssDto, error) {
	var (
		err              error
		ResourcesOssDtos []*resources.ResourcesOssDto
	)

	ResourcesOssDtos, err = resourcesOssService.resourcesOssDao.GetResourcesOsss(ResourcesOssDto)
	if err != nil {
		return nil, err
	}

	return ResourcesOssDtos, nil

}

/**
查询 系统信息
*/
func (resourcesOssService *ResourcesOssService) GetResourcesOsss(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesOssDto  = resources.ResourcesOssDto{}
		resourcesOssDtos []*resources.ResourcesOssDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	resourcesOssDto.Row = row * page

	resourcesOssDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesOssDto.TenantId = user.TenantId

	total, err = resourcesOssService.resourcesOssDao.GetResourcesOssCount(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesOssDtos, err = resourcesOssService.resourcesOssDao.GetResourcesOsss(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesOssService *ResourcesOssService) SaveResourcesOsss(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto resources.ResourcesOssDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesOssDto.TenantId = user.TenantId
	resourcesOssDto.OssId = seq.Generator()
	//ResourcesOssDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesOssService.resourcesOssDao.SaveResourcesOss(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
修改 系统信息
*/
func (resourcesOssService *ResourcesOssService) UpdateResourcesOsss(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto resources.ResourcesOssDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesOssDto.Id = ctx.FormValue("id")

	resourcesOssDto.TenantId = user.TenantId
	//resourcesOssDto.Name = ctx.FormValue("name")

	err = resourcesOssService.resourcesOssDao.UpdateResourcesOss(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
删除 系统信息
*/
func (resourcesOssService *ResourcesOssService) DeleteResourcesOsss(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto resources.ResourcesOssDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesOssService.resourcesOssDao.DeleteResourcesOss(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}
