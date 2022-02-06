package installAppService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/installAppDao"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/installApp"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type InstallAppService struct {
	installAppDao installAppDao.InstallAppDao
}

// get db link
// all db by this user
func (installAppService *InstallAppService) GetInstallAppAll(InstallAppDto installApp.InstallAppDto) ([]*installApp.InstallAppDto, error) {
	var (
		err        error
		InstallAppDtos []*installApp.InstallAppDto
	)

	InstallAppDtos, err = installAppService.installAppDao.GetInstallApps(InstallAppDto)
	if err != nil {
		return nil, err
	}

	return InstallAppDtos, nil

}

/**
查询 系统信息
*/
func (installAppService *InstallAppService) GetInstallApps(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		installAppDto  = installApp.InstallAppDto{}
		installAppDtos []*installApp.InstallAppDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	installAppDto.Row = row * page

	installAppDto.Page = (page - 1) * row

	total, err = installAppService.installAppDao.GetInstallAppCount(installAppDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	installAppDtos, err = installAppService.installAppDao.GetInstallApps(installAppDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(installAppDtos, total, row)

}

/**
保存 系统信息
*/
func (installAppService *InstallAppService) SaveInstallApps(param string) result.ResultDto {
	var (
		err       error
		installAppDto installApp.InstallAppDto
	)
	json.Unmarshal([]byte(param), &installAppDto)

	//object convert
	objectConvert.Struct2Struct(installAppDto,installAppDto)

	installAppDto.AppId = seq.Generator()
	//InstallAppDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = installAppService.installAppDao.SaveInstallApp(installAppDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(installAppDto)

}

/**
修改 系统信息
*/
func (installAppService *InstallAppService) UpdateInstallApps(ctx iris.Context) result.ResultDto {
	var (
		err       error
		installAppDto installApp.InstallAppDto
	)
	if err = ctx.ReadJSON(&installAppDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = installAppService.installAppDao.UpdateInstallApp(installAppDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(installAppDto)

}

/**
删除 系统信息
*/
func (installAppService *InstallAppService) DeleteInstallApps(ctx iris.Context) result.ResultDto {
	var (
		err       error
		installAppDto installApp.InstallAppDto
	)
	if err = ctx.ReadJSON(&installAppDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = installAppService.installAppDao.DeleteInstallApp(installAppDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(installAppDto)

}
