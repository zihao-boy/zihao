package log

import "github.com/zihao-boy/zihao/entity/dto"

// trace dto
type LogTraceAnnotationsDto struct {
	dto.PageDto
	Id          string `json:"id"`
	SpanId      string `json:"spanId"`
	ServiceName string `json:"spanId" sql:"-"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Value       string `json:"value"`
	CreateTime  string `json:"createTime" sql:"-"`
	StatusCd    string `json:"statusCd" sql:"-"`
}
