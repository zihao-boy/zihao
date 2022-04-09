package main

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/innerNet/client"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"sync"
)

func main() {
	//加载配置文件

	wg := &sync.WaitGroup{}
	wg.Add(2)
	config.InitProp("./conf/innerNetClient.properties")
	serverAddr,_ := config.Prop.Property("serverAddr")
	username,_ := config.Prop.Property("username")
	password,_ := config.Prop.Property("password")
	tunName,_ := config.Prop.Property("tunName")
	innerNetClientDto := innerNet.InnerNetClientDto{
		ServerAddr: serverAddr,
		Username: username,
		Password: password,
		TunName: tunName,
	}
	err := client.StartClient(&innerNetClientDto)
	if err != nil{
		fmt.Println(err)
	}
	wg.Wait()

}
