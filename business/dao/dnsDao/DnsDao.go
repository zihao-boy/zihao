package dnsDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/dns"
	"gorm.io/gorm"
)

const (
	query_dns_count string = `
	select count(1) total
from dns t
					where t.status_cd = '0'
					$if DnsIp != '' then
					and t.dnsIp = #DnsIp#
					$endif
					$if DnsPort != '' then
					and t.dns_port = #DnsPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if DnsId != '' then
					and t.dns_id = #DnsId#
					$endif

	`
	query_dns string = `
select t.*
from dns t
					where t.status_cd = '0'
					$if DnsIp != '' then
					and t.dnsIp = #DnsIp#
					$endif
					$if DnsPort != '' then
					and t.dns_port = #DnsPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if DnsId != '' then
					and t.dns_id = #DnsId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_dns string = `
	insert into dns(dns_id, dns_port, dns_ip, state)
VALUES(#DnsId#,#DnsPort#,#DnsIp#,#State#)
`

	update_dns string = `
	update dns set
		$if DnsPort != '' then
		dns_port = #DnsPort#,
		$endif
		$if DnsIp != '' then
		dns_ip = #DnsIp#,
		$endif
        $if State != '' then
		state = #State#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if DnsId != '' then
		and dns_id = #DnsId#
		$endif
	`
	delete_dns string = `
	update dns  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if DnsId != '' then
						  and dns_id = #DnsId#
						  $endif
	`
)

type DnsDao struct {
}

/**
查询用户
*/
func (*DnsDao) GetDnsCount(dnsDto dns.DnsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_dns_count, objectConvert.Struct2Map(dnsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*DnsDao) GetDnss(dnsDto dns.DnsDto) ([]*dns.DnsDto, error) {
	var dnsDtos []*dns.DnsDto
	sqlTemplate.SelectList(query_dns, objectConvert.Struct2Map(dnsDto), func(db *gorm.DB) {
		db.Scan(&dnsDtos)
	}, false)

	return dnsDtos, nil
}

/**
保存服务sql
*/
func (*DnsDao) SaveDns(dnsDto dns.DnsDto) error {
	return sqlTemplate.Insert(insert_dns, objectConvert.Struct2Map(dnsDto), false)
}

/**
修改服务sql
*/
func (*DnsDao) UpdateDns(dnsDto dns.DnsDto) error {
	return sqlTemplate.Update(update_dns, objectConvert.Struct2Map(dnsDto), false)
}

/**
删除服务sql
*/
func (*DnsDao) DeleteDns(dnsDto dns.DnsDto) error {
	return sqlTemplate.Delete(delete_dns, objectConvert.Struct2Map(dnsDto), false)
}
