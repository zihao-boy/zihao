package logTraceDbDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"gorm.io/gorm"
)

const (
	query_logTraceDb_count string = `
	select count(1) total
from log_trace_db t
left join log_trace lt on t.span_id = lt.id and lt.status_cd = '0'
					where t.status_cd = '0'
					$if SpanId != '' then
					and t.span_id = #SpanId#
					$endif
					$if TraceId != '' then
					and t.trace_id = #TraceId#
					$endif
					$if ServiceName != '' then
					and t.service_name = #ServiceName#
					$endif
					$if Id != '' then
					and t.id = #Id#
					$endif

	`
	query_logTraceDb string = `
select t.*,lt.service_name,lt.trace_id
from log_trace_db t
left join log_trace lt on t.span_id = lt.id and lt.status_cd = '0'
					where t.status_cd = '0'
					$if SpanId != '' then
					and t.span_id = #SpanId#
					$endif
					$if TraceId != '' then
					and t.trace_id = #TraceId#
					$endif
					$if ServiceName != '' then
					and t.service_name = #ServiceName#
					$endif
					$if Id != '' then
					and t.id = #Id#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_logTraceDb string = `
	insert into log_trace_db(id, span_id, db_sql, param,duration)
VALUES(#Id#,#SpanId#,#DbSql#,#Param#,#Duration#)
`

	update_logTraceDb string = `
	update log_trace_db set
		$if DbSql != '' then
		db_sql = #DbSql#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_logTraceDb string = `
	update log_trace_db  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type LogTraceDbDao struct {
}

/**
查询用户
*/
func (*LogTraceDbDao) GetLogTraceDbCount(logTraceDbDto log.LogTraceDbDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_logTraceDb_count, objectConvert.Struct2Map(logTraceDbDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*LogTraceDbDao) GetLogTraceDbs(logTraceDbDto log.LogTraceDbDto) ([]*log.LogTraceDbDto, error) {
	var logTraceDbDtos []*log.LogTraceDbDto
	sqlTemplate.SelectList(query_logTraceDb, objectConvert.Struct2Map(logTraceDbDto), func(db *gorm.DB) {
		db.Scan(&logTraceDbDtos)
	}, false)

	return logTraceDbDtos, nil
}

/**
保存服务sql
*/
func (*LogTraceDbDao) SaveLogTraceDb(logTraceDbDto log.LogTraceDbDto) error {
	return sqlTemplate.Insert(insert_logTraceDb, objectConvert.Struct2Map(logTraceDbDto), false)
}

/**
修改服务sql
*/
func (*LogTraceDbDao) UpdateLogTraceDb(logTraceDbDto log.LogTraceDbDto) error {
	return sqlTemplate.Update(update_logTraceDb, objectConvert.Struct2Map(logTraceDbDto), false)
}

/**
删除服务sql
*/
func (*LogTraceDbDao) DeleteLogTraceDb(logTraceDbDto log.LogTraceDbDto) error {
	return sqlTemplate.Delete(delete_logTraceDb, objectConvert.Struct2Map(logTraceDbDto), false)
}
