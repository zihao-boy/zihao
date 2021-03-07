package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	rcover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/host"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/menu"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/monitor"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/system"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/tenant"
	"github.com/zihao-boy/zihao/zihao-service/app/controller/user"
	"github.com/zihao-boy/zihao/zihao-service/common/aop"
)

// 所有的路由
func Hub(app *iris.Application) {

	party := preSettring(app)

	//系统信息
	system.SystemControllerRouter(party)

	//菜单信息
	menu.MenuControllerRouter(party)

	//用户类控制类
	user.UserControllerRouter(party)

	//主机信息
	host.HostControllerRouter(party)

	//租户控制类
	tenant.TenantControllerRouter(party)

	//监控控制类
	monitor.MonitorControllerRouter(party)
}

func preSettring(app *iris.Application) (party iris.Party) {
	//app.Logger().SetLevel(config.G_AppConfig.LogLevel)
	//logger, close := .NewRequestLogger()
	//defer close()
	app.Use(
		rcover.New(), // 恢复
		//logger,       // 记录请求
		//aop.ServeHTTP
	)
	//// 定义错误处理
	//app.OnErrorCode(iris.StatusNotFound, logger, func(ctx iris.Context) {
	//	support.Error(ctx, iris.StatusNotFound, support.CODE_NOTFOUNT)
	//})
	//app.OnErrorCode(iris.StatusInternalServerError, logger, func(ctx iris.Context) {
	//	support.Error(ctx, iris.StatusInternalServerError, support.CODE_FAILUR)
	//})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})

	// 设置跨域
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		//Debug:            true,
	})
	party = app.Party("/app", crs).AllowMethods(iris.MethodOptions)
	party.Use(aop.ServeHTTP)
	return party
}

