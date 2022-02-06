package service

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12"
	installApp "github.com/zihao-boy/zihao/business/dao/installAppDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/costTime"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/queue/dockerfileQueue"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	installApp2 "github.com/zihao-boy/zihao/entity/dto/installApp"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
	"path/filepath"
	"strconv"
	"time"
)

type BusinessImagesService struct {
	businessImagesDao     dao.BusinessImagesDao
	businessImagesVerDao  dao.BusinessImagesVerDao
	businessDockerfileDao dao.BusinessDockerfileDao
	installAppDao         installApp.InstallAppDao
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
		err             error
		resultDto       result.ResultDto
		imagesPoolsDto  businessImages.ImagesPoolsDto
		imagesPoolsDtos []businessImages.ImagesPoolsDto
		message         string = ""
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	if err = ctx.ReadJSON(&imagesPoolsDto); err != nil {
		return result.Error("解析入参失败")
	}

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}

	resp, err := httpReq.Get(config.Remote_Images_Url+"?page=1&row=1&appId="+imagesPoolsDto.AppId, headers)
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
			message += (businessImagesDto.Name + ":" + imagesPoolsDtos[0].Version + "已存在，")
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
