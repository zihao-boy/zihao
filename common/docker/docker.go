package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// 镜像结构
type Image struct {
	Created  uint64
	Id string
	ParentId string
	RepoTags []string
	Size uint64
	VirtualSize uint64
}

// 容器结构
type Container struct {
	Id string `json:"Id"`
	Names []string `json:"Names"`
	Image string `json:"Image"`
	ImageID string `json:"ImageID"`
	Command string `json:"Command"`
	Created uint64 `json:"Created"`
	State string `json:"State"`
	Status string `json:"Status"`
	Ports []Port `json:"Ports"`
	Labels map[string]string `json:"Labels"`
	HostConfig map[string]string `json:"HostConfig"`
	NetworkSettings map[string]interface{} `json:"NetworkSettings"`
	Mounts []Mount `json:"Mounts"`
}

// docker 端口映射
type Port struct {
	IP string `json:"IP"`
	PrivatePort int `json:"PrivatePort"`
	PublicPort int `json:"PublicPort"`
	Type string `json:"Type"`
}

// docker 挂载
type Mount struct {
	Type string `json:"Type"`
	Source string `json:"Source"`
	Destination string `json:"Destination"`
	Mode string `json:"Mode"`
	RW bool `json:"RW"`
	Propatation string `json:"Propagation"`
}

// 连接列表
var SockAddr = "/var/run/docker.sock"
var imagesSock = "GET /images/json HTTP/1.0\r\n\r\n"
var containerSock = "GET /containers/json?all=true HTTP/1.0\r\n\r\n"
var startContainerSock = "POST /containers/%s/start HTTP/1.0\r\n\r\n"

// 获取 unix sock 连接
func connectDocker() (*net.UnixConn, error) {
	addr := net.UnixAddr{SockAddr, "unix"}
	return net.DialUnix("unix", nil, &addr)
}

// 启动容器
func StartContainer(id string) error {
	conn, err := connectDocker()
	if err != nil {
		return errors.New("connect error: " + err.Error())
	}

	start := fmt.Sprintf(startContainerSock, id)
	fmt.Println(start)
	cmd := []byte(start)
	code, err := conn.Write(cmd)
	if err != nil {
		return err
	}
	log.Println("start container response code: ", code)
	// 启动容器等待20秒，防止数据重发
	//time.Sleep(20*time.Second)
	return nil
}

// 获取容器列表
func ReadContainer() ([]Container, error) {
	conn, err := connectDocker()
	if err != nil {
		return nil, errors.New("connectError")
	}

	_, err = conn.Write([]byte(containerSock))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	body := getBody(result)

	var containers []Container
	err = json.Unmarshal(body, &containers)
	if err != nil {
		return nil, err
	}

	log.Println("len of containers: ", len(containers))
	if len(containers) == 0 {
		return nil, errors.New("noContainers")
	}
	return containers, nil
}

// 获取镜像列表
func ReadImage() ([]Image, error) {
	conn, err := connectDocker()
	_, err = conn.Write([]byte(imagesSock))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	body := getBody(result[:])

	var images []Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// 从返回的 http 响应中提取 body
func getBody(result []byte) (body []byte) {
	for i:=0; i<=len(result)-4; i++ {
		if result[i] == 13 && result[i+1] == 10 && result[i+2] == 13 && result[i+3] == 10 {
			body = result[i+4:]
			break
		}
	}
	return
}