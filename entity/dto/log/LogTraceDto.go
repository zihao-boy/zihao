package log

import "github.com/zihao-boy/zihao/entity/dto"

// trace dto
type LogTraceDto struct {
	dto.PageDto
	Id                string                    `json:"id"`
	Name              string                    `json:"name"`
	ParentId          string                    `json:"parentSpanId" sql:"-"`
	TraceId           string                    `json:"traceId" sql:"-"`
	Timestamp         int64                     `json:"timestamp"`
	Duration          int64                     `json:"duration"`
	ServiceName       string                    `json:"serviceName" sql:"-"`
	ParentServiceName string                    `json:"parentServiceName" sql:"-"`
	Ip                string                    `json:"ipv4"`
	Port              string                    `json:"port"`
	CreateTime        string                    `json:"createTime" sql:"-"`
	StatusCd          string                    `json:"statusCd" sql:"-"`
	Annotations       []*LogTraceAnnotationsDto `json:"annotations"`
}

type LogTraceDataDto struct {
	LogTraceDto
	Annotations []*LogTraceAnnotationsDataDto `json:"annotations"`
	Dbs         []*LogTraceDbDto              `json:"dbs"`
	Param       *LogTraceParamDto             `json:"param"`
}
