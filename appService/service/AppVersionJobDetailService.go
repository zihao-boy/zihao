package service

import (
	"fmt"
	"os/exec"
	systemUser "os/user"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appVersionJob"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
)

type AppVersionJobDetailService struct {
	appVersionJobDetailDao dao.AppVersionJobDetailDao
}

/**
查询 系统信息
*/
func (appVersionJobDetailService *AppVersionJobDetailService) GetAppVersionJobDetailAll(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) ([]*appVersionJob.AppVersionJobDetailDto, error) {
	var (
		err                     error
		appVersionJobDetailDtos []*appVersionJob.AppVersionJobDetailDto
	)

	appVersionJobDetailDtos, err = appVersionJobDetailService.appVersionJobDetailDao.GetAppVersionJobDetails(appVersionJobDetailDto)
	if err != nil {
		return nil, err
	}

	return appVersionJobDetailDtos, nil

}

/**
查询 系统信息
*/
func (appVersionJobDetailService *AppVersionJobDetailService) GetAppVersionJobDetails(ctx iris.Context) result.ResultDto {
	var (
		err                     error
		page                    int64
		row                     int64
		total                   int64
		appVersionJobDetailDto  = appVersionJob.AppVersionJobDetailDto{}
		appVersionJobDetailDtos []*appVersionJob.AppVersionJobDetailDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionJobDetailDto.Row = row * page

	appVersionJobDetailDto.Page = (page - 1) * row

	appVersionJobDetailDto.JobId = ctx.URLParam("jobId")
	appVersionJobDetailDto.DetailId = ctx.URLParam("detailId")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobDetailDto.TenantId = user.TenantId

	total, err = appVersionJobDetailService.appVersionJobDetailDao.GetAppVersionJobDetailCount(appVersionJobDetailDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionJobDetailDtos, err = appVersionJobDetailService.appVersionJobDetailDao.GetAppVersionJobDetails(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDetailDtos, total, row)

}

/**
保存 系统信息
*/
func (appVersionJobDetailService *AppVersionJobDetailService) SaveAppVersionJobDetails(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto
	)

	if err = ctx.ReadJSON(&appVersionJobDetailDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDetailDto.TenantId = user.TenantId
	appVersionJobDetailDto.DetailId = seq.Generator()
	appVersionJobDetailDto.State = appVersionJob.STATE_wait

	err = appVersionJobDetailService.appVersionJobDetailDao.SaveAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDetailDto)

}

/**
修改 系统信息
*/
func (appVersionJobDetailService *AppVersionJobDetailService) UpdateAppVersionJobDetails(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto
	)

	if err = ctx.ReadJSON(&appVersionJobDetailDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDetailDto.TenantId = user.TenantId
	err = appVersionJobDetailService.appVersionJobDetailDao.UpdateAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDetailDto)

}

/**
删除 系统信息
*/
func (appVersionJobDetailService *AppVersionJobDetailService) DoJob(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto
	)

	if err = ctx.ReadJSON(&appVersionJobDetailDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDetailDto.TenantId = user.TenantId
	appVersionJobDetailDto.State = appVersionJob.STATE_doing
	err = appVersionJobDetailService.appVersionJobDetailDao.UpdateAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	tmpUser, _ := systemUser.Current()
	var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDetailDto.JobId + "/"
	var fileName string = path + appVersionJobDetailDto.JobId + ".sh"

	jobShell := "nohup sh " + fileName + " >" + path + appVersionJobDetailDto.JobId + ".log &"
	cmd := exec.Command("bash", "-c", jobShell)
	//fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		appVersionJobDetailDto.State = appVersionJob.STATE_error
		err = appVersionJobDetailService.appVersionJobDetailDao.UpdateAppVersionJobDetail(appVersionJobDetailDto)
	}
	return result.SuccessData(appVersionJobDetailDto)

}

/**
构建
*/
func (appVersionJobDetailService *AppVersionJobDetailService) DeleteAppVersionJobDetails(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto
	)

	if err = ctx.ReadJSON(&appVersionJobDetailDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionJobDetailService.appVersionJobDetailDao.DeleteAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDetailDto)

}
