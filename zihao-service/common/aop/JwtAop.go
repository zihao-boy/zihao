package aop

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/jwt"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"strings"
)

type Aop struct {
}

func ServeHTTP(ctx iris.Context) {
	var err  error
	golog.Infof("==> host=[%s], method=[%s], path=[%s]", ctx.Host(), ctx.Method(), ctx.Path())

	// 不需要验证的特殊接口，如：登录
	if func(path string) bool {
		if strings.Contains(path, "/static") {
			return true
		}
		for _, v := range config.G_AppConfig.IgnoreURLs {
			if path == v {
				return true
			}
		}
		return false
	}(ctx.Path()) {
		ctx.Next()
		return
	}
	// 检查回话
	if err = jwt.G_JWT.ServeHTTP(*ctx); err != nil {
		golog.Errorf("中间件token检验失败，错误：%s", err)
		return
	}
	// 验证权限
	//if !icasbin.G_Casbin.Enforce(ctx.Values().GetString(support.UID), common.G_AppConfig.Domain, ctx.Path(), ctx.Method(), ".*") {
	//	support.Error(ctx, iris.StatusForbidden, support.CODE_PERMISSION_NIL)
	//	return
	//}
	// Pass to real API
	ctx.Next()
}

