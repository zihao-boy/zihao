package workflowDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/workflow"
	"gorm.io/gorm"
)

const (
	query_workflowStepParam_count string = `
	select count(1) total
from workflow_step_param t
					where t.status_cd = '0'
					$if ParamId != '' then
					and t.param_id = #ParamId#
					$endif
					$if ParamName != '' then
					and t.param_name = #ParamName#
					$endif
					$if StepId != '' then
					and t.step_id = #StepId#
					$endif

	`
	query_workflowStepParam string = `
select t.*
from workflow_step_param t
					where t.status_cd = '0'
					$if ParamId != '' then
					and t.param_id = #ParamId#
					$endif
					$if ParamName != '' then
					and t.param_name = #ParamName#
					$endif
					$if StepId != '' then
					and t.step_id = #StepId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_workflowStepParam string = `
	insert into workflow_step_param(param_id,step_id, param_name, param_spec,param_value,seq)
VALUES(#ParamId#,#StepId#,#ParamName#,#ParamSpec#,#ParamValue#,#Seq#)
`

	update_workflowStepParam string = `
	update workflow_step_param set
		$if ParamSpec != '' then
					 param_spec = #ParamSpec#,
					$endif
					$if ParamName != '' then
					 param_name = #ParamName#,
					$endif
		status_cd = '0'
		where status_cd = '0'
					$if ParamId != '' then
					and param_id = #ParamId#
					$endif
	`
	delete_workflowStepParam string = `
	update workflow_step_param  set
                          status_cd = '1'
                          where status_cd = '0'
					$if ParamId != '' then
					and param_id = #ParamId#
					$endif
	`
)

type WorkflowStepParamDao struct {
}

/**
查询用户
*/
func (*WorkflowStepParamDao) GetWorkflowStepParamCount(workflowStepParamDto workflow.WorkflowStepParamDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_workflowStepParam_count, objectConvert.Struct2Map(workflowStepParamDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*WorkflowStepParamDao) GetWorkflowStepParams(workflowStepParamDto workflow.WorkflowStepParamDto) ([]*workflow.WorkflowStepParamDto, error) {
	var workflowStepParamDtos []*workflow.WorkflowStepParamDto
	sqlTemplate.SelectList(query_workflowStepParam, objectConvert.Struct2Map(workflowStepParamDto), func(db *gorm.DB) {
		db.Scan(&workflowStepParamDtos)
	}, false)

	return workflowStepParamDtos, nil
}

/**
保存服务sql
*/
func (*WorkflowStepParamDao) SaveWorkflowStepParam(workflowStepParamDto workflow.WorkflowStepParamDto) error {
	return sqlTemplate.Insert(insert_workflowStepParam, objectConvert.Struct2Map(workflowStepParamDto), false)
}

/**
修改服务sql
*/
func (*WorkflowStepParamDao) UpdateWorkflowStepParam(workflowStepParamDto workflow.WorkflowStepParamDto) error {
	return sqlTemplate.Update(update_workflowStepParam, objectConvert.Struct2Map(workflowStepParamDto), false)
}

/**
删除服务sql
*/
func (*WorkflowStepParamDao) DeleteWorkflowStepParam(workflowStepParamDto workflow.WorkflowStepParamDto) error {
	return sqlTemplate.Delete(delete_workflowStepParam, objectConvert.Struct2Map(workflowStepParamDto), false)
}
