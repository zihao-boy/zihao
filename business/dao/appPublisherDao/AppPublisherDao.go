package appPublisherDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	appPublisher "github.com/zihao-boy/zihao/entity/dto/appPublisherDto"
	"gorm.io/gorm"
)

const (
	query_appPublisher_count string = `
	select count(1) total
	from app_publisher t
	where t.status_cd = '0'
	$if PublisherId != '' then
	and t.publisher_id = #PublisherId#
	$endif
	$if Username != '' then
	and t.username = #Username#
	$endif
	$if ExtPublisherId != '' then
	and t.ext_publisher_id = #ExtPublisherId#
	$endif
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
    	
	`
	query_appPublisher string = `
				select t.*
				from app_publisher t
				where t.status_cd = '0'
				$if PublisherId != '' then
				and t.publisher_id = #PublisherId#
				$endif
				$if Username != '' then
				and t.username = #Username#
				$endif
				$if ExtPublisherId != '' then
				and t.ext_publisher_id = #ExtPublisherId#
				$endif
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appPublisher string = `
	insert into app_publisher(publisher_id, username, email, token,phone,state,tenant_id,ext_publisher_id)
VALUES(#PublisherId#,#Username#,#Email#,#Token#,#Phone#,#State#,#TenantId#,#ExtPublisherId#)
`

	update_appPublisher string = `
	update app_publisher set
		$if Token != '' then
		token = #Token#,
		$endif
		$if Username != '' then
		username = #Username#,
		$endif
		$if Email != '' then
		email = #Email#,
		$endif
		$if Phone != '' then
		phone = #Phone#,
		$endif
		  $if ExtPublisherId != '' then
		   ext_publisher_id = #ExtPublisherId#,
		  $endif
		status_cd = '0'
		where status_cd = '0'
		 $if PublisherId != '' then
		  and publisher_id = #PublisherId#
		  $endif
	`
	delete_appPublisher string = `
	update app_publisher  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if PublisherId != '' then
						  and publisher_id = #PublisherId#
						  $endif
                          $if ExtPublisherId != '' then
						  and ext_publisher_id = #ExtPublisherId#
						  $endif
	`
)

type AppPublisherDao struct {
}

/**
查询用户
*/
func (*AppPublisherDao) GetAppPublisherCount(appPublisherDto appPublisher.AppPublisherDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appPublisher_count, objectConvert.Struct2Map(appPublisherDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppPublisherDao) GetAppPublishers(appPublisherDto appPublisher.AppPublisherDto) ([]*appPublisher.AppPublisherDto, error) {
	var appPublisherDtos []*appPublisher.AppPublisherDto
	sqlTemplate.SelectList(query_appPublisher, objectConvert.Struct2Map(appPublisherDto), func(db *gorm.DB) {
		db.Scan(&appPublisherDtos)
	}, false)

	return appPublisherDtos, nil
}

/**
保存服务sql
*/
func (*AppPublisherDao) SaveAppPublisher(appPublisherDto appPublisher.AppPublisherDto) error {
	return sqlTemplate.Insert(insert_appPublisher, objectConvert.Struct2Map(appPublisherDto), false)
}

/**
修改服务sql
*/
func (*AppPublisherDao) UpdateAppPublisher(appPublisherDto appPublisher.AppPublisherDto) error {
	return sqlTemplate.Update(update_appPublisher, objectConvert.Struct2Map(appPublisherDto), false)
}

/**
删除服务sql
*/
func (*AppPublisherDao) DeleteAppPublisher(appPublisherDto appPublisher.AppPublisherDto) error {
	return sqlTemplate.Delete(delete_appPublisher, objectConvert.Struct2Map(appPublisherDto), false)
}
