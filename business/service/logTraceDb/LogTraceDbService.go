package logTraceDbService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/logTraceDbDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type LogTraceDbService struct {
	logTraceDbDao logTraceDbDao.LogTraceDbDao
}

// get db link
// all db by this user
func (logTraceDbService *LogTraceDbService) GetLogTraceDbAll(LogTraceDbDto log.LogTraceDbDto) ([]*log.LogTraceDbDto, error) {
	var (
		err        error
		LogTraceDbDtos []*log.LogTraceDbDto
	)

	LogTraceDbDtos, err = logTraceDbService.logTraceDbDao.GetLogTraceDbs(LogTraceDbDto)
	if err != nil {
		return nil, err
	}

	return LogTraceDbDtos, nil

}

/**
查询 系统信息
*/
func (logTraceDbService *LogTraceDbService) GetLogTraceDbs(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		logTraceDbDto  = log.LogTraceDbDto{}
		logTraceDbDtos []*log.LogTraceDbDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDbDto.Row = row * page

	logTraceDbDto.Page = (page - 1) * row

	total, err = logTraceDbService.logTraceDbDao.GetLogTraceDbCount(logTraceDbDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceDbDtos, err = logTraceDbService.logTraceDbDao.GetLogTraceDbs(logTraceDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDbDtos, total, row)

}

/**
保存 系统信息
*/
func (logTraceDbService *LogTraceDbService) SaveLogTraceDbs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceDbDto log.LogTraceDbDto
	)
	if err = ctx.ReadJSON(&logTraceDbDto); err != nil {
		return result.Error("解析入参失败")
	}
	logTraceDbDto.Id = seq.Generator()
	//LogTraceDbDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = logTraceDbService.logTraceDbDao.SaveLogTraceDb(logTraceDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDbDto)

}

/**
修改 系统信息
*/
func (logTraceDbService *LogTraceDbService) UpdateLogTraceDbs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceDbDto log.LogTraceDbDto
	)
	if err = ctx.ReadJSON(&logTraceDbDto); err != nil {
		return result.Error("解析入参失败")
	}

	//logTraceDbDto.Id = ctx.FormValue("id")

	//logTraceDbDto.Name = ctx.FormValue("name")

	err = logTraceDbService.logTraceDbDao.UpdateLogTraceDb(logTraceDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDbDto)

}

/**
删除 系统信息
*/
func (logTraceDbService *LogTraceDbService) DeleteLogTraceDbs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceDbDto log.LogTraceDbDto
	)
	if err = ctx.ReadJSON(&logTraceDbDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceDbService.logTraceDbDao.DeleteLogTraceDb(logTraceDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDbDto)

}
