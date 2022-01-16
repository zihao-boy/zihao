package log

// trace dto
type LogTraceDto struct {
	Id string `json:"id"`
	Name  string `json:"name"`
	ParentId string `json:"parentId" sql:"-"`
	TraceId string `json:"traceId" sql:"-"`
	Timestamp  string `json:"timestamp"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
}