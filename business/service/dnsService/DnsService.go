package dnsService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/dnsDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
	"strings"
)

type DnsService struct {
	dnsDao      dnsDao.DnsDao
	dnsHostsDao dnsDao.DnsHostsDao
}

// get db link
// all db by this user
func (dnsService *DnsService) GetDnsAll(DnsDto dns.DnsDto) ([]*dns.DnsDto, error) {
	var (
		err     error
		DnsDtos []*dns.DnsDto
	)

	DnsDtos, err = dnsService.dnsDao.GetDnss(DnsDto)
	if err != nil {
		return nil, err
	}

	return DnsDtos, nil

}

/**
查询 系统信息
*/
func (dnsService *DnsService) GetDnss(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		dnsDto  = dns.DnsDto{}
		dnsDtos []*dns.DnsDto
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

	total, err = dnsService.dnsDao.GetDnsCount(dnsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	dnsDtos, err = dnsService.dnsDao.GetDnss(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, dnsDto := range dnsDtos {

		dnsHostsDto := dns.DnsHostsDto{
			DnsId: dnsDto.DnsId,
		}
		dnsHostsDtos, _ := dnsService.dnsHostsDao.GetDnsHostss(dnsHostsDto)
		dnsDto.DnsHosts = dnsHostsDtos
	}

	return result.SuccessData(dnsDtos, total, row)

}

/**
保存 系统信息
*/
func (dnsService *DnsService) SaveDnss(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}
	dnsDto.DnsId = seq.Generator()
	//DnsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = dnsService.dnsDao.SaveDns(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

/**
修改 系统信息
*/
func (dnsService *DnsService) UpdateDnss(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	//dnsDto.Id = ctx.FormValue("id")

	//dnsDto.Name = ctx.FormValue("name")

	err = dnsService.dnsDao.UpdateDns(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if utils.IsEmpty(dnsDto.HostIds) {
		return result.SuccessData(dnsDto)
	}

	dnsHostsDto := dns.DnsHostsDto{
		DnsId: dnsDto.DnsId,
	}
	dnsService.dnsHostsDao.DeleteDnsHosts(dnsHostsDto)

	for _, hostId := range strings.Split(dnsDto.HostIds, ",") {

		dnsHostsDto = dns.DnsHostsDto{
			DnsId:     dnsDto.DnsId,
			DnsHostId: seq.Generator(),
			HostId:    hostId,
		}
		dnsService.dnsHostsDao.SaveDnsHosts(dnsHostsDto)
	}

	return result.SuccessData(dnsDto)

}

/**
删除 系统信息
*/
func (dnsService *DnsService) DeleteDnss(ctx iris.Context) result.ResultDto {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = dnsService.dnsDao.DeleteDns(dnsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dnsDto)

}

func (dnsService *DnsService) StartDnsf(ctx iris.Context) interface{} {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start dns
	tmpDnsDto := dns.DnsDto{
		DnsId: dnsDto.DnsId,
	}
	dnsDtos, err := dnsService.dnsDao.GetDnss(tmpDnsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(dnsDtos) < 1 {
		return result.Error("未查询到数据")
	}

	resultDto, _ := shell.ExecStartDns(dnsService.getDnsConfig(*dnsDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpDnsDto = dns.DnsDto{
		DnsId: dnsDto.DnsId,
		State: dns.Dns_state_start,
	}

	err = dnsService.dnsDao.UpdateDns(tmpDnsDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(dnsDto)

}

func (dnsService *DnsService) StopDnsf(ctx iris.Context) interface{} {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start dns
	tmpDnsDto := dns.DnsDto{
		DnsId: dnsDto.DnsId,
	}
	dnsDtos, err := dnsService.dnsDao.GetDnss(tmpDnsDto)

	if err != nil || len(dnsDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecStopDns(dnsService.getDnsConfig(*dnsDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpDnsDto = dns.DnsDto{
		DnsId: dnsDto.DnsId,
		State: dns.Dns_state_stop,
	}

	err = dnsService.dnsDao.UpdateDns(tmpDnsDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(dnsDto)

}

func (dnsService *DnsService) RefreshDnsConfig(ctx iris.Context) interface{} {
	var (
		err    error
		dnsDto dns.DnsDto
	)
	if err = ctx.ReadJSON(&dnsDto); err != nil {
		return result.Error("解析入参失败")
	}
	// start dns
	tmpDnsDto := dns.DnsDto{
		DnsId: dnsDto.DnsId,
	}
	dnsDtos, err := dnsService.dnsDao.GetDnss(tmpDnsDto)

	if err != nil || len(dnsDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecRefreshDnsConfig(dnsService.getDnsConfig(*dnsDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	return result.SuccessData(dnsDto)
}

// get dns config
func (dnsService *DnsService) getDnsConfig(dnsDto dns.DnsDto) dns.DnsDataDto {
	var (
		dnsMapDao dnsDao.DnsMapDao
	)

	dnsHostsDto := dns.DnsHostsDto{
		DnsId: dnsDto.DnsId,
	}
	dnsHostsDtos, _ := dnsService.dnsHostsDao.GetDnsHostss(dnsHostsDto)
	dnsDto.DnsHosts = dnsHostsDtos

	// start dns
	tmpDnsMapDto := dns.DnsMapDto{

	}
	maps, _ := dnsMapDao.GetDnsMaps(tmpDnsMapDto)

	return dns.DnsDataDto{
		DnsIp: dnsDto.DnsIp,
		DnsPort:  dnsDto.DnsPort,
		Maps:  maps,
	}

}
