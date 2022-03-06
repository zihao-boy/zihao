package task

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shopspring/decimal"
	"github.com/zihao-boy/zihao/common/docker"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"strconv"
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
	cpuPercent, _ := cpu.Percent(time.Second, false)

	cpuPercentDec := decimal.NewFromFloat(cpuPercent[0])
	cpuPercentDec = cpuPercentDec.Mul(decimal.NewFromInt(int64(cpuCore)))

	//useCpu, _ := cpuPercentDec.Float64()
	// 获取内存
	totalMem, _ := mem.VirtualMemory()
	totalMemDec := decimal.NewFromInt(int64(totalMem.Total))
	totalMemDec = totalMemDec.Div(decimal.NewFromInt(1024 * 1024))
	totalMemValue, _ := totalMemDec.Float64()

	totalMemUseDec := decimal.NewFromInt(int64(totalMem.Used))
	totalMemUseDec = totalMemUseDec.Div(decimal.NewFromInt(1024 * 1024))
	totalMemUseValue, _ := totalMemUseDec.Float64()
	// 获取磁盘
	totalDisk, _ := disk.Usage("/")

	totalDiskDec := decimal.NewFromInt(int64(totalDisk.Total))
	totalDiskDec = totalDiskDec.Div(decimal.NewFromInt(1024 * 1024))
	totalDiskValue, _ := totalDiskDec.Float64()

	totalDiskUseDec := decimal.NewFromInt(int64(totalDisk.Used))
	totalDiskUseDec = totalDiskUseDec.Div(decimal.NewFromInt(1024 * 1024))
	totalDiskUseValue, _ := totalMemUseDec.Float64()

	relContainers, _ := docker.ReadContainer()

	data := map[string]interface{}{
		"hostId": slaveId,
		"cpu":    strconv.FormatInt(int64(cpuCore), 10),
		"mem":    fmt.Sprintf("%.2f", totalMemValue),
		"disk":   fmt.Sprintf("%.2f", totalDiskValue),
		//"useCpu":     fmt.Sprintf("%.2f", useCpu),
		"useCpu":     fmt.Sprintf("%.2f",cpuPercent),
		"useMem":     fmt.Sprintf("%.2f", totalMemUseValue),
		"useDisk":    fmt.Sprintf("%.2f", totalDiskUseValue),
		"containers": relContainers,
	}
	resp, err := httpReq.Post(url, data, nil)
	if err != nil {
		fmt.Print(err.Error(), url, data)
	}
	fmt.Print(resp)
	//var (
	//	resultDto  result.ResultDto
	//	containers []appService.AppServiceContainerDto
	//)
	//json.Unmarshal([]byte(resp), resultDto)
	//
	//dataByte, err := json.Marshal(resultDto.Data)
	//if err != nil {
	//	fmt.Print("解析报文异常", err)
	//	return
	//}
	//
	//json.Unmarshal(dataByte, containers)
	//
	//if len(containers) < 1 {
	//	fmt.Print("主机没有容器")
	//	return
	//}

	//relContainers, err := docker.ReadContainer()

	//if err != nil && err.Error() == "noContainers" {
	//	fmt.Print(err.Error())
	//	for _, container := range containers {
	//		docker.StartContainer(container.ContainerId)
	//	}
	//	return
	//}
	//if err != nil {
	//	fmt.Print(err.Error())
	//	return
	//}
	//for _, container := range containers {
	//	if hasInRealContainers(container, relContainers) {
	//		continue
	//	}
	//	docker.StartContainer(container.ContainerId)
	//}

}

func hasInRealContainers(container appService.AppServiceContainerDto, realContainers []docker.Container) bool {
	for _, realContainer := range realContainers {
		//container exists include restarting and running
		if realContainer.Id == container.ContainerId {
			return true
		}
	}
	return false
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
