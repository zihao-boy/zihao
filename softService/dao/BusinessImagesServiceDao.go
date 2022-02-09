package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"gorm.io/gorm"
)

const (
	query_businessImages_count string = `
	select count(1) total
	from business_images t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name like '%' || #Name# || '%'
	$endif
	$if Version != '' then
	and t.version like '%' || #Version# || '%'
	$endif
	$if ImagesType != '' then
	and t.images_type =  #ImagesType# 
	$endif
	$if ImagesFlag != '' then
	and t.images_flag =  #ImagesFlag# 
	$endif
	$if CreateUserId != '' then
	and t.create_user_id = #CreateUserId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_businessImages string = `
				select t.*,uu.username,bie.ext_images_id,bie.ext_publisher_id,bie.app_id,bie.app_name,ap.publisher_id
				from business_images t
				left join u_user uu on t.create_user_id = uu.user_id and uu.status_cd = '0'
				left join business_images_ext bie on t.id = bie.images_id and bie.status_cd = '0'
				left join app_publisher ap on bie.ext_publisher_id = ap.ext_publisher_id and ap.status_cd = '0'
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name like '%' || #Name# || '%'
				$endif
				$if Version != '' then
				and t.version like '%' || #Version# || '%'
				$endif
				$if CreateUserId != '' then
				and t.create_user_id = #CreateUserId#
				$endif
				$if ImagesType != '' then
				and t.images_type =  #ImagesType# 
				$endif
				$if ImagesFlag != '' then
				and t.images_flag =  #ImagesFlag# 
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`
	insert_businessImages string = `
	insert into business_images(id, name, version, images_type,type_url,images_flag, create_user_id,tenant_id)
VALUES(#Id#,#Name#,#Version#,#ImagesType#,#TypeUrl#,#ImagesFlag#,#CreateUserId#,#TenantId#)
`

	update_businessImages string = `
	update business_images set
		$if Name != '' then
		name = #Name#,
		$endif
		$if Version != '' then
		version = #Version#,
		$endif
		$if TypeUrl != '' then
		type_url = #TypeUrl#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_businessImages string = `
	update business_images  set
                          status_cd = '1'
                          where status_cd = '0'
						  and id = #Id#
	`
)

type BusinessImagesDao struct {
}

/**
查询用户
*/
func (*BusinessImagesDao) GetBusinessImagesCount(businessImagesDto businessImages.BusinessImagesDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessImages_count, objectConvert.Struct2Map(businessImagesDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*BusinessImagesDao) GetBusinessImagess(businessImagesDto businessImages.BusinessImagesDto) ([]*businessImages.BusinessImagesDto, error) {
	var businessImagesDtos []*businessImages.BusinessImagesDto
	sqlTemplate.SelectList(query_businessImages, objectConvert.Struct2Map(businessImagesDto), func(db *gorm.DB) {
		db.Scan(&businessImagesDtos)
	}, false)

	return businessImagesDtos, nil
}

/**
保存服务sql
*/
func (*BusinessImagesDao) SaveBusinessImages(businessImagesDto businessImages.BusinessImagesDto) error {
	return sqlTemplate.Insert(insert_businessImages, objectConvert.Struct2Map(businessImagesDto), false)
}

/**
修改服务sql
*/
func (*BusinessImagesDao) UpdateBusinessImages(businessImagesDto businessImages.BusinessImagesDto) error {
	return sqlTemplate.Update(update_businessImages, objectConvert.Struct2Map(businessImagesDto), false)
}

/**
删除服务sql
*/
func (*BusinessImagesDao) DeleteBusinessImages(businessImagesDto businessImages.BusinessImagesDto) error {
	return sqlTemplate.Delete(delete_businessImages, objectConvert.Struct2Map(businessImagesDto), false)
}
