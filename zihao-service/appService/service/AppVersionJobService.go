package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/appService/dao"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/date"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVersionJob"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"os"
	"strconv"
)

type AppVersionJobService struct {
	appVersionJobDao dao.AppVersionJobDao

}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobAll(appVersionJobDto appVersionJob.AppVersionJobDto)  ([]*appVersionJob.AppVersionJobDto,error) {
	var (
		err       error
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)

	appVersionJobDtos,err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if(err != nil){
		return nil,err
	}

	return appVersionJobDtos,nil

}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobs(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		page int64
		row int64
		total int64
		appVersionJobDto = appVersionJob.AppVersionJobDto{}
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)


	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	appVersionJobDto.Row = row * page

	appVersionJobDto.Page = (page -1) * row

	appVersionJobDto.JobId = ctx.URLParam("jobId")
	appVersionJobDto.JobName = ctx.URLParam("jobName")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobDto.TenantId = user.TenantId

	total,err = appVersionJobService.appVersionJobDao.GetAppVersionJobCount(appVersionJobDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}

	appVersionJobDtos,err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDtos,total,row)

}


/**
保存 系统信息
*/
func (appVersionJobService *AppVersionJobService) SaveAppVersionJobs(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		appVersionJobDto appVersionJob.AppVersionJobDto

	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.JobId = seq.Generator()
	appVersionJobDto.State = appVersionJob.STATE_wait
	appVersionJobDto.PreJobTime = date.GetNowTimeString()
	appVersionJobDto.CurJobTime = date.GetNowTimeString()

	err = appVersionJobService.appVersionJobDao.SaveAppVersionJob(appVersionJobDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	var pathFile string = "~/.zihao/"+appVersionJobDto.JobId+"/"+appVersionJobDto.JobId+".sh"

	//当前用户目录下生成 文件夹
	file,er:=os.OpenFile(pathFile,os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer func(){file.Close()}()
	if er!=nil && os.IsNotExist(err){
		file,_ = os.Create(pathFile)
	}

	file.WriteString(appVersionJobDto.JobShell)

	return result.SuccessData(appVersionJobDto)

}


/**
修改 系统信息
*/
func (appVersionJobService *AppVersionJobService) UpdateAppVersionJobs(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if(err != nil){
		return result.Error(err.Error())
	}


	var pathFile string = "~/.zihao/"+appVersionJobDto.JobId+"/"+appVersionJobDto.JobId+".sh"

	//当前用户目录下生成 文件夹
	file,er:=os.OpenFile(pathFile,os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer func(){file.Close()}()
	if er!=nil && os.IsNotExist(err){
		file,_ = os.Create(pathFile)
	}

	file.WriteString(appVersionJobDto.JobShell)

	return result.SuccessData(appVersionJobDto)

}


/**
删除 系统信息
*/
func (appVersionJobService *AppVersionJobService) DeleteAppVersionJobs(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJob(appVersionJobDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDto)

}