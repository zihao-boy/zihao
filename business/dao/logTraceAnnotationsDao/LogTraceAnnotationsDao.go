package logTraceAnnotationsDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/log"
	"gorm.io/gorm"
)

const (
	query_logTraceAnnotations_count string = `
	select count(1) total
	from log_trace_annotations t
	where t.status_cd = '0'
	$if TraceId != '' then
	and t.trace_id = #TraceId#
	$endif
	$if SpanId != '' then
	and t.span_id = #SpanId#
	$endif
	$if ServiceName != '' then
	and t.service_name = #ServiceName#
	$endif
	$if Value != '' then
	and t.value = #Value#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_logTraceAnnotations string = `
				select t.*
				from log_trace_annotations t
				where t.status_cd = '0'
				$if TraceId != '' then
				and t.trace_id = #TraceId#
				$endif
				$if SpanId != '' then
				and t.span_id = #SpanId#
				$endif
				$if ServiceName != '' then
				and t.service_name = #ServiceName#
				$endif
				$if Value != '' then
				and t.value = #Value#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_logTraceAnnotations string = `
	insert into log_trace_annotations(id, span_id, service_name, ip,port,value,timestamp)
VALUES(#Id#,#SpanId#,#ServiceName#,#Ip#,#Port#,#Value#,#Timestamp#)
`

	update_logTraceAnnotations string = `
	update log_trace_annotations set
		$if ServiceName != '' then
		service_name = #ServiceName#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_logTraceAnnotations string = `
	update log_trace_annotations  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type LogTraceAnnotationsDao struct {
}

/**
查询用户
*/
func (*LogTraceAnnotationsDao) GetLogTraceAnnotationsCount(logTraceAnnotationsDto log.LogTraceAnnotationsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_logTraceAnnotations_count, objectConvert.Struct2Map(logTraceAnnotationsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*LogTraceAnnotationsDao) GetLogTraceAnnotationss(logTraceAnnotationsDto log.LogTraceAnnotationsDto) ([]*log.LogTraceAnnotationsDto, error) {
	var logTraceAnnotationsDtos []*log.LogTraceAnnotationsDto
	sqlTemplate.SelectList(query_logTraceAnnotations, objectConvert.Struct2Map(logTraceAnnotationsDto), func(db *gorm.DB) {
		db.Scan(&logTraceAnnotationsDtos)
	}, false)

	return logTraceAnnotationsDtos, nil
}

/**
保存服务sql
*/
func (*LogTraceAnnotationsDao) SaveLogTraceAnnotations(logTraceAnnotationsDto log.LogTraceAnnotationsDto) error {
	return sqlTemplate.Insert(insert_logTraceAnnotations, objectConvert.Struct2Map(logTraceAnnotationsDto), false)
}

/**
修改服务sql
*/
func (*LogTraceAnnotationsDao) UpdateLogTraceAnnotations(logTraceAnnotationsDto log.LogTraceAnnotationsDto) error {
	return sqlTemplate.Update(update_logTraceAnnotations, objectConvert.Struct2Map(logTraceAnnotationsDto), false)
}

/**
删除服务sql
*/
func (*LogTraceAnnotationsDao) DeleteLogTraceAnnotations(logTraceAnnotationsDto log.LogTraceAnnotationsDto) error {
	return sqlTemplate.Delete(delete_logTraceAnnotations, objectConvert.Struct2Map(logTraceAnnotationsDto), false)
}
