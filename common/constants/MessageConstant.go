package constants

const (
	CODE_OK = 200
	CODE_FAILUR = 500
	CODE_FORBIDDEN = 403
	CODE_NOTFOUNT = 404
)

const (
	CODE_USER_PHONE_FAILUR Code = iota + 1000
	CODE_USER_PHONE_DUPLICATE
	CODE_USER_PASSWORD_FAILUR
	CODE_USER_REGISTE_OK
	CODE_USER_REGISTE_FAILUR
	CODE_USER_LOGIN_OK
	CODE_USER_NOT_ENABLED
)
const (
	CODE_SYS_OK Code = iota + 2000
	CODE_SYS_FAILUR
	CODE_SYS_ERROR
	CODE_SYS_PARSE_PARAMS_ERROR
)
const (
	CODE_TOKEN_CREATE_FAILUR Code = iota + 3000
	CODE_TOKEN_EXPIRE         // token过期，如：redis没有用户的token，jwt解析token过期
	CODE_TOKEN_EMPTY          // 请求头中没有携带token，或为空串
	CODE_TOKEN_INVALID        // 无效token，比如：伪造token，算法错误，header的关键字错误等等
	CODE_TOKEN_CACHE_ERROR    // 令牌缓存错误，redis
	CODE_TOKEN_NOT_MAPCLAIMS    // 令牌不是MapClaims类型
)
const (
	CODE_PERMISSION_NIL Code = iota + 4000 // 没有权限
)

type Code int
// 定义系统内部消息，模拟枚举类型
func (code Code) String() string {
	switch code {
	case CODE_OK: return "成功"
	case CODE_FAILUR: return "失败"
	case CODE_FORBIDDEN: return "未登录"
	case CODE_NOTFOUNT: return "页面不存在"

	case CODE_USER_PHONE_FAILUR: return "用户名错误，或账号被冻结"
	case CODE_USER_PHONE_DUPLICATE: return "该号码已注册"
	case CODE_USER_PASSWORD_FAILUR: return "密码错误"
	case CODE_USER_REGISTE_OK: return "注册成功"
	case CODE_USER_LOGIN_OK: return "恭喜，登录成功"
	case CODE_USER_NOT_ENABLED: return "改用户已被冻结，请联系系统管理员"

	case CODE_SYS_OK: return "恭喜，成功"
	case CODE_SYS_FAILUR: return "失败"
	case CODE_SYS_ERROR: return "服务器内部错误"
	case CODE_SYS_PARSE_PARAMS_ERROR: return "请求参数错误"

	case CODE_TOKEN_CREATE_FAILUR: return "创建令牌失败"
	case CODE_TOKEN_EXPIRE: return "当前回话已过期"
	case CODE_TOKEN_EMPTY: return "令牌不存在"
	case CODE_TOKEN_INVALID: return "无效令牌"
	case CODE_TOKEN_CACHE_ERROR: return "令牌缓存错误"
	case CODE_TOKEN_NOT_MAPCLAIMS: return "令牌不是MapClaims类型"

	case CODE_PERMISSION_NIL: return "没有操作权限"


	default: return "未知错误"
	}
}
