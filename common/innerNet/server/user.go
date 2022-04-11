package server

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/innerNet/encrypt"
	"github.com/zihao-boy/zihao/common/innerNet/header"
	"github.com/zihao-boy/zihao/common/innerNet/io"
	"net"
	"strings"
	"time"
)

var USERCHANBUFFERSIZE = 1024
var READBUFFERSIZE = 65535

type User struct {
	Client        string
	Protocol      string
	RemoteTunIp   string
	LocalTunIp    string
	Token         string
	Key           string
	TunToConnChan chan string
	ConnToTunChan chan string
	Conn          net.Conn
	Logout        func(client string)
	HeartbeatTime time.Time
}

func NewUser(client string, protocol string, tun string, token string, conn net.Conn, logout func(string)) *User {
	key := string(encrypt.GetAESKey([]byte(token)))
	return &User{
		Client:        client,
		Protocol:      protocol,
		LocalTunIp:    tun,
		RemoteTunIp:   tun,
		Token:         token,
		Key:           key,
		TunToConnChan: make(chan string, USERCHANBUFFERSIZE),
		ConnToTunChan: make(chan string, USERCHANBUFFERSIZE),
		Conn:          conn,
		Logout:        logout,
		HeartbeatTime: time.Now().Add(10*time.Second),
	}
}

func (user *User) Start() {
	encryptKey := encrypt.GetAESKey([]byte(user.Token))
	//read from client, write to channel
	go func() {
		for {
			var err error
			var data []byte

			if user.Protocol == "tcp" {
				data, err = io.ReadPacket(user.Conn)
			}

			if err != nil {
				fmt.Println("io.ReadPacket err",err)
				user.Close()
				return
			}

			if ln := len(data); ln > 0 {
				if data, err = encrypt.DecryptAES(data, encryptKey); err == nil {
					if proto, src, dst, err := header.GetBase(data); err == nil {
						remoteIp, _ := header.ParseAddr(src)
						user.RemoteTunIp = remoteIp
						//user, ok := lm.Users[client];
						Snat(data, user.LocalTunIp)
						user.ConnToTunChan <- string(data)
						fmt.Println("From %v client: client:%v, protocol:%v, len:%v, src:%v, dst:%v", user.Protocol, user.Client, proto, ln, src, dst)
						continue
					}
				}
				ipData := string(data)
				if !strings.HasPrefix(ipData, "ping") {
					continue
				}
				user.HeartbeatTime = time.Now().Add(10*time.Second)
				_, err = io.WritePacket(user.Conn, data)
				if err != nil{
					fmt.Println("deal data err",err)
				}

			}
		}
	}()



	//read from channel, write to client
	go func() {
		for {
			datas, ok := <-user.TunToConnChan
			if !ok {
				user.Close()
				return
			}
			data := []byte(datas)
			if ln := len(data); ln > 0 {
				if proto, src, dst, err := header.GetBase(data); err == nil {
					Dnat(data, user.RemoteTunIp)
					if endata, err := encrypt.EncryptAES(data, encryptKey); err == nil {
						if user.Protocol == "tcp" {
							_, err = io.WritePacket(user.Conn, endata)

						}

						if err != nil {
							user.Close()
							return
						}
						fmt.Println("To %v client: client:%v, protocol:%v, len:%v, src:%v, dst:%v", user.Protocol, user.Client, proto, ln, src, dst)
					}
				}
			}
		}
	}()

	//check heart beat
	go func() {
		for{
			time.Sleep(10 * time.Second)
			if user.HeartbeatTime.After(time.Now()){
				continue
			}
			user.Close()
		}

	}()
}

func (user *User) Close() {
	go func() {
		defer func() {
			recover()
		}()
		close(user.TunToConnChan)
	}()

	go func() {
		defer func() {
			recover()
		}()
		close(user.ConnToTunChan)
	}()

	go func() {
		defer func() {
			recover()
		}()
		user.Conn.Close()
	}()

	user.Logout(user.Client)
}
