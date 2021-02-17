package user

import (
	"time"
)

type UserDto struct {
	Id         int64         `json:"id" form:"id" gorm:"primary_key"`
	Username   string        `json:"username" form:"username"`
	Password   string        `json:"password" form:"password"`
	Gender     int           `json:"gender" form:"gender"` // 1=男 2=女
	Enable     bool          `json:"enable" form:"enable"`
	Name       string        `json:"name" form:"name"`
	Age        int           `json:"age" form:"age"`
	Phone      string        `json:"phone" form:"phone"`
	Email      string        `json:"email" form:"email"`
	Userface   string        `json:"userface" form:"userface"`
	CreateTime time.Time     `json:"createTime" form:"createTime"`
	UpdateTime time.Time     `json:"updateTime" form:"updateTime"`
	Label      string        `json:"label" form:"label"`
	Online     bool          `json:"online" gorm:"-"` // 是否在线
}
