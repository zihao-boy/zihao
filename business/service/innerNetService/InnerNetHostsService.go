package innerNetService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/innerNetDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"strconv"
)

type InnerNetHostsService struct {
	innerNetDao innerNetDao.InnerNetHostsDao
}

// get db link
// all db by this user
func (innerNetService *InnerNetHostsService) GetInnerNetHostsAll(InnerNetHostsDto innerNet.InnerNetHostsDto) ([]*innerNet.InnerNetHostsDto, error) {
	var (
		err        error
		InnerNetHostsDtos []*innerNet.InnerNetHostsDto
	)

	InnerNetHostsDtos, err = innerNetService.innerNetDao.GetInnerNetHostss(InnerNetHostsDto)
	if err != nil {
		return nil, err
	}

	return InnerNetHostsDtos, nil

}

/**
查询 系统信息
*/
func (innerNetService *InnerNetHostsService) GetInnerNetHostss(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		innerNetDto  = innerNet.InnerNetHostsDto{}
		innerNetDtos []*innerNet.InnerNetHostsDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	innerNetDto.Row = row * page

	innerNetDto.Page = (page - 1) * row

	total, err = innerNetService.innerNetDao.GetInnerNetHostsCount(innerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	innerNetDtos, err = innerNetService.innerNetDao.GetInnerNetHostss(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDtos, total, row)

}

/**
保存 系统信息
*/
func (innerNetService *InnerNetHostsService) SaveInnerNetHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		innerNetDto innerNet.InnerNetHostsDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	innerNetDto.InnerNetHostId = seq.Generator()
	//InnerNetHostsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = innerNetService.innerNetDao.SaveInnerNetHosts(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/**
修改 系统信息
*/
func (innerNetService *InnerNetHostsService) UpdateInnerNetHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		innerNetDto innerNet.InnerNetHostsDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	//innerNetDto.Id = ctx.FormValue("id")

	//innerNetDto.Name = ctx.FormValue("name")

	err = innerNetService.innerNetDao.UpdateInnerNetHosts(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/**
删除 系统信息
*/
func (innerNetService *InnerNetHostsService) DeleteInnerNetHostss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		innerNetDto innerNet.InnerNetHostsDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = innerNetService.innerNetDao.DeleteInnerNetHosts(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}
