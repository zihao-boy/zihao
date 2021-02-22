package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/entity/vo"
)

const(
	query_user string = "userDao.GetUser"
	update_user string ="userDao.UpdateUser"
	)

type UserDao struct {

}

/**
查询用户
 */
func (*UserDao) GetUser(userVo vo.LoginUserVo) (*user.UserDto,error){
	 var (
	 	userDto = user.UserDto{}
	 	err error
	 )
	sqlTemplate.SelectOne(query_user,objectConvert.Struct2Map(userVo), func(db *gorm.DB) {
		err=db.Scan(&userDto).Error
	},true)
	 return &userDto,err
}


/**
查询用户
*/
func (*UserDao) UpdateUser(userDto user.UserDto) error{
	return sqlTemplate.Update(update_user,objectConvert.Struct2Map(userDto),true)
}