package crontab

import (
	"github.com/robfig/cron"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
)

var hostMonitorTaks *MonitorJob

type MonitorJob struct {
	cron *cron.Cron
	monitorHostGroupDao dao.MonitorHostGroupDao
}


//启动多个任务
func (task *MonitorJob)Start() error{
	var (
		hostGroups []*monitor.MonitorHostGroupDto
		err error
	)

	//查询host_group
	var monitorHostGroup = monitor.MonitorHostGroupDto{
		State:"3302",
	}
	 hostGroups,err = task.monitorHostGroupDao.GetMonitorHostGroups(monitorHostGroup)
	 if err != nil{
	 	return err
	 }
	 //没有任务时不启动
	 if len(hostGroups) < 1{
	 	return nil
	 }

	 c := cron.New()

	 for _,item := range hostGroups{
		 //AddJob方法
		 c.AddJob(item.MonCron, HostGroupTask{
			 MonitorHostGroupDto: item,
		 })
	 }

	//启动计划任务
	c.Start()

	hostMonitorTaks = &MonitorJob{
		cron: c,
	}
	select{}
}


//启动多个任务
func (task *MonitorJob)Stop(_time string) {
	//停止 所有定时器
	task.cron.Stop()
	//启动还没有停止的任务
	task.cron.Start()
	select{}
}
