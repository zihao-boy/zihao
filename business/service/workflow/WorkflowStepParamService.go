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

type WorkflowStepParamService struct {
	workflowDao workflowDao.WorkflowStepParamDao
	hostDao         hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesOssService *WorkflowStepParamService) GetWorkflowStepParamAll(WorkflowStepParamDto workflow.WorkflowStepParamDto) ([]*workflow.WorkflowStepParamDto, error) {
	var (
		err              error
		WorkflowStepParamDtos []*workflow.WorkflowStepParamDto
	)

	WorkflowStepParamDtos, err = resourcesOssService.workflowDao.GetWorkflowStepParams(WorkflowStepParamDto)
	if err != nil {
		return nil, err
	}

	return WorkflowStepParamDtos, nil

}

/**
查询 系统信息
*/
func (resourcesOssService *WorkflowStepParamService) GetWorkflowStepParams(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesOssDto  = workflow.WorkflowStepParamDto{}
		resourcesOssDtos []*workflow.WorkflowStepParamDto
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

	total, err = resourcesOssService.workflowDao.GetWorkflowStepParamCount(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesOssDtos, err = resourcesOssService.workflowDao.GetWorkflowStepParams(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesOssService *WorkflowStepParamService) SaveWorkflowStepParams(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepParamDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	resourcesOssDto.StepId = seq.Generator()
	//WorkflowStepParamDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesOssService.workflowDao.SaveWorkflowStepParam(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
修改 系统信息
*/
func (resourcesOssService *WorkflowStepParamService) UpdateWorkflowStepParams(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepParamDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}
	//resourcesOssDto.Name = ctx.FormValue("name")

	err = resourcesOssService.workflowDao.UpdateWorkflowStepParam(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
删除 系统信息
*/
func (resourcesOssService *WorkflowStepParamService) DeleteWorkflowStepParams(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowStepParamDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesOssService.workflowDao.DeleteWorkflowStepParam(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}
