package task

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/queue/monitorHostQueue"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/monitor/dao"
	"github.com/zihao-boy/zihao/monitor/service"
	"golang.org/x/crypto/ssh"
)

type HostGroupTask struct {
	MonitorHostGroupDto *monitor.MonitorHostGroupDto
}

func (h HostGroupTask) Run() {
	var (
		group           *monitor.MonitorHostGroupDto
		hostService     service.MonitorHostService
		monitorHostDto  monitor.MonitorHostDto
		monitorHostDtos []*monitor.MonitorHostDto
		err             error
	)
	group = h.MonitorHostGroupDto

	//查询监控组下的主机
	monitorHostDto = monitor.MonitorHostDto{MhgId: group.MhgId}
	monitorHostDtos, err = hostService.GetMonitorHostAll(monitorHostDto)

	if err != nil {
		return
	}

	for _, monitorHostDto := range monitorHostDtos {
		h.checkHost(monitorHostDto, *group)
	}
}

/**
检查host
cpu_rate,cpu 使用率
mem_rate,内存使用率
disk_rate,磁盘使用率
free_mem,空闲内存，单位为G
free_disk,空闲磁盘单位为G

*/
func (h *HostGroupTask) checkHost(host *monitor.MonitorHostDto, group monitor.MonitorHostGroupDto) {

	client, err := ssh.Dial("tcp", host.Ip, &ssh.ClientConfig{
		User:            host.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(host.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		fmt.Print("连接" + host.Ip + "出错" + err.Error())
		return
	}

	session, err := client.NewSession()

	defer session.Close()
	if err != nil {
		fmt.Print("连接" + host.Ip + "出错+err.Error()")
		return
	}

	//总内存
	//totalMem, _ := session.Output("free -m | sed -n '2p' | awk '{print $2}'")
	// 使用内存
	userMem, _ := session.Output(strings.ReplaceAll(constants.Check_host_resource_shell, "$1", host.MonDisk))
	//userMem, _ := session.Output("top -b -n 1 | grep Cpu | awk '{print $2}'")
	//userMem, _ := session.Output("df -m /dev/sda1 | grep /dev/sda1 | awk '{print $2}'")

	var (
		monitorCheckHostInfoDto *monitor.MonitorCheckHostInfoDto
		monitorHostDao          = dao.MonitorHostDao{}
		outData                 = strings.ReplaceAll(string(userMem), "'", "\"")
	)

	json.Unmarshal([]byte(outData), &monitorCheckHostInfoDto)

	host.CpuRate = strconv.FormatFloat(monitorCheckHostInfoDto.CpuRate/100, 'f', -1, 64)
	if monitorCheckHostInfoDto.MemTotal != 0 {
		host.MemRate = strconv.FormatFloat(monitorCheckHostInfoDto.MemUsed/monitorCheckHostInfoDto.MemTotal,
			'f', -1, 64)
	}
	host.FreeMem = strconv.FormatFloat(monitorCheckHostInfoDto.MemTotal-monitorCheckHostInfoDto.MemUsed,
		'f', -1, 64)

	if monitorCheckHostInfoDto.DiskTotal != 0 {
		host.DiskRate = strconv.FormatFloat(monitorCheckHostInfoDto.DiskUsed/monitorCheckHostInfoDto.DiskTotal,
			'f', -1, 64)
	}

	host.FreeDisk = strconv.FormatFloat(monitorCheckHostInfoDto.DiskTotal-monitorCheckHostInfoDto.DiskUsed,
		'f', -1, 64)

	host.MonDate = date.GetNowDateString()

	monitorHostDao.UpdateMonitorHost(*host)

	host.NoticeType = group.NoticeType
	//存储告警记录
	saveMonitorHostLog(host)
	//告警
	monitorHostQueue.SendData(*host)

}

func saveMonitorHostLog(host *monitor.MonitorHostDto) {

	var monitorHostLogDto = monitor.MonitorHostLogDto{}
	var monitorHostDao dao.MonitorHostLogDao
	monitorHostLogDto.TenantId = host.TenantId
	monitorHostLogDto.HostId = host.HostId
	monitorHostLogDto.LogId = seq.Generator()
	monitorHostLogDto.CpuRate = host.CpuRate
	monitorHostLogDto.DiskRate = host.DiskRate
	monitorHostLogDto.MemRate = host.MemRate

	_ = monitorHostDao.SaveMonitorHostLog(monitorHostLogDto)
}
