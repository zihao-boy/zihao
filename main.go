package main


import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
		app.Logger().Info("我在学习Iris噢")
	})
	app.Run(iris.Addr(":8009"))
}