package dnsDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"gorm.io/gorm"
)

const (
	query_dnsMap_count string = `
	select count(1) total
from dns_map t
					where t.status_cd = '0'
					$if DnsMapId != '' then
					and t.dns_map_id = #DnsMapId#
					$endif
					$if Host != '' then
					and t.host = #Host#
					$endif
					$if Type != '' then
					and t.type = #Type#
					$endif
					$if Value != '' then
					and t.value = #Value#
					$endif
				
	`
	query_dnsMap string = `
select t.*
from dns_map t
					where t.status_cd = '0'
					$if DnsMapId != '' then
					and t.dns_map_id = #DnsMapId#
					$endif
					$if Host != '' then
					and t.host = #Host#
					$endif
					$if Type != '' then
					and t.type = #Type#
					$endif
					$if Value != '' then
					and t.value = #Value#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_dnsMap string = `
	insert into dns_map(dns_map_id,host,type,value)
VALUES(#DnsMapId#,#Host#,#Type#,#Value#)
`

	update_dnsMap string = `
	update dns_map set
					$if Host != '' then
					 host = #Host#,
					$endif
					$if Type != '' then
					 type = #Type#,
					$endif
					$if Value != '' then
					 value = #Value#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if DnsMapId != '' then
					and dns_map_id = #DnsMapId#
					$endif
	`
	delete_dnsMap string = `
	update dns_map  set
                          status_cd = '1'
                          where status_cd = '0'
					$if DnsMapId != '' then
					and dns_map_id = #DnsMapId#
					$endif
	`
)

type DnsMapDao struct {
}

/**
查询用户
*/
func (*DnsMapDao) GetDnsMapCount(dnsMapDto dns.DnsMapDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_dnsMap_count, objectConvert.Struct2Map(dnsMapDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*DnsMapDao) GetDnsMaps(dnsMapDto dns.DnsMapDto) ([]*dns.DnsMapDto, error) {
	var dnsMapDtos []*dns.DnsMapDto
	sqlTemplate.SelectList(query_dnsMap, objectConvert.Struct2Map(dnsMapDto), func(db *gorm.DB) {
		db.Scan(&dnsMapDtos)
	}, false)

	return dnsMapDtos, nil
}

/**
保存服务sql
*/
func (*DnsMapDao) SaveDnsMap(dnsMapDto dns.DnsMapDto) error {
	return sqlTemplate.Insert(insert_dnsMap, objectConvert.Struct2Map(dnsMapDto), false)
}

/**
修改服务sql
*/
func (*DnsMapDao) UpdateDnsMap(dnsMapDto dns.DnsMapDto) error {
	return sqlTemplate.Update(update_dnsMap, objectConvert.Struct2Map(dnsMapDto), false)
}

/**
删除服务sql
*/
func (*DnsMapDao) DeleteDnsMap(dnsMapDto dns.DnsMapDto) error {
	return sqlTemplate.Delete(delete_dnsMap, objectConvert.Struct2Map(dnsMapDto), false)
}
