package appPublisherService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/appPublisherDao"
	"github.com/zihao-boy/zihao/common/seq"
	appPublisher "github.com/zihao-boy/zihao/entity/dto/appPublisherDto"
	"github.com/zihao-boy/zihao/entity/dto/result"
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
	)
	if err = ctx.ReadJSON(&appPublisherDto); err != nil {
		return result.Error("解析入参失败")
	}
	appPublisherDto.PublisherId = seq.Generator()

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
	)
	if err = ctx.ReadJSON(&appPublisherDto); err != nil {
		return result.Error("解析入参失败")
	}

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
