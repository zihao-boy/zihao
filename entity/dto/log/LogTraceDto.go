package log

import "github.com/zihao-boy/zihao/entity/dto"

// trace dto
type LogTraceDto struct {
	dto.PageDto
	Id string `json:"id"`
	Name  string `json:"name"`
	ParentId string `json:"parentId" sql:"-"`
	TraceId string `json:"traceId" sql:"-"`
	Timestamp  string `json:"timestamp"`
	Duration  string `json:"duration"`
	ServiceName string `json:"serviceName" sql:"-"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
}

type LogTraceDataDto struct {
	LogTraceDto
	Annotations []*LogTraceAnnotationsDataDto
	Param *LogTraceParamDto
}