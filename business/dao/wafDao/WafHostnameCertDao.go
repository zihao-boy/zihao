package wafDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafHostnameCert_count string = `
	select count(1) total
from waf_hostname_cert t
					where t.status_cd = '0'
					$if CertId != '' then
					and t.cert_id = #CertId#
					$endif
					$if Hostname != '' then
					and t.hostname = #Hostname#
					$endif

	`
	query_wafHostnameCert string = `
select t.*
from waf_hostname_cert t
					where t.status_cd = '0'
					$if CertId != '' then
					and t.cert_id = #CertId#
					$endif
					$if Hostname != '' then
					and t.hostname = #Hostname#
					$endif
	`

	insert_wafHostnameCert string = `
	insert into waf_hostname_cert(cert_id, hostname,cert_content,priv_key_content)
VALUES(#CertId#,#Hostname#,#CertContent#,#PrivKeyContent#)
`

	update_wafHostnameCert string = `
	update waf_hostname_cert set
		$if CertContent != '' then
		cert_content = #CertContent#,
		$endif
		$if PrivKeyContent != '' then
		priv_key_content = #PrivKeyContent#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if CertId != '' then
		and cert_id = #CertId#
		$endif
		$if Hostname != '' then
		and hostname = #Hostname#
		$endif
	`
	delete_wafHostnameCert string = `
	update waf_hostname_cert  set
                          status_cd = '1'
                          where status_cd = '0'
		$if CertId != '' then
		and cert_id = #CertId#
		$endif
		$if Hostname != '' then
		and hostname = #Hostname#
		$endif
	`
)

type WafHostnameCertDao struct {
}

/**
查询用户
*/
func (*WafHostnameCertDao) GetWafHostnameCertCount(wafHostnameCertDto waf.WafHostnameCertDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafHostnameCert_count, objectConvert.Struct2Map(wafHostnameCertDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WafHostnameCertDao) GetWafHostnameCerts(wafHostnameCertDto waf.WafHostnameCertDto) ([]*waf.WafHostnameCertDto, error) {
	var wafHostnameCertDtos []*waf.WafHostnameCertDto
	sqlTemplate.SelectList(query_wafHostnameCert, objectConvert.Struct2Map(wafHostnameCertDto), func(db *gorm.DB) {
		db.Scan(&wafHostnameCertDtos)
	}, false)

	return wafHostnameCertDtos, nil
}

/**
保存服务sql
*/
func (*WafHostnameCertDao) SaveWafHostnameCert(wafHostnameCertDto waf.WafHostnameCertDto) error {
	return sqlTemplate.Insert(insert_wafHostnameCert, objectConvert.Struct2Map(wafHostnameCertDto), false)
}

/**
修改服务sql
*/
func (*WafHostnameCertDao) UpdateWafHostnameCert(wafHostnameCertDto waf.WafHostnameCertDto) error {
	return sqlTemplate.Update(update_wafHostnameCert, objectConvert.Struct2Map(wafHostnameCertDto), false)
}

/**
删除服务sql
*/
func (*WafHostnameCertDao) DeleteWafHostnameCert(wafHostnameCertDto waf.WafHostnameCertDto) error {
	return sqlTemplate.Delete(delete_wafHostnameCert, objectConvert.Struct2Map(wafHostnameCertDto), false)
}
