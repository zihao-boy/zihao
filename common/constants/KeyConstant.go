package constants

// 帖子的状态
const (
	Checking     = iota + 1 // 审核中
	Check_Ok                // 通过
	Check_Fail              // 未通过
	Check_Unpaid            // 未付费
	Check_Finish            // 帖子完成状态
	Check_Down              // 帖子下架
)

// 店铺状态
const (
	Shop_Checking = iota + 1 // 审核中
	Shop_Ok                  // 通过
	Shop_Fail                // 未通过
	Shop_
)

// 聊天消息状态
const (
	CHAT_STATUS_UNREAD = iota + 1 // 未读
	CHAT_STATUS_READ              // 已读
)

// 聊天消息类型
const (
	CHAT_TYPEZ_SYSTEM = iota + 1 // 系统类型
	CHAT_TYPEZ_O2O               // 聊天消息类型
)

// 访问七牛云的图片前缀
const IMAGE_PREFIX = "http://pxjp230z7.bkt.clouddn.com/"

const UID = "uid"

const UINFO = "uinfo"

// resid
const (
	REDIS_ADMIN_FORMAT = "admin.token.id:%s"
	REDIS_USER_FORMAT  = "user.token.id:%s"
)

// websocket topic define
const (
	WS_TOPIC_LOGINED = "LOGINED" // 已登录
	WS_TOPIC_CHAT    = "CHAT"
	WS_TOPIC_INING   = "INING" // 在聊天界面，同步的消息，意思是在聊天可以直接发送消息
	WS_TOPIC_UNREAD  = "UNREAD"
)

