package defalutContainerScheduling

import (
	"encoding/json"
	"github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
	"strings"
	"time"
)

//default  scheduling
// base on mem
//add by wuxw 2021-12-07
func Scheduling(hosts []*host.HostDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {

	var (
		resultDto     result.ResultDto
		err           error
		hostDto       *host.HostDto
		appServiceDao dao.AppServiceDao
	)

	//最优主机
	hostDto, err = getOptimalHost(hosts)
	resultDto, err = doStartContainer(hostDto, appServiceDto)

	if err != nil {
		return resultDto, err
	}


	if resultDto.Code != result.CODE_SUCCESS{
		return resultDto, err
	}

	data := map[string]string{
		"ContainerId": "",
	}
	dateByte, _ := json.Marshal(resultDto.Data)
	json.Unmarshal(dateByte, &data)

	appServiceContainerDto := appService.AppServiceContainerDto{
		ContainerId:       seq.Generator(),
		AsId:              appServiceDto.AsId,
		TenantId:          appServiceDto.TenantId,
		HostId:            hostDto.HostId,
		State:             "",
		Message:           "创建成功",
		UpdateTime:        time.Now().Format("2006-01-02 15:04:05"),
		DockerContainerId: data["ContainerId"],
	}

	appServiceDao.SaveAppServiceContainer(appServiceContainerDto)

	return resultDto, err
}

//目前根据 内存最低 来时实现
func getOptimalHost(hosts []*host.HostDto) (*host.HostDto, error) {
	if len(hosts) == 1 {
		return hosts[0], nil
	}

	var (
		hostDto *host.HostDto
	)

	useMem, err := strconv.ParseFloat(hosts[0].UseMem, 64)
	if err != nil {
		return nil, err
	}

	for _, host := range hosts {
		hostUseMem, err := strconv.ParseFloat(host.UseMem, 64)
		if err != nil {
			return nil, err
		}
		if useMem > hostUseMem {
			useMem = hostUseMem
			hostDto = host
		}
	}

	return hostDto, nil

}

func doStartContainer(host *host.HostDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&appServiceDto)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + string(config.Slave))

	resp, err := httpReq.Post(ip+"/app/slave/startContainer", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil
}

//default  stop
// base on mem
//add by wuxw 2021-12-07
func StopContainer(containerDto *appService.AppServiceContainerDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {

	var (
		resultDto     result.ResultDto
		err           error
		appServiceDao dao.AppServiceDao
	)

	resultDto, err = doStopContainer(containerDto, appServiceDto)

	if err != nil {
		return resultDto, err
	}

	if resultDto.Code != result.CODE_SUCCESS{
		return resultDto, err
	}

	appServiceContainerDto := appService.AppServiceContainerDto{
		ContainerId: seq.Generator(),
		AsId:        appServiceDto.AsId,
		TenantId:    appServiceDto.TenantId,
	}

	appServiceDao.DeleteAppServiceContainer(appServiceContainerDto)

	return resultDto, err
}

func doStopContainer(containerDto *appService.AppServiceContainerDto, dto2 *appService.AppServiceDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := containerDto.Ip
	var resultDto result.ResultDto

	containerDtoData, _ := json.Marshal(&containerDto)
	json.Unmarshal([]byte(containerDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + string(config.Slave))

	resp, err := httpReq.Post(ip+"/app/slave/startContainer", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil
}
