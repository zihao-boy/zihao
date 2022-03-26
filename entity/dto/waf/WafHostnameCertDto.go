package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

type WafHostnameCertDto struct {
	dto.PageDto
	CertId         string    `json:"certId" sql:"-" `
	Hostname       string    `json:"hostname" `
	CertContent    string    `json:"certContent" sql:"-"`
	PrivKeyContent string    `json:"privKeyContent"  sql:"-"`
	CreateTime     time.Time `json:"createTime" sql:"-"`
	StatusCd       string    `json:"statusCd" sql:"-"`
}
