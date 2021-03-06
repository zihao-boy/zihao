package task

import (
	"encoding/json"
	"fmt"
	"github.com/zihao-boy/zihao/zihao-service/common/date"
	"github.com/zihao-boy/zihao/zihao-service/common/queue/monitorHostQueue"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
	"github.com/zihao-boy/zihao/zihao-service/monitor/service"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
)

const check_shell="#!/bin/bash\n\nmem_total=`free -m | awk '/Mem/ {print $2}'`\n\nmem_used=`free -m | awk '/Mem/ {print $3}'`\n\ndisk_total=`df -m $1 | grep $1 | awk '{print $2}'`\n\nif [ ! -n \"$disk_total\" ]; then\n        disk_total=0\nfi\n\ndisk_used=`df -m $1 | grep $1 | awk '{print $3}'`\n\nif [ ! -n \"$disk_used\" ]; then\n        disk_used=0\nfi\n\ncpu_us=`top -b -n 1 | grep Cpu | awk '{print $2}'`\n\necho \"{'cpuRate':$cpu_us,'memTotal':$mem_total,'memUsed':$mem_used,'diskTotal':$disk_total,'diskUsed':$disk_used}\"\n"


type HostGroupTask struct {
	MonitorHostGroupDto *monitor.MonitorHostGroupDto
}

func (h HostGroupTask) Run() {
	var (
		group *monitor.MonitorHostGroupDto
		hostService service.MonitorHostService
		monitorHostDto monitor.MonitorHostDto
		monitorHostDtos []*monitor.MonitorHostDto
		err error
	)
	group = h.MonitorHostGroupDto

	//查询监控组下的主机
	monitorHostDto = monitor.MonitorHostDto{MhgId: group.MhgId}
	monitorHostDtos,err = hostService.GetMonitorHostAll(monitorHostDto)

	if err != nil{
		return ;
	}

	for _,monitorHostDto := range monitorHostDtos{
		h.checkHost(monitorHostDto)
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
func (h *HostGroupTask) checkHost(host *monitor.MonitorHostDto){

	client, err := ssh.Dial("tcp", host.Ip, &ssh.ClientConfig{
		User: host.Username,
		Auth: []ssh.AuthMethod{ssh.Password(host.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil{
		fmt.Print("连接"+host.Ip+"出错"+err.Error())
		return;
	}

	session, err := client.NewSession()

	defer session.Close()
	if err != nil{
		fmt.Print("连接"+host.Ip+"出错+err.Error()")
		return;
	}

	//总内存
	//totalMem, _ := session.Output("free -m | sed -n '2p' | awk '{print $2}'")
	// 使用内存
	userMem, _ := session.Output(strings.ReplaceAll(check_shell,"$1",host.MonDisk))
	//userMem, _ := session.Output("top -b -n 1 | grep Cpu | awk '{print $2}'")
	//userMem, _ := session.Output("df -m /dev/sda1 | grep /dev/sda1 | awk '{print $2}'")


	var (
		monitorCheckHostInfoDto *monitor.MonitorCheckHostInfoDto
		monitorHostDao=dao.MonitorHostDao{}
		outData = strings.ReplaceAll(string(userMem),"'","\"")
	)

	json.Unmarshal([]byte(outData),&monitorCheckHostInfoDto)

	host.CpuRate = strconv.FormatFloat(monitorCheckHostInfoDto.CpuRate,'f',-1,64)
	if monitorCheckHostInfoDto.MemTotal != 0{
		host.MemRate = strconv.FormatFloat(monitorCheckHostInfoDto.MemUsed/monitorCheckHostInfoDto.MemTotal,
			'f',-1,64)
	}
	host.FreeMem = strconv.FormatFloat(monitorCheckHostInfoDto.MemTotal-monitorCheckHostInfoDto.MemUsed,
		'f',-1,64)

	if monitorCheckHostInfoDto.DiskTotal != 0{
		host.DiskRate = strconv.FormatFloat(monitorCheckHostInfoDto.DiskUsed/monitorCheckHostInfoDto.DiskTotal,
			'f',-1,64)
	}

	host.FreeDisk = strconv.FormatFloat(monitorCheckHostInfoDto.DiskTotal-monitorCheckHostInfoDto.DiskUsed,
		'f',-1,64)

	host.MonDate = date.GetNowDateString()

	monitorHostDao.UpdateMonitorHost(*host)

	//存储告警记录

	//告警
	monitorHostQueue.SendData(*host)

}

