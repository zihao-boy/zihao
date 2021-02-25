package crontab


import (
	"github.com/robfig/cron"
	"log"
	"fmt"
)

type HostMonitorTask struct {
}

func (this HostMonitorTask)Run() {
	fmt.Println("testJob1...")
}

//启动多个任务
func main() {
	i := 0
	c := cron.New()

	//AddFunc
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	//AddJob方法
	c.AddJob(spec, HostMonitorTask{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select{}
}
