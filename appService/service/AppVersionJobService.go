package service

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/queue/dockerfileQueue"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appVersionJob"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	softDao "github.com/zihao-boy/zihao/softService/dao"
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

	if appVersionJobDto.WorkDir == "" || appVersionJobDto.WorkDir == "/" {
		return result.Error("工作目录错误，不能为空或者/")
	}

	err = appVersionJobService.appVersionJobDao.SaveAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(appVersionJobDto.AppVersionJobImages) > 0 {
		err = appVersionJobService.saveAppVersionJobImage(appVersionJobDto, appVersionJobDto.AppVersionJobImages)
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
	if len(appVersionJobDto.AppVersionJobImages) < 1 {
		return result.SuccessData(appVersionJobDto)
	}
	var appVersionJobImagesDto = appVersionJob.AppVersionJobImagesDto{
		JobId: appVersionJobDto.JobId,
	}
	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJobImages(appVersionJobImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}
	err = appVersionJobService.saveAppVersionJobImage(appVersionJobDto, appVersionJobDto.AppVersionJobImages)
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
	appVersionJobDtos, err := appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)

	if len(appVersionJobDtos) < 1 {
		return result.Error("构建不存在")
	}

	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.State = appVersionJob.STATE_doing
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	appVersionJobDto = *appVersionJobDtos[0]

	workDir := path.Join(appVersionJobDto.WorkDir, appVersionJobDto.JobId)

	//判断是否是 /开头

	if !strings.HasPrefix(workDir, "/") {
		workDir = "/" + workDir
	}

	//删除目录
	if utils.IsDir(workDir) {
		err = os.RemoveAll(workDir)
	}
	utils.CreateDir(workDir)

	dest := path.Join(workDir, "build.sh")
	// remove file that exists
	if utils.IsFile(dest) {
		os.Remove(dest)
	}

	file, err := os.Create(dest)
	defer func() {
		file.Close()
	}()

	//git 拉代码
	var git_url string = "cd " + workDir + " \n git clone "
	if len(appVersionJobDto.GitUsername) < 1 {
		git_url = git_url + appVersionJobDto.GitUrl
	} else {
		git_url = git_url + strings.Replace(appVersionJobDto.GitUrl, ":\\", ":\\"+appVersionJobDto.GitUsername+":"+appVersionJobDto.GitPasswd+"@", 1)
	}
	git_url += "\n"
	var build_hook string = "\ncurl -H \"Content-Type: application/json\" -X POST -d '{\"jobId\": \"JOB_ID\"}' \"MASTER_SERVER/app/appVersion/doJobHook\""

	build_hook = strings.Replace(build_hook, "JOB_ID", appVersionJobDto.JobId, 1)
	build_hook = strings.Replace(build_hook, "MASTER_SERVER", "http://127.0.0.1:"+strconv.FormatInt(int64(config.G_AppConfig.Port), 10), 1)

	_, err = file.WriteString(git_url + appVersionJobDto.JobShell + build_hook)

	if err != nil {
		fmt.Print("err=", err.Error())
		return result.Error(err.Error())
	}

	//插入构建记录
	var appVersionJobDetailDto = appVersionJob.AppVersionJobDetailDto{
		JobId:    appVersionJobDto.JobId,
		State:    appVersionJob.STATE_wait,
		LogPath:  path.Join(workDir, appVersionJobDto.JobId+".log"),
		TenantId: user.TenantId,
		DetailId: seq.Generator(),
	}

	err = appVersionJobService.appVersionJobDetailDao.SaveAppVersionJobDetail(appVersionJobDetailDto)
	if err != nil {
		return result.Error(err.Error())
	}

	jobShell := "nohup sh " + dest + " >" + path.Join(workDir, appVersionJobDto.JobId+".log") + " &"
	cmd := exec.Command("bash", "-c", jobShell)
	fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
		appVersionJobDto.State = appVersionJob.STATE_error
		err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	}
	return result.SuccessData(appVersionJobDto)

}

func (appVersionJobService *AppVersionJobService) DoJobHook(ctx iris.Context) interface{} {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDtos, err := appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)

	appVersionJobDto = *appVersionJobDtos[0]

	if len(appVersionJobDtos) < 1 {
		return result.Error("构建不存在")
	}

	//插入构建记录
	var appVersionJobDetailDto = appVersionJob.AppVersionJobDetailDto{
		JobId: appVersionJobDto.JobId,
	}

	appVersionJobDetailDtos, err := appVersionJobService.appVersionJobDetailDao.GetAppVersionJobDetails(appVersionJobDetailDto)
	if len(appVersionJobDetailDtos) < 1 {
		return result.Error("构建日志不存在")
	}
	appVersionJobDetailDto = *appVersionJobDetailDtos[0]
	//插入构建记录
	var appVersionJobImagesDto = appVersionJob.AppVersionJobImagesDto{
		JobId: appVersionJobDto.JobId,
	}
	appVersionJobImagesDtos, _ := appVersionJobService.appVersionJobDao.GetAppVersionJobImages(appVersionJobImagesDto)

	if len(appVersionJobImagesDtos) < 1 {
		return result.Success()
	}

	for _, appVersionJobImagesDto := range appVersionJobImagesDtos {
		appVersionJobService.doGeneratorImages(appVersionJobImagesDto, appVersionJobDetailDto, appVersionJobDto)
	}
	return result.Success()
}

func (appVersionJobService *AppVersionJobService) doGeneratorImages(jobImagesDto *appVersionJob.AppVersionJobImagesDto,
	jobDetailDto appVersionJob.AppVersionJobDetailDto,
	appVersionJobDto appVersionJob.AppVersionJobDto) {
	workDir := path.Join(appVersionJobDto.WorkDir, appVersionJobDto.JobId)
	//判断是否是 /开头
	if !strings.HasPrefix(workDir, "/") {
		workDir = "/" + workDir
	}
	businessJar := path.Join(workDir, jobImagesDto.PackageUrl)

	//查询业务包
	var businessPackageDao softDao.BusinessPackageDao

	businessPackageDto := businessPackage.BusinessPackageDto{
		Id: jobImagesDto.BusinessPackageId,
	}
	businessPackageDtos, _ := businessPackageDao.GetBusinessPackages(businessPackageDto)
	if len(businessPackageDtos) < 1 {
		return
	}

	targetPath := path.Join("/zihao/master/businessPackage", jobImagesDto.TenantId, businessPackageDtos[0].Path)

	input, err := ioutil.ReadFile(businessJar)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(targetPath, input, 0644)
	if err != nil {
		fmt.Println("Error creating", targetPath)
		fmt.Println(err)
		return
	}
	//查询业务包
	var businessDockerfileDao softDao.BusinessDockerfileDao

	businessDockerfileDto := businessDockerfile.BusinessDockerfileDto{
		Id: jobImagesDto.BusinessDockerfileId,
	}
	businessDockerfileDtos, _ := businessDockerfileDao.GetBusinessDockerfiles(businessDockerfileDto)
	if len(businessDockerfileDtos) < 1 {
		return
	}
	//消息队列
	dockerfileQueue.SendData(businessDockerfileDtos[0])
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
