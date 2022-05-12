package workflowDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/workflow"
	"gorm.io/gorm"
)

const (
	query_workflow_count string = `
	select count(1) total
from workflow t
					where t.status_cd = '0'
					$if State != '' then
					and t.state = #State#
					$endif
	$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if Name != '' then
					and t.name = #Name#
					$endif
					$if WorkflowId != '' then
					and t.workflow_id = #WorkflowId#
					$endif

	`
	query_workflow string = `
select t.*
from workflow t
					where t.status_cd = '0'
						$if State != '' then
					and t.state = #State#
					$endif
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if Name != '' then
					and t.name = #Name#
					$endif
					$if WorkflowId != '' then
					and t.workflow_id = #WorkflowId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_workflow string = `
	insert into workflow(workflow_id, name,yaml,tenant_id, state,job_time)
VALUES(#WorkflowId#,#Name#,#Yaml#,#State#,#TenantId#,#JobTime#)
`

	update_workflow string = `
	update workflow set
		$if State != '' then
					 state = #State#,
					$endif
					$if Yaml != '' then
					 yaml = #Yaml#,
					$endif
					$if JobTime != '' then
					 job_time = #JobTime#,
					$endif
					$if Name != '' then
					 name = #Name#,
					$endif
		status_cd = '0'
		where status_cd = '0'
		$if WorkflowId != '' then
		and workflow_id = #WorkflowId#
		$endif
	`
	delete_workflow string = `
	update workflow  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if WorkflowId != '' then
						  and workflow_id = #WorkflowId#
						  $endif
	`
)

type WorkflowDao struct {
}

/**
查询用户
*/
func (*WorkflowDao) GetWorkflowCount(workflowDto workflow.WorkflowDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_workflow_count, objectConvert.Struct2Map(workflowDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WorkflowDao) GetWorkflows(workflowDto workflow.WorkflowDto) ([]*workflow.WorkflowDto, error) {
	var workflowDtos []*workflow.WorkflowDto
	sqlTemplate.SelectList(query_workflow, objectConvert.Struct2Map(workflowDto), func(db *gorm.DB) {
		db.Scan(&workflowDtos)
	}, false)

	return workflowDtos, nil
}

/**
保存服务sql
*/
func (*WorkflowDao) SaveWorkflow(workflowDto workflow.WorkflowDto) error {
	return sqlTemplate.Insert(insert_workflow, objectConvert.Struct2Map(workflowDto), false)
}

/**
修改服务sql
*/
func (*WorkflowDao) UpdateWorkflow(workflowDto workflow.WorkflowDto) error {
	return sqlTemplate.Update(update_workflow, objectConvert.Struct2Map(workflowDto), false)
}

/**
删除服务sql
*/
func (*WorkflowDao) DeleteWorkflow(workflowDto workflow.WorkflowDto) error {
	return sqlTemplate.Delete(delete_workflow, objectConvert.Struct2Map(workflowDto), false)
}
