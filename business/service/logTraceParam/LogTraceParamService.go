package logTraceParamService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/logTraceParamDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type LogTraceParamService struct {
	logTraceParamDao logTraceParamDao.LogTraceParamDao
}

// get db link
// all db by this user
func (logTraceParamService *LogTraceParamService) GetLogTraceParamAll(LogTraceParamDto log.LogTraceParamDto) ([]*log.LogTraceParamDto, error) {
	var (
		err        error
		LogTraceParamDtos []*log.LogTraceParamDto
	)

	LogTraceParamDtos, err = logTraceParamService.logTraceParamDao.GetLogTraceParams(LogTraceParamDto)
	if err != nil {
		return nil, err
	}

	return LogTraceParamDtos, nil

}

/**
查询 系统信息
*/
func (logTraceParamService *LogTraceParamService) GetLogTraceParams(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		logTraceParamDto  = log.LogTraceParamDto{}
		logTraceParamDtos []*log.LogTraceParamDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceParamDto.Row = row * page

	logTraceParamDto.Page = (page - 1) * row

	total, err = logTraceParamService.logTraceParamDao.GetLogTraceParamCount(logTraceParamDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceParamDtos, err = logTraceParamService.logTraceParamDao.GetLogTraceParams(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDtos, total, row)

}

/**
保存 系统信息
*/
func (logTraceParamService *LogTraceParamService) SaveLogTraceParams(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceParamDto log.LogTraceParamDto
	)
	if err = ctx.ReadJSON(&logTraceParamDto); err != nil {
		return result.Error("解析入参失败")
	}
	logTraceParamDto.Id = seq.Generator()
	//LogTraceParamDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = logTraceParamService.logTraceParamDao.SaveLogTraceParam(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDto)

}

/**
修改 系统信息
*/
func (logTraceParamService *LogTraceParamService) UpdateLogTraceParams(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceParamDto log.LogTraceParamDto
	)
	if err = ctx.ReadJSON(&logTraceParamDto); err != nil {
		return result.Error("解析入参失败")
	}

	//logTraceParamDto.Id = ctx.FormValue("id")

	//logTraceParamDto.Name = ctx.FormValue("name")

	err = logTraceParamService.logTraceParamDao.UpdateLogTraceParam(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDto)

}

/**
删除 系统信息
*/
func (logTraceParamService *LogTraceParamService) DeleteLogTraceParams(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceParamDto log.LogTraceParamDto
	)
	if err = ctx.ReadJSON(&logTraceParamDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceParamService.logTraceParamDao.DeleteLogTraceParam(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDto)

}
