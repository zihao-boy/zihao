package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVersionJob"
	"gorm.io/gorm"
)

const (
	query_appVersionJobDetail_count string = `
		select count(1) total
			from app_version_job_detail t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id= #TenantId#
					$endif
					$if JobId != '' then
					and t.job_id = #JobId#
					$endif
					$if DetailId != '' then
					and t.detail_id = #DetailId#
					$endif
    	
	`
	query_appVersionJobDetail string = `
				select * from app_version_job_detail t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id= #TenantId#
					$endif
					$if JobId != '' then
					and t.job_id = #JobId#
					$endif
					$if DetailId != '' then
					and t.detail_id = #DetailId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersionJobDetail string = `
insert into app_version_job_detail(detail_id, job_id, log_path, tenant_id, state)
VALUES (#DetailId#, #JobId#, #LogPath#, #TenantId#, #State#)
`

	update_appVersionJobDetail string = `
	update app_version_job_detail t set
	$if State != '' then 
     t.state =  #State#,
     $endif
     t.status_cd = '0'
	where
		t.status_cd = '0'
	and t.detail_id = #DetailId#
	and t.tenant_id = #TenantId#
	`
	delete_appVersionJobDetail string = `
	update app_version_job_detail t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if DetailId != '' then
			t.detail_id = #DetailId#
		$endif
	`
)

type AppVersionJobDetailDao struct {
}

/**
查询用户
*/
func (*AppVersionJobDetailDao) GetAppVersionJobDetailCount(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appVersionJobDetail_count, objectConvert.Struct2Map(appVersionJobDetailDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppVersionJobDetailDao) GetAppVersionJobDetails(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) ([]*appVersionJob.AppVersionJobDetailDto, error) {
	var appVersionJobDetailDtos []*appVersionJob.AppVersionJobDetailDto
	sqlTemplate.SelectList(query_appVersionJobDetail, objectConvert.Struct2Map(appVersionJobDetailDto), func(db *gorm.DB) {
		db.Scan(&appVersionJobDetailDtos)
	}, false)

	return appVersionJobDetailDtos, nil
}

/**
保存服务sql
*/
func (*AppVersionJobDetailDao) SaveAppVersionJobDetail(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) error {
	return sqlTemplate.Insert(insert_appVersionJobDetail, objectConvert.Struct2Map(appVersionJobDetailDto), false)
}

/**
修改服务sql
*/
func (*AppVersionJobDetailDao) UpdateAppVersionJobDetail(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) error {
	return sqlTemplate.Update(update_appVersionJobDetail, objectConvert.Struct2Map(appVersionJobDetailDto), false)
}

/**
删除服务sql
*/
func (*AppVersionJobDetailDao) DeleteAppVersionJobDetail(appVersionJobDetailDto appVersionJob.AppVersionJobDetailDto) error {
	return sqlTemplate.Delete(delete_appVersionJobDetail, objectConvert.Struct2Map(appVersionJobDetailDto), false)
}
