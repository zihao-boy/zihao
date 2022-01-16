package logTraceAnnotationsService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/logTraceAnnotationsDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type LogTraceAnnotationsService struct {
	logTraceAnnotationsDao logTraceAnnotationsDao.LogTraceAnnotationsDao
}

// get db link
// all db by this user
func (logTraceAnnotationsService *LogTraceAnnotationsService) GetLogTraceAnnotationsAll(LogTraceAnnotationsDto log.LogTraceAnnotationsDto) ([]*log.LogTraceAnnotationsDto, error) {
	var (
		err        error
		LogTraceAnnotationsDtos []*log.LogTraceAnnotationsDto
	)

	LogTraceAnnotationsDtos, err = logTraceAnnotationsService.logTraceAnnotationsDao.GetLogTraceAnnotationss(LogTraceAnnotationsDto)
	if err != nil {
		return nil, err
	}

	return LogTraceAnnotationsDtos, nil

}

/**
查询 系统信息
*/
func (logTraceAnnotationsService *LogTraceAnnotationsService) GetLogTraceAnnotationss(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		logTraceAnnotationsDto  = log.LogTraceAnnotationsDto{}
		logTraceAnnotationsDtos []*log.LogTraceAnnotationsDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceAnnotationsDto.Row = row * page

	logTraceAnnotationsDto.Page = (page - 1) * row

	total, err = logTraceAnnotationsService.logTraceAnnotationsDao.GetLogTraceAnnotationsCount(logTraceAnnotationsDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceAnnotationsDtos, err = logTraceAnnotationsService.logTraceAnnotationsDao.GetLogTraceAnnotationss(logTraceAnnotationsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceAnnotationsDtos, total, row)

}

/**
保存 系统信息
*/
func (logTraceAnnotationsService *LogTraceAnnotationsService) SaveLogTraceAnnotationss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceAnnotationsDto log.LogTraceAnnotationsDto
	)
	if err = ctx.ReadJSON(&logTraceAnnotationsDto); err != nil {
		return result.Error("解析入参失败")
	}
	logTraceAnnotationsDto.Id = seq.Generator()
	//LogTraceAnnotationsDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = logTraceAnnotationsService.logTraceAnnotationsDao.SaveLogTraceAnnotations(logTraceAnnotationsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceAnnotationsDto)

}

/**
修改 系统信息
*/
func (logTraceAnnotationsService *LogTraceAnnotationsService) UpdateLogTraceAnnotationss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceAnnotationsDto log.LogTraceAnnotationsDto
	)
	if err = ctx.ReadJSON(&logTraceAnnotationsDto); err != nil {
		return result.Error("解析入参失败")
	}

	//logTraceAnnotationsDto.Id = ctx.FormValue("id")

	//logTraceAnnotationsDto.Name = ctx.FormValue("name")

	err = logTraceAnnotationsService.logTraceAnnotationsDao.UpdateLogTraceAnnotations(logTraceAnnotationsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceAnnotationsDto)

}

/**
删除 系统信息
*/
func (logTraceAnnotationsService *LogTraceAnnotationsService) DeleteLogTraceAnnotationss(ctx iris.Context) result.ResultDto {
	var (
		err       error
		logTraceAnnotationsDto log.LogTraceAnnotationsDto
	)
	if err = ctx.ReadJSON(&logTraceAnnotationsDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceAnnotationsService.logTraceAnnotationsDao.DeleteLogTraceAnnotations(logTraceAnnotationsDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceAnnotationsDto)

}
