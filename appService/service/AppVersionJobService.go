package service

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/costTime"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/notifyMessage"
	"github.com/zihao-boy/zihao/common/queue/dockerfileQueue"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

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
		err                error
		appVersionJobParam appVersionJob.AppVersionJobParam
	)

	if err = ctx.ReadJSON(&appVersionJobParam); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	return appVersionJobService.commonJob(appVersionJobParam,user)

}

func (appVersionJobService *AppVersionJobService) commonJob(appVersionJobParam appVersionJob.AppVersionJobParam,user *user.UserDto) result.ResultDto{

	appVersionJobDto := appVersionJob.AppVersionJobDto{
	}
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.JobId = appVersionJobParam.JobId
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
	if !utils.IsDir(workDir) {
		//err = os.RemoveAll(workDir)
		utils.CreateDir(workDir)
	}

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
	var git_url string = ""
	if appVersionJobDto.GitUsername == "" || strings.Trim(appVersionJobDto.GitUsername, " ") == "无" {
		git_url = appVersionJobDto.GitUrl
	} else {
		git_url = strings.Replace(appVersionJobDto.GitUrl,
			"://",
			"://"+strings.ReplaceAll(appVersionJobDto.GitUsername, "@", "%40")+":"+strings.ReplaceAll(appVersionJobDto.GitPasswd, "@", "%40")+"@",
			1)
	}

	if !utils.IsDir(path.Join(workDir, "job")) {
		git_url = "cd " + workDir + " \n git clone " + git_url + " job \n cd job"
	} else {
		git_url = "cd " + path.Join(workDir, "job") + "\n git pull " + git_url
	}

	git_url += "\n"
	var build_hook string = "\ncurl -H \"Content-Type: application/json\" -X POST -d '{\"jobId\": \"JOB_ID\",\"action\":\"ACTION\",\"images\":\"IMAGES\"}' \"MASTER_SERVER/app/appVersion/doJobHook\""

	build_hook = strings.Replace(build_hook, "JOB_ID", appVersionJobDto.JobId, 1)
	build_hook = strings.Replace(build_hook, "ACTION", appVersionJobParam.Action, 1)
	build_hook = strings.Replace(build_hook, "IMAGES", appVersionJobParam.Images, 1)
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
		fmt.Println(err)
		appVersionJobDto.State = appVersionJob.STATE_error
		err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
		return result.Error(err.Error())
	}

	jobShell := "nohup sh " + dest + " >" + path.Join(workDir, appVersionJobDto.JobId+".log") + " &"
	notifyMessage.SendMsg(user.TenantId, "开始构建>"+appVersionJobDto.JobName)

	go shell.ExecLocalShell(jobShell)

	return result.SuccessData(appVersionJobDto)
}

// do job build
func (appVersionJobService *AppVersionJobService) DoJobHook(ctx iris.Context) interface{} {
	var (
		err                error
		appVersionJobParam appVersionJob.AppVersionJobParam
	)

	if err = ctx.ReadJSON(&appVersionJobParam); err != nil {
		return result.Error("解析入参失败")
	}
	appVersionJobDto := appVersionJob.AppVersionJobDto{
		JobId: appVersionJobParam.JobId,
	}
	appVersionJobDtos, err := appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	appVersionJobDto = *appVersionJobDtos[0]

	if len(appVersionJobDtos) < 1 {
		appVersionJobService.updateAppVersionJobState(appVersionJobDto, appVersionJob.STATE_error)
		return result.Error("构建不存在")
	}
	notifyMessage.SendMsg(appVersionJobDto.TenantId, "代码编译完成>"+appVersionJobDto.JobName)

	//插入构建记录
	var appVersionJobDetailDto = appVersionJob.AppVersionJobDetailDto{
		JobId: appVersionJobDto.JobId,
	}

	appVersionJobDetailDtos, err := appVersionJobService.appVersionJobDetailDao.GetAppVersionJobDetails(appVersionJobDetailDto)
	if len(appVersionJobDetailDtos) < 1 {
		appVersionJobService.updateAppVersionJobState(appVersionJobDto, appVersionJob.STATE_error)
		return result.Error("构建日志不存在")
	}
	appVersionJobDetailDto = *appVersionJobDetailDtos[0]
	// save build log
	var appVersionJobImagesDto = appVersionJob.AppVersionJobImagesDto{
		JobId: appVersionJobDto.JobId,
	}
	appVersionJobImagesDtos, _ := appVersionJobService.appVersionJobDao.GetAppVersionJobImages(appVersionJobImagesDto)

	if len(appVersionJobImagesDtos) < 1 {
		appVersionJobService.updateAppVersionJobState(appVersionJobDto, appVersionJob.STATE_success)
		return result.Success()
	}

	for _, appVersionJobImagesDto := range appVersionJobImagesDtos {
		if !hasAppVersionJob(appVersionJobImagesDto, appVersionJobParam) {
			continue
		}
		appVersionJobImagesDto.Action = appVersionJobParam.Action
		appVersionJobService.doGeneratorImages(appVersionJobImagesDto, appVersionJobDetailDto, appVersionJobDto)
	}

	//appVersionJobDto.State = appVersionJob.STATE_success
	//err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	//if err != nil {
	//	appVersionJobService.updateAppVersionJobState(appVersionJobDto.JobId,appVersionJob.STATE_error)
	//	return result.Error(err.Error())
	//}
	appVersionJobService.updateAppVersionJobState(appVersionJobDto, appVersionJob.STATE_success)
	return result.Success()
}

// container images
func hasAppVersionJob(jobImagesDto *appVersionJob.AppVersionJobImagesDto, jobParam appVersionJob.AppVersionJobParam) bool {
	for _, images := range strings.Split(jobParam.Images, ",") {
		if images == jobImagesDto.JobImagesId {
			return true
		}
	}

	return false
}

func (appVersionJobService *AppVersionJobService) updateAppVersionJobState(job appVersionJob.AppVersionJobDto, state string) {
	var (
		appVersionJobDto appVersionJob.AppVersionJobDto
	)
	appVersionJobDto.JobId = job.JobId
	appVersionJobDto.TenantId = job.TenantId
	appVersionJobDto.State = state
	appVersionJobDto.JobTime= date.GetNowTimeString()
	appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
}

// to do generator images
func (appVersionJobService *AppVersionJobService) doGeneratorImages(jobImagesDto *appVersionJob.AppVersionJobImagesDto,
	jobDetailDto appVersionJob.AppVersionJobDetailDto,
	appVersionJobDto appVersionJob.AppVersionJobDto) {

	defer costTime.TimeoutWarning("AppVersionJobService", "doGeneratorImages", time.Now())

	workDir := path.Join(appVersionJobDto.WorkDir, appVersionJobDto.JobId)
	//判断是否是 /开头
	if !strings.HasPrefix(workDir, "/") {
		workDir = "/" + workDir
	}

	workDir = path.Join(workDir, "job")
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
	if !utils.IsFile(targetPath) {
		targetPathDir := path.Dir(targetPath)
		if !utils.IsDir(targetPathDir) {
			os.MkdirAll(targetPathDir, 0777)
		}
		file, _ := os.Create(targetPath)
		defer func() {
			file.Close()
		}()
	}
	err = ioutil.WriteFile(targetPath, input, 0777)
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
	businessDockerfileDtos[0].LogPath = jobDetailDto.LogPath
	// action
	businessDockerfileDtos[0].Action = jobImagesDto.Action
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

func (appVersionJobService *AppVersionJobService) GetJobLog(ctx iris.Context) interface{} {
	var (
		appVersionJobDetailDto = appVersionJob.AppVersionJobDetailDto{}
	)

	appVersionJobDetailDto.Row = 1
	appVersionJobDetailDto.Page = 0

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobDetailDto.TenantId = user.TenantId
	appVersionJobDetailDto.JobId = ctx.URLParam("jobId")

	appVersionJobDetailDtos, _ := appVersionJobService.appVersionJobDetailDao.GetAppVersionJobDetails(appVersionJobDetailDto)

	if len(appVersionJobDetailDtos) < 1 {
		return result.Error("没有日志")
	}

	return result.SuccessData(appVersionJobDetailDtos[0].LogPath)
}

// webhooks
// add by wuxw 2022-01-05
func (appVersionJobService *AppVersionJobService) Payload(ctx iris.Context) interface{} {

	var (
		imagesIds string
	)

	//come from
	event := ctx.URLParam("event")
	//job Id
	jobId := ctx.URLParam("jobId")

	appVersionJobDto := appVersionJob.AppVersionJobDto{
		JobId: jobId,
	}
	appVersionJobDtos, _ := appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if len(appVersionJobDtos) < 1 {
		return result.Error("构建不存在")
	}

	// save build log
	var appVersionJobImagesDto = appVersionJob.AppVersionJobImagesDto{
		JobId: appVersionJobDto.JobId,
	}
	appVersionJobImagesDtos, _ := appVersionJobService.appVersionJobDao.GetAppVersionJobImages(appVersionJobImagesDto)

	if len(appVersionJobImagesDtos) > 0{
		for _,appversionJobImagesDto := range appVersionJobImagesDtos{
			imagesIds += (appversionJobImagesDto.JobImagesId+",")
		}
	}

	if strings.HasSuffix(imagesIds,","){
		imagesIds = imagesIds[0:len(imagesIds)-1]
	}


	appVersionJobParam := appVersionJob.AppVersionJobParam{
		JobId: jobId,
		Images: imagesIds,
	}

	user := user.UserDto{
		TenantId: appVersionJobDtos[0].TenantId,
		UserId: "1",
	}

	if event == appVersionJob.EVENT_PUSH{
		appVersionJobParam.Action = businessDockerfile.ActionBuild
		return appVersionJobService.commonJob(appVersionJobParam,&user)
	}else if event == appVersionJob.EVENT_PUSH_AND_RESTART{
		appVersionJobParam.Action = businessDockerfile.ActionBuildStart
		return appVersionJobService.commonJob(appVersionJobParam,&user)
	}

	return result.Error("no support")
}
