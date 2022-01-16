package logTraceService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/logTraceDao"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type LogTraceService struct {
	logTraceDao logTraceDao.LogTraceDao
}

// get db link
// all db by this user
func (logTraceService *LogTraceService) GetLogTraceAll(LogTraceDto log.LogTraceDto) ([]*log.LogTraceDto, error) {
	var (
		err        error
		LogTraceDtos []*log.LogTraceDto
	)

	LogTraceDtos, err = logTraceService.logTraceDao.GetLogTraces(LogTraceDto)
	if err != nil {
		return nil, err
	}

	return LogTraceDtos, nil

}

/**
查询 系统信息
*/
func (logTraceService *LogTraceService) GetLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		logTraceDto  = log.LogTraceDto{}
		logTraceDtos []*log.LogTraceDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDto.Row = row * page

	logTraceDto.Page = (page - 1) * row

	total, err = logTraceService.logTraceDao.GetLogTraceCount(logTraceDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceDtos, err = logTraceService.logTraceDao.GetLogTraces(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDtos, total, row)

}

/**
保存 系统信息
*/
func (logTraceService *LogTraceService) SaveLogTraces(param string) result.ResultDto {
	var (
		err       error
		logTraceDto log.LogTraceDto
		logTraceDataDto log.LogTraceDataDto
	)
	json.Unmarshal([]byte(param), &logTraceDataDto)

	//object convert
	objectConvert.Struct2Struct(logTraceDataDto,logTraceDto)

	logTraceDto.Id = seq.Generator()
	//LogTraceDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = logTraceService.logTraceDao.SaveLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDto)

}

/**
修改 系统信息
*/
func (logTraceService *LogTraceService) UpdateLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceDto log.LogTraceDto
	)
	if err = ctx.ReadJSON(&logTraceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceService.logTraceDao.UpdateLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDto)

}

/**
删除 系统信息
*/
func (logTraceService *LogTraceService) DeleteLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceDto log.LogTraceDto
	)
	if err = ctx.ReadJSON(&logTraceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceService.logTraceDao.DeleteLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDto)

}
