package shell

import (
	"encoding/json"
	"fmt"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"strconv"
	"strings"

	"github.com/zihao-boy/zihao/entity/dto/host"
	"golang.org/x/crypto/ssh"
)

func ExecShell(host host.HostDto, cmd string) error {
	client, err := ssh.Dial("tcp", host.Ip, &ssh.ClientConfig{
		User:            host.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(host.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		fmt.Print("链接主机失败", err)
		return err
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存
	fmt.Print("主机执行指令", cmd)
	session.Output(cmd)

	return nil

}

func ExecListFiles(host host.HostDto) (result.ResultDto, error){
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/listFiles", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}
