package task

import (
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/host"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"golang.org/x/crypto/ssh"
)

type MonitorCommonTaskImpl struct {
	HostDto host.HostDto
	TaskDto monitor.MonitorTaskDto
}

func (h *MonitorCommonTaskImpl) getSession() (*ssh.Session,error) {
	var (
		err error
	)
	client, err := ssh.Dial("tcp", h.HostDto.Ip, &ssh.ClientConfig{
		User: h.HostDto.Username,
		Auth: []ssh.AuthMethod{ssh.Password(h.HostDto.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil{
		return nil,err
	}

	session, err := client.NewSession()

	//defer session.Close()
	if err != nil{
		return nil,err
	}

	return session,nil

}
