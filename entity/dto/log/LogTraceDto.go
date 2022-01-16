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
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
}

type LogTraceDataDto struct {
	LogTraceDto
	annotations *[]LogTraceAnnotationsDto
	param *LogTraceParamDto
}