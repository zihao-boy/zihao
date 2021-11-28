package vo

type ChangeUserPwdVo struct {
	UserId string `json:"userId"`
	Username string `json:"username"`
	OldPwd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`

}
