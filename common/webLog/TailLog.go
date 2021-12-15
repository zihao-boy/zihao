package webLog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

type TailLog struct {
	Path   string
	ConnId string
	Conn   bool
}

var (
	tailLogs = make(map[string]*TailLog)
)

func WebSocketHandler(data []byte, connId string, nsConn *websocket.NSConn) {

	path := string(data)

	tails, err := tail.TailFile(path, tail.Config{
		ReOpen: true, // 文件被移除或被打包，需要重新打开,基础库会检测，如果文件有改变，会重新打开
		Follow: true, // 实时跟踪
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取
		MustExist: false, //flase日志文件不存在也监控
		Poll:      true,  //不断的去查询
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	tailLog := TailLog{
		ConnId: connId,
		Path:   path,
		Conn:   true,
	}
	tailLogs[connId] = &tailLog

	go func() {
		var msg *tail.Line
		var ok bool
		for tailLogs[connId].Conn {
			msg, ok = <-tails.Lines //chan
			if !ok {
				// ok 是判断管道是否被关闭，如果关闭就是文件被重置了，需要重新读取新的管道
				fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			fmt.Println("msg:", msg)
			mg := websocket.Message{
				Body:     []byte(msg.Text + "\n"),
				IsNative: true,
			}
			nsConn.Conn.Write(mg)
		}
		fmt.Println("--------------> tail finish")
	}()
}

func CloseTail(connId string) {
	tailLog := tailLogs[connId]
	if tailLog.ConnId == "" {
		return
	}

	tailLog.Conn = false
}
