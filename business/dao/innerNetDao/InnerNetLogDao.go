package innerNetDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"gorm.io/gorm"
)

const (
	query_innerNetLog_count string = `
	select count(1) total
from inner_net_log t
					where t.status_cd = '0'
					$if LogId != '' then
					and t.log_id = #LogId#
					$endif
					$if Username != '' then
					and t.username = #Username#
					$endif
					$if SrcIp != '' then
					and t.src_ip = #SrcIp#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif

	`
	query_innerNetLog string = `
select t.*
from inner_net_log t
					where t.status_cd = '0'
					$if LogId != '' then
					and t.log_id = #LogId#
					$endif
					$if Username != '' then
					and t.username = #Username#
					$endif
					$if SrcIp != '' then
					and t.src_ip = #SrcIp#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_innerNetLog string = `
	insert into inner_net_log(log_id,username,src_ip,ip, state)
VALUES(#LogId#,#Username#,#SrcIp#,#Ip#,#State#)
`

	update_innerNetLog string = `
	update inner_net_log set
					$if State != '' then
					 state = #State#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if LogId != '' then
					and log_id = #LogId#
					$endif
	`
	delete_innerNetLog string = `
	update inner_net_log  set
                          status_cd = '1'
                          where status_cd = '0'
					$if LogId != '' then
					and log_id = #LogId#
					$endif
	`
)

type InnerNetLogDao struct {
}

/**
查询用户
*/
func (*InnerNetLogDao) GetInnerNetLogCount(innerNetLogDto innerNet.InnerNetLogDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_innerNetLog_count, objectConvert.Struct2Map(innerNetLogDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InnerNetLogDao) GetInnerNetLogs(innerNetLogDto innerNet.InnerNetLogDto) ([]*innerNet.InnerNetLogDto, error) {
	var innerNetLogDtos []*innerNet.InnerNetLogDto
	sqlTemplate.SelectList(query_innerNetLog, objectConvert.Struct2Map(innerNetLogDto), func(db *gorm.DB) {
		db.Scan(&innerNetLogDtos)
	}, false)

	return innerNetLogDtos, nil
}

/**
保存服务sql
*/
func (*InnerNetLogDao) SaveInnerNetLog(innerNetLogDto innerNet.InnerNetLogDto) error {
	return sqlTemplate.Insert(insert_innerNetLog, objectConvert.Struct2Map(innerNetLogDto), false)
}

/**
修改服务sql
*/
func (*InnerNetLogDao) UpdateInnerNetLog(innerNetLogDto innerNet.InnerNetLogDto) error {
	return sqlTemplate.Update(update_innerNetLog, objectConvert.Struct2Map(innerNetLogDto), false)
}

/**
删除服务sql
*/
func (*InnerNetLogDao) DeleteInnerNetLog(innerNetLogDto innerNet.InnerNetLogDto) error {
	return sqlTemplate.Delete(delete_innerNetLog, objectConvert.Struct2Map(innerNetLogDto), false)
}
