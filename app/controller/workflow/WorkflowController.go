package workflow

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/workflow"
)

type WorkflowController struct {
	workflowService workflow.WorkflowService

	workflowStepService workflow.WorkflowStepService
	workflowStepParamService workflow.WorkflowStepParamService

}

func WorkflowControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/workflow")
		aus      = WorkflowController{workflowStepService: workflow.WorkflowStepService{},
			workflowStepParamService:workflow.WorkflowStepParamService{},
			workflowService:workflow.WorkflowService{},

		}
	)

	//查询sql
	adinUser.Get("/getWorkflows", hero.Handler(aus.getWorkflows))

	//保存sql
	adinUser.Post("/saveWorkflow", hero.Handler(aus.saveWorkflow))

	//保存sql
	adinUser.Post("/updateWorkflow", hero.Handler(aus.updateWorkflow))

	//保存sql
	adinUser.Post("/deleteWorkflow", hero.Handler(aus.deleteWorkflow))


	//查询sql
	adinUser.Get("/getWorkflowSteps", hero.Handler(aus.getWorkflowSteps))

	//保存sql
	adinUser.Post("/saveWorkflowStep", hero.Handler(aus.saveWorkflowStep))

	//保存sql
	adinUser.Post("/updateWorkflowStep", hero.Handler(aus.updateWorkflowStep))

	//保存sql
	adinUser.Post("/deleteWorkflowStep", hero.Handler(aus.deleteWorkflowStep))

	//查询sql
	adinUser.Get("/getWorkflowStepParams", hero.Handler(aus.getWorkflowStepParams))

	//保存sql
	adinUser.Post("/saveWorkflowStepParam", hero.Handler(aus.saveWorkflowStepParam))

	//保存sql
	adinUser.Post("/updateWorkflowStepParam", hero.Handler(aus.updateWorkflowStepParam))

	//保存sql
	adinUser.Post("/deleteWorkflowStepParam", hero.Handler(aus.deleteWorkflowStepParam))

}



func (aus *WorkflowController) getWorkflows(ctx iris.Context) {
	relustDto := aus.workflowService.GetWorkflows(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) saveWorkflow(ctx iris.Context) {
	relustDto := aus.workflowService.SaveWorkflows(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) updateWorkflow(ctx iris.Context) {
	relustDto := aus.workflowService.UpdateWorkflows(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) deleteWorkflow(ctx iris.Context) {
	relustDto := aus.workflowService.DeleteWorkflows(ctx)
	ctx.JSON(relustDto)
}


func (aus *WorkflowController) getWorkflowSteps(ctx iris.Context) {
	relustDto := aus.workflowStepService.GetWorkflowSteps(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) saveWorkflowStep(ctx iris.Context) {
	relustDto := aus.workflowStepService.SaveWorkflowSteps(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) updateWorkflowStep(ctx iris.Context) {
	relustDto := aus.workflowStepService.UpdateWorkflowSteps(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) deleteWorkflowStep(ctx iris.Context) {
	relustDto := aus.workflowStepService.DeleteWorkflowSteps(ctx)
	ctx.JSON(relustDto)
}


func (aus *WorkflowController) getWorkflowStepParams(ctx iris.Context) {
	relustDto := aus.workflowStepParamService.GetWorkflowStepParams(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) saveWorkflowStepParam(ctx iris.Context) {
	relustDto := aus.workflowStepParamService.SaveWorkflowStepParams(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) updateWorkflowStepParam(ctx iris.Context) {
	relustDto := aus.workflowStepParamService.UpdateWorkflowStepParams(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *WorkflowController) deleteWorkflowStepParam(ctx iris.Context) {
	relustDto := aus.workflowStepParamService.DeleteWorkflowStepParams(ctx)
	ctx.JSON(relustDto)
}
