package service

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12"
	appDao "github.com/zihao-boy/zihao/appService/dao"
	dao2 "github.com/zihao-boy/zihao/assets/dao"
	installApp "github.com/zihao-boy/zihao/business/dao/installAppDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/costTime"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/queue/dockerfileQueue"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/composeYaml"
	"github.com/zihao-boy/zihao/entity/dto/host"
	installApp2 "github.com/zihao-boy/zihao/entity/dto/installApp"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type BusinessImagesService struct {
	businessImagesDao     dao.BusinessImagesDao
	businessImagesVerDao  dao.BusinessImagesVerDao
	businessDockerfileDao dao.BusinessDockerfileDao
	installAppDao         installApp.InstallAppDao
	hostDao               dao2.HostDao
	appServiceDao         appDao.AppServiceDao
}

/**
查询 系统信息
*/
func (businessImagesService *BusinessImagesService) GetBusinessImagesAll(businessImagesDto businessImages.BusinessImagesDto) ([]*businessImages.BusinessImagesDto, error) {
	var (
		err                error
		businessImagesDtos []*businessImages.BusinessImagesDto
	)

	businessImagesDtos, err = businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
	if err != nil {
		return nil, err
	}

	return businessImagesDtos, nil

}

/**
查询 系统信息
*/
func (businessImagesService *BusinessImagesService) GetBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err                error
		page               int64
		row                int64
		total              int64
		businessImagesDto  = businessImages.BusinessImagesDto{}
		businessImagesDtos []*businessImages.BusinessImagesDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessImagesDto.Row = row * page

	businessImagesDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesDto.TenantId = user.TenantId
	businessImagesDto.Name = ctx.URLParam("name")
	businessImagesDto.ImagesType = ctx.URLParam("imagesType")
	businessImagesDto.Version = ctx.URLParam("version")
	businessImagesDto.ImagesFlag = ctx.URLParam("imagesFlag")

	total, err = businessImagesService.businessImagesDao.GetBusinessImagesCount(businessImagesDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessImagesDtos, err = businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDtos, total, row)

}

/**
保存 系统信息
*/
func (businessImagesService *BusinessImagesService) SaveBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)
	ctx.SetMaxRequestBodySize(maxSize)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesDto.Id = seq.Generator()
	dest := filepath.Join("/businessImages", user.TenantId, businessImagesDto.Id)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest = filepath.Join(dest, fileHeader.Filename)

	_, err = ctx.SaveFormFile(fileHeader, config.G_AppConfig.DataPath+dest)
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	businessImagesDto.TenantId = user.TenantId
	businessImagesDto.CreateUserId = user.UserId
	businessImagesDto.Version = "V" + date.GetNowAString()
	businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_IMPORT
	businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_CUSTOM
	businessImagesDto.TypeUrl = dest
	businessImagesDto.Name = ctx.FormValue("name")

	err = businessImagesService.businessImagesDao.SaveBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//save images version
	businessImagesVerDto := businessImages.BusinessImagesVerDto{
		Id:       seq.Generator(),
		ImagesId: businessImagesDto.Id,
		Version:  businessImagesDto.Version,
		TypeUrl:  dest,
		TenantId: user.TenantId,
	}
	err = businessImagesService.businessImagesVerDao.SaveBusinessImagesVer(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(businessImagesDto)

}

/**
修改 系统信息
*/
func (businessImagesService *BusinessImagesService) UpdateBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)

	if err = ctx.ReadJSON(&businessImagesDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesService.businessImagesDao.UpdateBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDto)

}

/**
删除 系统信息
*/
func (businessImagesService *BusinessImagesService) DeleteBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)

	if err = ctx.ReadJSON(&businessImagesDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesService.businessImagesDao.DeleteBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDto)

}

/**
删除 系统信息
*/
func (businessImagesService *BusinessImagesService) GeneratorImages(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)
	defer costTime.TimeoutWarning("BusinessImagesService", "GeneratorImages", time.Now())

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	tmpBusinessDockerfileDto := businessDockerfile.BusinessDockerfileDto{
		Id:       businessDockerfileDto.Id,
		TenantId: user.TenantId,
	}
	businessDockerfileDtos, err := businessImagesService.businessDockerfileDao.GetBusinessDockerfiles(tmpBusinessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(businessDockerfileDtos) < 1 {
		return result.Error("dockerfile 不存在")
	}

	doGeneratorImage(businessDockerfileDtos[0], user)

	return result.Success()

}

func doGeneratorImage(businessDockerfileDto *businessDockerfile.BusinessDockerfileDto, user *user.UserDto) {
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	//消息队列
	dockerfileQueue.SendData(businessDockerfileDto)
}

/**
查询镜像 系统信息
*/
func (businessImagesService *BusinessImagesService) GetImagesPool(ctx iris.Context) result.ResultDto {
	var (
		err       error
		resultDto result.ResultDto
	)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	values := ctx.URLParams()

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}
	value := Urlencode(values)
	resp, err := httpReq.Get(config.Remote_Images_Url+"?"+value, headers)
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	datas := resultDto.Data.([]interface{})

	if datas == nil || len(datas) < 1 {
		return resultDto
	}

	installAppDto := installApp2.InstallAppDto{
		TenantId: user.TenantId,
	}
	installAppDtos, _ := businessImagesService.installAppDao.GetInstallApps(installAppDto)

	if installAppDtos == nil || len(installAppDtos) < 1 {
		return resultDto
	}
	for _, data := range datas {
		freshAppState(&data, installAppDtos)
	}

	return resultDto

}

func freshAppState(data *interface{}, installAppDtos []*installApp2.InstallAppDto) {
	dataMap := (*data).(map[string]interface{})
	extAppId := dataMap["appId"].(string)
	for _, installAppDto := range installAppDtos {
		if installAppDto.ExtAppId == extAppId {
			dataMap["state"] = "001"
		}
	}
}

//安装镜像
func (businessImagesService *BusinessImagesService) InstallImages(ctx iris.Context) result.ResultDto {

	var (
		err               error
		resultDto         result.ResultDto
		imagesPoolsDtos   []businessImages.ImagesPoolsDto
		message           string = ""
		installAppPageDto installApp2.InstallAppPageDto
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	if err = ctx.ReadJSON(&installAppPageDto); err != nil {
		return result.Error("解析入参失败")
	}

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}

	resp, err := httpReq.Get(config.Remote_Images_Url+"?page=1&row=1&appId="+installAppPageDto.AppId, headers)
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	data, _ := json.Marshal(resultDto.Data)
	json.Unmarshal(data, &imagesPoolsDtos)

	if len(imagesPoolsDtos) < 1 {
		return result.Error("未查询到应用")
	}

	for _, zihaoAppImagesDto := range imagesPoolsDtos[0].ZihaoAppImagesDtos {
		//if exits images
		businessImagesDto := businessImages.BusinessImagesDto{}
		businessImagesDto.TenantId = user.TenantId
		businessImagesDto.Name = zihaoAppImagesDto.Name
		//
		businessImagesDtos, _ := businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
		if len(businessImagesDtos) < 1 { // no exits images
			businessImagesDto.CreateUserId = user.UserId
			businessImagesDto.Id = seq.Generator()
			businessImagesDto.Version = imagesPoolsDtos[0].Version
			businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_REMOTE
			businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_PUBLIC
			businessImagesDto.TypeUrl = zihaoAppImagesDto.Url
			err = businessImagesService.businessImagesDao.SaveBusinessImages(businessImagesDto)
		} else {
			businessImagesDto.Id = businessImagesDtos[0].Id
			businessImagesDto.Version = imagesPoolsDtos[0].Version
			businessImagesDto.TypeUrl = zihaoAppImagesDto.Url
			err = businessImagesService.businessImagesDao.UpdateBusinessImages(businessImagesDto)
			businessImagesDto = *businessImagesDtos[0]
		}

		//if exits images version
		businessImagesVerDto := businessImages.BusinessImagesVerDto{
			TenantId: user.TenantId,
			ImagesId: businessImagesDto.Id,
			Version:  imagesPoolsDtos[0].Version,
		}
		businessImagesVerDtos, _ := businessImagesService.businessImagesVerDao.GetBusinessImagesVers(businessImagesVerDto)

		if len(businessImagesVerDtos) > 0 {
			//message += (businessImagesDto.Name + ":" + imagesPoolsDtos[0].Version + "已存在，")
			continue
		}

		businessImagesVerDto = businessImages.BusinessImagesVerDto{
			Id:       seq.Generator(),
			ImagesId: businessImagesDto.Id,
			Version:  imagesPoolsDtos[0].Version,
			TypeUrl:  zihaoAppImagesDto.Url,
			TenantId: user.TenantId,
		}
		err = businessImagesService.businessImagesVerDao.SaveBusinessImagesVer(businessImagesVerDto)
		if err != nil {
			return result.Error(err.Error())
		}

	}
	if message != "" {
		return result.Error(message)
	}
	// save app info
	resultDto = businessImagesService.installApp(installAppPageDto, imagesPoolsDtos[0].Compose, user)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	installAppDto := installApp2.InstallAppDto{
		AppId:        seq.Generator(),
		AppName:      imagesPoolsDtos[0].Name,
		Version:      imagesPoolsDtos[0].Version,
		TenantId:     user.TenantId,
		ExtAppId:     imagesPoolsDtos[0].AppId,
		CreateUserId: user.UserId,
	}
	businessImagesService.installAppDao.SaveInstallApp(installAppDto)

	return resultDto

}

//install App
func (businessImagesService *BusinessImagesService) installApp(installAppPageDto installApp2.InstallAppPageDto,
	compose string,
	user *user.UserDto) result.ResultDto {
	var (
		composeYamlDto = composeYaml.ComposeYamlZiHaoDto{
		}
		serviceName   string
		appServiceDto appService.AppServiceDto
	)
	// if is empty
	if utils.IsEmpty(compose) {
		return result.Success()
	}
	yaml.Unmarshal([]byte(compose), &composeYamlDto)

	if len(composeYamlDto.Services) < 1 {
		return result.Error("yml文件中不包含应用信息")
	}

	for _, service := range composeYamlDto.Services {
		serviceMap := service.(map[string]interface{})

		for key := range serviceMap {
			serviceName = key
		}
		//
		//verId := seq.Generator()
		appServiceDto = appService.AppServiceDto{
			AsId:         seq.Generator(),
			AsName:       serviceName,
			AsType:       installAppPageDto.AsType,
			TenantId:     user.TenantId,
			AsDesc:       serviceName,
			State:        "10012",
			AsCount:      "1",
			AsGroupId:    installAppPageDto.AsGroupId,
			AsDeployType: installAppPageDto.AsDeployType,
			AsDeployId:   installAppPageDto.AsDeployId,
			//ImagesId: imagesId,
			//VerId:verId,
		}
		resultDto := businessImagesService.doImportAppService(appServiceDto, serviceMap[serviceName], user)
		if resultDto.Code != result.CODE_SUCCESS {
			return resultDto
		}
	}

	// zihao cmd
	if utils.IsEmpty(composeYamlDto.ZihaoCmd) {
		return result.Success()
	}

	var (
		hosts []*host.HostDto
	)

	if installAppPageDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST {
		hostDto := host.HostDto{
			HostId: installAppPageDto.AsDeployId,
		}
		hosts, _ = businessImagesService.hostDao.GetHosts(hostDto)
	} else {
		hostDto := host.HostDto{
			GroupId: installAppPageDto.AsDeployId,
		}
		hosts, _ = businessImagesService.hostDao.GetHosts(hostDto)
	}
	for _, host := range hosts {
		shell.ExecCommonShell(*host, composeYamlDto.ZihaoCmd)
	}

	return result.Success()
}

func (businessImagesService *BusinessImagesService) doImportAppService(appServiceDto appService.AppServiceDto, info interface{}, user *user.UserDto) result.ResultDto {

	var (
		serviceDto composeYaml.ServiceDto
		err        error
	)
	//objectConvert.Map2Struct(info.(map[string]interface{}), &serviceDto)
	data, _ := json.Marshal(info)
	json.Unmarshal(data, &serviceDto)
	//get images
	businessImagesDto := businessImages.BusinessImagesDto{
		Name:     appServiceDto.AsName,
		TenantId: appServiceDto.TenantId,
	}
	images, _ := businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)

	if images == nil || len(images) < 1 {
		return result.Error("镜像不存在")
	}

	businessImagesVerDto := businessImages.BusinessImagesVerDto{
		TenantId: user.TenantId,
		ImagesId: businessImagesDto.Id,
		Version:  images[0].Version,
	}
	businessImagesVerDtos, _ := businessImagesService.businessImagesVerDao.GetBusinessImagesVers(businessImagesVerDto)

	if businessImagesVerDtos == nil || len(businessImagesVerDtos) < 1 {
		return result.Error("镜像版本不存在")
	}
	appServiceDto.ImagesId = images[0].Id
	appServiceDto.VerId = businessImagesVerDtos[0].Id

	//save app service
	err = businessImagesService.appServiceDao.SaveAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(serviceDto.Ports) > 0 {
		for _, port := range serviceDto.Ports {
			if !strings.Contains(port, ":") {
				continue
			}
			appServicePort := appService.AppServicePortDto{
				PortId:     seq.Generator(),
				AsId:       appServiceDto.AsId,
				TenantId:   appServiceDto.TenantId,
				SrcPort:    strings.Split(port, ":")[0],
				TargetPort: strings.Split(port, ":")[1],
			}
			businessImagesService.appServiceDao.SaveAppServicePort(appServicePort)
		}
	}

	if len(serviceDto.ExtraHosts) > 0 {
		for _, host := range serviceDto.ExtraHosts {
			if !strings.Contains(host, ":") {
				continue
			}
			appServiceHosts := appService.AppServiceHostsDto{
				HostsId:  seq.Generator(),
				AsId:     appServiceDto.AsId,
				TenantId: appServiceDto.TenantId,
				Hostname: strings.Split(host, ":")[0],
				Ip:       strings.Split(host, ":")[1],
			}
			businessImagesService.appServiceDao.SaveAppServiceHosts(appServiceHosts)
		}
	}

	if len(serviceDto.Volumes) > 0 {
		for _, volume := range serviceDto.Volumes {
			if !strings.Contains(volume, ":") {
				continue
			}
			appServiceDir := appService.AppServiceDirDto{
				DirId:     seq.Generator(),
				AsId:      appServiceDto.AsId,
				TenantId:  appServiceDto.TenantId,
				SrcDir:    strings.Split(volume, ":")[0],
				TargetDir: strings.Split(volume, ":")[1],
			}
			businessImagesService.appServiceDao.SaveAppServiceDir(appServiceDir)
		}
	}

	if len(serviceDto.Environment) > 0 {
		for _, env := range serviceDto.Environment {
			if !strings.Contains(env, ":") {
				continue
			}
			appServiceVar := appService.AppServiceVarDto{
				AvId:     seq.Generator(),
				AsId:     appServiceDto.AsId,
				TenantId: appServiceDto.TenantId,
				VarSpec:  strings.Split(env, ":")[0],
				VarValue: strings.Split(env, ":")[1],
				VarName:  strings.Split(env, ":")[0],
			}
			businessImagesService.appServiceDao.SaveAppServiceVar(appServiceVar)
		}
	}

	return result.Success()
}

// uninstall app
func (businessImagesService *BusinessImagesService) UninstallImages(ctx iris.Context) result.ResultDto {

	var (
		err               error
		resultDto         result.ResultDto
		imagesPoolsDtos   []businessImages.ImagesPoolsDto
		message           string = ""
		installAppPageDto installApp2.InstallAppPageDto
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	if err = ctx.ReadJSON(&installAppPageDto); err != nil {
		return result.Error("解析入参失败")
	}

	if utils.IsEmpty(installAppPageDto.AppId) {
		return result.Error("未包含appId")
	}

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}

	resp, err := httpReq.Get(config.Remote_Images_Url+"?page=1&row=1&appId="+installAppPageDto.AppId, headers)
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	data, _ := json.Marshal(resultDto.Data)
	json.Unmarshal(data, &imagesPoolsDtos)

	if len(imagesPoolsDtos) < 1 {
		return result.Error("未查询到应用")
	}
	// no delete
	//for _, zihaoAppImagesDto := range imagesPoolsDtos[0].ZihaoAppImagesDtos {
	//	//if exits images
	//	businessImagesDto := businessImages.BusinessImagesDto{}
	//	businessImagesDto.TenantId = user.TenantId
	//	businessImagesDto.Name = zihaoAppImagesDto.Name
	//	//
	//	businessImagesDtos, _ := businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
	//	if len(businessImagesDtos) > 1 { // no exits images
	//		businessImagesDto.Id = businessImagesDtos[0].Id
	//		err = businessImagesService.businessImagesDao.DeleteBusinessImages(businessImagesDto)
	//	}
	//	//if exits images version
	//	businessImagesVerDto := businessImages.BusinessImagesVerDto{
	//		TenantId: user.TenantId,
	//		ImagesId: businessImagesDto.Id,
	//		Version:  imagesPoolsDtos[0].Version,
	//	}
	//	businessImagesVerDtos, _ := businessImagesService.businessImagesVerDao.GetBusinessImagesVers(businessImagesVerDto)
	//	if len(businessImagesVerDtos) < 1 {
	//		//message += (businessImagesDto.Name + ":" + imagesPoolsDtos[0].Version + "已存在，")
	//		continue
	//	}
	//	businessImagesVerDto = businessImages.BusinessImagesVerDto{
	//		Id:       businessImagesVerDtos[0].Id,
	//		TenantId: user.TenantId,
	//	}
	//	err = businessImagesService.businessImagesVerDao.DeleteBusinessImagesVer(businessImagesVerDto)
	//	if err != nil {
	//		return result.Error(err.Error())
	//	}
	//}
	if message != "" {
		return result.Error(message)
	}
	// delete app info
	resultDto = businessImagesService.uninstallApp(installAppPageDto, imagesPoolsDtos[0].Compose, user)

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	installAppDto := installApp2.InstallAppDto{
		TenantId: user.TenantId,
		ExtAppId: installAppPageDto.AppId,
	}

	businessImagesService.installAppDao.DeleteInstallApp(installAppDto)

	return resultDto

}

//install App
func (businessImagesService *BusinessImagesService) uninstallApp(installAppPageDto installApp2.InstallAppPageDto,
	compose string,
	user *user.UserDto) result.ResultDto {
	var (
		composeYamlDto = composeYaml.ComposeYamlZiHaoDto{
		}
		serviceName          string
		appServiceDto        appService.AppServiceDto
		deleteAppServiceDtos []*appService.AppServiceDto
	)
	// if is empty
	if utils.IsEmpty(compose) {
		return result.Success()
	}
	yaml.Unmarshal([]byte(compose), &composeYamlDto)

	if len(composeYamlDto.Services) < 1 {
		return result.Error("yml文件中不包含应用信息")
	}

	for _, service := range composeYamlDto.Services {
		serviceMap := service.(map[string]interface{})

		for key := range serviceMap {
			serviceName = key
		}
		//
		//verId := seq.Generator()
		appServiceDto = appService.AppServiceDto{
			AsName:   serviceName,
			TenantId: user.TenantId,
		}
		appServiceDtos, _ := businessImagesService.appServiceDao.GetAppServices(appServiceDto)
		if appServiceDtos == nil || len(appServiceDtos) < 1 {
			continue
		}
		if appServiceDtos[0].State != appService.STATE_STOP {
			return result.Error("请先关闭应用（" + appServiceDtos[0].AsName + ")")
		}

		deleteAppServiceDtos = append(deleteAppServiceDtos, appServiceDtos[0])
	}

	if len(deleteAppServiceDtos) < 1 {
		return result.Error("没有可卸载的应用")
	}

	for _, service := range deleteAppServiceDtos {
		resultDto := businessImagesService.doDeleteAppService(*service)
		if resultDto.Code != result.CODE_SUCCESS {
			return resultDto
		}
	}

	return result.Success()
}

// delete app service

func (businessImagesService *BusinessImagesService) doDeleteAppService(appServiceDto appService.AppServiceDto) result.ResultDto {

	var (
		err error
	)

	//delete app service
	err = businessImagesService.appServiceDao.DeleteAppService(appServiceDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.Success()
}

func Urlencode(data map[string]string) string {
	var buf bytes.Buffer
	for k, v := range data {
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0 : len(s)-1]
}
