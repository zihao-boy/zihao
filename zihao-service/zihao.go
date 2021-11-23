package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/app/router"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/factory"
	"github.com/zihao-boy/zihao/zihao-service/common/crontab"
	"github.com/zihao-boy/zihao/zihao-service/common/db/dbFactory"
	"github.com/zihao-boy/zihao/zihao-service/common/jwt"
	"github.com/zihao-boy/zihao/zihao-service/config"
)

/**
 * 项目地址：https://github.com/zihao-boy/zihao.git
 *  作者：吴学文
 */
func main() {
	config.InitConfig()
	//support.InitLog()
	//support.InitValidator()
	//mysql.InitGorm()
	dbFactory.Init()
	factory.Init()
	//auth.InitAuth()
	jwt.InitJWT()

	//初始化缓存信息
	factory.InitServiceSql()

	//启动定时任务
	var (
		monitorJob = crontab.MonitorJob{}
	)
	monitorJob.Restart()

	app := iris.New()

	router.Hub(app)
	app.HandleDir("/", "../zihao-front/public")

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！</h1>")
	// 	app.Logger().Info("欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！")
	// })

	app.Run(iris.Addr(":7000"))

}
