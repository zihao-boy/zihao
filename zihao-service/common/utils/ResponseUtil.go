package utils


import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
)

// common error define
func Error(ctx iris.Context, status int, code constants.Code) {
	ctx.StatusCode(status)
	ctx.JSON(json(code, nil))
}

// 200 define
func Ok_(ctx iris.Context, code constants.Code) {
	Ok(ctx, code, nil)
}
func Ok(ctx iris.Context, code constants.Code, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(json(code, data))
}

// 401 error define
func Unauthorized(ctx context.Context, code constants.Code) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(json(code, nil))
}

func InternalServerError(ctx iris.Context, code constants.Code) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(json(code, nil))
}

func PaginationTableData(ctx iris.Context, total int64, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(json(constants.CODE_OK, iris.Map{
		"total": total,
		"data":  data,
	}))
}

// json包装
func json(code constants.Code, data interface{}) context.Map {
	return iris.Map{
		"code": fmt.Sprintf("%d", code),
		"msg":  code.String(),
		"data": data,
	}
}

