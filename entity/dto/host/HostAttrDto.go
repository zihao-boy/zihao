package host

import "github.com/zihao-boy/zihao/entity/dto"

const Spec_cd_osName= "1"

type HostAttrDto struct {
	dto.PageDto
	AttrId     string `json:"attrId" sql:"-"`
	HostId       string `json:"hostId" sql:"-"`
	SpecCd    string `json:"specCd" sql:"-"`
	Value    string `json:"value" `
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
