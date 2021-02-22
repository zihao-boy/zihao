package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
	"github.com/zihao-boy/zihao/zihao-service/system/dao"
)

type ServiceSqlService struct {
	serviceSqlDao dao.ServiceSqlDao
}

/**
查询 系统信息
*/
func (serviceSqlService *ServiceSqlService) GetServiceSqls(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		serviceSqlDto = serviceSql.ServiceSqlDto{SqlCode: ctx.URLParam("sqlCode")}
		serviceSqlDtos []*serviceSql.ServiceSqlDto
	)

	serviceSqlDtos,err = serviceSqlService.serviceSqlDao.GetServiceSqls(serviceSqlDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDtos)

}


/**
保存 系统信息
*/
func (serviceSqlService *ServiceSqlService) SaveServiceSqls(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = serviceSqlService.serviceSqlDao.SaveServiceSql(serviceSqlDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}


/**
修改 系统信息
*/
func (serviceSqlService *ServiceSqlService) UpdateServiceSqls(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = serviceSqlService.serviceSqlDao.UpdateServiceSql(serviceSqlDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}


/**
删除 系统信息
*/
func (serviceSqlService *ServiceSqlService) DeleteServiceSqls(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = serviceSqlService.serviceSqlDao.DeleteServiceSql(serviceSqlDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}