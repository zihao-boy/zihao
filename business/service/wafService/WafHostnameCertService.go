package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafHostnameCertService struct {
	wafDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafHostnameCertService) GetWafHostnameCertAll(WafHostnameCertDto waf.WafHostnameCertDto) ([]*waf.WafHostnameCertDto, error) {
	var (
		err        error
		WafHostnameCertDtos []*waf.WafHostnameCertDto
	)

	WafHostnameCertDtos, err = wafService.wafDao.GetWafHostnameCerts(WafHostnameCertDto)
	if err != nil {
		return nil, err
	}

	return WafHostnameCertDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafHostnameCertService) GetWafHostnameCerts(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafHostnameCertDto{}
		wafDtos []*waf.WafHostnameCertDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	wafDto.Row = row * page

	wafDto.Page = (page - 1) * row

	total, err = wafService.wafDao.GetWafHostnameCertCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafHostnameCerts(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafHostnameCertService) SaveWafHostnameCerts(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostnameCertDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.CertId = seq.Generator()
	//WafHostnameCertDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafHostnameCert(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafHostnameCertService) UpdateWafHostnameCerts(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostnameCertDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafHostnameCert(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafHostnameCertService) DeleteWafHostnameCerts(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafHostnameCertDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafHostnameCert(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
