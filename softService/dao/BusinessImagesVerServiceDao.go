package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"gorm.io/gorm"
)

const (
	query_businessImagesVer_count string = `
	select count(1) total
	from business_images_ver t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Version != '' then
	and t.version like '%' || #Version# || '%'
	$endif
	$if ImagesId != '' then
	and t.images_id =  #ImagesId# 
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_businessImagesVer string = `
				select t.*
				from business_images_ver t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if Version != '' then
					and t.version like '%' || #Version# || '%'
					$endif
					$if ImagesId != '' then
					and t.images_id =  #ImagesId# 
					$endif
					$if Id != '' then
					and t.id = #Id#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`
	insert_businessImagesVer string = `
	insert into business_images_ver(id, images_id, version,type_url,tenant_id)
VALUES(#Id#,#ImagesId#,#Version#,#TypeUrl#,#TenantId#)
`

	update_businessImagesVer string = `
	update business_images_ver set
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
	delete_businessImagesVer string = `
	update business_images_ver  set
                          status_cd = '1'
                          where status_cd = '0'
						  and id = #Id#
	`
)

type BusinessImagesVerDao struct {
}

/**
查询用户
*/
func (*BusinessImagesVerDao) GetBusinessImagesVerCount(businessImagesVerDto businessImages.BusinessImagesVerDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessImagesVer_count, objectConvert.Struct2Map(businessImagesVerDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*BusinessImagesVerDao) GetBusinessImagesVers(businessImagesVerDto businessImages.BusinessImagesVerDto) ([]*businessImages.BusinessImagesVerDto, error) {
	var businessImagesVerDtos []*businessImages.BusinessImagesVerDto
	sqlTemplate.SelectList(query_businessImagesVer, objectConvert.Struct2Map(businessImagesVerDto), func(db *gorm.DB) {
		db.Scan(&businessImagesVerDtos)
	}, false)

	return businessImagesVerDtos, nil
}

/**
保存服务sql
*/
func (*BusinessImagesVerDao) SaveBusinessImagesVer(businessImagesVerDto businessImages.BusinessImagesVerDto) error {
	return sqlTemplate.Insert(insert_businessImagesVer, objectConvert.Struct2Map(businessImagesVerDto), false)
}

/**
修改服务sql
*/
func (*BusinessImagesVerDao) UpdateBusinessImagesVer(businessImagesVerDto businessImages.BusinessImagesVerDto) error {
	return sqlTemplate.Update(update_businessImagesVer, objectConvert.Struct2Map(businessImagesVerDto), false)
}

/**
删除服务sql
*/
func (*BusinessImagesVerDao) DeleteBusinessImagesVer(businessImagesVerDto businessImages.BusinessImagesVerDto) error {
	return sqlTemplate.Delete(delete_businessImagesVer, objectConvert.Struct2Map(businessImagesVerDto), false)
}
