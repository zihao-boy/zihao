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
insert into app_version_job(job_id, job_name, job_shell, tenant_id, state,git_url,git_username,git_passwd,work_dir,job_time) 
VALUES(#JobId#, #JobName#, #JobShell#, #TenantId#, #State#,#GitUrl#,#GitUsername#,#GitPasswd#,#WorkDir#,#JobTime#) 
`

	update_appVersionJob string = `
	update app_version_job set
    $if Name != '' then 
    job_name = #Name#,
    $endif
    $if JobShell != '' then 
     job_shell =  #JobShell#,
     $endif
     $if State != '' then 
     state =  #State#,
     $endif
  $if GitUrl != '' then 
     git_url =  #GitUrl#,
     $endif
  $if GitUsername != '' then 
     git_username =  #GitUsername#,
     $endif
  $if WorkDir != '' then 
     work_dir =  #WorkDir#,
     $endif
$if GitPasswd != '' then 
     git_passwd =  #GitPasswd#,
     $endif
        $if JobTime != '' then 
     job_time =  #JobTime#,
     $endif
     status_cd = '0'
	where
		status_cd = '0'
	and job_id = #JobId#
	and tenant_id = #TenantId#
	`
	delete_appVersionJob string = `
	update app_version_job  set
                          status_cd = '1'
                          where status_cd = '0'
		$if JobId != '' then
			and job_id = #JobId#
		$endif
	`
	query_appVersionJobImages_count string = `
		select count(1) total
			from app_version_job_images t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id= #TenantId#
			$endif
			$if JobImagesId != '' then
			and t.job_images_id = #JobImagesId#
			$endif
			$if JobId != '' then
			and t.job_id = #JobId#
			$endif
	`

	query_appVersionJobImages string = `
				select t.*,bd.name business_dockerfile_name,bp.name business_package_name
from app_version_job_images t
left join business_dockerfile bd on t.business_dockerfile_id = bd.id and bd.status_cd = '0'
left join business_package bp on t.business_package_id = bp.id and bp.status_cd = '0'
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id= #TenantId#
					$endif
					$if JobImagesId != '' then
					and t.job_images_id = #JobImagesId#
					$endif
					$if JobId != '' then
					and t.job_id = #JobId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersionJobImages string = `
insert into app_version_job_images(job_images_id, tenant_id, package_url, business_package_id, business_dockerfile_id,job_id) 
VALUES(#JobImagesId#, #TenantId#, #PackageUrl#, #BusinessPackageId#, #BusinessDockerfileId#,#JobId#) 
`

	update_appVersionJobImages string = `
	update app_version_job_images  set
    $if PackageUrl != '' then 
    package_url = #PackageUrl#,
    $endif
    $if BusinessPackageId != '' then 
     business_package_id =  #BusinessPackageId#,
     $endif
    $if BusinessDockerfileId != '' then 
     business_dockerfile_id =  #BusinessDockerfileId#,
     $endif
     status_cd = '0'
	where
		status_cd = '0'
	and job_images_id = #JobImagesId#
	and tenant_id = #TenantId#
	$if JobId != '' then
	and job_id = #JobId#
	$endif
	`
	delete_appVersionJobImages string = `
	update app_version_job_images  set
                          status_cd = '1'
                          where status_cd = '0'
$if JobImagesId != '' then
		and job_images_id = #JobImagesId#
	$endif
$if JobId != '' then
	and job_id = #JobId#
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
