package monitorHostQueue

import (
	"fmt"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"sync"
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
	que = make(chan monitor.MonitorHostDto, 1)
	lock.Unlock()

	go readData(que)

}

func SendData(host monitor.MonitorHostDto)  {
	initQueue()
	que <- host
}

func readData(que chan monitor.MonitorHostDto){
	for{
		select {
			case data := <- que:
				dealData(data)
		}
	}
}

func dealData(host monitor.MonitorHostDto)  {
	fmt.Print(host.Ip)
}
