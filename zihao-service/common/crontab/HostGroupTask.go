package crontab

import (
	"fmt"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/monitor/service"
	"golang.org/x/crypto/ssh"
)



type HostGroupTask struct {
	MonitorHostGroupDto *monitor.MonitorHostGroupDto
}

func (h *HostGroupTask) Run() {
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
		fmt.Print("连接"+host.Ip+"出错")
		return;
	}

	session, err := client.NewSession()

	defer session.Close()
	if err != nil{
		fmt.Print("连接"+host.Ip+"出错")
		return;
	}

	result, err := session.Output("ls")
	fmt.Print(result)

}

