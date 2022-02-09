package businessImagesExtDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	businessImagesExt "github.com/zihao-boy/zihao/entity/dto/businessImages"
	"gorm.io/gorm"
)

const (
	query_businessImagesExt_count string = `
	select count(1) total
	from business_images_ext t
	where t.status_cd = '0'
	$if Id != '' then
	and t.id = #Id#
	$endif
	$if ImagesId != '' then
	and t.images_id = #ImagesId#
	$endif
	$if ExtPublisherId != '' then
	and t.ext_publisher_id = #ExtPublisherId#
	$endif
	$if AppId != '' then
	and t.app_id = #AppId#
	$endif
	$if ExtImagesId != '' then
	and t.ext_images_id = #ExtImagesId#
	$endif
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
    	
	`
	query_businessImagesExt string = `
				select t.*
				from business_images_ext t
				where t.status_cd = '0'
				$if Id != '' then
				and t.id = #Id#
				$endif
				$if ImagesId != '' then
				and t.images_id = #ImagesId#
				$endif
				$if ExtPublisherId != '' then
				and t.ext_publisher_id = #ExtPublisherId#
				$endif
				$if AppId != '' then
				and t.app_id = #AppId#
				$endif
				$if ExtImagesId != '' then
				and t.ext_images_id = #ExtImagesId#
				$endif
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_businessImagesExt string = `
	insert into business_images_ext(id, images_id, app_id, app_name,ext_images_id,ext_publisher_id,tenant_id)
VALUES(#Id#,#ImagesId#,#AppId#,#AppName#,#ExtImagesId#,#ExtPublisherId#,#TenantId#)
`

	update_businessImagesExt string = `
	update business_images_ext set
		$if AppName != '' then
		app_name = #AppName#,
		$endif
		$if AppId != '' then
		app_id = #AppId#,
		$endif
		$if ExtImagesId != '' then
		ext_images_id = #ExtImagesId#,
		$endif
		  $if ExtPublisherId != '' then
		   ext_publisher_id = #ExtPublisherId#,
		  $endif
		status_cd = '0'
		where status_cd = '0'
		 $if Id != '' then
		  and id = #Id#
		  $endif
          $if ImagesId != '' then
		  and images_id = #ImagesId#
		  $endif
	`
	delete_businessImagesExt string = `
	update business_images_ext  set
                          status_cd = '1'
                          where status_cd = '0'
						 $if Id != '' then
						  and id = #Id#
						  $endif
						  $if ImagesId != '' then
						  and images_id = #ImagesId#
						  $endif
	`
)

type BusinessImagesExtDao struct {
}

/**
查询用户
*/
func (*BusinessImagesExtDao) GetBusinessImagesExtCount(businessImagesExtDto businessImagesExt.BusinessImagesExtDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessImagesExt_count, objectConvert.Struct2Map(businessImagesExtDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*BusinessImagesExtDao) GetBusinessImagesExts(businessImagesExtDto businessImagesExt.BusinessImagesExtDto) ([]*businessImagesExt.BusinessImagesExtDto, error) {
	var businessImagesExtDtos []*businessImagesExt.BusinessImagesExtDto
	sqlTemplate.SelectList(query_businessImagesExt, objectConvert.Struct2Map(businessImagesExtDto), func(db *gorm.DB) {
		db.Scan(&businessImagesExtDtos)
	}, false)

	return businessImagesExtDtos, nil
}

/**
保存服务sql
*/
func (*BusinessImagesExtDao) SaveBusinessImagesExt(businessImagesExtDto businessImagesExt.BusinessImagesExtDto) error {
	return sqlTemplate.Insert(insert_businessImagesExt, objectConvert.Struct2Map(businessImagesExtDto), false)
}

/**
修改服务sql
*/
func (*BusinessImagesExtDao) UpdateBusinessImagesExt(businessImagesExtDto businessImagesExt.BusinessImagesExtDto) error {
	return sqlTemplate.Update(update_businessImagesExt, objectConvert.Struct2Map(businessImagesExtDto), false)
}

/**
删除服务sql
*/
func (*BusinessImagesExtDao) DeleteBusinessImagesExt(businessImagesExtDto businessImagesExt.BusinessImagesExtDto) error {
	return sqlTemplate.Delete(delete_businessImagesExt, objectConvert.Struct2Map(businessImagesExtDto), false)
}
