package service

import (
	"github.com/zihao-boy/zihao/config"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
)

type BusinessDockerfileService struct {
	businessDockerfileDao dao.BusinessDockerfileDao
}

/**
查询 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) GetBusinessDockerfileAll(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) ([]*businessDockerfile.BusinessDockerfileDto, error) {
	var (
		err                    error
		businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	)

	businessDockerfileDtos, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfiles(businessDockerfileDto)
	if err != nil {
		return nil, err
	}

	return businessDockerfileDtos, nil

}

/**
查询 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) GetBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		page                   int64
		row                    int64
		total                  int64
		businessDockerfileDto  = businessDockerfile.BusinessDockerfileDto{}
		businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessDockerfileDto.Row = row * page

	businessDockerfileDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.Name = ctx.URLParam("name")
	businessDockerfileDto.Version = ctx.URLParam("version")
	businessDockerfileDto.Id = ctx.URLParam("id")

	total, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfileCount(businessDockerfileDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessDockerfileDtos, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfiles(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDtos, total, row)

}

/**
保存 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) SaveBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)
	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	businessDockerfileDto.Id = seq.Generator()
	businessDockerfileDto.Version = "V" + date.GetNowAString()

	//save log

	logPath := filepath.Join(config.WorkSpace, "businessPackage/"+user.TenantId, businessDockerfileDto.Id+".log")

	businessDockerfileDto.LogPath = logPath

	err = businessDockerfileService.businessDockerfileDao.SaveBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}

/**
修改 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) UpdateBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessDockerfileService.businessDockerfileDao.UpdateBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}

/**
删除 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) DeleteBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessDockerfileService.businessDockerfileDao.DeleteBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}
