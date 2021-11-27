package task

import (
	"fmt"
	"time"

	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/config"
)

func doSlaveHealth() {
	mastIp, isExist := config.Prop.Property("mastIp")
	if !isExist {
		mastIp = "127.0.0.1"
	}
	url := "http://" + mastIp + ":7000/"

	slaveId, isExist := config.Prop.Property("slaveId")
	if !isExist {
		slaveId = "-1"
	}

	data := map[string]interface{}{
		"hostId": slaveId,
	}
	resp, err := httpReq.Post(url, data, nil)
	if err != nil {
		fmt.Print(err.Error(), url, data)
	}
	fmt.Print(resp)
}

// slave 心跳
func SlaveHealth() {
	heartbeat := time.Tick(30 * time.Second)
	for {
		select {
		// … do some stuff
		case <-heartbeat:
			fmt.Println("*")
			//… do heartbeat stuff
			doSlaveHealth()
		}
	}
}
