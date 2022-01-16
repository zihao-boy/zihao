package log

// trace dto
type LogTraceAnnotationsDto struct {
	Id string `json:"id"`
	SpanId  string `json:"spanId"`
	ParentId string `json:"parentId" sql:"-"`
	TraceId string `json:"traceId" sql:"-"`
	Timestamp  string `json:"timestamp"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
}