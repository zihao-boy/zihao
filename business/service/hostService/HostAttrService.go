package service

import (
	dao "github.com/zihao-boy/zihao/business/dao/host"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
)

type HostAttrService struct {
	hostAttrDao dao.HostAttrDao
}

/**
查询 系统信息
*/
func (hostAttrService *HostAttrService) GetHostAttrAll(hostAttrDto host.HostAttrDto) ([]*host.HostAttrDto, error) {
	var (
		err                error
		hostAttrDtos []*host.HostAttrDto
	)

	hostAttrDtos, err = hostAttrService.hostAttrDao.GetHostAttrs(hostAttrDto)
	if err != nil {
		return nil, err
	}

	return hostAttrDtos, nil

}

/**
查询 系统信息
*/
func (hostAttrService *HostAttrService) GetHostAttrs(ctx iris.Context) result.ResultDto {
	var (
		err                error
		page               int64
		row                int64
		total              int64
		hostAttrDto  = host.HostAttrDto{}
		hostAttrDtos []*host.HostAttrDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	hostAttrDto.Row = row * page

	hostAttrDto.Page = (page - 1) * row

	hostAttrDto.HostId = ctx.URLParam("hostId")

	total, err = hostAttrService.hostAttrDao.GetHostAttrCount(hostAttrDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	hostAttrDtos, err = hostAttrService.hostAttrDao.GetHostAttrs(hostAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostAttrDtos, total, row)

}

/**
保存 系统信息
*/
func (hostAttrService *HostAttrService) SaveHostAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		hostAttrDto host.HostAttrDto
	)

	if err = ctx.ReadJSON(&hostAttrDto); err != nil {
		return result.Error("解析入参失败")
	}
	hostAttrDto.HostId = seq.Generator()

	err = hostAttrService.hostAttrDao.SaveHostAttr(hostAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostAttrDto)

}

/**
修改 系统信息
*/
func (hostAttrService *HostAttrService) UpdateHostAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		hostAttrDto host.HostAttrDto
	)

	if err = ctx.ReadJSON(&hostAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostAttrService.hostAttrDao.UpdateHostAttr(hostAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostAttrDto)

}

/**
删除 系统信息
*/
func (hostAttrService *HostAttrService) DeleteHostAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		hostAttrDto host.HostAttrDto
	)

	if err = ctx.ReadJSON(&hostAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostAttrService.hostAttrDao.DeleteHostAttr(hostAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostAttrDto)

}
