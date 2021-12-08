package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/appVersionJob"
	"gorm.io/gorm"
)

const (
	query_appVersionJob_count string = `
		select count(1) total
			from app_version_job t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id= #TenantId#
			$endif
			$if JobId != '' then
			and t.job_id = #JobId#
			$endif
			$if JobName != '' then
			and t.job_name = #JobName#
			$endif
			$if State != '' then
			and t.state = #State#
			$endif
    	
	`
	query_appVersionJob string = `
				select t.*,td.name state_name from app_version_job t
left join t_dict td on t.state = td.status_cd and td.table_name = 'app_version_job' and td.table_columns = 'state'
where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id= #TenantId#
					$endif
					$if JobId != '' then
					and t.job_id = #JobId#
					$endif
					$if JobName != '' then
					and t.job_name = #JobName#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersionJob string = `
insert into app_version_job(job_id, job_name, job_shell, tenant_id, pre_job_time, cur_job_time,state) 
VALUES(#JobId#, #JobName#, #JobShell#, #TenantId#, #PreJobTime#, #CurJobTime#,#State#) 
`

	update_appVersionJob string = `
	update app_version_job t set
    $if Name != '' then 
    t.job_name = #Name#,
    $endif
    $if JobShell != '' then 
     t.job_shell =  #JobShell#,
     $endif
    $if CurJobTime != '' then 
     t.cur_job_time =  #CurJobTime#,
     $endif
        $if PreJobTime != '' then 
     t.pre_job_time =  #PreJobTime#,
     $endif
        $if State != '' then 
     t.state =  #State#,
     $endif
     t.status_cd = '0'
	where
		t.status_cd = '0'
	and t.job_id = #JobId#
	and t.tenant_id = #TenantId#
	`
	delete_appVersionJob string = `
	update app_version_job t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if JobId != '' then
			and t.job_id = #JobId#
		$endif
	`

	query_appVersionJobImages_count string = `
		select count(1) total
			from app_version_job t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id= #TenantId#
			$endif
			$if JobId != '' then
			and t.job_id = #JobId#
			$endif
			$if JobName != '' then
			and t.job_name = #JobName#
			$endif
			$if State != '' then
			and t.state = #State#
			$endif
    	
	`
	query_appVersionJobImages string = `
				select t.*,td.name state_name from app_version_job t
left join t_dict td on t.state = td.status_cd and td.table_name = 'app_version_job' and td.table_columns = 'state'
where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id= #TenantId#
					$endif
					$if JobId != '' then
					and t.job_id = #JobId#
					$endif
					$if JobName != '' then
					and t.job_name = #JobName#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersionJobImages string = `
insert into app_version_job(job_id, job_name, job_shell, tenant_id, pre_job_time, cur_job_time,state) 
VALUES(#JobId#, #JobName#, #JobShell#, #TenantId#, #PreJobTime#, #CurJobTime#,#State#) 
`

	update_appVersionJobImages string = `
	update app_version_job t set
    $if Name != '' then 
    t.job_name = #Name#,
    $endif
    $if JobShell != '' then 
     t.job_shell =  #JobShell#,
     $endif
    $if CurJobTime != '' then 
     t.cur_job_time =  #CurJobTime#,
     $endif
        $if PreJobTime != '' then 
     t.pre_job_time =  #PreJobTime#,
     $endif
        $if State != '' then 
     t.state =  #State#,
     $endif
     t.status_cd = '0'
	where
		t.status_cd = '0'
	and t.job_id = #JobId#
	and t.tenant_id = #TenantId#
	`
	delete_appVersionJobImages string = `
	update app_version_job t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if JobId != '' then
			and t.job_id = #JobId#
		$endif
	`
)

type AppVersionJobDao struct {
}

/**
查询用户
*/
func (*AppVersionJobDao) GetAppVersionJobCount(appVersionJobDto appVersionJob.AppVersionJobDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appVersionJob_count, objectConvert.Struct2Map(appVersionJobDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppVersionJobDao) GetAppVersionJobs(appVersionJobDto appVersionJob.AppVersionJobDto) ([]*appVersionJob.AppVersionJobDto, error) {
	var appVersionJobDtos []*appVersionJob.AppVersionJobDto
	sqlTemplate.SelectList(query_appVersionJob, objectConvert.Struct2Map(appVersionJobDto), func(db *gorm.DB) {
		db.Scan(&appVersionJobDtos)
	}, false)

	return appVersionJobDtos, nil
}

/**
保存服务sql
*/
func (*AppVersionJobDao) SaveAppVersionJob(appVersionJobDto appVersionJob.AppVersionJobDto) error {
	return sqlTemplate.Insert(insert_appVersionJob, objectConvert.Struct2Map(appVersionJobDto), false)
}

/**
修改服务sql
*/
func (*AppVersionJobDao) UpdateAppVersionJob(appVersionJobDto appVersionJob.AppVersionJobDto) error {
	return sqlTemplate.Update(update_appVersionJob, objectConvert.Struct2Map(appVersionJobDto), false)
}

/**
删除服务sql
*/
func (*AppVersionJobDao) DeleteAppVersionJob(appVersionJobDto appVersionJob.AppVersionJobDto) error {
	return sqlTemplate.Delete(delete_appVersionJob, objectConvert.Struct2Map(appVersionJobDto), false)
}

func (d *AppVersionJobDao) SaveAppVersionJobImages(imagesDto appVersionJob.AppVersionJobImagesDto) error {
	return sqlTemplate.Insert(insert_appVersionJobImages, objectConvert.Struct2Map(imagesDto), false)
}

func (d *AppVersionJobDao) UpdateAppVersionJobImages(imagesDto appVersionJob.AppVersionJobImagesDto) error {
	return sqlTemplate.Update(update_appVersionJobImages, objectConvert.Struct2Map(imagesDto), false)
}

func (d *AppVersionJobDao) GetAppVersionJobImagesCount(imagesDto appVersionJob.AppVersionJobImagesDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appVersionJobImages_count, objectConvert.Struct2Map(imagesDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

func (d *AppVersionJobDao) GetAppVersionJobImages(imagesDto appVersionJob.AppVersionJobImagesDto) ([]*appVersionJob.AppVersionJobImagesDto, error) {
	var appVersionJobImagesDtos []*appVersionJob.AppVersionJobImagesDto
	sqlTemplate.SelectList(query_appVersionJobImages, objectConvert.Struct2Map(imagesDto), func(db *gorm.DB) {
		db.Scan(&appVersionJobImagesDtos)
	}, false)

	return appVersionJobImagesDtos, nil
}

func (d *AppVersionJobDao) DeleteAppVersionJobImages(imagesDto appVersionJob.AppVersionJobImagesDto) error {
	return sqlTemplate.Delete(delete_appVersionJobImages, objectConvert.Struct2Map(imagesDto), false)
}
