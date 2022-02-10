package service

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/appPublisherDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/config"
	appPublisher "github.com/zihao-boy/zihao/entity/dto/appPublisherDto"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	installApp2 "github.com/zihao-boy/zihao/entity/dto/installApp"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
	"strconv"
)

type BusinessImagesVerService struct {
	businessImagesVerDao     dao.BusinessImagesVerDao
	businessDockerfileDao dao.BusinessDockerfileDao

	appPublisherDao appPublisherDao.AppPublisherDao
}

/**
查询 系统信息
*/
func (businessImagesVerService *BusinessImagesVerService) GetBusinessImagesVerAll(businessImagesVerDto businessImages.BusinessImagesVerDto) ([]*businessImages.BusinessImagesVerDto, error) {
	var (
		err                error
		businessImagesVerDtos []*businessImages.BusinessImagesVerDto
	)

	businessImagesVerDtos, err = businessImagesVerService.businessImagesVerDao.GetBusinessImagesVers(businessImagesVerDto)
	if err != nil {
		return nil, err
	}

	return businessImagesVerDtos, nil

}

/**
查询 系统信息
*/
func (businessImagesVerService *BusinessImagesVerService) GetBusinessImagesVer(ctx iris.Context) result.ResultDto {
	var (
		err                error
		page               int64
		row                int64
		total              int64
		businessImagesVerDto  = businessImages.BusinessImagesVerDto{}
		businessImagesVerDtos []*businessImages.BusinessImagesVerDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessImagesVerDto.Row = row * page

	businessImagesVerDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesVerDto.TenantId = user.TenantId
	businessImagesVerDto.ImagesId=ctx.URLParam("imagesId")
	businessImagesVerDto.Version=ctx.URLParam("version")

	total, err = businessImagesVerService.businessImagesVerDao.GetBusinessImagesVerCount(businessImagesVerDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessImagesVerDtos, err = businessImagesVerService.businessImagesVerDao.GetBusinessImagesVers(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesVerDtos, total, row)

}

/**
保存 系统信息
*/
func (businessImagesVerService *BusinessImagesVerService) SaveBusinessImagesVer(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesVerDto businessImages.BusinessImagesVerDto
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	ctx.ReadJSON(&businessImagesVerDto)
	tmpBusinessImagesVerDto := businessImages.BusinessImagesVerDto{
		ImagesId: businessImagesVerDto.ImagesId,
		Version: businessImagesVerDto.Version,
	}
	businessImagesVerDtos, _ := businessImagesVerService.businessImagesVerDao.GetBusinessImagesVers(tmpBusinessImagesVerDto)

	if businessImagesVerDtos != nil && len(businessImagesVerDtos) > 0{
		return result.Error("版本已存在")
	}
	businessImagesVerDto.Id = seq.Generator()

	businessImagesVerDto.TenantId = user.TenantId
	//businessImagesVerDto.Version = "V" + date.GetNowAString()

	err = businessImagesVerService.businessImagesVerDao.SaveBusinessImagesVer(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesVerDto)

}

/**
修改 系统信息
*/
func (businessImagesVerService *BusinessImagesVerService) UpdateBusinessImagesVer(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesVerDto businessImages.BusinessImagesVerDto
	)

	if err = ctx.ReadJSON(&businessImagesVerDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesVerService.businessImagesVerDao.UpdateBusinessImagesVer(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesVerDto)

}

/**
删除 系统信息
*/
func (businessImagesVerService *BusinessImagesVerService) DeleteBusinessImagesVer(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesVerDto businessImages.BusinessImagesVerDto
	)

	if err = ctx.ReadJSON(&businessImagesVerDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesVerService.businessImagesVerDao.DeleteBusinessImagesVer(businessImagesVerDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesVerDto)

}

// get remote images version

func (businessImagesVerService *BusinessImagesVerService) GetRemoteBusinessImagesVer(ctx iris.Context) interface{} {

	var (
		err       error
		resultDto result.ResultDto
	)

	values := ctx.URLParams()

	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}
	value := Urlencode(values)
	resp, err := httpReq.Get(config.Remote_get_images_version_Url+"?"+value, headers)
	if err != nil {
		return result.Error(err.Error())
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto
}

func (businessImagesVerService *BusinessImagesVerService) SaveRemoteBusinessImagesVer(ctx iris.Context) interface{} {

	var (
		err                error
		remoteBusinessImagesVerDto businessImages.RemoteBusinessImagesVerDto

		resultDto result.ResultDto
	)
	if err = ctx.ReadJSON(&remoteBusinessImagesVerDto); err != nil {
		return result.Error("解析入参失败")
	}

	// get publisher id

	appPublisherDto := appPublisher.AppPublisherDto{
		PublisherId: remoteBusinessImagesVerDto.PublisherId,
	}

	appPublisherDtos, err := businessImagesVerService.appPublisherDao.GetAppPublishers(appPublisherDto)

	if len(appPublisherDtos) < 1 {
		return result.Error("发布者不存在")
	}

	remoteBusinessImagesVerDto.PublisherId = appPublisherDtos[0].ExtPublisherId


	headers := map[string]string{
		"APP-ID":         config.Hc_cloud_app_id,
		"REQ-TIME":       date.GetNowAString(),
		"SIGN":           "",
		"TRANSACTION-ID": seq.Generator(),
		"USER-ID":        "-1",
	}
	mjson, _ := json.Marshal(remoteBusinessImagesVerDto)
	dataEn, _ := encrypt.Encrypt(appPublisherDtos[0].Token, string(mjson))

	applyAppDto := installApp2.ApplyAppDto{
		PublisherId: appPublisherDtos[0].ExtPublisherId,
		Data:        dataEn,
	}
	resp, err := httpReq.SendRequest(config.Remote_Apply_Publish_App_Version_Url, applyAppDto, headers, "POST")
	if err != nil {
		return result.Error(err.Error())
	}

	//resps, _ := encrypt.Decrypt(appPublisherDtos[0].Token, string(resp))
	json.Unmarshal(resp, &resultDto)

	return resultDto
}

