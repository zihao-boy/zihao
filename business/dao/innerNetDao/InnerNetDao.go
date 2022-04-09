package innerNetDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"gorm.io/gorm"
)

const (
	query_inner_net_count string = `
	select count(1) total
from inner_net t
					where t.status_cd = '0'
					$if Tun != '' then
					and t.tun = #Tun#
					$endif
					$if TunName != '' then
					and t.tun_name = #TunName#
					$endif
					$if Dns != '' then
					and t.dns = #Dns#
					$endif
					$if Protocol != '' then
					and t.protocol = #Protocol#
					$endif
					$if InnerNetPort != '' then
					and t.inner_net_port = #InnerNetPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if InnerNetId != '' then
					and t.inner_net_id = #InnerNetId#
					$endif

	`
	query_inner_net string = `
select t.*
from inner_net t
					where t.status_cd = '0'
					$if Tun != '' then
					and t.tun = #Tun#
					$endif
					$if TunName != '' then
					and t.tun_name = #TunName#
					$endif
					$if Dns != '' then
					and t.dns = #Dns#
					$endif
					$if Protocol != '' then
					and t.protocol = #Protocol#
					$endif
					$if InnerNetPort != '' then
					and t.inner_net_port = #InnerNetPort#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if InnerNetId != '' then
					and t.inner_net_id = #InnerNetId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_inner_net string = `
	insert into inner_net(inner_net_id, inner_net_port, tun,tun_name,dns,protocol, state)
VALUES(#InnerNetId#,#InnerNetPort#,#Tun#,#TunName#,#Dns#,#Protocol#,#State#)
`

	update_inner_net string = `
	update inner_net set
		$if InnerNetPort != '' then
		inner_net_port = #InnerNetPort#,
		$endif
		$if Tun != '' then
		tun = #Tun#,
		$endif
		$if TunName != '' then
		tun_name = #TunName#,
		$endif
		$if Dns != '' then
		dns = #Dns#,
		$endif
		$if Protocol != '' then
		protocol = #Protocol#,
		$endif
        $if State != '' then
		state = #State#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if InnerNetId != '' then
		and inner_net_id = #InnerNetId#
		$endif
	`
	delete_inner_net string = `
	update inner_net  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if InnerNetId != '' then
						  and inner_net_id = #InnerNetId#
						  $endif
	`
)

type InnerNetDao struct {
}

/**
查询用户
*/
func (*InnerNetDao) GetInnerNetCount(inner_netDto innerNet.InnerNetDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_inner_net_count, objectConvert.Struct2Map(inner_netDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InnerNetDao) GetInnerNets(inner_netDto innerNet.InnerNetDto) ([]*innerNet.InnerNetDto, error) {
	var inner_netDtos []*innerNet.InnerNetDto
	sqlTemplate.SelectList(query_inner_net, objectConvert.Struct2Map(inner_netDto), func(db *gorm.DB) {
		db.Scan(&inner_netDtos)
	}, false)

	return inner_netDtos, nil
}

/**
保存服务sql
*/
func (*InnerNetDao) SaveInnerNet(inner_netDto innerNet.InnerNetDto) error {
	return sqlTemplate.Insert(insert_inner_net, objectConvert.Struct2Map(inner_netDto), false)
}

/**
修改服务sql
*/
func (*InnerNetDao) UpdateInnerNet(inner_netDto innerNet.InnerNetDto) error {
	return sqlTemplate.Update(update_inner_net, objectConvert.Struct2Map(inner_netDto), false)
}

/**
删除服务sql
*/
func (*InnerNetDao) DeleteInnerNet(inner_netDto innerNet.InnerNetDto) error {
	return sqlTemplate.Delete(delete_inner_net, objectConvert.Struct2Map(inner_netDto), false)
}
