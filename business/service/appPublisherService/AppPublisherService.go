package appPublisherService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/appPublisherDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/config"
	appPublisher "github.com/zihao-boy/zihao/entity/dto/appPublisherDto"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type AppPublisherService struct {
	appPublisherDao appPublisherDao.AppPublisherDao
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
