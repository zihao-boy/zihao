package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"gorm.io/gorm"
)

const (
	query_appService_count string = `
		select count(1) total
			from app_service t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if AsType != '' then
			and t.as_type = #AsType#
			$endif
			$if State != '' then
			and t.state = #State#
			$endif
			$if AsName != '' then
			and t.as_name like '%'|| #AsName# || '%'
			$endif
			$if AsGroupId != '' then
			and t.as_group_id = #AsGroupId#
			$endif
	`

	query_appService string = `
				select t.*,td.name as_type_name,td1.name state_name,bi.name images_name,bi.version images_version,bi.type_url images_url,
asv.avg_name,hg.name host_group_name,h.name host_name
from app_service t
left join t_dict td on t.as_type = td.status_cd and td.table_name = 'app_service' and td.table_columns = 'as_type'
left join t_dict td1 on t.state = td1.status_cd and td1.table_name = 'app_service' and td1.table_columns = 'state'
left join business_images bi on bi.id = t.images_id and bi.status_cd = '0'
left join app_var_group asv on asv.avg_id = t.as_group_id and asv.status_cd ='0'
left join host_group hg on hg.group_id = t.as_deploy_id and hg.status_cd ='0'
left join host h on h.host_id = t.as_deploy_id and h.status_cd ='0'
where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if AsType != '' then
					and t.as_type = #AsType#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if AsName != '' then
					and t.as_name like '%'|| #AsName# || '%'
					$endif
					$if AsGroupId != '' then
					and t.as_group_id = #AsGroupId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appService string = `
	insert into app_service(as_id, as_name, as_type, tenant_id, as_desc,state,as_count,as_group_id,as_deploy_type,as_deploy_id,images_id)
VALUES(#AsId#,#AsName#,#AsType#,#TenantId#,#AsDesc#,#State#,#AsCount#,#AsGroupId#,#AsDeployType#,#AsDeployId#,#ImagesId#)
`

	update_appService string = `
	update app_service set
		$if AsType != '' then
		 as_type = #AsType#,
		$endif
		$if AsName != '' then
		 as_name = #AsName#,
		$endif
		$if State != '' then
		 state = #State#,
		$endif
		$if AsCount != '' then
		 as_count = #AsCount#,
		$endif
		$if AsDesc != '' then
		 as_desc = #AsDesc#,
		$endif
		$if AsGroupId != '' then
		 as_group_id = #AsGroupId#,
		$endif
		$if AsDeployType != '' then
		 as_deploy_type = #AsDeployType#,
		$endif
		$if AsDeployId != '' then
		 as_deploy_id = #AsDeployId#,
		$endif
		$if ImagesId != '' then
		 images_id = #ImagesId#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if AsId != '' then
		and as_id = #AsId#
		$endif
	`
	delete_appService string = `
	update app_service  set
                          status_cd = '1'
                          where status_cd = '0'

		$if AsId != '' then
		and as_id = #AsId#
		$endif
	`



	query_appServiceVar_count string = `
		select count(1) total
			from app_service_var t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if AvId != '' then
			and t.av_id = #AvId#
			$endif
			$if VarName != '' then
			and t.var_name = #VarName#
			$endif
			$if VarSpec != '' then
			and t.var_spec = #VarSpec#
			$endif
    	
	`
	query_appServiceVar string = `
				select t.*
				from app_service_var t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if AvId != '' then
					and t.av_id = #AvId#
					$endif
					$if VarName != '' then
					and t.var_name = #VarName#
					$endif
					$if VarSpec != '' then
					and t.var_spec = #VarSpec#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appServiceVar string = `
	insert into app_service_var(av_id, as_id, tenant_id, var_spec, var_name,var_value)
VALUES(#AvId#,#AsId#,#TenantId#,#VarSpec#,#VarName#,#VarValue#)
`

	update_appServiceVar string = `
	update app_service_var set
		$if VarSpec != '' then
		 var_spec = #VarSpec#,
		$endif
		$if VarName != '' then
		 var_name = #VarName#,
		$endif
		$if VarValue != '' then
		 var_value = #VarValue#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if AvId != '' then
		and av_id = #AvId#
		$endif
	`
	delete_appServiceVar string = `
	update app_service_var  set
                          status_cd = '1'
                          where status_cd = '0'

		$if AvId != '' then
		and av_id = #AvId#
		$endif
	`

	query_appServiceHosts_count string = `
		select count(1) total
			from app_service_hosts t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if Ip != '' then
			and t.ip = #Ip#
			$endif
			$if Hostname != '' then
			and t.hostname = #Hostname#
			$endif
			$if HostsId != '' then
			and t.hosts_id = #HostsId#
			$endif
    	
	`
	query_appServiceHosts string = `
				select t.*
				from app_service_hosts t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if Ip != '' then
					and t.ip = #Ip#
					$endif
					$if Hostname != '' then
					and t.hostname = #Hostname#
					$endif
					$if HostsId != '' then
					and t.hosts_id = #HostsId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appServiceHosts string = `
	insert into app_service_hosts(hosts_id, as_id, tenant_id, hostname, ip)
	VALUES(#HostsId#,#AsId#,#TenantId#,#Hostname#,#Ip#)
`

	update_appServiceHosts string = `
	update app_service_hosts set
		$if Ip != '' then
		 ip = #Ip#,
		$endif
		$if Hostname != '' then
		 hostname = #Hostname#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if HostsId != '' then
		and hosts_id = #HostsId#
		$endif
	`
	delete_appServiceHosts string = `
	update app_service_hosts  set
                          status_cd = '1'
                          where status_cd = '0'
		$if HostsId != '' then
			and hosts_id = #HostsId#
		$endif
	`


	//查询应用目录映射
	query_appServiceDir_count string = `
		select count(1) total
			from app_service_dir t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if DirId != '' then
			and t.dir_id = #DirId#
			$endif
    	
	`
	query_appServiceDir string = `
				select t.*
					from app_service_dir t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if DirId != '' then
					and t.dir_id = #DirId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`
	insert_appServiceDir string = `
	insert into app_service_dir(dir_id, as_id, tenant_id, src_dir, target_dir)
	VALUES(#DirId#,#AsId#,#TenantId#,#SrcDir#,#TargetDir#)
`

	update_appServiceDir string = `
	update app_service_dir set
		$if SrcDir != '' then
		 src_dir = #SrcDir#,
		$endif
		$if TargetDir != '' then
		 target_dir = #TargetDir#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if DirId != '' then
		and dir_id = #DirId#
		$endif
	`
	delete_appServiceDir string = `
	update app_service_dir  set
                          status_cd = '1'
                          where status_cd = '0'
		$if DirId != '' then
		and dir_id = #DirId#
		$endif
	`


	//查询应用目录映射
	query_appServicePort_count string = `
		select count(1) total
			from app_service_port t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if PortId != '' then
			and t.port_id = #PortId#
			$endif

	`
	query_appServicePort string = `
				select t.*
					from app_service_port t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if PortId != '' then
					and t.port_id = #PortId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`
	insert_appServicePort string = `
	insert into app_service_port(port_id, as_id, tenant_id, src_port, target_port)
	VALUES(#PortId#,#AsId#,#TenantId#,#SrcPort#,#TargetPort#)
`

	update_appServicePort string = `
	update app_service_port set
		$if SrcPort != '' then
		 src_port = #SrcPort#,
		$endif
		$if TargetPort != '' then
		 target_port = #TargetPort#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if PortId != '' then
		and port_id = #PortId#
		$endif
	`
	delete_appServicePort string = `
	update app_service_port  set
            status_cd = '1'
	    where status_cd = '0'
		$if PortId != '' then
		and port_id = #PortId#
		$endif
	`

	//查询应用目录映射
	query_appServiceContainer_count string = `
		select count(1) total
			from app_service_container t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if ContainerId != '' then
			and t.container_id = #ContainerId#
			$endif

	`

	query_appServiceContainer string = `
				select t.*,h.name hostname,h.ip
					from app_service_container t
          left join host h on t.host_id = h.host_id and h.status_cd = '0'
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if ContainerId != '' then
					and t.container_id = #ContainerId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`
	insert_appServiceContainer string = `
	insert into app_service_container(container_id, as_id, tenant_id, host_id, docker_container_id, state,message,update_time)
	VALUES(#ContainerId#,#AsId#,#TenantId#,#HostId#,#DockerContainerId#,#State#,#Message#,#UpdateTime#)
`

	update_appServiceContainer string = `
	update app_service_container set
		
		$if State != '' then
		 state = #State#,
		$endif
		$if UpdateTime != '' then
		 update_time = #UpdateTime#,
		$endif
		$if Message != '' then
		 message = #Message#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if ContainerId != '' then
		and container_id = #ContainerId#
		$endif
		$if DockerContainerId != '' then
		and docker_container_id = #DockerContainerId#,
		$endif
		$if HostId != '' then
		 and host_id = #HostId#,
		$endif
	`
	delete_appServiceContainer string = `
	update app_service_container  set
            status_cd = '1'
	    where status_cd = '0'
		$if ContainerId != '' then
		and container_id = #ContainerId#
		$endif
	`
)

type AppServiceDao struct {

}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceCount(appServiceDto appService.AppServiceDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appService_count, objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServices(appServiceDto appService.AppServiceDto) ([]*appService.AppServiceDto, error) {
	var appServiceDtos []*appService.AppServiceDto
	sqlTemplate.SelectList(query_appService, objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		db.Scan(&appServiceDtos)
	}, false)

	return appServiceDtos, nil
}

/**
保存服务sql
*/
func (*AppServiceDao) SaveAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Insert(insert_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
修改服务sql
*/
func (*AppServiceDao) UpdateAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Update(update_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
删除服务sql
*/
func (*AppServiceDao) DeleteAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Delete(delete_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceVarCount(appServiceVarDto appService.AppServiceVarDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServiceVar_count, objectConvert.Struct2Map(appServiceVarDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceVars(appServiceVarDto appService.AppServiceVarDto) ([]*appService.AppServiceVarDto, error) {
	var appServiceVarDtos []*appService.AppServiceVarDto
	sqlTemplate.SelectList(query_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), func(db *gorm.DB) {
		db.Scan(&appServiceVarDtos)
	}, false)

	return appServiceVarDtos, nil
}

/**
保存服务sql
*/
func (*AppServiceDao) SaveAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Insert(insert_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}

/**
修改服务sql
*/
func (*AppServiceDao) UpdateAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Update(update_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}

/**
删除服务sql
*/
func (*AppServiceDao) DeleteAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Delete(delete_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}

func (d *AppServiceDao) GetAppServiceHostsCount(appServiceHostsDto appService.AppServiceHostsDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServiceHosts_count, objectConvert.Struct2Map(appServiceHostsDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

func (d *AppServiceDao) GetAppServiceHosts(hostsDto appService.AppServiceHostsDto) ([]*appService.AppServiceHostsDto, error) {

	var appServiceHostsDtos []*appService.AppServiceHostsDto
	sqlTemplate.SelectList(query_appServiceHosts, objectConvert.Struct2Map(hostsDto), func(db *gorm.DB) {
		db.Scan(&appServiceHostsDtos)
	}, false)

	return appServiceHostsDtos, nil
}

func (d *AppServiceDao) SaveAppServiceHosts(hostsDto appService.AppServiceHostsDto) error {
	return sqlTemplate.Insert(insert_appServiceHosts, objectConvert.Struct2Map(hostsDto), false)
}

func (d *AppServiceDao) UpdateAppServiceHost(hostsDto appService.AppServiceHostsDto) error {
	return sqlTemplate.Update(update_appServiceHosts, objectConvert.Struct2Map(hostsDto), false)

}

func (d *AppServiceDao) DeleteAppServiceHosts(hostsDto appService.AppServiceHostsDto) error {
	return sqlTemplate.Delete(delete_appServiceHosts, objectConvert.Struct2Map(hostsDto), false)

}

func (d *AppServiceDao) GetAppServiceDirCount(dirDto appService.AppServiceDirDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServiceDir_count, objectConvert.Struct2Map(dirDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

//查询服务磁盘路径
func (d *AppServiceDao) GetAppServiceDir(dirDto appService.AppServiceDirDto) ([]*appService.AppServiceDirDto, error) {
	var appServiceDirDtos []*appService.AppServiceDirDto
	sqlTemplate.SelectList(query_appServiceDir, objectConvert.Struct2Map(dirDto), func(db *gorm.DB) {
		db.Scan(&appServiceDirDtos)
	}, false)

	return appServiceDirDtos, nil
}

func (d *AppServiceDao) SaveAppServiceDir(dirDto appService.AppServiceDirDto) error {
	return sqlTemplate.Insert(insert_appServiceDir, objectConvert.Struct2Map(dirDto), false)
}

//修改服务映射目录
func (d *AppServiceDao) UpdateAppServiceDir(dirDto appService.AppServiceDirDto) error {
	return sqlTemplate.Update(update_appServiceDir, objectConvert.Struct2Map(dirDto), false)
}

//删除应用映射目录
func (d *AppServiceDao) DeleteAppServiceDir(dirDto appService.AppServiceDirDto) error {
	return sqlTemplate.Delete(delete_appServiceDir, objectConvert.Struct2Map(dirDto), false)
}

//查询端口
func (d *AppServiceDao) GetAppServicePortCount(portDto appService.AppServicePortDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServicePort_count, objectConvert.Struct2Map(portDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

//query app service port mapping
func (d *AppServiceDao) GetAppServicePort(portDto appService.AppServicePortDto) ([]*appService.AppServicePortDto, error) {
	var appServicePortDtos []*appService.AppServicePortDto
	sqlTemplate.SelectList(query_appServicePort, objectConvert.Struct2Map(portDto), func(db *gorm.DB) {
		db.Scan(&appServicePortDtos)
	}, false)

	return appServicePortDtos, nil
}

func (d *AppServiceDao) SaveAppServicePort(portDto appService.AppServicePortDto) error {
	return sqlTemplate.Insert(insert_appServicePort, objectConvert.Struct2Map(portDto), false)
}

func (d *AppServiceDao) UpdateAppServicePort(portDto appService.AppServicePortDto) error {
	return sqlTemplate.Update(update_appServicePort, objectConvert.Struct2Map(portDto), false)
}

func (d *AppServiceDao) DeleteAppServicePort(portDto appService.AppServicePortDto) error {
	return sqlTemplate.Delete(delete_appServicePort, objectConvert.Struct2Map(portDto), false)
}

func (d *AppServiceDao) GetAppServiceContainerCount(containerDto appService.AppServiceContainerDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServiceContainer_count, objectConvert.Struct2Map(containerDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

func (d *AppServiceDao) GetAppServiceContainer(containerDto appService.AppServiceContainerDto) ([]*appService.AppServiceContainerDto, error) {
	var appServiceContainerDtos []*appService.AppServiceContainerDto
	sqlTemplate.SelectList(query_appServiceContainer, objectConvert.Struct2Map(containerDto), func(db *gorm.DB) {
		db.Scan(&appServiceContainerDtos)
	}, false)

	return appServiceContainerDtos, nil
}

func (d *AppServiceDao) SaveAppServiceContainer(containerDto appService.AppServiceContainerDto) error {
	return sqlTemplate.Insert(insert_appServiceContainer, objectConvert.Struct2Map(containerDto), false)
}

func (d *AppServiceDao) UpdateAppServiceContainer(containerDto appService.AppServiceContainerDto) error {
	return sqlTemplate.Update(update_appServiceContainer, objectConvert.Struct2Map(containerDto), false)
}

func (d *AppServiceDao) DeleteAppServiceContainer(containerDto appService.AppServiceContainerDto) error {
	return sqlTemplate.Delete(delete_appServiceContainer, objectConvert.Struct2Map(containerDto), false)
}