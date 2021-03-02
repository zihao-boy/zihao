package crontab

import (
	"github.com/robfig/cron"
	task2 "github.com/zihao-boy/zihao/zihao-service/common/task"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
	"sync"
)


var lock sync.Mutex

var c *cron.Cron

type MonitorJob struct {

	monitorHostGroupDao dao.MonitorHostGroupDao
}

func (task MonitorJob)init(){

	lock.Lock()
	defer lock.Unlock()
	if c != nil{
		return
	}

	c = cron.New()

}


//启动多个任务
func (task MonitorJob)Start() error{
	var (
		hostGroups []*monitor.MonitorHostGroupDto
		err error
	)

	task.init()

	//查询host_group
	var monitorHostGroup = monitor.MonitorHostGroupDto{
		State:"3301",
	}
	 hostGroups,err = task.monitorHostGroupDao.GetMonitorHostGroups(monitorHostGroup)
	 if err != nil{
	 	return err
	 }
	 //没有任务时不启动
	 if len(hostGroups) < 1{
	 	return nil
	 }


	 for _,item := range hostGroups{
		 //AddJob方法
		 c.AddJob(item.MonCron, task2.HostGroupTask{
			 MonitorHostGroupDto: item,
		 })
	 }

	//启动计划任务
	c.Start()


	return nil
}


//启动多个任务
func (task MonitorJob)Restart() {
	//停止 所有定时器
	if c != nil{
		c.Stop()
	}
	//启动还没有停止的任务
	task.Start()

}
