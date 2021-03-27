package controller

import "github.com/kataras/iris/v12"

func GetRequestMap(ctx iris.Context) *map[string]interface{} {
	var params map[string]interface{}
	_ = ctx.ReadJSON(&params)
	return &params
}
