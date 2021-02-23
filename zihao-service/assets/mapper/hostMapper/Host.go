package hostMapper

const (
	//查询主机组 数量sql
	QueryHostGroupCount = `select count(1) total 
						from host_group t
						where 
						1=1
						and t.status_cd = '0'
						$if Name != '' then
						and t.name = #Name#
						$endif
						$if TenantId != '' then
						and t.tenant_id = #TenantId#
						$endif
	`
	//查询主机组 sql
	QueryHostGroups = `select * from host_group t
						where 
						1=1
						and t.status_cd = '0'
						$if Name != '' then
						and t.name = #Name#
						$endif
						$if TenantId != '' then
						and t.tenant_id = #TenantId#
						$endif
						$if Page != -1 then
						limit #Page#,#Row#
						$endif
						
	`
)
