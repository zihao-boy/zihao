package workflow

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/workflowDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/entity/dto/workflow"
	"strconv"
)

type WorkflowService struct {
	workflowDao workflowDao.WorkflowDao
	hostDao         hostDao.HostDao
}

// get db link
// all db by this user
func (resourcesOssService *WorkflowService) GetWorkflowAll(WorkflowDto workflow.WorkflowDto) ([]*workflow.WorkflowDto, error) {
	var (
		err              error
		WorkflowDtos []*workflow.WorkflowDto
	)

	WorkflowDtos, err = resourcesOssService.workflowDao.GetWorkflows(WorkflowDto)
	if err != nil {
		return nil, err
	}

	return WorkflowDtos, nil

}

/**
查询 系统信息
*/
func (resourcesOssService *WorkflowService) GetWorkflows(ctx iris.Context) result.ResultDto {
	var (
		err              error
		page             int64
		row              int64
		total            int64
		resourcesOssDto  = workflow.WorkflowDto{}
		resourcesOssDtos []*workflow.WorkflowDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesOssDto.TenantId = user.TenantId

	resourcesOssDto.Row = row * page

	resourcesOssDto.Page = (page - 1) * row

	total, err = resourcesOssService.workflowDao.GetWorkflowCount(resourcesOssDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	resourcesOssDtos, err = resourcesOssService.workflowDao.GetWorkflows(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDtos, total, row)

}

/**
保存 系统信息
*/
func (resourcesOssService *WorkflowService) SaveWorkflows(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	resourcesOssDto.WorkflowId = seq.Generator()
	resourcesOssDto.JobTime = date.GetNowTimeString()
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	resourcesOssDto.TenantId = user.TenantId
	//WorkflowDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = resourcesOssService.workflowDao.SaveWorkflow(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
修改 系统信息
*/
func (resourcesOssService *WorkflowService) UpdateWorkflows(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}
	//resourcesOssDto.Name = ctx.FormValue("name")

	err = resourcesOssService.workflowDao.UpdateWorkflow(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}

/**
删除 系统信息
*/
func (resourcesOssService *WorkflowService) DeleteWorkflows(ctx iris.Context) result.ResultDto {
	var (
		err             error
		resourcesOssDto workflow.WorkflowDto
	)
	if err = ctx.ReadJSON(&resourcesOssDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = resourcesOssService.workflowDao.DeleteWorkflow(resourcesOssDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(resourcesOssDto)

}
