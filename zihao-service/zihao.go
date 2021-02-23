package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/app/router"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/jwt"
	"github.com/zihao-boy/zihao/zihao-service/config"

)
/**
 * 项目地址：https://github.com/zihao-boy/zihao.git
 *  作者：吴学文
 */
func main() {
	config.InitConfig(config.Asset)
	//support.InitLog()
	//support.InitValidator()
	mysql.InitGorm()
	redis.InitRedis()
	//auth.InitAuth()
	jwt.InitJWT()

	//初始化缓存信息
	redis.InitServiceSql()

	app := iris.New()
	router.Hub(app)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！</h1>")
		app.Logger().Info("欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！")
	})
	app.Run(iris.Addr(":7000"))
}