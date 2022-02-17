package log

import "github.com/zihao-boy/zihao/entity/dto"

// trace dto
type LogTraceAnnotationsDto struct {
	dto.PageDto
	Id          string `json:"id"`
	SpanId      string `json:"spanId" sql:"-"`
	ServiceName string `json:"serviceName" sql:"-"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Value       string `json:"value"`
	Timestamp   string `json:"timestamp"`
	CreateTime  string `json:"createTime" sql:"-"`
	StatusCd    string `json:"statusCd" sql:"-"`
}

type EndpointDto struct {
	ServiceName string `json:"serviceName" sql:"-"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
}

type LogTraceAnnotationsDataDto struct {
	dto.PageDto
	Id         string      `json:"id"`
	SpanId     string      `json:"spanId" sql:"-"`
	Value      string      `json:"value"`
	Timestamp  string      `json:"timestamp"`
	CreateTime string      `json:"createTime" sql:"-"`
	StatusCd   string      `json:"statusCd" sql:"-"`
	Endpoint   EndpointDto `json:"endpoint"`
}
