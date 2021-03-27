package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/tenant"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/user/dao"
	"strconv"
)

type TenantSettingService struct {
	tenantSettingDao dao.TenantSettingDao
	userDao dao.UserDao
}

/**
查询 系统信息
*/
func (tenantSettingService *TenantSettingService) GetTenantSettings(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		page int64
		row int64
		total int64
		tenantSettingDto = tenant.TenantSettingDto{}
		tenantSettingDtos []*tenant.TenantSettingDto
	)


	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	tenantSettingDto.Row = row * page

	tenantSettingDto.Page = (page -1) * row

	total,err = tenantSettingService.tenantSettingDao.GetTenantSettingCount(tenantSettingDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}

	tenantSettingDtos,err = tenantSettingService.tenantSettingDao.GetTenantSettings(tenantSettingDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantSettingDtos,total,row)

}

/**
查询 系统信息
*/
func (tenantSettingService *TenantSettingService) GetTenantSettingAll(tenantSettingDto tenant.TenantSettingDto)  ([] *tenant.TenantSettingDto,error) {
	var (
		err       error
		tenantSettingDtos []*tenant.TenantSettingDto
	)


	tenantSettingDtos,err = tenantSettingService.tenantSettingDao.GetTenantSettings(tenantSettingDto)
	if(err != nil){
		return nil,err
	}
	return tenantSettingDtos,nil

}


/**
保存 系统信息
*/
func (tenantSettingService *TenantSettingService) SaveTenantSettings(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		tenantSettingDto tenant.TenantSettingDto
	)

	if err = ctx.ReadJSON(&tenantSettingDto); err != nil {
		return result.Error("解析入参失败")
	}
	tenantSettingDto.SettingId = seq.Generator()

	// 校验用户是否存在
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	tenantSettingDto.TenantId = user.TenantId

	err = tenantSettingService.tenantSettingDao.SaveTenantSetting(tenantSettingDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantSettingDto)

}


/**
修改 系统信息
*/
func (tenantSettingService *TenantSettingService) UpdateTenantSettings(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		tenantSettingDto tenant.TenantSettingDto
	)

	if err = ctx.ReadJSON(&tenantSettingDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = tenantSettingService.tenantSettingDao.UpdateTenantSetting(tenantSettingDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantSettingDto)

}


/**
删除 系统信息
*/
func (tenantSettingService *TenantSettingService) DeleteTenantSettings(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		tenantSettingDto tenant.TenantSettingDto
	)

	if err = ctx.ReadJSON(&tenantSettingDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = tenantSettingService.tenantSettingDao.DeleteTenantSetting(tenantSettingDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantSettingDto)

}