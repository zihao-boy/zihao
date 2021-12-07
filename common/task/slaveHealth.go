package task

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shopspring/decimal"
	"time"

	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/config"
)

func doSlaveHealth() {
	mastIp, isExist := config.Prop.Property("mastIp")
	if !isExist {
		mastIp = "127.0.0.1:7000"
	}
	url := "http://" + mastIp + "/app/host/slaveHealth"

	slaveId, isExist := config.Prop.Property("slaveId")
	if !isExist {
		slaveId = "-1"
	}

	//获取cpu
	cpuCore, _ := cpu.Counts(true)
	cpuPercent, _:= cpu.Percent(time.Second, false)

	cpuPercentDec := decimal.NewFromFloat(cpuPercent[0])
	cpuPercentDec = cpuPercentDec.Mul(decimal.NewFromInt(int64(cpuCore)))

	useCpu,_ :=cpuPercentDec.Float64()
	// 获取内存
	totalMem, _ := mem.VirtualMemory()

	// 获取磁盘
	totalDisk, _ := disk.Usage("/")


	data := map[string]interface{}{
		"hostId":  slaveId,
		"cpu":     cpuCore,
		"mem":     totalMem.Total,
		"disk":    totalDisk.Total,
		"useCpu":  useCpu,
		"useMem":  totalMem.Used,
		"useDisk": totalDisk.Used,
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
