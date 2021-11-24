package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/entity/dto/mapping"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/system/dao"
)

type MappingService struct {
	mappingDao dao.MappingDao
}

/**
查询 系统信息
*/
func (mappingService *MappingService) GetMappings(ctx iris.Context) result.ResultDto {
	var (
		err         error
		page        int64
		row         int64
		total       int64
		mappingDto  = mapping.MappingDto{}
		mappingDtos []*mapping.MappingDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	mappingDto.Row = row * page

	mappingDto.Page = (page - 1) * row

	total, err = mappingService.mappingDao.GetMappingCount(mappingDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	mappingDtos, err = mappingService.mappingDao.GetMappings(mappingDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(mappingDtos, total, row)

}

/**
查询 系统信息
*/
func (mappingService *MappingService) GetDicts(ctx iris.Context) result.ResultDto {
	var (
		err      error
		dictDto  = mapping.DictDto{}
		dictDtos []*mapping.DictDto
	)
	dictDto.TableName = ctx.URLParam("tableName")
	dictDto.TableColumns = ctx.URLParam("tableColumns")
	dictDtos, err = mappingService.mappingDao.GetDicts(dictDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dictDtos)

}

/**
保存 系统信息
*/
func (mappingService *MappingService) SaveMappings(ctx iris.Context) result.ResultDto {
	var (
		err        error
		mappingDto mapping.MappingDto
	)

	if err = ctx.ReadJSON(&mappingDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = mappingService.mappingDao.SaveMapping(mappingDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(mappingDto)

}

/**
修改 系统信息
*/
func (mappingService *MappingService) UpdateMappings(ctx iris.Context) result.ResultDto {
	var (
		err        error
		mappingDto mapping.MappingDto
	)

	if err = ctx.ReadJSON(&mappingDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = mappingService.mappingDao.UpdateMapping(mappingDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(mappingDto)

}

/**
删除 系统信息
*/
func (mappingService *MappingService) DeleteMappings(ctx iris.Context) result.ResultDto {
	var (
		err        error
		mappingDto mapping.MappingDto
	)

	if err = ctx.ReadJSON(&mappingDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = mappingService.mappingDao.DeleteMapping(mappingDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(mappingDto)

}
