package serviceSql

import "github.com/zihao-boy/zihao/zihao-service/entity/dto"

type ServiceSqlDto struct {
	dto.PageDto
	SqlText string `json:"sqlText" sql:"-"`
	SqlId string `json:"sqlId" sql:"-"`
	SqlCode string `json:"sqlCode" sql:"-"`
	Remark string `json:"remark"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
