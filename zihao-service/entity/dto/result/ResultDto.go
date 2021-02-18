package result

const(
	CODE_SUCCESS int = 0
	CODE_ERROR int = -1
)

type ResultDto struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

/**
默认成功封装
 */
func Success() ResultDto {
	res :=  ResultDto{Code: CODE_SUCCESS,Msg: "成功"}
	return res
}

func SuccessData(data interface{}) ResultDto {
	res :=  ResultDto{Code: CODE_SUCCESS,Msg: "成功",Data: data}
	return res
}

/**
默认失败封装
 */
func Error(msg string) ResultDto {
	res := ResultDto{Code: CODE_ERROR,Msg: msg}
	return res
}

/**
默认失败封装
*/
func ErrorData(data interface{}) ResultDto {
	res := ResultDto{Code: CODE_ERROR,Msg: "未知异常",Data: data}
	return res
}

