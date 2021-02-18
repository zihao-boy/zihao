package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/entity/vo"
	dao2 "github.com/zihao-boy/zihao/zihao-service/user/dao"
)

type UserService struct {
	userDao dao2.UserDao

}
/**
   用户登录处理
 */
func (userService *UserService) Login(ctx iris.Context) (result.ResultDto,*user.UserDto) {
	var (
		err       error
		userVo = new(vo.LoginUserVo)
		userDto *user.UserDto
	)
	if err = ctx.ReadJSON(&userVo); err != nil {
		return result.Error("解析入参失败"),nil
	}

	userDto,err = userService.userDao.GetUser(*userVo)
	if(err != nil){
		return result.Error("用户名密码错误"),nil
	}

	return result.SuccessData(userDto),userDto
}