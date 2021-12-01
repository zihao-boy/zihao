package vo

type LoginUserVo struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}
