package sysError

import "errors"

var (
	ERR_HEADER_NON_BEARER = errors.New("请求头必须包含Authorization: bearer token")
	TOKEN_PARSE_IS_EMPTY  = errors.New("token为空,解析失败")
	TOKEN_PARSE_INVALID   = errors.New("无效token")

	ERR_LOCK_ALREADY_REQUIRED = errors.New("锁已被占用")
	ERR_NO_LOCAL_IP_FOUND     = errors.New("没有找到网卡IP")

	ERR_PARSE_TOKEN_PHONE_FIELD = errors.New("解析token中的[id]字段错误")
	ERR_PARSE_TOKEN_ID_FIELD    = errors.New("解析token中的[phone]字段错误")
)
