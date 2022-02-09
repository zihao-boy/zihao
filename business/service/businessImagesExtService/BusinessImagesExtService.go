package businessImagesExtService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/business/dao/businessImagesExtDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	businessImagesExt "github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type BusinessImagesExtService struct {
	businessImagesExtDao businessImagesExtDao.BusinessImagesExtDao
	appServiceDao        dao.AppServiceDao
}

// get db link
// all db by this user
func (businessImagesExtService *BusinessImagesExtService) GetBusinessImagesExtAll(BusinessImagesExtDto businessImagesExt.BusinessImagesExtDto) ([]*businessImagesExt.BusinessImagesExtDto, error) {
	var (
		err                   error
		BusinessImagesExtDtos []*businessImagesExt.BusinessImagesExtDto
	)

	BusinessImagesExtDtos, err = businessImagesExtService.businessImagesExtDao.GetBusinessImagesExts(BusinessImagesExtDto)
	if err != nil {
		return nil, err
	}

	return BusinessImagesExtDtos, nil

}

/**
查询 系统信息
*/
func (businessImagesExtService *BusinessImagesExtService) GetBusinessImagesExts(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		page                  int64
		row                   int64
		total                 int64
		businessImagesExtDto  = businessImagesExt.BusinessImagesExtDto{}
		businessImagesExtDtos []*businessImagesExt.BusinessImagesExtDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessImagesExtDto.Row = row * page

	businessImagesExtDto.Page = (page - 1) * row

	total, err = businessImagesExtService.businessImagesExtDao.GetBusinessImagesExtCount(businessImagesExtDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessImagesExtDtos, err = businessImagesExtService.businessImagesExtDao.GetBusinessImagesExts(businessImagesExtDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesExtDtos, total, row)

}

/**
保存 系统信息
*/
func (businessImagesExtService *BusinessImagesExtService) SaveBusinessImagesExts(ctx iris.Context) result.ResultDto {
	var (
		err                  error
		businessImagesExtDto businessImagesExt.BusinessImagesExtDto
	)
	if err = ctx.ReadJSON(&businessImagesExtDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesExtDto.TenantId = user.TenantId

	businessImagesExtDto.Id = seq.Generator()

	err = businessImagesExtService.businessImagesExtDao.SaveBusinessImagesExt(businessImagesExtDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesExtDto)

}

/**
修改 系统信息
*/
func (businessImagesExtService *BusinessImagesExtService) UpdateBusinessImagesExts(ctx iris.Context) result.ResultDto {
	var (
		err                  error
		businessImagesExtDto businessImagesExt.BusinessImagesExtDto
	)
	if err = ctx.ReadJSON(&businessImagesExtDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesExtService.businessImagesExtDao.UpdateBusinessImagesExt(businessImagesExtDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesExtDto)

}

/**
删除 系统信息
*/
func (businessImagesExtService *BusinessImagesExtService) DeleteBusinessImagesExts(ctx iris.Context) result.ResultDto {
	var (
		err                  error
		businessImagesExtDto businessImagesExt.BusinessImagesExtDto
	)
	if err = ctx.ReadJSON(&businessImagesExtDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesExtService.businessImagesExtDao.DeleteBusinessImagesExt(businessImagesExtDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesExtDto)

}
