package appPublisherService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/business/dao/appPublisherDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/config"
	appPublisher "github.com/zihao-boy/zihao/entity/dto/appPublisherDto"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/composeYaml"
	installApp2 "github.com/zihao-boy/zihao/entity/dto/installApp"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"gopkg.in/yaml.v3"
	"strconv"
)

type AppPublisherService struct {
	appPublisherDao appPublisherDao.AppPublisherDao
	appServiceDao   dao.AppServiceDao
}

// get db link
// all db by this user
func (appPublisherService *AppPublisherService) GetAppPublisherAll(AppPublisherDto appPublisher.AppPublisherDto) ([]*appPublisher.AppPublisherDto, error) {
	var (
		err              error
		AppPublisherDtos []*appPublisher.AppPublisherDto
	)

	AppPublisherDtos, err = appPublisherService.appPublisherDao.GetAppPublishers(AppPublisherDto)
	if err != nil {
		return nil, err
	}

	return AppPublisherDtos, nil

}

/**
查询 系统信息
*/
func (appPublisherService *AppPublisherService) GetAppPublishers(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		appPublisherDto  = appPublisher.AppPublisherDto{}
		appPublisherDtos []*appPublisher.AppPublisherDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appPublisherDto.Row = row * page

	appPublisherDto.Page = (page - 1) * row

	total, err = appPublisherService.appPublisherDao.GetAppPublisherCount(appPublisherDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appPublisherDtos, err = appPublisherService.appPublisherDao.GetAppPublishers(appPublisherDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appPublisherDtos, total, row)

}

/**
保存 系统信息
*/
func (appPublisherService *AppPublisherService) SaveAppPublishers(ctx iris.Context) result.ResultDto {
	var (
		err             error
		appPublisherDto appPublisher.AppPublisherDto
		resultDto       result.ResultDto
	)
	if err = ctx.ReadJSON(&appPublisherDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appPublisherDto.TenantId = user.TenantId

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}

	resp, err := httpReq.SendRequest(config.Remote_Save_Publisher, appPublisherDto, headers, "POST")
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	remoteData := resultDto.Data.(map[string]interface{})

	appPublisherDto.PublisherId = seq.Generator()
	appPublisherDto.State = appPublisher.StateNormal
	appPublisherDto.ExtPublisherId = remoteData["publisherId"].(string)
	appPublisherDto.Token = remoteData["token"].(string)

	err = appPublisherService.appPublisherDao.SaveAppPublisher(appPublisherDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appPublisherDto)

}

/**
修改 系统信息
*/
func (appPublisherService *AppPublisherService) UpdateAppPublishers(ctx iris.Context) result.ResultDto {
	var (
		err             error
		appPublisherDto appPublisher.AppPublisherDto
		resultDto       result.ResultDto
		publisherId     string
	)
	if err = ctx.ReadJSON(&appPublisherDto); err != nil {
		return result.Error("解析入参失败")
	}

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}

	publisherId = appPublisherDto.PublisherId
	appPublisherDto.PublisherId = appPublisherDto.ExtPublisherId
	resp, err := httpReq.SendRequest(config.Remote_Update_Publisher, appPublisherDto, headers, "POST")
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	remoteData := resultDto.Data.(map[string]interface{})
	appPublisherDto.PublisherId = publisherId
	appPublisherDto.Token = remoteData["token"].(string)

	err = appPublisherService.appPublisherDao.UpdateAppPublisher(appPublisherDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appPublisherDto)

}

/**
删除 系统信息
*/
func (appPublisherService *AppPublisherService) DeleteAppPublishers(ctx iris.Context) result.ResultDto {
	var (
		err             error
		appPublisherDto appPublisher.AppPublisherDto
	)
	if err = ctx.ReadJSON(&appPublisherDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appPublisherService.appPublisherDao.DeleteAppPublisher(appPublisherDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appPublisherDto)

}

//apply publish app
func (appPublisherService *AppPublisherService) ApplyPublishApp(ctx iris.Context) result.ResultDto {
	var (
		err                error
		applyPublishAppDto appPublisher.ApplyPublishAppDto
		serviceDto         composeYaml.ServiceDto
		services           []interface{}
		composeYamlDto     = composeYaml.ComposeYamlDto{
			Version: "2",
		}
		resultDto result.ResultDto
	)
	if err = ctx.ReadJSON(&applyPublishAppDto); err != nil {
		return result.Error("解析入参失败")
	}

	// get publisher id

	appPublisherDto := appPublisher.AppPublisherDto{
		PublisherId: applyPublishAppDto.PublisherId,
	}

	appPublisherDtos, err := appPublisherService.appPublisherDao.GetAppPublishers(appPublisherDto)

	if len(appPublisherDtos) < 1 {
		return result.Error("发布者不存在")
	}

	applyPublishAppDto.PublisherId = appPublisherDtos[0].ExtPublisherId

	for _, appS := range applyPublishAppDto.Apps {

		serviceDto = appPublisherService.getServiceDto(appS)
		servicesDto := composeYaml.ServicesDto{
			Name:    appS.AsName,
			Service: serviceDto,
		}
		services = append(services, servicesDto.ToMap())
		composeYamlDto.Services = services
	}

	data, _ := yaml.Marshal(composeYamlDto)

	applyPublishAppDto.Compose = string(data)

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}
	mjson, _ := json.Marshal(applyPublishAppDto)
	dataEn, _ := encrypt.Encrypt(appPublisherDtos[0].Token, string(mjson))

	applyAppDto := installApp2.ApplyAppDto{
		PublisherId: appPublisherDtos[0].ExtPublisherId,
		Data:        dataEn,
	}
	resp, err := httpReq.SendRequest(config.Remote_Apply_Publish_App_Url, applyAppDto, headers, "POST")
	if err != nil {
		return result.Error(err.Error())
	}

	//resps, _ := encrypt.Decrypt(appPublisherDtos[0].Token, string(resp))
	json.Unmarshal(resp, &resultDto)
	return resultDto
}

func (appPublisherService *AppPublisherService) getServiceDto(appServiceDto *appService.AppServiceDto) composeYaml.ServiceDto {
	var (
		serviceDto composeYaml.ServiceDto
	)

	serviceDto.Image = appServiceDto.ImagesUrl

	portDto := appService.AppServicePortDto{
		AsId: appServiceDto.AsId,
	}
	portDtos, _ := appPublisherService.appServiceDao.GetAppServicePort(portDto)

	if len(portDtos) > 0 {
		for _, port := range portDtos {
			portStr := port.SrcPort + ":" + port.TargetPort
			serviceDto.Ports = append(serviceDto.Ports, portStr)
		}
	}

	hostsDto := appService.AppServiceHostsDto{
		AsId: appServiceDto.AsId,
	}
	hostsDtos, _ := appPublisherService.appServiceDao.GetAppServiceHosts(hostsDto)

	if len(hostsDtos) > 0 {
		for _, host := range hostsDtos {
			hostStr := host.Hostname + ":" + host.Ip
			serviceDto.ExtraHosts = append(serviceDto.ExtraHosts, hostStr)
		}
	}
	dirDto := appService.AppServiceDirDto{
		AsId: appServiceDto.AsId,
	}
	dirDtos, _ := appPublisherService.appServiceDao.GetAppServiceDir(dirDto)

	if len(dirDtos) > 0 {
		for _, dir := range dirDtos {
			dirStr := dir.SrcDir + ":" + dir.TargetDir
			serviceDto.Volumes = append(serviceDto.Volumes, dirStr)
		}
	}

	varDto := appService.AppServiceVarDto{
		AsId: appServiceDto.AsId,
	}
	varDtos, _ := appPublisherService.appServiceDao.GetAppServiceVars(varDto)
	if len(varDtos) > 0 {
		for _, vari := range varDtos {
			variStr := vari.VarSpec + ":" + vari.VarValue
			serviceDto.Environment = append(serviceDto.Environment, variStr)
		}
	}

	return serviceDto
}
