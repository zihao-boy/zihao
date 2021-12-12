package host

import "github.com/zihao-boy/zihao/entity/dto"

const (
	State_w string = "1001" // 未管理
	State_N string = "2002" // 正常
	State_D string = "3003" // 处理中
)

/**
主机 实体
*/
type HostDto struct {
	dto.PageDto
	HostId        string `json:"hostId" sql:"-"`
	GroupId       string `json:"groupId" sql:"-"`
	GroupName     string `json:"groupName" sql:"-"`
	Name          string `json:"name" `
	Ip            string `json:"ip" `
	Username      string `json:"username" `
	Passwd        string `json:"passwd" `
	Cpu           string `json:"cpu" `
	UseCpu        string `json:"useCpu"`
	Mem           string `json:"mem" `
	UseMem        string `json:"useMem"`
	Disk          string `json:"disk" `
	UseDisk       string `json:"useDisk"`
	TenantId      string `json:"tenantId" sql:"-"`
	CreateTime    string `json:"createTime" sql:"-"`
	StatusCd      string `json:"statusCd" sql:"-"`
	State         string `json:"state"`
	HeartbeatTime string `json:"heartbeatTime" sql:"-"`
	CurPath       string `json:"curPath" `
}
