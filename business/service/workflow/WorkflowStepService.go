package workflow

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/workflowDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/workflow"
	"strconv"
)

const maxSize = 1000 * iris.MB // 第二种方法
type WorkflowStepService struct {
	workflowDao workflowDao.WorkflowStepDao
	hostDao         hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesOssService *WorkflowStepService) GetWorkflowStepAll(WorkflowStepDto workflow.WorkflowStepDto) ([]*workflow.WorkflowStepDto, error) {
	var (
		err              error
		WorkflowStepDtos []*workflow.WorkflowStepDto
	)

	WorkflowStepDtos, err = resourcesOssService.workflowDao.GetWorkflowSteps(WorkflowStepDto)
	if err != nil {
		return nil, err
	}

	return WorkflowStepDtos, nil

}

/**
查询 系统信息
*/
func (resourcesOssService *WorkflowStepService) GetWorkflowSteps(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesOssDto  = workflow.WorkflowStepDto{}
		resourcesOssDtos []*workflow.WorkflowStepDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	resourcesOssDto.Row = row * page

	resourcesOssDto.Page = (page - 1) * row

	total, err = resourcesOssService.workflowDao.GetWorkflowStepCount(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesOssDtos, err = resourcesOssService.workflowDao.GetWorkflowSteps(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesOssService *WorkflowStepService) SaveWorkflowSteps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	resourcesOssDto.StepId = seq.Generator()
	//WorkflowStepDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesOssService.workflowDao.SaveWorkflowStep(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
修改 系统信息
*/
func (resourcesOssService *WorkflowStepService) UpdateWorkflowSteps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}
	//resourcesOssDto.Name = ctx.FormValue("name")

	err = resourcesOssService.workflowDao.UpdateWorkflowStep(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
删除 系统信息
*/
func (resourcesOssService *WorkflowStepService) DeleteWorkflowSteps(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesOssService.workflowDao.DeleteWorkflowStep(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}
