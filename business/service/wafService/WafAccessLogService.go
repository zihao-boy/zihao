package wafService

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/ip"
	"github.com/zihao-boy/zihao/common/queue/wafAccessLogQueue"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"io"
	"os"
	"strconv"
	"time"
)

const default_access_log_count = 10000

type WafAccessLogService struct {
	wafDao wafDao.WafAccessLogDao
}

// get db link
// all db by this user
func (wafService *WafAccessLogService) GetWafAccessLogAll(WafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogDto, error) {
	var (
		err        error
		WafAccessLogDtos []*waf.WafAccessLogDto
	)

	WafAccessLogDtos, err = wafService.wafDao.GetWafAccessLogs(WafAccessLogDto)
	if err != nil {
		return nil, err
	}

	return WafAccessLogDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafAccessLogService) GetWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafAccessLogDto{}
		wafDtos []*waf.WafAccessLogDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	wafDto.Row = row * page

	wafDto.Page = (page - 1) * row

	total, err = wafService.wafDao.GetWafAccessLogCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafAccessLogs(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	ip.GetAccessLogByIp(wafDtos)

	return result.SuccessData(wafDtos, total, row)

}

/**
查询 系统信息
*/
func (wafService *WafAccessLogService) GetWafAccessLogMap(ctx iris.Context) result.ResultDto {
	var (
		err        error
		wafDto  = waf.WafAccessLogDto{}
		wafDtos []*waf.WafAccessLogMapDto
	)

	wafDto.StartTime = date.GetTimeString(time.Now().Add(-time.Minute * 30))
	wafDtos, err = wafService.wafDao.GetWafAccessLogMap(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	ip.GetAccessLogMapByIp(wafDtos)

	return result.SuccessData(wafDtos)

}

func (wafService *WafAccessLogService) GetWafAccessLogTop5(ctx iris.Context) result.ResultDto {
	var (
		err        error
		wafDto  = waf.WafAccessLogDto{}
		wafDtos []*waf.WafAccessLogMapDto
	)

	wafDto.StartTime = date.GetTimeString(time.Now().Add(-time.Minute * 30))
	wafDtos, err = wafService.wafDao.GetWafAccessLogTop5(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	ip.GetAccessLogMapByIp(wafDtos)

	return result.SuccessData(wafDtos)

}


func (wafService *WafAccessLogService) GetWafAccessLogIntercept(ctx iris.Context) result.ResultDto {
	var (
		err        error
		wafDto  = waf.WafAccessLogDto{}

		wafDtos []*waf.WafAccessLogDto
	)

	wafDto.StartTime = date.GetNowDateString()
	wafDtos, err = wafService.wafDao.GetWafAccessLogIntercept(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	ip.GetAccessLogByIp(wafDtos)

	return result.SuccessData(wafDtos)

}



/**
保存 系统信息
*/
func (wafService *WafAccessLogService) SaveWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	go wafAccessLogQueue.SendData(wafDto)

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafAccessLogService) UpdateWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafAccessLog(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafAccessLogService) DeleteWafAccessLogs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafAccessLogDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafAccessLog(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

// load ips
func (wafService *WafAccessLogService) LoadIps(ctx iris.Context) {
	var (
		err     error
	)

	file, err := os.Open(config.G_AppConfig.IpData)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	stat, _ := file.Stat()
	responseWriter := ctx.ResponseWriter()
	responseWriter.Header().Set("Content-Type", "application/octet-stream")
	responseWriter.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))
	responseWriter.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	//responseWriter.Write(content)
	io.Copy(responseWriter, file)
	responseWriter.Flush()
}
