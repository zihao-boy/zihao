package logTraceService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/logTraceAnnotationsDao"
	"github.com/zihao-boy/zihao/business/dao/logTraceDao"
	"github.com/zihao-boy/zihao/business/dao/logTraceParamDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
)

type LogTraceService struct {
	logTraceDao            logTraceDao.LogTraceDao
	logTraceParamDao            logTraceParamDao.LogTraceParamDao
	logTraceAnnotationsDao logTraceAnnotationsDao.LogTraceAnnotationsDao
}

// get db link
// all db by this user
func (logTraceService *LogTraceService) GetLogTraceAll(LogTraceDto log.LogTraceDto) ([]*log.LogTraceDto, error) {
	var (
		err          error
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
		err          error
		page         int64
		row          int64
		total        int64
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

	logTraceDto.Name = ctx.URLParam("name")

	logTraceDto.TraceId = ctx.URLParam("traceId")

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
查询 系统信息
*/
func (logTraceService *LogTraceService) GetLogTraceDetail(ctx iris.Context) result.ResultDto {
	var (
		err          error
		page         int64
		row          int64
		total        int64
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

	logTraceDto.Name = ctx.URLParam("name")

	logTraceDto.TraceId = ctx.URLParam("traceId")

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

	logTraceDtos = logTraceService.addAnn(logTraceDtos)

	return result.SuccessData(logTraceDtos, total, row)

}
/**
保存 系统信息
*/
func (logTraceService *LogTraceService) SaveLogTraces(param string) result.ResultDto {
	var (
		err                    error
		logTraceDto            log.LogTraceDto
		logTraceDataDto        log.LogTraceDataDto
		logTraceAnnotationsDto log.LogTraceAnnotationsDto
		crTimestame int64
		csTimestame int64

	)
	json.Unmarshal([]byte(param), &logTraceDataDto)

	//object convert
	json.Unmarshal([]byte(param), &logTraceDto)


	logTraceDto.Id = seq.Generator()
	//LogTraceDto.Path = filepath.Join(curDest, fileHeader.Filename)

	if logTraceDataDto.Annotations == nil || len(logTraceDataDto.Annotations) < 1 {
		return result.Error("未包含Annotations")
	}

	logTraceDto.ServiceName = logTraceDataDto.Annotations[0].Endpoint.ServiceName
	logTraceDto.Ip = logTraceDataDto.Annotations[0].Endpoint.Ip
	logTraceDto.Port = logTraceDataDto.Annotations[0].Endpoint.Port
	logTraceDto.Duration = 0
	//compute Duration cr - cs
	if len(logTraceDataDto.Annotations) == 4{

		for _, annotation := range logTraceDataDto.Annotations {
			if annotation.Value == "cr"{
				crTimestame = annotation.Timestamp
			}

			if annotation.Value == "cs"{
				csTimestame = annotation.Timestamp
			}
		}

		logTraceDto.Duration = crTimestame - csTimestame

	}

	err = logTraceService.logTraceDao.SaveLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, annotation := range logTraceDataDto.Annotations {
		logTraceAnnotationsDto = log.LogTraceAnnotationsDto{
			Id:          seq.Generator(),
			SpanId:      logTraceDto.Id,
			ServiceName: annotation.Endpoint.ServiceName,
			Ip:          annotation.Endpoint.Ip,
			Port:        annotation.Endpoint.Port,
			Value:       annotation.Value,
			Timestamp:   annotation.Timestamp,
		}
		err = logTraceService.logTraceAnnotationsDao.SaveLogTraceAnnotations(logTraceAnnotationsDto)
		if err != nil {
			return result.Error(err.Error())
		}
	}

	return result.SuccessData(logTraceDto)
}

/**
修改 系统信息
*/
func (logTraceService *LogTraceService) UpdateLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err         error
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
		err         error
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

func (logTraceService *LogTraceService) addAnn(dtos []*log.LogTraceDto) []*log.LogTraceDto {

	for _, trace := range dtos{

		anno := log.LogTraceAnnotationsDto{
			SpanId: trace.Id,
		}

		annos,err := logTraceService.logTraceAnnotationsDao.GetLogTraceAnnotationss(anno)
		if err != nil{
			continue
		}
		trace.Annotations = annos
	}

	return dtos

}

func (logTraceService *LogTraceService) GetLogTraceParam(ctx iris.Context) interface{} {
	var (
		err          error
		page         int64
		row          int64
		total        int64
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

	logTraceParamDto.SpanId = ctx.URLParam("spanId")


	total, err = logTraceService.logTraceParamDao.GetLogTraceParamCount(logTraceParamDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceParamDtos, err = logTraceService.logTraceParamDao.GetLogTraceParams(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDtos, total, row)
}
