package iface

import (
	"fmt"
	"github.com/songgao/water"
	"github.com/zihao-boy/zihao/common/vpn/cache"
	"github.com/zihao-boy/zihao/common/vpn/header"
	"github.com/zihao-boy/zihao/common/vpn/server"
	"runtime"
	"time"
)

var TUNCHANBUFFSIZE = 1024

type TunServer struct {
	TunConn *water.Interface
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
		iface *water.Interface
		err error
	)

	sysType := runtime.GOOS
	if sysType == "windows" {
		iface, err = water.New(water.Config{
			DeviceType: water.TAP,
		})
	} else {
		config := water.Config{
			DeviceType: water.TAP,
		}
		config.Name = "O_O"
		iface, err = water.New(config)
	}
	if err != nil {
		return nil, err
	}
		ts.TunConn = iface

	return ts, nil
}

func (ts *TunServer) Start() {
	//tun to client
	go func() {
		defer func() {
			recover()
		}()

		for {
			data := make([]byte, server.READBUFFERSIZE)
			if n, err := ts.TunConn.Read(data); err == nil && n > 0 {
				if proto, src, dst, err := header.GetBase(data); err == nil {
					key := proto + ":" + dst + ":" + src
					if outputChan := ts.RouteMap.Get(key); outputChan != nil {
						go func() {
							defer func() {
								recover()
							}()
							outputChan.(chan string) <- string(data[:n])
						}()
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
			}
		}
	}()
}

func (ts *TunServer) StartClient(client string, connToTunChan chan string, tunToConnChan chan string) {
	go func() {
		defer func() {
			recover()
		}()

		for {
			data, ok := <-connToTunChan
			if !ok {
				return
			}
			if proto, src, dst, err := header.GetBase([]byte(data)); err == nil {
				key := proto + ":" + src + ":" + dst
				ts.RouteMap.Put(key, tunToConnChan)
				ts.InputChan <- data
				fmt.Println("ToTun: protocol:%v, src:%v, dst:%v", proto, src, dst)
			}
		}
	}()
}

func (ts *TunServer) Stop() {
	fmt.Println("TunServer stopped")
	defer func() {
		recover()
	}()

	close(ts.InputChan)
	ts.RouteMap.Clear()
}

