package hostMapper

const (
	//查询主机组 数量sql
	QueryHostGroupCount = `select count(1) count from host_group t
						where 
						1=1
	`
	//查询主机组 sql
	QueryHostGroups = `select * count from host_group t
						where 
						1=1
	`
)
