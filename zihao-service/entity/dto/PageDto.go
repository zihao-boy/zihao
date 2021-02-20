package dto

/**
页面对象
 */
type PageDto struct {
	Page int `json:"page"`
	Row int `json:"row"`
	Records int `json:"records"`
	Total int `json:"total"`
}
