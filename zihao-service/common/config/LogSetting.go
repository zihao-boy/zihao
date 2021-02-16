package config

import (
	"fmt"
"github.com/kataras/golog"
"github.com/kataras/iris/v12"
"github.com/kataras/iris/v12/middleware/logger"
"log"
"os"
"strings"
"time"
)

const (
	deleteFileOnExit = false // 进程退出是否关闭当前日志文件
	sysTimeform      = "2006/01/02 - 15:04:05"
	sysTimeformShort = "2006-01-02"
	dir              = "logs"
	reqLogDir        = dir + "/request/"
	otherLogDir      = dir + "/other/"
)

func InitLog() {
	var (
		err error
	)
	if err = os.MkdirAll(reqLogDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err = os.MkdirAll(otherLogDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	newOtherLogger()
	return
}

func newLogFile(filename string) *os.File {
	var (
		file *os.File
		err  error
	)
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

// 记录请求的 logger
func NewRequestLogger() (h iris.Handler, close func() error) {
	var (
		err               error
		conf              logger.Config
		logFile           = newLogFile(reqLogDir + time.Now().Format(sysTimeformShort) + ".log")
		excludeExtensions = [...]string{
			".js",
			".css",
			".jpg",
			".png",
			".ico",
			".svg",
		}
	)

	close = func() error { return nil }
	conf = logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	close = func() error {
		err = logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}
	conf.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		var output = logger.Columnize(now.Format(sysTimeform), latency, status, ip, method, path, message, headerMessage)
		fmt.Println(output)
		logFile.Write([]byte(output))
	}
	// 不想使用记录器，一些静态请求等
	conf.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(conf)
	return
}

// 其他打日志的 logger
func newOtherLogger() {
	var (
		logFile = newLogFile(otherLogDir + time.Now().Format(sysTimeformShort) + ".log")
	)
	golog.AddOutput(logFile)
}

