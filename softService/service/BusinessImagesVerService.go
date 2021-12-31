package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
	"strconv"
)

type BusinessImagesVerService struct {
	businessImagesVerDao     dao.BusinessImagesVerDao
	businessDockerfileDao dao.BusinessDockerfileDao
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
	businessImagesVerDto.Id = seq.Generator()


	businessImagesVerDto.TenantId = user.TenantId
	businessImagesVerDto.Version = "V" + date.GetNowAString()

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

