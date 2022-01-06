package service

import (
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"os"
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
	hostDao            hostDao.HostDao
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

	curDest := filepath.Join("businessPackage", user.TenantId)
	dest := filepath.Join(config.WorkSpace, curDest)

	if businessPackageDtos == nil || len(businessPackageDtos) < 1 {
		return result.SuccessData(businessPackageDtos, total, row)
	}

	for _, tmpBusinessPackageDto := range businessPackageDtos {
		tmpBusinessPackageDto.BasePath = dest
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
	defer file.Close()
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessPackageDto.Id = seq.Generator()

	curDest := filepath.Join("businessPackage", user.TenantId, businessPackageDto.Id)
	dest := filepath.Join(config.WorkSpace, curDest)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest = filepath.Join(dest, fileHeader.Filename)

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	//businessPackageDto.Path = filepath.Join(curDest, fileHeader.Filename)
	businessPackageDto.Path = filepath.Join(businessPackageDto.Id, fileHeader.Filename)
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

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessPackageDto.Id = ctx.FormValue("id")

	curDest := filepath.Join("businessPackage", user.TenantId, businessPackageDto.Id)
	dest := filepath.Join(config.WorkSpace, curDest)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest = filepath.Join(dest, fileHeader.Filename)

	// remove file that exists
	if utils.IsFile(dest) {
		os.Remove(dest)
	}

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	businessPackageDto.Name = ctx.FormValue("name")

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

func (businessPackageService *BusinessPackageService) Upload(ctx iris.Context) interface{} {
	var (
		err                error
		businessPackageDto businessPackage.BusinessPackageDto
	)

	ctx.SetMaxRequestBodySize(maxSize)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessPackageDto.Id = seq.Generator()

	curDest := filepath.Join("businessPackage", user.TenantId, businessPackageDto.Id)
	dest := filepath.Join(config.WorkSpace, curDest)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest = filepath.Join(dest, fileHeader.Filename)

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	//businessPackageDto.Path = filepath.Join(curDest, fileHeader.Filename)
	businessPackageDto.Path = filepath.Join(businessPackageDto.Id, fileHeader.Filename)
	businessPackageDto.Varsion = "V" + date.GetNowAString()
	businessPackageDto.Name = fileHeader.Filename

	err = businessPackageService.businessPackageDao.SaveBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessPackageDto)
}

func (businessPackageService *BusinessPackageService) ListBusinessPackageContext(ctx iris.Context) interface{} {

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	shellPackageId := ctx.URLParam("shellPackageId")
	businessPackageDto := businessPackage.BusinessPackageDto{
		Id: shellPackageId,
	}
	businessPackageDtos, err := businessPackageService.businessPackageDao.GetBusinessPackages(businessPackageDto)
	if err != nil || len(businessPackageDtos) < 1 {
		return result.Error(err.Error())
	}

	curDest := filepath.Join("businessPackage", user.TenantId)
	dest := filepath.Join(config.WorkSpace, curDest, businessPackageDtos[0].Path)
	var (
		hostDto = host.HostDto{
			HostId:   host.MASTER_HOST_ID,
			TenantId: user.TenantId,
			FileName: dest,
		}
	)
	hostDtos, err := businessPackageService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecListFileContext(hostDto)
	return resultDto
}

func (businessPackageService *BusinessPackageService) EditBusinessPackageContext(ctx iris.Context) interface{} {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto         host.HostDto
		fasterDeployDto appService.FasterDeployDto
	)

	if err := ctx.ReadJSON(&fasterDeployDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}
	businessPackageDto := businessPackage.BusinessPackageDto{
		Id: fasterDeployDto.ShellPackageId,
	}

	businessPackageDtos, err := businessPackageService.businessPackageDao.GetBusinessPackages(businessPackageDto)
	if err != nil || len(businessPackageDtos) < 1 {
		return result.Error(err.Error())
	}

	curDest := filepath.Join("businessPackage", user.TenantId)
	dest := filepath.Join(config.WorkSpace, curDest, businessPackageDtos[0].Path)

	hostDto = host.HostDto{
		HostId:      host.MASTER_HOST_ID,
		TenantId:    user.TenantId,
		FileName:    dest,
		FileContext: fasterDeployDto.ShellContext,
	}

	hostDtos, err := businessPackageService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	hostDto.Ip = hostDtos[0].Ip

	resultDto, _ := shell.ExecEditFile(hostDto)
	return resultDto
}
