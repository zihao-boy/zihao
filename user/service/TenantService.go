package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/encrypt"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/privilege"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/tenant"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/entity/vo"
	"github.com/zihao-boy/zihao/user/dao"
)

type TenantService struct {
	tenantDao dao.TenantDao
	userDao   dao.UserDao
}

/**
查询 系统信息
*/
func (tenantService *TenantService) GetTenants(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		tenantDto  = tenant.TenantDto{}
		tenantDtos []*tenant.TenantDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	tenantDto.Row = row * page

	tenantDto.Page = (page - 1) * row

	total, err = tenantService.tenantDao.GetTenantCount(tenantDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	tenantDtos, err = tenantService.tenantDao.GetTenants(tenantDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantDtos, total, row)

}

/**
保存 系统信息
*/
func (tenantService *TenantService) SaveTenants(ctx iris.Context) result.ResultDto {
	var (
		err              error
		tenantDto        tenant.TenantDto
		userVo           vo.LoginUserVo
		userDto          user.UserDto
		privilegeUserDto privilege.PrivilegeUserDto
	)

	if err = ctx.ReadJSON(&tenantDto); err != nil {
		return result.Error("解析入参失败")
	}
	tenantDto.TenantId = seq.Generator()

	// 校验用户是否存在
	userVo = vo.LoginUserVo{Username: tenantDto.Username}
	tmpUserDto, _ := tenantService.userDao.GetUser(userVo)

	if tmpUserDto != nil && tmpUserDto.UserId != "" {
		return result.Error("账户已经存在，请更换")
	}
	userDto = user.UserDto{Username: tenantDto.Username,
		Passwd:   encrypt.Md5(tenantDto.Passwd),
		RealName: tenantDto.PersonName,
		Sex:      1,
		Phone:    tenantDto.Phone,
		State:    "100201",
		TenantId: tenantDto.TenantId,
		Email:    tenantDto.Username + "@zihao.com",
		UserId:   seq.Generator(),
		UserRole: "1001",
	}
	err = tenantService.userDao.SaveUser(userDto)
	if err != nil {
		return result.Error(err.Error())
	}

	//将用户 和管理员权限绑定
	privilegeUserDto = privilege.PrivilegeUserDto{
		PrivilegeFlag: "1",
		Pid:           "600201904002",
		UserId:        userDto.UserId,
		TenantId:      userDto.TenantId,
	}
	err = tenantService.userDao.SaveUserPrivilege(privilegeUserDto)
	if err != nil {
		return result.Error(err.Error())
	}
	tenantDto.State = "3003"
	err = tenantService.tenantDao.SaveTenant(tenantDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantDto)

}

/**
修改 系统信息
*/
func (tenantService *TenantService) UpdateTenants(ctx iris.Context) result.ResultDto {
	var (
		err       error
		tenantDto tenant.TenantDto
	)

	if err = ctx.ReadJSON(&tenantDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = tenantService.tenantDao.UpdateTenant(tenantDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantDto)

}

/**
删除 系统信息
*/
func (tenantService *TenantService) DeleteTenants(ctx iris.Context) result.ResultDto {
	var (
		err       error
		tenantDto tenant.TenantDto
	)

	if err = ctx.ReadJSON(&tenantDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = tenantService.tenantDao.DeleteTenant(tenantDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(tenantDto)

}
