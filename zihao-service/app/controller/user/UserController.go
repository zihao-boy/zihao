package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/user/service"
)

type UserController struct {
	userService service.UserService
}


func UserControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/user")
		aus      = UserController{userService: service.UserService{}}
	)
	adinUser.Post("/login", hero.Handler(aus.login))
}

func (aus *UserController) login(ctx iris.Context) {
	aus.userService.Login("","")
}