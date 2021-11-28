package crontab

import (
	"sync"

	"github.com/robfig/cron"
	task2 "github.com/zihao-boy/zihao/common/task"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/monitor/dao"
)

var lock sync.Mutex

var c *cron.Cron

type MonitorJob struct {
	monitorHostGroupDao dao.MonitorHostGroupDao
	monitorTaskDao      dao.MonitorTaskDao
}

func (task MonitorJob) init() {

	lock.Lock()
	defer lock.Unlock()
	if c != nil {
		return
	}

	c = cron.New()

}

//启动多个任务
func (task MonitorJob) Start() error {
	var (
		hostGroups []*monitor.MonitorHostGroupDto
		taskDtos   []*monitor.MonitorTaskDto
	)

	task.init()

	//查询host_group
	var monitorHostGroup = monitor.MonitorHostGroupDto{
		State: "3301",
	}
	hostGroups, _ = task.monitorHostGroupDao.GetMonitorHostGroups(monitorHostGroup)

	for _, item := range hostGroups {
		//AddJob方法
		c.AddJob(item.MonCron, task2.HostGroupTask{
			MonitorHostGroupDto: item,
		})
	}
	var taskDto = monitor.MonitorTaskDto{
		State: "002",
	}
	taskDtos, _ = task.monitorTaskDao.GetMonitorTasks(taskDto)

	for _, item := range taskDtos {
		//AddJob方法
		c.AddJob(item.TaskCron, task2.MonitorCommonTask{
			MonitorTaskDto: item,
		})
	}
	//启动计划任务
	c.Start()

	return nil
}

//启动多个任务
func (task MonitorJob) Restart() {
	//停止 所有定时器
	if c != nil {
		c.Stop()
	}

	//启动还没有停止的任务
	task.Start()

}
