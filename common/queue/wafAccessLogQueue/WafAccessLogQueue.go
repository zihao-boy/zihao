package wafAccessLogQueue

import (
	"fmt"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"sync"
)

var lock sync.Mutex
var que chan waf.WafAccessLogDto
const default_access_log_count = 10000
/**
初始化
*/
func initQueue() {
	if que != nil {
		return
	}
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()
	if que != nil {
		return
	}
	que = make(chan waf.WafAccessLogDto, 200)

	go readData(que)

}

func SendData(wafDto waf.WafAccessLogDto) {
	initQueue()

	que <- wafDto
}

func readData(que chan waf.WafAccessLogDto) {
	for {
		select {
		case data := <-que:
			dealData(data)
		}
	}
}

func dealData(wafDto waf.WafAccessLogDto) {
	var wafDao wafDao.WafAccessLogDao
	tmpWafAccessLog := waf.WafAccessLogDto{

	}
	count ,_ := wafDao.GetWafAccessLogCount(tmpWafAccessLog)

	if count > default_access_log_count{
		wafDao.DeleteWafAccessLog(tmpWafAccessLog)
	}

	err := wafDao.SaveWafAccessLog(wafDto)
	if err != nil {
		fmt.Println(err)
	}
}
