package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto/privilege"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/entity/vo"
	"gorm.io/gorm"
)

const (
	query_user          string = "userDao.GetUser"
	update_user         string = "userDao.UpdateUser"
	save_user           string = "userDao.SaveUser"
	save_user_privilege string = "userDao.SaveUserPrivilege"
)

type UserDao struct {
}

/**
查询用户
*/
func (*UserDao) GetUser(userVo vo.LoginUserVo) (*user.UserDto, error) {
	var (
		userDto = user.UserDto{}
		err     error
	)
	sqlTemplate.SelectOne(query_user, objectConvert.Struct2Map(userVo), func(db *gorm.DB) {
		err = db.Scan(&userDto).Error
	}, true)
	return &userDto, err
}

/**
查询用户
*/
func (*UserDao) UpdateUser(userDto user.UserDto) error {
	return sqlTemplate.Update(update_user, objectConvert.Struct2Map(userDto), true)
}

/**
查询用户
*/
func (*UserDao) SaveUser(userDto user.UserDto) error {
	return sqlTemplate.Insert(save_user, objectConvert.Struct2Map(userDto), true)
}

/**
查询用户
*/
func (*UserDao) SaveUserPrivilege(privilegeUserDto privilege.PrivilegeUserDto) error {
	return sqlTemplate.Insert(save_user_privilege, objectConvert.Struct2Map(privilegeUserDto), true)
}
