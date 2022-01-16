package log

// trace dto
type LogTraceParamDto struct {
	Id string `json:"id"`
	SpanId  string `json:"spanId"`
	ReqParam     string `json:"reqParam" sql:"-"`
	ResParam     string `json:"resParam" sql:"-"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
}