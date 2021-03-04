package monitorHostQueue

import (
	"fmt"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"sync"
	"time"
)
var lock sync.Mutex
var que chan monitor.MonitorHostDto


/**
初始化
 */
func initQueue()  {
	lock.Lock()
	if que != nil{
		lock.Unlock()
		return
	}
	que = make(chan monitor.MonitorHostDto, 200)
	lock.Unlock()

	go func() {
		select {
			case data := <- que:
				dealData(data)
			case <-time.After(time.Second * 2):
				fmt.Println("timeout 2 seconds")
		}
	}()

}

func SendData(host monitor.MonitorHostDto)  {
	initQueue()

	que <- host

}

func dealData(host monitor.MonitorHostDto)  {

	fmt.Print(host.Ip)

}
