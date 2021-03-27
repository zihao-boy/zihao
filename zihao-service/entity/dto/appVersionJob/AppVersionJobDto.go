package appVersionJob

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

const (
	STATE_wait string = "1001"
	STATE_doing string = "2002"
	STATE_success string = "4004"
	STATE_error string = "3003"
)

type AppVersionJobDto struct {
	dto.PageDto
	JobId string `json:"jobId" sql:"-"`
	JobName string `json:"jobName" sql:"-"`
	JobShell string `json:"jobShell" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	PreJobTime string `json:"preJobTime" sql:"-"`
	CurJobTime string `json:"curJobTime" sql:"-"`
	State string `json:"state"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
	StateName string `json:"stateName" sql:"-"`

}