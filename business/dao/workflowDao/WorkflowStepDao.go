package workflowDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/workflow"
	"gorm.io/gorm"
)

const (
	query_workflowStep_count string = `
	select count(1) total
from workflow_step t
					where t.status_cd = '0'
					$if Step != '' then
					and t.step = #Step#
					$endif
					$if Name != '' then
					and t.name = #Name#
					$endif
					$if StepId != '' then
					and t.step_id = #StepId#
					$endif

	`
	query_workflowStep string = `
select t.*
from workflow_step t
					where t.status_cd = '0'
					$if Step != '' then
					and t.step = #Step#
					$endif
					$if Name != '' then
					and t.name = #Name#
					$endif
					$if StepId != '' then
					and t.step_id = #StepId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_workflowStep string = `
	insert into workflow_step(step_id, name, step)
VALUES(#StepId#,#Name#,#Step#)
`

	update_workflowStep string = `
	update workflow_step set
		$if Step != '' then
					 t.step = #Step#,
					$endif
					$if Name != '' then
					 t.name = #Name#,
					$endif
		status_cd = '0'
		where status_cd = '0'
		$if StepId != '' then
		and step_id = #StepId#
		$endif
	`
	delete_workflowStep string = `
	update workflow_step  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if StepId != '' then
						  and step_id = #StepId#
						  $endif
	`
)

type WorkflowStepDao struct {
}

/**
查询用户
*/
func (*WorkflowStepDao) GetWorkflowStepCount(workflowStepDto workflow.WorkflowStepDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_workflowStep_count, objectConvert.Struct2Map(workflowStepDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WorkflowStepDao) GetWorkflowSteps(workflowStepDto workflow.WorkflowStepDto) ([]*workflow.WorkflowStepDto, error) {
	var workflowStepDtos []*workflow.WorkflowStepDto
	sqlTemplate.SelectList(query_workflowStep, objectConvert.Struct2Map(workflowStepDto), func(db *gorm.DB) {
		db.Scan(&workflowStepDtos)
	}, false)

	return workflowStepDtos, nil
}

/**
保存服务sql
*/
func (*WorkflowStepDao) SaveWorkflowStep(workflowStepDto workflow.WorkflowStepDto) error {
	return sqlTemplate.Insert(insert_workflowStep, objectConvert.Struct2Map(workflowStepDto), false)
}

/**
修改服务sql
*/
func (*WorkflowStepDao) UpdateWorkflowStep(workflowStepDto workflow.WorkflowStepDto) error {
	return sqlTemplate.Update(update_workflowStep, objectConvert.Struct2Map(workflowStepDto), false)
}

/**
删除服务sql
*/
func (*WorkflowStepDao) DeleteWorkflowStep(workflowStepDto workflow.WorkflowStepDto) error {
	return sqlTemplate.Delete(delete_workflowStep, objectConvert.Struct2Map(workflowStepDto), false)
}
