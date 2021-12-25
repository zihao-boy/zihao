package service

import (
	"encoding/json"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	appServiceDao "github.com/zihao-boy/zihao/appService/dao"
	"github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/entity/dto/container"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"golang.org/x/crypto/ssh"
)

const maxSize = 1000 * iris.MB // 第二种方法

type HostService struct {
	hostDao dao.HostDao
	appServiceDao appServiceDao.AppServiceDao
}

const (
	host_token    string = "host_token"
	get_container string = "#!/bin/bash\n\ncontainer_id=`docker ps -a|awk '{if (NR>1){print $1}}'`\n\ncontainer_name=`docker ps -a|awk '{if (NR>1){print $(NF)}}'`\n\nimage=`docker ps -a|awk '{if (NR>1){print $2}}'`\n\nport=`docker ps -a|awk '{if (NR>1){print $(NF-1)}}'`\n\necho \"$container_id&&$container_name&&$image&&$port\""
)

/**
查询 系统信息
*/
func (hostService *HostService) GetHostGroups(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err           error
		page          int64
		row           int64
		total         int64
		hostGroupDto  = host.HostGroupDto{TenantId: user.TenantId}
		hostGroupDtos []*host.HostGroupDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	hostGroupDto.Row = row * page

	hostGroupDto.Page = (page - 1) * row

	total, err = hostService.hostDao.GetHostGroupCount(hostGroupDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}
	hostGroupDtos, err = hostService.hostDao.GetHostGroups(hostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDtos, total, row)

}

/**
保存 系统信息
*/
func (hostService *HostService) SaveHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err          error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	hostGroupDto.TenantId = user.TenantId
	hostGroupDto.GroupId = seq.Generator()

	err = hostService.hostDao.SaveHostGroup(hostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}

/**
修改 系统信息
*/
func (hostService *HostService) UpdateHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err          error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.UpdateHostGroup(hostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}

/**
删除 系统信息
*/
func (hostService *HostService) DeleteHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err          error
		hostGroupDto host.HostGroupDto
	)

	if err = ctx.ReadJSON(&hostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.DeleteHostGroup(hostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostGroupDto)

}

/**
查询 系统信息
*/
func (hostService *HostService) GetHosts(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err      error
		page     int64
		row      int64
		total    int64
		hostDto  = host.HostDto{TenantId: user.TenantId}
		hostDtos []*host.HostDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	hostDto.Row = row * page

	hostDto.Page = (page - 1) * row

	groupId := ctx.URLParam("groupId")

	hostDto.GroupId = groupId

	hostDto.HostId = ctx.URLParam("hostId")

	total, err = hostService.hostDao.GetHostCount(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}
	hostDtos, err = hostService.hostDao.GetHosts(hostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, item := range hostDtos {
		item.Passwd = ""
	}

	return result.SuccessData(hostDtos, total, row)

}

/**
保存 系统信息
*/
func (hostService *HostService) SaveHost(ctx iris.Context) result.ResultDto {
	var (
		err     error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	hostDto.TenantId = user.TenantId
	hostDto.HostId = seq.Generator()
	hostDto.State = "1001"
	hostDto.HeartbeatTime = time.Now().Format("2006-01-02 15:04:05")

	err = hostService.hostDao.SaveHost(hostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}

/**
修改 系统信息
*/
func (hostService *HostService) UpdateHost(ctx iris.Context) result.ResultDto {
	var (
		err     error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.UpdateHost(hostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}

/**
删除 系统信息
*/
func (hostService *HostService) DeleteHost(ctx iris.Context) result.ResultDto {
	var (
		err     error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = hostService.hostDao.DeleteHost(hostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(hostDto)

}

/**
主机生成token
*/
func (hostService *HostService) GetHostToken(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	var hostToken string = seq.Generator()
	factory.SetValue(host_token, hostToken)

	return result.SuccessData(hostToken)

}

/**
主机生成token
*/
func (hostService *HostService) GetContainers(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	client, err := ssh.Dial("tcp", hostDtos[0].Ip, &ssh.ClientConfig{
		User:            hostDtos[0].Username,
		Auth:            []ssh.AuthMethod{ssh.Password(hostDtos[0].Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		return result.Error("连接主机失败")
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存
	processName, _ := session.Output(get_container)
	var (
		processDtos []container.ContainerDto = make([]container.ContainerDto, 0)
		outData                              = strings.ReplaceAll(string(processName), "'", "\"")
	)

	containers := strings.Split(outData, "&&")

	containerIds := strings.Split(containers[0], "\n")
	containerNames := strings.Split(containers[1], "\n")
	containerImages := strings.Split(containers[2], "\n")
	containerPorts := strings.Split(containers[3], "\n")

	if len(containerIds) == 1 && containerIds[0] == "" {
		return result.Success()
	}

	for index, item := range containerIds {
		tmpProcessDto := container.ContainerDto{
			ContainerId:   item,
			ContainerName: containerNames[index],
			Image:         containerImages[index],
			Port:          containerPorts[index],
		}
		processDtos = append(processDtos, tmpProcessDto)
	}

	return result.SuccessData(processDtos)

}

/**
主机生成token
*/
func (hostService *HostService) GetHostPort(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	client, err := ssh.Dial("tcp", hostDtos[0].Ip, &ssh.ClientConfig{
		User:            hostDtos[0].Username,
		Auth:            []ssh.AuthMethod{ssh.Password(hostDtos[0].Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		return result.Error("连接主机失败")
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存
	processName, _ := session.Output(constants.Get_host_port_shell)
	var (
		portDtos []host.HostPortDto = make([]host.HostPortDto, 0)
		outData                     = strings.ReplaceAll(string(processName), "'", "\"")
	)

	ports := strings.Split(outData, "&&")

	protocols := strings.Split(ports[0], "\n")
	tPorts := strings.Split(ports[1], "\n")
	programNames := strings.Split(ports[2], "\n")

	if len(protocols) == 1 && protocols[0] == "" {
		return result.Success()
	}

	for index, item := range protocols {
		tmpPortDto := host.HostPortDto{
			Protocol:    item,
			Port:        tPorts[index],
			ProgramName: programNames[index],
		}
		portDtos = append(portDtos, tmpPortDto)
	}

	return result.SuccessData(portDtos)

}

/**
查询主机资源
*/
func (hostService *HostService) GetHostResource(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	client, err := ssh.Dial("tcp", hostDtos[0].Ip, &ssh.ClientConfig{
		User:            hostDtos[0].Username,
		Auth:            []ssh.AuthMethod{ssh.Password(hostDtos[0].Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		return result.Error("连接主机失败")
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存

	processName, _ := session.Output(strings.ReplaceAll(constants.Check_host_resource_shell, "$1", "/"))

	var (
		monitorCheckHostInfoDto *monitor.MonitorCheckHostInfoDto
		outData                 = strings.ReplaceAll(string(processName), "'", "\"")
	)

	json.Unmarshal([]byte(outData), &monitorCheckHostInfoDto)

	return result.SuccessData(monitorCheckHostInfoDto)

}

/**
查询主机资源
*/
func (hostService *HostService) GetHostDisk(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	client, err := ssh.Dial("tcp", hostDtos[0].Ip, &ssh.ClientConfig{
		User:            hostDtos[0].Username,
		Auth:            []ssh.AuthMethod{ssh.Password(hostDtos[0].Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		return result.Error("连接主机失败")
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存

	processName, _ := session.Output(constants.Get_host_disk_shell)

	var (
		diskDtos []host.HostDiskDto = make([]host.HostDiskDto, 0)
		outData                     = strings.ReplaceAll(string(processName), "'", "\"")
	)

	disks := strings.Split(outData, "&&")

	diskNames := strings.Split(disks[0], "\n")
	sizes := strings.Split(disks[1], "\n")
	freeSizes := strings.Split(disks[2], "\n")
	dirs := strings.Split(disks[3], "\n")

	if len(diskNames) == 1 && diskNames[0] == "" {
		return result.Success()
	}

	for index, item := range diskNames {
		tmpPortDto := host.HostDiskDto{
			DiskName: item,
			Size:     sizes[index],
			FreeSize: freeSizes[index],
			Dir:      dirs[index],
		}
		diskDtos = append(diskDtos, tmpPortDto)
	}

	return result.SuccessData(diskDtos)

}

/**
  控制主机
*/
func (hostService *HostService) ControlHost(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{}
	)
	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败")
	}
	hostDto.TenantId = user.TenantId
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	mapping, err := factory.GetMapping("MASTER_IP")
	var masterIp string
	if err != nil {
		masterIp = "127.0.0.1:7000"
	} else {
		masterIp = mapping.Value
	}

	shellScript := strings.ReplaceAll(constants.DownLoad_slave_shell, "$zihao_ip", masterIp)

	shellScript = strings.ReplaceAll(shellScript, "$zihao_hostId", hostDtos[0].HostId)

	go shell.ExecShell(*hostDtos[0], shellScript)

	hostDto.State = host.State_D
	hostService.hostDao.UpdateHost(hostDto)

	return result.Success()

}

/**
  控制主机
*/
func (hostService *HostService) SlaveHealth(ctx iris.Context) result.ResultDto {

	var (
		hostDto = host.HostDto{}
	)
	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	//
	//      "hostId":  slaveId,
	//		"cpu":     cpuCore,
	//		"mem":     totalMem.Total,
	//		"disk":    totalDisk.Total,
	//		"useCpu":  cpuPercentDec.Float64(),
	//		"useMem":  totalMem.Used,
	//		"useDisk": totalDisk.Used,
	hostDto.State = host.State_N

	hostDto.HeartbeatTime = time.Now().Format("2006-01-02 15:04:05")
	hostService.hostDao.UpdateHost(hostDto)

	// container := appService.AppServiceContainerDto{
	// 	HostId: hostDtos[0].HostId,
	// }
	//containers , _ :=hostService.appServiceDao.GetAppServiceContainer(container)

	if len(hostDto.Containers) < 1 {
		return result.Success();
	}

	for _,tempContainer := range hostDto.Containers{
		container := appService.AppServiceContainerDto{
			HostId: hostDtos[0].HostId,
			DockerContainerId:tempContainer.Id,
			State: tempContainer.Status,
			UpdateTime: hostDto.HeartbeatTime,
		}
		hostService.appServiceDao.UpdateAppServiceContainer(container)
	}


	return result.Success()

}

func (hostService *HostService) ListFiles(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
			CurPath:  ctx.URLParam("curPath"),
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecListFiles(hostDto)
	return resultDto
}

func (hostService *HostService) RemoveFile(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto host.HostDto
	)

	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}

	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecRemoveFile(hostDto)
	return resultDto
}

func (hostService *HostService) NewFile(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto host.HostDto
	)

	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}

	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecNewFile(hostDto)
	return resultDto
}

func (hostService *HostService) RenameFile(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto host.HostDto
	)

	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}

	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecRenameFile(hostDto)
	return resultDto
}

func (hostService *HostService) ListFileContext(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto = host.HostDto{
			HostId:   ctx.URLParam("hostId"),
			TenantId: user.TenantId,
			FileName: ctx.URLParam("fileName"),
		}
	)
	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecListFileContext(hostDto)
	return resultDto
}

func (hostService *HostService) EditFile(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		hostDto host.HostDto
	)

	if err := ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败" + err.Error())
	}

	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	hostDto.Ip = hostDtos[0].Ip

	resultDto, _ := shell.ExecEditFile(hostDto)
	return resultDto
}

func (hostService *HostService) UploadFile(ctx iris.Context) result.ResultDto {

	var (
		hostDto host.HostDto
	)

	ctx.SetMaxRequestBodySize(maxSize)
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	hostDto.HostId = ctx.FormValue("hostId")
	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if len(hostDtos) < 1 {
		return result.Error("主机不存在")
	}

	hostDto.CurPath = ctx.FormValue("curPath")
	hostDto.Ip = hostDtos[0].Ip
	resultDto, _ := shell.ExecUploadFile(hostDto, file, fileHeader)
	return resultDto

}

func (hostService *HostService) DownloadFile(ctx iris.Context) {

	var (
		hostDto host.HostDto
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	hostDto.HostId = ctx.URLParam("hostId")
	hostDto.TenantId = user.TenantId

	hostDtos, err := hostService.hostDao.GetHosts(hostDto)

	if err != nil {
		ctx.WriteString(err.Error())
		return
	}

	if len(hostDtos) < 1 {
		ctx.WriteString("主机不存在")
		return
	}

	responseWriter := ctx.ResponseWriter()
	hostDto.FileName = path.Join(ctx.FormValue("curPath"), ctx.URLParam("fileName"))
	hostDto.CurPath = ctx.FormValue("curPath")
	hostDto.Ip = hostDtos[0].Ip
	responseWriter.Header().Set("Content-Disposition", "attachment; filename="+ctx.URLParam("fileName"))
	shell.ExecDownloadFile(hostDto,responseWriter)

}
