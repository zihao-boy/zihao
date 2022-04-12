package iface

import (
	"errors"
	"fmt"
	"github.com/songgao/water"
	"github.com/zihao-boy/zihao/common/innerNet/cache"
	"github.com/zihao-boy/zihao/common/innerNet/header"
	"golang.zx2c4.com/wireguard/tun"
	"runtime"
	"strings"
	"time"
)

var TUNCHANBUFFSIZE = 1024
var READBUFFERSIZE = 65535

type TunServer struct {
	TunConn    *water.Interface
	WinTunConn tun.Device
	//Key: clientProtocol:clientIP:clientPort  Value: chan string
	RouteMap *cache.Cache
	//write to tun
	InputChan chan string
}

func NewTunServer(tname string, mtu int) (*TunServer, error) {
	ts := &TunServer{
		RouteMap:  cache.NewCache(time.Minute * 10),
		InputChan: make(chan string, TUNCHANBUFFSIZE),
	}

	var (
		iface    *water.Interface
		iwinface tun.Device
		err      error
	)

	sysType := runtime.GOOS
	if sysType == "windows" {
		//iface, err = water.New(water.Config{
		//	DeviceType: water.TUN,
		//})
		ifname := tname
		iwinface, err = tun.CreateTUN(ifname, 1500)
	} else if sysType == "darwin" {
		config := water.Config{
			DeviceType: water.TUN,
		}
		//config.Name = tname
		iface, err = water.New(config)
	} else {
		config := water.Config{
			DeviceType: water.TUN,
		}
		//config.Name = tname
		iface, err = water.New(config)
	}
	if err != nil {
		return nil, err
	}

	// 保存原始设备句柄
	if sysType == "windows" {
		ts.WinTunConn = iwinface
	} else {
		ts.TunConn = iface
	}

	return ts, nil
}

func (ts *TunServer) Start() {
	//tun to client
	go func() {
		defer func() {
			recover()
		}()

		var (
			n   int
			err error
		)

		for {
			data := make([]byte, 1500)
			sysType := runtime.GOOS
			if sysType == "windows" {
				n, err = ts.WinTunConn.Read(data, 0)
			} else {
				n, err = ts.TunConn.Read(data)
			}

			fmt.Println("网卡中读取数据", n, err, string(data))
			if err == nil && n > 0 {
				proto, src, dst, err := header.GetBase(data)
				fmt.Println("解析数据", src, dst, err)
				if err == nil {
					dstClient := proto + ":" + strings.Split(dst, ":")[0]
					fmt.Println("dstClient", dstClient)
					dstTunToConnChan := ts.RouteMap.Get(dstClient)
					if dstTunToConnChan != nil {
						dstTunToConnChan.(chan string) <- string(data)
					}
				}
			}
		}
	}()

	//chan to tun
	go func() {
		defer func() {
			recover()
		}()
		for {
			if data, ok := <-ts.InputChan; ok && len(data) > 0 {
				ts.TunConn.Write([]byte(data))
				sysType := runtime.GOOS
				if sysType == "windows" {
					ts.WinTunConn.Write([]byte(data), 0)
				} else {
					ts.TunConn.Write([]byte(data))
				}
			}
		}
	}()
}

func (ts *TunServer) StartClient(client string, connToTunChan chan string, tunToConnChan chan string, localTunIp string, proto string) {
	go func() {
		defer func() {
			recover()
		}()

		key := proto + ":" + localTunIp
		fmt.Println("localTunIp", key)
		oldConn := ts.RouteMap.Get(key)
		if oldConn != nil {
			ts.RouteMap.Delete(key)
		}
		ts.RouteMap.Put(key, tunToConnChan)

		for {
			fmt.Println("读取数据到tun")
			data, ok := <-connToTunChan
			if !ok {
				return
			}
			if proto, src, dst, err := header.GetBase([]byte(data)); err == nil {
				//key := proto + ":" + src + ":" + dst
				// 检查是否目标用户是否存在
				dstClient := "tcp:" + strings.Split(dst, ":")[0]
				fmt.Println("dstClient", dstClient)
				dstTunToConnChan := ts.RouteMap.Get(dstClient)
				if dstTunToConnChan != nil {
					//dstTunToConnChan.(chan string) <- string(data)
					err = ts.pushData(dstTunToConnChan.(chan string), string(data), 3)
					if err != nil{
						fmt.Println("通道超时",err)
					}
					continue
				}
				//if ts.RouteMap.GetAndSendData(dstClient, string(data)) {
				//	continue
				//}
				ts.InputChan <- data
				fmt.Println("ToTun: protocol:%v, src:%v, dst:%v", proto, src, dst)
			}
		}
	}()
}

func (ts *TunServer) pushData(q chan string, item string, timeoutSecs int) error {

	select {
	case q <- item:
		return nil
	case <-time.After(time.Duration(timeoutSecs) * time.Second):
		return errors.New("queue full, wait timeout")

	}
}

func (ts *TunServer) Stop() {
	fmt.Println("TunServer stopped")
	defer func() {
		recover()
	}()

	ts.TunConn.Close()

	close(ts.InputChan)
	ts.RouteMap.Clear()
}
