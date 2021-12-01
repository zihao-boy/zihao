package service

import (
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"path"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
)

type BusinessPackageService struct {
	businessPackageDao dao.BusinessPackageDao
}

const maxSize = 1000 * iris.MB // 第二种方法

/**
查询 系统信息
*/
func (businessPackageService *BusinessPackageService) GetBusinessPackageAll(businessPackageDto businessPackage.BusinessPackageDto) ([]*businessPackage.BusinessPackageDto, error) {
	var (
		err                 error
		businessPackageDtos []*businessPackage.BusinessPackageDto
	)

	businessPackageDtos, err = businessPackageService.businessPackageDao.GetBusinessPackages(businessPackageDto)
	if err != nil {
		return nil, err
	}

	return businessPackageDtos, nil

}

/**
查询 系统信息
*/
func (businessPackageService *BusinessPackageService) GetBusinessPackages(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		page                int64
		row                 int64
		total               int64
		businessPackageDto  = businessPackage.BusinessPackageDto{}
		businessPackageDtos []*businessPackage.BusinessPackageDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessPackageDto.Row = row * page

	businessPackageDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessPackageDto.TenantId = user.TenantId

	total, err = businessPackageService.businessPackageDao.GetBusinessPackageCount(businessPackageDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessPackageDtos, err = businessPackageService.businessPackageDao.GetBusinessPackages(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessPackageDtos, total, row)

}

/**
保存 系统信息
*/
func (businessPackageService *BusinessPackageService) SaveBusinessPackages(ctx iris.Context) result.ResultDto {
	var (
		err                error
		businessPackageDto businessPackage.BusinessPackageDto
	)

	ctx.SetMaxRequestBodySize(maxSize)

	file, fileHeader, err := ctx.FormFile("uploadFile")

	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	defer file.Close()
	dest := filepath.Join(config.G_AppConfig.DataPath,"/businessPackage")

	if(!utils.IsDir(dest)){
		utils.CreateDir(dest)
	}

	fileName := path.Ext(fileHeader.Filename)

	dest = filepath.Join(dest,seq.Generator()+fileName)

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败"+err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	businessPackageDto.Id = seq.Generator()
	businessPackageDto.Path = dest
	businessPackageDto.Varsion = "V" + date.GetNowAString()
	businessPackageDto.Name = ctx.FormValue("name")

	err = businessPackageService.businessPackageDao.SaveBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessPackageDto)

}

/**
修改 系统信息
*/
func (businessPackageService *BusinessPackageService) UpdateBusinessPackages(ctx iris.Context) result.ResultDto {
	var (
		err                error
		businessPackageDto businessPackage.BusinessPackageDto
	)

	if err = ctx.ReadJSON(&businessPackageDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessPackageService.businessPackageDao.UpdateBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessPackageDto)

}

/**
删除 系统信息
*/
func (businessPackageService *BusinessPackageService) DeleteBusinessPackages(ctx iris.Context) result.ResultDto {
	var (
		err                error
		businessPackageDto businessPackage.BusinessPackageDto
	)

	if err = ctx.ReadJSON(&businessPackageDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessPackageService.businessPackageDao.DeleteBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessPackageDto)

}
