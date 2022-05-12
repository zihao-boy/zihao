package system

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/workflow"
)

type WorkflowController struct {
	workflowStepService workflow.WorkflowStepService
}

func WorkflowControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/workflow")
		aus      = WorkflowController{workflowStepService: workflow.WorkflowStepService{}}
	)
	//查询sql
	adinUser.Get("/getWorkflowSteps", hero.Handler(aus.getWorkflowSteps))

	//保存sql
	adinUser.Post("/saveWorkflowStep", hero.Handler(aus.saveWorkflowStep))

	//保存sql
	adinUser.Post("/updateWorkflowStep", hero.Handler(aus.updateWorkflowStep))

	//保存sql
	adinUser.Post("/deleteWorkflowStep", hero.Handler(aus.deleteWorkflowStep))

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

