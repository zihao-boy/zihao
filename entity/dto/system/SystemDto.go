package system

/**
系统信息对象
add by wuxw 2021-02-18
*/
type SystemDto struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Time    string `json:"time"`
}
