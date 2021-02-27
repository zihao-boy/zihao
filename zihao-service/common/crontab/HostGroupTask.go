package crontab

import (
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

type HostGroupTask struct {
	MonitorHostGroupDto *monitor.MonitorHostGroupDto
}

func (h HostGroupTask) Run() {
	panic("implement me")
}

