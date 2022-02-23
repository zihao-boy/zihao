package log

import "github.com/zihao-boy/zihao/entity/dto"

type LogTraceDbDto struct {
	dto.PageDto
	Id         string `json:"id"`
	SpanId     string `json:"spanId" sql:"-"`
	DbSql      string `json:"dbSql" sql:"-"`
	Param      string `json:"param"`
	Duration   string `json:"duration"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`

	TraceId     string `json:"traceId" sql:"-"`
	ServiceName string `json:"serviceName" sql:"-"`
}
