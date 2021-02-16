package main


import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	if !common.G_AppConfig.Separate {
		app.RegisterView(iris.HTML("./resources", ".html").Binary(static.Asset, static.AssetNames))
		app.HandleDir("/", "./resources", iris.DirOptions{
			//IndexName:  "/index.html", // default "/index.html"
			Asset:      static.Asset,
			AssetNames: static.AssetNames,
			AssetInfo:  static.AssetInfo,
		})
	}
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
		app.Logger().Info("我在学习Iris噢")
	})
	app.Run(iris.Addr(":8009"))
}