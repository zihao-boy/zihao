package crontab

import (
	"github.com/robfig/cron"
	"github.com/zihao-boy/zihao/business/dao/resourcesBackUpDao"
	task2 "github.com/zihao-boy/zihao/common/task"
	"github.com/zihao-boy/zihao/entity/dto/resources"
)

//var lock sync.Mutex

var backUpCron *cron.Cron

type BackUpJob struct {
	resourcesBackUpDao resourcesBackUpDao.ResourcesBackUpDao
}

func (job BackUpJob) init() {

	if backUpCron != nil {
		return
	}
	lock.Lock()
	defer lock.Unlock()
	if backUpCron != nil {
		return
	}

	backUpCron = cron.New()

}

//启动多个任务
func (job BackUpJob) Start() error {
	var (
		backUps []*resources.ResourcesBackUpDto
	)

	job.init()

	//停止 所有定时器
	backUpCron.Stop()

	//查询host_group
	var resourcesBackUpDto = resources.ResourcesBackUpDto{
		State: resources.Back_up_state_START,
	}
	backUps, _ = job.resourcesBackUpDao.GetResourcesBackUps(resourcesBackUpDto)

	if backUps == nil || len(backUps) < 1 {
		return nil
	}

	for _, item := range backUps {
		//AddJob方法
		backUpCron.AddJob(item.ExecTime, task2.ResourcesBackUpTask{
			ResourcesBackUpDto: item,
		})
	}

	//启动计划任务
	backUpCron.Start()

	return nil
}

//启动多个任务
func (job BackUpJob) Restart() {
	//停止 所有定时器
	//if backUpCron != nil {
	//	backUpCron.Stop()
	//}

	//启动还没有停止的任务
	job.Start()

}

func (job BackUpJob) Stop() {
	job.Restart()
}