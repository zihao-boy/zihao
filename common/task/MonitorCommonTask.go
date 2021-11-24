package task

import (
	"fmt"
	"reflect"

	"github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/monitor/service"
)

type MonitorCommonTask struct {
	MonitorTaskDto *monitor.MonitorTaskDto
}

func (h MonitorCommonTask) Run() {
	var (
		taskDto *monitor.MonitorTaskDto
	)
	taskDto = h.MonitorTaskDto

	h.checkHost(taskDto)

}

/**
检查host
cpu_rate,cpu 使用率
mem_rate,内存使用率
disk_rate,磁盘使用率
free_mem,空闲内存，单位为G
free_disk,空闲磁盘单位为G

*/
func (h *MonitorCommonTask) checkHost(taskDto *monitor.MonitorTaskDto) {

	//根据hostId 查询主机信息
	var (
		hostDao                dao.HostDao
		hostDto                host.HostDto
		hostDtos               []*host.HostDto
		impl                   MonitorCommonTaskImpl
		monitorTaskAttrService service.MonitorTaskAttrService
	)

	fmt.Print(taskDto.TaskName + ",监控执行")

	hostDto = host.HostDto{
		HostId: taskDto.HostId,
	}
	hostDtos, _ = hostDao.GetHosts(hostDto)

	if len(hostDtos) < 1 {
		return
	}
	var monitorTaskAttrDto = monitor.MonitorTaskAttrDto{
		TaskId: taskDto.TaskId,
	}
	monitorTaskAttrDtos, _ := monitorTaskAttrService.GetMonitorTaskAttrAll(monitorTaskAttrDto)
	taskDto.Attr = monitorTaskAttrDtos
	impl = MonitorCommonTaskImpl{
		HostDto: *hostDtos[0],
		TaskDto: *taskDto,
	}

	refImpl := reflect.ValueOf(&impl)

	method := refImpl.MethodByName(taskDto.ClassBean)

	method.Call(make([]reflect.Value, 0))

}
