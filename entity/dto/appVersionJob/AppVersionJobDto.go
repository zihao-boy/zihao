package appVersionJob

import "github.com/zihao-boy/zihao/entity/dto"

const (
	STATE_wait    string = "1001"
	STATE_doing   string = "2002"
	STATE_success string = "4004"
	STATE_error   string = "3003"
	WEBHOOKS_FROM_GITEE string ="gitee" // gitee
	WEBHOOKS_FROM_GITHUB string ="github" // github
	WEBHOOKS_FROM_GITLAB string ="gitlab" // gitlab

	EVENT_PUSH string ="push"
	EVENT_PUSH_AND_RESTART = "pushAndRestart"
)

type AppVersionJobDto struct {
	dto.PageDto
	JobId      string `json:"jobId" sql:"-"`
	JobName    string `json:"jobName" sql:"-"`
	JobShell   string `json:"jobShell" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
	State      string `json:"state"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
	StateName  string `json:"stateName" sql:"-"`
	GitUrl  string `json:"gitUrl" sql:"-"`
	GitPasswd  string `json:"gitPasswd" sql:"-"`
	GitUsername  string `json:"gitUsername" sql:"-"`
	WorkDir  string `json:"workDir" sql:"-"`
	JobTime  string `json:"jobTime" sql:"-"`
	AppVersionJobImages[] AppVersionJobImagesDto `json:"appVersionJobImages"`
}
// do build param
type AppVersionJobParam struct {
	JobId      string `json:"jobId" `
	Action     string `json:"action"`
	Images string `json:"images"`
}
// gitee webhooks
type GiteeWebHooksDto struct {
	HookName string `json:"hook_name"`
}
