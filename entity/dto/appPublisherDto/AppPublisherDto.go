package appPublisherDto

import "github.com/zihao-boy/zihao/entity/dto"

const (
	// normal state

	StateNormal string = "001"
)

type AppPublisherDto struct {
	dto.PageDto
	PublisherId string `json:"publisherId" sql:"-"`
	Username    string `json:"username" sql:"-"`
	Email       string `json:"email" sql:"-"`
	Token       string `json:"token" sql:"-"`
	Phone       string `json:"phone" sql:"-"`
	State       string `json:"state" sql:"-"`
	CreateTime  string `json:"createTime" sql:"-"`
	StatusCd    string `json:"statusCd" sql:"-"`
	TenantId    string `json:"tenantId" sql:"-"`
	ExtPublisherId string `json:"extPublisherId" sql:"-"`
}

