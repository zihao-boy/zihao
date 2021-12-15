package service

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/date"
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

type AppVersionJobService struct {
	appVersionJobDao       dao.AppVersionJobDao
	appVersionJobDetailDao dao.AppVersionJobDetailDao
}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobAll(appVersionJobDto appVersionJob.AppVersionJobDto) ([]*appVersionJob.AppVersionJobDto, error) {
	var (
		err               error
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)

	appVersionJobDtos, err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if err != nil {
		return nil, err
	}

	return appVersionJobDtos, nil

}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		page              int64
		row               int64
		total             int64
		appVersionJobDto  = appVersionJob.AppVersionJobDto{}
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionJobDto.Row = row * page

	appVersionJobDto.Page = (page - 1) * row

	appVersionJobDto.JobId = ctx.URLParam("jobId")
	appVersionJobDto.JobName = ctx.URLParam("jobName")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobDto.TenantId = user.TenantId

	total, err = appVersionJobService.appVersionJobDao.GetAppVersionJobCount(appVersionJobDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionJobDtos, err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDtos, total, row)

}

/**
保存 系统信息
*/
func (appVersionJobService *AppVersionJobService) SaveAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.JobId = seq.Generator()
	appVersionJobDto.State = appVersionJob.STATE_wait
	appVersionJobDto.JobTime = date.GetNowTimeString()

	err = appVersionJobService.appVersionJobDao.SaveAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(appVersionJobDto.AppVersionJobImages) > 0 {
		err = appVersionJobService.saveAppVersionJobImage(appVersionJobDto,appVersionJobDto.AppVersionJobImages)
		if err != nil {
			return result.Error(err.Error())
		}
	}

	//var jobPath string = path.Join(appVersionJobDto.WorkDir, appVersionJobDto.JobId)
	//var fileName string = "job.sh"
	//
	//_, err = os.Stat(jobPath)
	//
	//if err != nil && os.IsNotExist(err) {
	//	err = os.MkdirAll(jobPath, 0777)
	//}
	//
	////当前用户目录下生成 文件夹
	//file, err := os.OpenFile(path.Join(jobPath, fileName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	//defer func() {
	//	file.Close()
	//}()
	//if err != nil && os.IsNotExist(err) {
	//	file, err = os.Create(path.Join(jobPath, fileName))
	//}
	//_, err = file.WriteString("cd " + jobPath + "\n")
	//_, err = file.WriteString(appVersionJobDto.JobShell)
	//
	//if err != nil {
	//	fmt.Print("err=", err.Error())
	//}

	return result.SuccessData(appVersionJobDto)

}

// 保存 构建计划
func (appVersionJobService *AppVersionJobService) saveAppVersionJobImage(appVersionJobDto appVersionJob.AppVersionJobDto,
	imagess []appVersionJob.AppVersionJobImagesDto) error {

	for _, images := range imagess {
		images.TenantId = appVersionJobDto.TenantId
		images.JobId = appVersionJobDto.JobId
		images.JobImagesId = seq.Generator()
		err := appVersionJobService.appVersionJobDao.SaveAppVersionJobImages(images)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
修改 系统信息
*/
func (appVersionJobService *AppVersionJobService) UpdateAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}
	if len(appVersionJobDto.AppVersionJobImages) < 1{
		return result.SuccessData(appVersionJobDto)
	}
	var appVersionJobImagesDto = appVersionJob.AppVersionJobImagesDto{
		JobId: appVersionJobDto.JobId,
	}
	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}
	err = appVersionJobService.saveAppVersionJobImage(appVersionJobDto,appVersionJobDto.AppVersionJobImages)
	if err != nil {
		return result.Error(err.Error())
	}
	//tmpUser, _ := systemUser.Current()
	//var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDto.JobId + "/"
	//var fileName string = appVersionJobDto.JobId + ".sh"
	//
	//_, err = os.Stat(path)
	//
	//if err != nil && os.IsNotExist(err) {
	//	err = os.MkdirAll(path, 0777)
	//}
	//
	////当前用户目录下生成 文件夹
	//file, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	//defer func() { file.Close() }()
	//if err != nil && os.IsNotExist(err) {
	//	file, err = os.Create(path + fileName)
	//}
	//_, err = file.WriteString("cd " + path + "\n")
	//_, err = file.WriteString(appVersionJobDto.JobShell)
	//
	//if err != nil {
	//	fmt.Print("err=", err.Error())
	//}
	return result.SuccessData(appVersionJobDto)
}

/**
删除 系统信息
*/
func (appVersionJobService *AppVersionJobService) DoJob(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.State = appVersionJob.STATE_doing
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	tmpUser, _ := systemUser.Current()
	var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDto.JobId + "/"
	var fileName string = path + appVersionJobDto.JobId + ".sh"

	//插入构建记录
	var appVersionJobDetailDto = appVersionJob.AppVersionJobDetailDto{
		JobId:    appVersionJobDto.JobId,
		State:    appVersionJob.STATE_wait,
		LogPath:  path + appVersionJobDto.JobId + ".log",
		TenantId: user.TenantId,
		DetailId: seq.Generator(),
	}

	err = appVersionJobService.appVersionJobDetailDao.SaveAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	jobShell := "nohup sh " + fileName + " " + appVersionJobDetailDto.DetailId + " >" + path + appVersionJobDto.JobId + ".log &"
	cmd := exec.Command("bash", "-c", jobShell)
	fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		appVersionJobDto.State = appVersionJob.STATE_error
		err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	}
	return result.SuccessData(appVersionJobDto)

}

/**
构建
*/
func (appVersionJobService *AppVersionJobService) DeleteAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDto)

}

func (appVersionJobService *AppVersionJobService) GetAppVersionJobImages(ctx iris.Context) interface{} {
	var (
		err                     error
		page                    int64
		row                     int64
		total                   int64
		appVersionJobImagesDto  = appVersionJob.AppVersionJobImagesDto{}
		appVersionJobImagesDtos []*appVersionJob.AppVersionJobImagesDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionJobImagesDto.Row = row * page
	appVersionJobImagesDto.Page = (page - 1) * row

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobImagesDto.TenantId = user.TenantId
	appVersionJobImagesDto.JobId = ctx.URLParam("jobId")

	total, err = appVersionJobService.appVersionJobDao.GetAppVersionJobImagesCount(appVersionJobImagesDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionJobImagesDtos, err = appVersionJobService.appVersionJobDao.GetAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobImagesDtos, total, row)
}

func (appVersionJobService *AppVersionJobService) SaveAppVersionJobImages(ctx iris.Context) interface{} {
	var (
		err                    error
		appVersionJobImagesDto appVersionJob.AppVersionJobImagesDto
	)

	if err = ctx.ReadJSON(&appVersionJobImagesDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobImagesDto.TenantId = user.TenantId
	appVersionJobImagesDto.JobImagesId = seq.Generator()

	err = appVersionJobService.appVersionJobDao.SaveAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobImagesDto)
}

func (appVersionJobService *AppVersionJobService) UpdateAppVersionJobImages(ctx iris.Context) interface{} {
	var (
		err                    error
		appVersionJobImagesDto appVersionJob.AppVersionJobImagesDto
	)

	if err = ctx.ReadJSON(&appVersionJobImagesDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobImagesDto.TenantId = user.TenantId
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobImagesDto)
}

func (appVersionJobService *AppVersionJobService) DeleteAppVersionJobImages(ctx iris.Context) interface{} {
	var (
		err                    error
		appVersionJobImagesDto appVersionJob.AppVersionJobImagesDto
	)

	if err = ctx.ReadJSON(&appVersionJobImagesDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobImagesDto)
}
