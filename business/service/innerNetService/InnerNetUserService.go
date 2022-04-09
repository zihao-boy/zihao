package innerNetService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/innerNetDao"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"strconv"
)

type InnerNetUserService struct {
	innerNetDao             innerNetDao.InnerNetUserDao
}

// get db link
// all db by this user
func (innerNetService *InnerNetUserService) GetInnerNetUserAll(InnerNetUserDto innerNet.InnerNetUserDto) ([]*innerNet.InnerNetUserDto, error) {
	var (
		err          error
		InnerNetUserDtos []*innerNet.InnerNetUserDto
	)

	InnerNetUserDtos, err = innerNetService.innerNetDao.GetInnerNetUsers(InnerNetUserDto)
	if err != nil {
		return nil, err
	}

	return InnerNetUserDtos, nil

}

/**
查询 系统信息
*/
func (innerNetService *InnerNetUserService) GetInnerNetUsers(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		innerNetDto  = innerNet.InnerNetUserDto{}
		innerNetDtos []*innerNet.InnerNetUserDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	innerNetDto.Row = row * page

	innerNetDto.Page = (page - 1) * row

	total, err = innerNetService.innerNetDao.GetInnerNetUserCount(innerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	innerNetDtos, err = innerNetService.innerNetDao.GetInnerNetUsers(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDtos, total, row)

}

/**
保存 系统信息
*/
func (innerNetService *InnerNetUserService) SaveInnerNetUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetUserDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	innerNetDto.UserId = seq.Generator()
	//InnerNetUserDto.Path = filepath.Join(curDest, fileHeader.Filename)
	innerNetDto.LoginTime = date.GetNowTime()

	err = innerNetService.innerNetDao.SaveInnerNetUser(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/**
修改 系统信息
*/
func (innerNetService *InnerNetUserService) UpdateInnerNetUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetUserDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	//innerNetDto.Id = ctx.FormValue("id")

	//innerNetDto.Name = ctx.FormValue("name")

	err = innerNetService.innerNetDao.UpdateInnerNetUser(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}


	return result.SuccessData(innerNetDto)

}

/**
删除 系统信息
*/
func (innerNetService *InnerNetUserService) DeleteInnerNetUsers(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetUserDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = innerNetService.innerNetDao.DeleteInnerNetUser(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}
