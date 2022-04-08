package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/app/router"
	"github.com/zihao-boy/zihao/common/task"
	"github.com/zihao-boy/zihao/config"
)

func main() {
	//加载配置文件

	config.InitProp("conf/zihao.properties")
	go task.SlaveHealth()
	app := iris.New()
	router.HubSlave(app)
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>欢迎访问梓豪平台slave</h1>")
		app.Logger().Info("欢迎访问梓豪平台slave")
	})
	app.Run(iris.Addr(":7001"))

}
