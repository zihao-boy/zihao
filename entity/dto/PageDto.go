package dto

/**
页面对象
*/
type PageDto struct {
	Page    int64 `json:"page"`
	Row     int64 `json:"row"`
	Records int64 `json:"records"`
	Total   int64 `json:"total"`
}
