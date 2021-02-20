package serviceSql

type ServiceSqlDto struct {
	SqlText string `json:"sqlText" sql:"-"`
	SqlId string `json:"sqlId" sql:"-"`
	SqlCode string `json:"sqlCode" sql:"-"`
	Remark string `json:"remark"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd string `json:"statusCd" sql:"-"`
}
