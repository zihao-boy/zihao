package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/app/router"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/crontab"
	"github.com/zihao-boy/zihao/common/db/dbFactory"
	"github.com/zihao-boy/zihao/common/ip"
	"github.com/zihao-boy/zihao/common/jwt"
	"github.com/zihao-boy/zihao/common/kafka"
	"github.com/zihao-boy/zihao/config"
	"strconv"
)

/**
 * project address：https://github.com/zihao-boy/zihao.git
 * author ：wuxw
 * email: 928255095@qq.com
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

	// load qqwry ip data
	ip.InitIP()

	//初始化缓存信息
	factory.InitServiceSql()

	//初始化映射
	factory.InitMapping()

	// init kafka
	kafka.Init()

	//启动定时任务
	crontab.StartCrontab()

	app := iris.New()

	router.Hub(app)
	app.HandleDir("/", "./web")

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！</h1>")
	// 	app.Logger().Info("欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！")
	// })

	port := config.G_AppConfig.Port

	if port == 0 {
		port = 7000
	}

	app.Run(iris.Addr(":" + strconv.Itoa(port)))

}
