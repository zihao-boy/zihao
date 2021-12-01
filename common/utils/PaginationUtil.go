package utils

// pagination struct
type Pagination struct {
	Start int `json:"start" validate:"min=0"`
	Size  int `json:"size" validate:"min=0"`
}

// 初始化分页参数
func (page *Pagination) Init() {
	if page.Start-1 < 0 {
		page.Start = 1
	}
	page.Start = (page.Start - 1) * page.Size
}
