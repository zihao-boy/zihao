package waf

import (
	"github.com/zihao-boy/zihao/entity/dto"
	"time"
)

const (
	Waf_Rule_State_T = "T" //start
	Waf_Rule_State_F = "F" //stop
	Waf_Rule_Obj_Type_Ip  = "IP"
	Waf_Rule_Obj_Type_CC  = "CC"
	Waf_Rule_Obj_Type_Location  = "Location"
)

type WafRuleDto struct {
	dto.PageDto
	RuleId    string    `json:"ruleId" sql:"-" `
	GroupId    string    `json:"groupId" sql:"-" `
	RuleName  string    `json:"ruleName" sql:"-" `
	Scope  string    `json:"scope" `
	ObjId  string    `json:"objId" sql:"-" `
	ObjType  string    `json:"objType" sql:"-" `
	Seq string `json:"seq"`
	State      string    `json:"state" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}