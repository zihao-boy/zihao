package user

import (
	"time"
)

type UserDto struct {
	UserId     string    `json:"userId" form:"userId" sql:"-" gorm:"primary_key"`
	Username   string    `json:"username" form:"username"`
	Passwd     string    `json:"passwd" form:"passwd"`
	RealName   string    `json:"realName" form:"realName" sql:"-"`
	Sex        int       `json:"sex" form:"sex"`
	Phone      string    `json:"phone" form:"phone"`
	Email      string    `json:"email" form:"email"`
	CreateTime time.Time `json:"createTime" form:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" form:"statusCd"`
	State      string    `json:"state" form:"state"`
	TenantId   string    `json:"tenantId" sql:"-"`
	UserRole   string    `json:"userRole" sql:"-"`
}
