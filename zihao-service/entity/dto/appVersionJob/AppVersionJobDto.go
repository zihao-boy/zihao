package appVersionJob

type AppVersionJobDto struct {
	JobId string `json:"jobId" sql:"-"`
	JobName string `json:"jobName" sql:"-"`
	JobShell string `json:"jobShell" sql:"-"`
	TenantId string `json:"tenantId" sql:"-"`
	PreJobTime string `json:"preJobTime" sql:"-"`
	CurJobTime string `json:"curJobTime" sql:"-"`
	State string `json:"state"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`

}