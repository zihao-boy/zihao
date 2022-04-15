package dnsService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/dnsDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type DnsMapService struct {
	dnsDao dnsDao.DnsMapDao
}

// get db link
// all db by this user
func (dnsService *DnsMapService) GetDnsMapAll(DnsMapDto dns.DnsMapDto) ([]*dns.DnsMapDto, error) {
	var (
		err        error
		DnsMapDtos []*dns.DnsMapDto
	)

	DnsMapDtos, err = dnsService.dnsDao.GetDnsMaps(DnsMapDto)
	if err != nil {
		return nil, err
	}

	return DnsMapDtos, nil

}

/**
查询 系统信息
*/
func (dnsService *DnsMapService) GetDnsMaps(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		dnsDto  = dns.DnsMapDto{}
		dnsDtos []*dns.DnsMapDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	dnsDto.Row = row * page

	dnsDto.Page = (page - 1) * row

	total, err = dnsService.dnsDao.GetDnsMapCount(dnsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	dnsDtos, err = dnsService.dnsDao.GetDnsMaps(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDtos, total, row)

}

/**
保存 系统信息
*/
func (dnsService *DnsMapService) SaveDnsMaps(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsMapDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}
	dnsDto.DnsMapId = seq.Generator()
	//DnsMapDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = dnsService.dnsDao.SaveDnsMap(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

/**
修改 系统信息
*/
func (dnsService *DnsMapService) UpdateDnsMaps(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsMapDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	//dnsDto.Id = ctx.FormValue("id")

	//dnsDto.Name = ctx.FormValue("name")

	err = dnsService.dnsDao.UpdateDnsMap(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

/**
删除 系统信息
*/
func (dnsService *DnsMapService) DeleteDnsMaps(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsMapDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = dnsService.dnsDao.DeleteDnsMap(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}
