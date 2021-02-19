package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/jwt"
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

/**
登录处理类
 */
func (aus *UserController) login(ctx iris.Context) {
	 resultDto,userDto := aus.userService.Login(ctx);

	 if userDto != nil{
		 token, _ := jwt.G_JWT.GenerateToken(userDto)
		 //token 保存至redis
		 redis.G_Redis.SetToken(constants.REDIS_ADMIN_FORMAT, userDto.UserId,token)
		 ctx.SetCookieKV(jwt.DEFAULT_TOKEN,token);
	 }

	ctx.JSON(resultDto)
}