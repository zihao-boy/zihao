package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafRoute_count string = `
	select count(1) total
from waf_route t
					where t.status_cd = '0'
					$if WafId != '' then
					and t.waf_id = #WafId#
					$endif
					$if RouteId != '' then
					and t.route_id = #RouteId#
					$endif
					$if Hostname != '' then
					and t.hostname = #Hostname#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if Port != '' then
					and t.port = #Port#
					$endif

	`
	query_wafRoute string = `
select t.*
from waf_route t
					where t.status_cd = '0'
					$if WafId != '' then
					and t.waf_id = #WafId#
					$endif
					$if RouteId != '' then
					and t.route_id = #RouteId#
					$endif
					$if Hostname != '' then
					and t.hostname = #Hostname#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if Port != '' then
					and t.port = #Port#
					$endif
	`

	insert_wafRoute string = `
	insert into waf_route(route_id, waf_id, hostname,ip,port,scheme)
VALUES(#RouteId#,#WafId#,#Hostname#,#Ip#,#Port#,#Scheme#)
`

	update_wafRoute string = `
	update waf_route set
		$if Hostname != '' then
		hostname = #Hostname#,
		$endif
		$if Ip != '' then
		ip = #Ip#,
		$endif
		$if Port != '' then
		port = #Port#,
		$endif
		$if Scheme != '' then
        scheme = #Scheme#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if RouteId != '' then
		and route_id = #RouteId#
		$endif
		$if WafId != '' then
		and waf_id = #WafId#
		$endif
	`
	delete_wafRoute string = `
	update waf_route  set
                          status_cd = '1'
                          where status_cd = '0'
		$if RouteId != '' then
		and route_id = #RouteId#
		$endif
		$if WafId != '' then
		and waf_id = #WafId#
		$endif
	`
)

type WafRouteDao struct {
}

/**
查询用户
*/
func (*WafRouteDao) GetWafRouteCount(wafRouteDto waf.WafRouteDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafRoute_count, objectConvert.Struct2Map(wafRouteDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafRouteDao) GetWafRoutes(wafRouteDto waf.WafRouteDto) ([]*waf.WafRouteDto, error) {
	var wafRouteDtos []*waf.WafRouteDto
	sqlTemplate.SelectList(query_wafRoute, objectConvert.Struct2Map(wafRouteDto), func(db *gorm.DB) {
		db.Scan(&wafRouteDtos)
	}, false)

	return wafRouteDtos, nil
}

/**
保存服务sql
*/
func (*WafRouteDao) SaveWafRoute(wafRouteDto waf.WafRouteDto) error {
	return sqlTemplate.Insert(insert_wafRoute, objectConvert.Struct2Map(wafRouteDto), false)
}

/**
修改服务sql
*/
func (*WafRouteDao) UpdateWafRoute(wafRouteDto waf.WafRouteDto) error {
	return sqlTemplate.Update(update_wafRoute, objectConvert.Struct2Map(wafRouteDto), false)
}

/**
删除服务sql
*/
func (*WafRouteDao) DeleteWafRoute(wafRouteDto waf.WafRouteDto) error {
	return sqlTemplate.Delete(delete_wafRoute, objectConvert.Struct2Map(wafRouteDto), false)
}
