package task

import (
	"encoding/json"
	"fmt"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/monitor/service"
	"golang.org/x/crypto/ssh"
	"strings"
)

var check_process_shell = "#!/bin/bash\n\nprocess_name=`ps -ef | grep target_process_name | grep -v grep | awk '{print $1}'`\n\nif [ ! -n \"$process_name\" ]; then\n        process_name=0\nfi\n\necho \"{'processName':'$process_name'}\""

type ProcessDto struct {
	ProcessName string `json:"processName"`
}
/**
检查host
cpu_rate,cpu 使用率
mem_rate,内存使用率
disk_rate,磁盘使用率
free_mem,空闲内存，单位为G
free_disk,空闲磁盘单位为G

*/
func (h *MonitorCommonTaskImpl) CheckProcess(){

	var (
		session *ssh.Session
		err error
		monitorEventService service.MonitorEventService

	)

	session,err = h.getSession()

	if err !=nil{
		fmt.Print("连接ip="+h.HostDto.Ip+"出错"+err.Error())
		return
	}

	defer session.Close()
	var targetProcessName string = ""
	for _,item := range h.TaskDto.Attr{
		if item.SpecCd == "100101"{
			targetProcessName=item.Value
		}
	}

	if targetProcessName == ""{
		return
	}

	// 使用内存
	processName, _ := session.Output(strings.ReplaceAll(check_process_shell,"target_process_name",targetProcessName))


	var (
		processDto *ProcessDto
		outData = strings.ReplaceAll(string(processName),"'","\"")
	)

	json.Unmarshal([]byte(outData),&processDto)

	if processDto.ProcessName !="0"{
		return
	}

	var remark string ="主机【"+h.HostDto.Name+"】,ip="+h.HostDto.Ip+","+targetProcessName+"进程不存在"

	//告警
	monitorEventDto := monitor.MonitorEventDto{
		EventType:"1001",
		EventObjId: h.HostDto.HostId,
		EventObjName: h.HostDto.Name,
		TenantId:h.TaskDto.TenantId,
		ThresholdValue:"1",
		CurValue:"0",
		Remark:remark,
		NoticeType:"2002",
	}
	monitorEventService.SaveMonitorEvents(monitorEventDto)


}
