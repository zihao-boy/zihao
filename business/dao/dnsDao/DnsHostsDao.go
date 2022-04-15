package dnsDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"gorm.io/gorm"
)

const (
	query_dnsHosts_count string = `
	select count(1) total
from dns_hosts t
					where t.status_cd = '0'
					$if DnsId != '' then
					and t.dns_id = #DnsId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif

	`
	query_dnsHosts string = `
select t.*
from dns_hosts t
					where t.status_cd = '0'
					$if DnsId != '' then
					and t.dns_id = #DnsId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_dnsHosts string = `
	insert into dns_hosts(dns_host_id, dns_id, host_id)
VALUES(#DnsHostId#,#DnsId#,#HostId#)
`

	update_dnsHosts string = `
	update dns_hosts set
		$if HostId != '' then
		host_id = #HostId#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if DnsHostId != '' then
		and dns_host_id = #DnsHostId#
		$endif
		$if DnsId != '' then
		and dns_id = #DnsId#
		$endif
	`
	delete_dnsHosts string = `
	update dns_hosts  set
                          status_cd = '1'
                          where status_cd = '0'
		$if DnsHostId != '' then
		and dns_host_id = #DnsHostId#
		$endif
		$if DnsId != '' then
		and dns_id = #DnsId#
		$endif
	`
)

type DnsHostsDao struct {
}

/**
查询用户
*/
func (*DnsHostsDao) GetDnsHostsCount(dnsHostsDto dns.DnsHostsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_dnsHosts_count, objectConvert.Struct2Map(dnsHostsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*DnsHostsDao) GetDnsHostss(dnsHostsDto dns.DnsHostsDto) ([]*dns.DnsHostsDto, error) {
	var dnsHostsDtos []*dns.DnsHostsDto
	sqlTemplate.SelectList(query_dnsHosts, objectConvert.Struct2Map(dnsHostsDto), func(db *gorm.DB) {
		db.Scan(&dnsHostsDtos)
	}, false)

	return dnsHostsDtos, nil
}

/**
保存服务sql
*/
func (*DnsHostsDao) SaveDnsHosts(dnsHostsDto dns.DnsHostsDto) error {
	return sqlTemplate.Insert(insert_dnsHosts, objectConvert.Struct2Map(dnsHostsDto), false)
}

/**
修改服务sql
*/
func (*DnsHostsDao) UpdateDnsHosts(dnsHostsDto dns.DnsHostsDto) error {
	return sqlTemplate.Update(update_dnsHosts, objectConvert.Struct2Map(dnsHostsDto), false)
}

/**
删除服务sql
*/
func (*DnsHostsDao) DeleteDnsHosts(dnsHostsDto dns.DnsHostsDto) error {
	return sqlTemplate.Delete(delete_dnsHosts, objectConvert.Struct2Map(dnsHostsDto), false)
}
