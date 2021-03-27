package result

import "math"

const(
	CODE_SUCCESS int = 0
	CODE_ERROR int = -1
)

type ResultDto struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
	Total int64 `json:"total"`
	Records int64 `json:"records"`
	Row int64 `json:"row"`
}

/**
默认成功封装
 */
func Success() ResultDto {
	res :=  ResultDto{Code: CODE_SUCCESS,Msg: "成功"}
	return res
}

func SuccessData(data interface{} ,totals ...int64) ResultDto {
	var (
		total int64
		records int64
		row int64
	)
	if len(totals) == 2{
		total = totals[0]
		row = totals[1]
		records = int64(math.Ceil(float64(total)/float64(totals[1])))
	}
	res :=  ResultDto{Code: CODE_SUCCESS,Msg: "成功",Total:total,Records:records,Row:row,Data: data}
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

