package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafAccessLog_count string = `
	select count(1) total
from waf_access_log t
					where t.status_cd = '0'
					$if RequestId != '' then
					and t.request_id = #RequestId#
					$endif
					$if WafIp != '' then
					and t.waf_ip = #WafIp#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if XRealIp != '' then
					and t.x_real_ip = #XRealIp#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if ResponseCode != '' then
					and t.response_code = #ResponseCode#
					$endif
					$if Method != '' then
					and t.method = #Method#
					$endif
					$if HttpHost != '' then
					and t.http_host = #HttpHost#
					$endif
					$if UpstreamAddr != '' then
					and t.upstream_addr = #UpstreamAddr#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif

	`
	query_wafAccessLog string = `
select t.*
from waf_access_log t
					where t.status_cd = '0'
					$if RequestId != '' then
					and t.request_id = #RequestId#
					$endif
					$if WafIp != '' then
					and t.waf_ip = #WafIp#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if XRealIp != '' then
					and t.x_real_ip = #XRealIp#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if ResponseCode != '' then
					and t.response_code = #ResponseCode#
					$endif
					$if Method != '' then
					and t.method = #Method#
					$endif
					$if HttpHost != '' then
					and t.http_host = #HttpHost#
					$endif
					$if UpstreamAddr != '' then
					and t.upstream_addr = #UpstreamAddr#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_wafAccessLog string = `
	insert into waf_access_log(request_id, waf_ip, host_id,x_real_ip,scheme,response_code,method,http_host,upstream_addr,url,request_length,response_length,state,message)
VALUES(#RequestId#,#WafIp#,#HostId#,#XRealIp#,#Scheme#,#ResponseCode#,#Method#,#HttpHost#,#UpstreamAddr#,#Url#,#RequestLength#,#ResponseLength#,#State#,#Message#)
`

	update_wafAccessLog string = `
	update waf_access_log set
		$if State != '' then
		state = #State#,
		$endif
		$if Message != '' then
		message = #Message#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if RequestId != '' then
		and t.request_id = #RequestId#
		$endif
	`
	delete_wafAccessLog string = `
	update waf_access_log  set
                          status_cd = '1'
                          where status_cd = '0'
		$if RequestId != '' then
		and t.request_id = #RequestId#
		$endif
	`
)

type WafAccessLogDao struct {
}

/**
查询用户
*/
func (*WafAccessLogDao) GetWafAccessLogCount(wafAccessLogDto waf.WafAccessLogDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafAccessLog_count, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafAccessLogDao) GetWafAccessLogs(wafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogDto, error) {
	var wafAccessLogDtos []*waf.WafAccessLogDto
	sqlTemplate.SelectList(query_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		db.Scan(&wafAccessLogDtos)
	}, false)

	return wafAccessLogDtos, nil
}

/**
保存服务sql
*/
func (*WafAccessLogDao) SaveWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Insert(insert_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}

/**
修改服务sql
*/
func (*WafAccessLogDao) UpdateWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Update(update_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}

/**
删除服务sql
*/
func (*WafAccessLogDao) DeleteWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Delete(delete_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}
