package dbLinkService

import (
	"github.com/kataras/iris/v12"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/dbLinkDao"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/db/dbFactory"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/dbLink"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"strconv"
)

type DbLinkService struct {
	dbLinkDao dbLinkDao.DbLinkDao
	hostDao   hostDao.HostDao
}

// get db link
// all db by this user
func (dbLinkService *DbLinkService) GetDbLinkAll(DbLinkDto dbLink.DbLinkDto) ([]*dbLink.DbLinkDto, error) {
	var (
		err        error
		DbLinkDtos []*dbLink.DbLinkDto
	)

	DbLinkDtos, err = dbLinkService.dbLinkDao.GetDbLinks(DbLinkDto)
	if err != nil {
		return nil, err
	}

	return DbLinkDtos, nil

}

/**
查询 系统信息
*/
func (dbLinkService *DbLinkService) GetDbLinks(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		dbLinkDto  = dbLink.DbLinkDto{}
		dbLinkDtos []*dbLink.DbLinkDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	dbLinkDto.Row = row * page

	dbLinkDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	dbLinkDto.TenantId = user.TenantId
	dbLinkDto.CreateUserId = user.UserId

	total, err = dbLinkService.dbLinkDao.GetDbLinkCount(dbLinkDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	dbLinkDtos, err = dbLinkService.dbLinkDao.GetDbLinks(dbLinkDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dbLinkDtos, total, row)

}

/**
保存 系统信息
*/
func (dbLinkService *DbLinkService) SaveDbLinks(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dbLinkDto dbLink.DbLinkDto
	)
	if err = ctx.ReadJSON(&dbLinkDto); err != nil {
		return result.Error("解析入参失败")
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	dbLinkDto.TenantId = user.TenantId
	dbLinkDto.CreateUserId = user.UserId
	dbLinkDto.Id = seq.Generator()
	//DbLinkDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = dbLinkService.dbLinkDao.SaveDbLink(dbLinkDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dbLinkDto)

}

/**
修改 系统信息
*/
func (dbLinkService *DbLinkService) UpdateDbLinks(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dbLinkDto dbLink.DbLinkDto
	)
	if err = ctx.ReadJSON(&dbLinkDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	//dbLinkDto.Id = ctx.FormValue("id")

	dbLinkDto.TenantId = user.TenantId
	dbLinkDto.CreateUserId = user.UserId
	//dbLinkDto.Name = ctx.FormValue("name")

	err = dbLinkService.dbLinkDao.UpdateDbLink(dbLinkDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dbLinkDto)

}

/**
删除 系统信息
*/
func (dbLinkService *DbLinkService) DeleteDbLinks(ctx iris.Context) result.ResultDto {
	var (
		err       error
		dbLinkDto dbLink.DbLinkDto
	)
	if err = ctx.ReadJSON(&dbLinkDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = dbLinkService.dbLinkDao.DeleteDbLink(dbLinkDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(dbLinkDto)

}

func (dbLinkService *DbLinkService) ExecSql(ctx iris.Context) interface{} {
	var (
		err       error
		dbSqlDto  dbLink.DbSqlDto
		dbLinkDto dbLink.DbLinkDto
	)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	if err = ctx.ReadJSON(&dbSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	if dbSqlDto.DbId == "" {
		return result.Error("无效数据库")
	}

	if dbSqlDto.Sql == "" {
		return result.Error("无效sql")
	}

	dbLinkDto.Row = 1
	dbLinkDto.Page = 0
	dbLinkDto.TenantId = user.TenantId
	dbLinkDto.CreateUserId = user.UserId
	dbLinkDto.Id = dbSqlDto.DbId

	dblinkDtos, err := dbLinkService.dbLinkDao.GetDbLinks(dbLinkDto)

	if err != nil || len(dblinkDtos) < 1 {
		return result.Error("无效数据库")
	}

	// execute sql
	data := dbFactory.ExecSql(*dblinkDtos[0], dbSqlDto)

	return data
}
