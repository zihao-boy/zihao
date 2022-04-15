package dnsService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/dnsDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"strconv"
)

type DnsHostsService struct {
	dnsDao dnsDao.DnsHostsDao
}

// get db link
// all db by this user
func (dnsService *DnsHostsService) GetDnsHostsAll(DnsHostsDto dns.DnsHostsDto) ([]*dns.DnsHostsDto, error) {
	var (
		err        error
		DnsHostsDtos []*dns.DnsHostsDto
	)

	DnsHostsDtos, err = dnsService.dnsDao.GetDnsHostss(DnsHostsDto)
	if err != nil {
		return nil, err
	}

	return DnsHostsDtos, nil

}

/**
查询 系统信息
*/
func (dnsService *DnsHostsService) GetDnsHostss(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		dnsDto  = dns.DnsHostsDto{}
		dnsDtos []*dns.DnsHostsDto
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

	total, err = dnsService.dnsDao.GetDnsHostsCount(dnsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	dnsDtos, err = dnsService.dnsDao.GetDnsHostss(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDtos, total, row)

}

/**
保存 系统信息
*/
func (dnsService *DnsHostsService) SaveDnsHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dnsDto dns.DnsHostsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}
	dnsDto.DnsHostId = seq.Generator()
	//DnsHostsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = dnsService.dnsDao.SaveDnsHosts(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

/**
修改 系统信息
*/
func (dnsService *DnsHostsService) UpdateDnsHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dnsDto dns.DnsHostsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	//dnsDto.Id = ctx.FormValue("id")

	//dnsDto.Name = ctx.FormValue("name")

	err = dnsService.dnsDao.UpdateDnsHosts(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

/**
删除 系统信息
*/
func (dnsService *DnsHostsService) DeleteDnsHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dnsDto dns.DnsHostsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = dnsService.dnsDao.DeleteDnsHosts(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}
