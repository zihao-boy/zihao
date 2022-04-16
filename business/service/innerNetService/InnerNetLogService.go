package innerNetService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/innerNetDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type InnerNetLogService struct {
	innerNetDao             innerNetDao.InnerNetLogDao
}

// get db link
// all db by this user
func (innerNetService *InnerNetLogService) GetInnerNetLogAll(InnerNetLogDto innerNet.InnerNetLogDto) ([]*innerNet.InnerNetLogDto, error) {
	var (
		err          error
		InnerNetLogDtos []*innerNet.InnerNetLogDto
	)

	InnerNetLogDtos, err = innerNetService.innerNetDao.GetInnerNetLogs(InnerNetLogDto)
	if err != nil {
		return nil, err
	}

	return InnerNetLogDtos, nil

}

/**
查询 系统信息
*/
func (innerNetService *InnerNetLogService) GetInnerNetLogs(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		innerNetDto  = innerNet.InnerNetLogDto{}
		innerNetDtos []*innerNet.InnerNetLogDto
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

	total, err = innerNetService.innerNetDao.GetInnerNetLogCount(innerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	innerNetDtos, err = innerNetService.innerNetDao.GetInnerNetLogs(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDtos, total, row)

}

/**
保存 系统信息
*/
func (innerNetService *InnerNetLogService) SaveInnerNetLogs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetLogDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	innerNetDto.LogId = seq.Generator()
	//InnerNetLogDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = innerNetService.innerNetDao.SaveInnerNetLog(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/**
修改 系统信息
*/
func (innerNetService *InnerNetLogService) UpdateInnerNetLogs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetLogDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	//innerNetDto.Id = ctx.FormValue("id")

	//innerNetDto.Name = ctx.FormValue("name")

	err = innerNetService.innerNetDao.UpdateInnerNetLog(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(innerNetDto)

}

/**
删除 系统信息
*/
func (innerNetService *InnerNetLogService) DeleteInnerNetLogs(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetLogDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = innerNetService.innerNetDao.DeleteInnerNetLog(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}
