package main


import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
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
	//jwt.InitJWT()
	app := iris.New()


	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
		app.Logger().Info("我在学习Iris噢")
	})
	app.Run(iris.Addr(":8009"))
}