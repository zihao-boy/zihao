package resourcesFtpService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesFtpDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/ftp"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type ResourcesFtpService struct {
	resourcesFtpDao resourcesFtpDao.ResourcesFtpDao
	hostDao         hostDao.HostDao
}

const maxSize = 1000 * iris.MB // 第二种方法

// get db link
// all db by this user
func (resourcesFtpService *ResourcesFtpService) GetResourcesFtpAll(ResourcesFtpDto resources.ResourcesFtpDto) ([]*resources.ResourcesFtpDto, error) {
	var (
		err              error
		ResourcesFtpDtos []*resources.ResourcesFtpDto
	)

	ResourcesFtpDtos, err = resourcesFtpService.resourcesFtpDao.GetResourcesFtps(ResourcesFtpDto)
	if err != nil {
		return nil, err
	}

	return ResourcesFtpDtos, nil

}

/**
查询 系统信息
*/
func (resourcesFtpService *ResourcesFtpService) GetResourcesFtps(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesFtpDto  = resources.ResourcesFtpDto{}
		resourcesFtpDtos []*resources.ResourcesFtpDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	resourcesFtpDto.Row = row * page

	resourcesFtpDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesFtpDto.TenantId = user.TenantId

	total, err = resourcesFtpService.resourcesFtpDao.GetResourcesFtpCount(resourcesFtpDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesFtpDtos, err = resourcesFtpService.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesFtpDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesFtpService *ResourcesFtpService) SaveResourcesFtps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesFtpDto resources.ResourcesFtpDto
	)
	if err = ctx.ReadJSON(&resourcesFtpDto); err != nil {
		return result.Error("解析入参失败")
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesFtpDto.TenantId = user.TenantId
	resourcesFtpDto.FtpId = seq.Generator()
	//ResourcesFtpDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesFtpService.resourcesFtpDao.SaveResourcesFtp(resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesFtpDto)

}

/**
修改 系统信息
*/
func (resourcesFtpService *ResourcesFtpService) UpdateResourcesFtps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesFtpDto resources.ResourcesFtpDto
	)
	if err = ctx.ReadJSON(&resourcesFtpDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//resourcesFtpDto.Id = ctx.FormValue("id")

	resourcesFtpDto.TenantId = user.TenantId
	//resourcesFtpDto.Name = ctx.FormValue("name")

	err = resourcesFtpService.resourcesFtpDao.UpdateResourcesFtp(resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesFtpDto)

}

/**
删除 系统信息
*/
func (resourcesFtpService *ResourcesFtpService) DeleteResourcesFtps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesFtpDto resources.ResourcesFtpDto
	)
	if err = ctx.ReadJSON(&resourcesFtpDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesFtpService.resourcesFtpDao.DeleteResourcesFtp(resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesFtpDto)

}

func (resourcesFtpService *ResourcesFtpService) ListFtpFiles(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		resourcesFtpDto = resources.ResourcesFtpDto{
			FtpId:    ctx.URLParam("ftpId"),
			TenantId: user.TenantId,
		}
	)
	resourcesFtpDtos, err := resourcesFtpService.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesFtpDtos) < 1 {
		return result.Error("oss不存在")
	}

	resourcesFtpDto = *resourcesFtpDtos[0]
	resourcesFtpDto.CurPath = ctx.URLParam("curPath")
	resultDto := ftp.ListFile(resourcesFtpDto)
	return resultDto
}

func (resourcesFtpService *ResourcesFtpService) RemoveFtpFile(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		resourcesFtpDto resources.ResourcesFtpDto
	)

	if err := ctx.ReadJSON(&resourcesFtpDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}
	curPath := resourcesFtpDto.CurPath
	fileGroupName := resourcesFtpDto.FileGroupName

	resourcesFtpDto.TenantId = user.TenantId

	resourcesFtpDtos, err := resourcesFtpService.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesFtpDtos) < 1 {
		return result.Error("ftp不存在")
	}
	resourcesFtpDto = *resourcesFtpDtos[0]
	resourcesFtpDto.CurPath = curPath
	resourcesFtpDto.FileGroupName = fileGroupName
	err = ftp.DeleteFile(resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.Success()
}

func (resourcesFtpService *ResourcesFtpService) UploadFtpFile(ctx iris.Context) result.ResultDto {
	ctx.SetMaxRequestBodySize(maxSize)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	var (
		resourcesFtpDto = resources.ResourcesFtpDto{
			FtpId:    ctx.FormValue("ftpId"),
			TenantId: user.TenantId,
		}
	)
	resourcesFtpDtos, err := resourcesFtpService.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(resourcesFtpDtos) < 1 {
		return result.Error("oss不存在")
	}
	resourcesFtpDto = *resourcesFtpDtos[0]
	fileName := fileHeader.Filename
	if strings.Contains(fileName, "/") {
		fileName = filepath.Base(fileName)
	}
	resourcesFtpDto.CurPath = path.Join(ctx.FormValue("curPath"), fileName)
	err = ftp.UploadFileByReader(file, resourcesFtpDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.Success()
}

func (resourcesFtpService *ResourcesFtpService) DownloadFtpFile(ctx iris.Context) {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	resourcesFtpDto := resources.ResourcesFtpDto{
		FtpId:    ctx.URLParam("ftpId"),
		TenantId: user.TenantId,
	}

	resourcesFtpDtos, err := resourcesFtpService.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)

	if err != nil {
		ctx.WriteString(err.Error())
	}

	if len(resourcesFtpDtos) < 1 {
		ctx.WriteString("主机不存在")
	}
	resourcesFtpDto = *resourcesFtpDtos[0]

	responseWriter := ctx.ResponseWriter()
	resourcesFtpDto.CurPath = path.Join(ctx.URLParam("curPath"), ctx.URLParam("fileName"))
	//hostDto.CurPath = ctx.URLParam("curPath")
	responseWriter.Header().Set("Content-Disposition", "attachment; filename="+ctx.URLParam("fileName"))
	ftp.DownloadFile( responseWriter,resourcesFtpDto)
}
