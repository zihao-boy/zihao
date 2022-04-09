package innerNetService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/innerNetDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"strconv"
	"strings"
)

type InnerNetService struct {
	innerNetDao      innerNetDao.InnerNetDao
	innerNetHostsDao innerNetDao.InnerNetHostsDao
}

// get db link
// all db by this user
func (innerNetService *InnerNetService) GetInnerNetAll(InnerNetDto innerNet.InnerNetDto) ([]*innerNet.InnerNetDto, error) {
	var (
		err     error
		InnerNetDtos []*innerNet.InnerNetDto
	)

	InnerNetDtos, err = innerNetService.innerNetDao.GetInnerNets(InnerNetDto)
	if err != nil {
		return nil, err
	}

	return InnerNetDtos, nil

}

/**
查询 系统信息
*/
func (innerNetService *InnerNetService) GetInnerNets(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		innerNetDto  = innerNet.InnerNetDto{}
		innerNetDtos []*innerNet.InnerNetDto
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

	total, err = innerNetService.innerNetDao.GetInnerNetCount(innerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	innerNetDtos, err = innerNetService.innerNetDao.GetInnerNets(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, innerNetDto := range innerNetDtos {

		innerNetHostsDto := innerNet.InnerNetHostsDto{
			InnerNetId: innerNetDto.InnerNetId,
		}
		innerNetHostsDtos, _ := innerNetService.innerNetHostsDao.GetInnerNetHostss(innerNetHostsDto)
		innerNetDto.InnerNetHosts = innerNetHostsDtos
	}

	return result.SuccessData(innerNetDtos, total, row)

}

/**
保存 系统信息
*/
func (innerNetService *InnerNetService) SaveInnerNets(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	innerNetDto.InnerNetId = seq.Generator()
	//InnerNetDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = innerNetService.innerNetDao.SaveInnerNet(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/**
修改 系统信息
*/
func (innerNetService *InnerNetService) UpdateInnerNets(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	//innerNetDto.Id = ctx.FormValue("id")

	//innerNetDto.Name = ctx.FormValue("name")

	err = innerNetService.innerNetDao.UpdateInnerNet(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if utils.IsEmpty(innerNetDto.HostIds) {
		return result.SuccessData(innerNetDto)
	}

	innerNetHostsDto := innerNet.InnerNetHostsDto{
		InnerNetId: innerNetDto.InnerNetId,
	}
	innerNetService.innerNetHostsDao.DeleteInnerNetHosts(innerNetHostsDto)

	for _, hostId := range strings.Split(innerNetDto.HostIds, ",") {

		innerNetHostsDto = innerNet.InnerNetHostsDto{
			InnerNetId:     innerNetDto.InnerNetId,
			InnerNetHostId: seq.Generator(),
			HostId:    hostId,
		}
		innerNetService.innerNetHostsDao.SaveInnerNetHosts(innerNetHostsDto)
	}

	return result.SuccessData(innerNetDto)

}

/**
删除 系统信息
*/
func (innerNetService *InnerNetService) DeleteInnerNets(ctx iris.Context) result.ResultDto {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = innerNetService.innerNetDao.DeleteInnerNet(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

func (innerNetService *InnerNetService) StartInnerNetf(ctx iris.Context) interface{} {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start innerNet
	tmpInnerNetDto := innerNet.InnerNetDto{
		InnerNetId: innerNetDto.InnerNetId,
	}
	innerNetDtos, err := innerNetService.innerNetDao.GetInnerNets(tmpInnerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(innerNetDtos) < 1 {
		return result.Error("未查询到数据")
	}

	resultDto, _ := shell.ExecStartInnerNet(innerNetService.getInnerNetConfig(*innerNetDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpInnerNetDto = innerNet.InnerNetDto{
		InnerNetId: innerNetDto.InnerNetId,
		State: innerNet.InnerNet_state_start,
	}

	err = innerNetService.innerNetDao.UpdateInnerNet(tmpInnerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(innerNetDto)

}

func (innerNetService *InnerNetService) StopInnerNetf(ctx iris.Context) interface{} {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	// start innerNet
	tmpInnerNetDto := innerNet.InnerNetDto{
		InnerNetId: innerNetDto.InnerNetId,
	}
	innerNetDtos, err := innerNetService.innerNetDao.GetInnerNets(tmpInnerNetDto)

	if err != nil || len(innerNetDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecStopInnerNet(innerNetService.getInnerNetConfig(*innerNetDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}

	tmpInnerNetDto = innerNet.InnerNetDto{
		InnerNetId: innerNetDto.InnerNetId,
		State: innerNet.InnerNet_state_stop,
	}

	err = innerNetService.innerNetDao.UpdateInnerNet(tmpInnerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(innerNetDto)

}

func (innerNetService *InnerNetService) RefreshInnerNetConfig(ctx iris.Context) interface{} {
	var (
		err    error
		innerNetDto innerNet.InnerNetDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	// start innerNet
	tmpInnerNetDto := innerNet.InnerNetDto{
		InnerNetId: innerNetDto.InnerNetId,
	}
	innerNetDtos, err := innerNetService.innerNetDao.GetInnerNets(tmpInnerNetDto)

	if err != nil || len(innerNetDtos) < 1 {
		return result.Error(err.Error())
	}

	resultDto, _ := shell.ExecRefreshInnerNetConfig(innerNetService.getInnerNetConfig(*innerNetDtos[0]))

	if resultDto.Code != result.CODE_SUCCESS {
		return resultDto
	}
	return result.SuccessData(innerNetDto)
}

// get innerNet config
func (innerNetService *InnerNetService) getInnerNetConfig(innerNetDto innerNet.InnerNetDto) innerNet.SlaveInnerNetDataDto {
	var (
		innerNetUserDao        innerNetDao.InnerNetUserDao
	)

	innerNetHostsDto := innerNet.InnerNetHostsDto{
		InnerNetId: innerNetDto.InnerNetId,
	}
	innerNetHostsDtos, _ := innerNetService.innerNetHostsDao.GetInnerNetHostss(innerNetHostsDto)
	innerNetDto.InnerNetHosts = innerNetHostsDtos

	// start innerNet
	tmpInnerNetUserDto := innerNet.InnerNetUserDto{
		
	}
	users, _ := innerNetUserDao.GetInnerNetUsers(tmpInnerNetUserDto)


	return innerNet.SlaveInnerNetDataDto{
		ServerIpUrl: config.G_AppConfig.ServerIpUrl,
		InnerNet:         innerNetDto,
		Users:      users,
	}

}
