package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"time"
)

//https://colobu.com/2019/06/01/packet-capture-injection-and-analysis-gopacket/
var (
	device       string = "en0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)


func main() {

	que := make(chan int,10)
	go func() {
		for {
			select {
			case data := <-que:
				fmt.Println(data)
			}
			time.Sleep(time.Second*4)
		}
	}()

	for i := 0;i < 1000;i++{
		select {
		case que <- i:
			fmt.Println("正常写入")
			break
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
			break
		}
	}
}