package resourcesOssService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesOssDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/oss"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const maxSize = 1000 * iris.MB // 第二种方法
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

func (resourcesOssService *ResourcesOssService) ListOssFiles(ctx iris.Context) interface{} {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		resourcesOssDto = resources.ResourcesOssDto{
			OssId:    ctx.URLParam("ossId"),
			TenantId: user.TenantId,
		}
	)
	resourcesOssDtos, err := resourcesOssService.resourcesOssDao.GetResourcesOsss(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesOssDtos) < 1 {
		return result.Error("oss不存在")
	}

	resourcesOssDto = *resourcesOssDtos[0]
	resourcesOssDto.CurPath = ctx.URLParam("curPath")
	resultDto := oss.ListALiOss(resourcesOssDto)
	return resultDto
}

func (resourcesOssService *ResourcesOssService) RemoveOssFile(ctx iris.Context) interface{} {

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		resourcesOssDto  resources.ResourcesOssDto
	)


	if err := ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}
	curPath := resourcesOssDto.CurPath
	fileGroupName := resourcesOssDto.FileGroupName

	resourcesOssDto.TenantId = user.TenantId

	resourcesOssDtos, err := resourcesOssService.resourcesOssDao.GetResourcesOsss(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesOssDtos) < 1 {
		return result.Error("oss不存在")
	}
	resourcesOssDto = *resourcesOssDtos[0]
	resourcesOssDto.CurPath = curPath
	resourcesOssDto.FileGroupName = fileGroupName
	err = oss.DeleteALiOss(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.Success()
}

func (resourcesOssService *ResourcesOssService) UploadOssFile(ctx iris.Context) interface{} {
	ctx.SetMaxRequestBodySize(maxSize)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	var (
		resourcesOssDto = resources.ResourcesOssDto{
			OssId:    ctx.FormValue("ossId"),
			TenantId: user.TenantId,
		}
	)
	resourcesOssDtos, err := resourcesOssService.resourcesOssDao.GetResourcesOsss(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesOssDtos) < 1 {
		return result.Error("oss不存在")
	}
	resourcesOssDto = *resourcesOssDtos[0]
	fileName := fileHeader.Filename
	if strings.Contains(fileName, "/") {
		fileName = filepath.Base(fileName)
	}
	resourcesOssDto.CurPath = path.Join(ctx.FormValue("curPath"), fileName)
	err = oss.SaveALiOssByReader(file, resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.Success()
}

func (resourcesOssService *ResourcesOssService) DownloadOssFile(ctx iris.Context) {

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	resourcesOssDto := resources.ResourcesOssDto{
		OssId:    ctx.URLParam("ossId"),
		TenantId: user.TenantId,
	}

	resourcesOssDtos, err := resourcesOssService.resourcesOssDao.GetResourcesOsss(resourcesOssDto)

	if err != nil {
		ctx.WriteString(err.Error())
	}

	if len(resourcesOssDtos) < 1 {
		ctx.WriteString("主机不存在")
	}
	resourcesOssDto = *resourcesOssDtos[0]

	responseWriter := ctx.ResponseWriter()
	resourcesOssDto.CurPath = path.Join(ctx.URLParam("curPath"), ctx.URLParam("fileName"))
	//hostDto.CurPath = ctx.URLParam("curPath")
	responseWriter.Header().Set("Content-Disposition", "attachment; filename="+ctx.URLParam("fileName"))
	oss.DownloadALiOssByReader( responseWriter,resourcesOssDto)
}

