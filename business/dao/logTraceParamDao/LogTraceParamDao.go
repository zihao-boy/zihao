package logTraceParamDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"gorm.io/gorm"
)

const (
	query_logTraceParam_count string = `
	select count(1) total
	from log_trace_param t
	where t.status_cd = '0'
	$if SpanId != '' then
	and t.span_id = #SpanId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_logTraceParam string = `
				select t.*
					from log_trace_param t
					where t.status_cd = '0'
					$if SpanId != '' then
					and t.span_id = #SpanId#
					$endif
					$if Id != '' then
					and t.id = #Id#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_logTraceParam string = `
	insert into log_trace_param(id, span_id, req_param, res_param)
VALUES(#Id#,#SpanId#,#ReqParam#,#ResParam#)
`

	update_logTraceParam string = `
	update log_trace_param set
		$if ReqParam != '' then
		req_param = #ReqParam#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_logTraceParam string = `
	update log_trace_param  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type LogTraceParamDao struct {
}

/**
查询用户
*/
func (*LogTraceParamDao) GetLogTraceParamCount(logTraceParamDto log.LogTraceParamDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_logTraceParam_count, objectConvert.Struct2Map(logTraceParamDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*LogTraceParamDao) GetLogTraceParams(logTraceParamDto log.LogTraceParamDto) ([]*log.LogTraceParamDto, error) {
	var logTraceParamDtos []*log.LogTraceParamDto
	sqlTemplate.SelectList(query_logTraceParam, objectConvert.Struct2Map(logTraceParamDto), func(db *gorm.DB) {
		db.Scan(&logTraceParamDtos)
	}, false)

	return logTraceParamDtos, nil
}

/**
保存服务sql
*/
func (*LogTraceParamDao) SaveLogTraceParam(logTraceParamDto log.LogTraceParamDto) error {
	return sqlTemplate.Insert(insert_logTraceParam, objectConvert.Struct2Map(logTraceParamDto), false)
}

/**
修改服务sql
*/
func (*LogTraceParamDao) UpdateLogTraceParam(logTraceParamDto log.LogTraceParamDto) error {
	return sqlTemplate.Update(update_logTraceParam, objectConvert.Struct2Map(logTraceParamDto), false)
}

/**
删除服务sql
*/
func (*LogTraceParamDao) DeleteLogTraceParam(logTraceParamDto log.LogTraceParamDto) error {
	return sqlTemplate.Delete(delete_logTraceParam, objectConvert.Struct2Map(logTraceParamDto), false)
}
