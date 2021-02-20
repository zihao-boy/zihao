package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/encrypt"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/entity/vo"
)

const(
	query_user string = "select t.user_id ,t.username,t.real_name ,\nt.phone,t.email,t.state,t.create_time,t.tenant_id \nfrom u_user t\nwhere t.username = ? and t.passwd = ?\nand t.state = '100201'\nlimit 1"
)

type UserDao struct {

}

/**
查询用户
 */
func (*UserDao) GetUser(userVo vo.LoginUserVo) (*user.UserDto,error){
	 var userDto = user.UserDto{}
	 db := mysql.G_DB.Raw(query_user,userVo.Username,encrypt.Md5(userVo.Passwd))
	 if err:=db.Scan(&userDto).Error; err !=nil{
	 	return nil,err
	 }

	 return &userDto,nil
}