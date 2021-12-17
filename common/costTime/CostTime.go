package costTime

import (
	"fmt"
	"log"
	"os"
	"time"
)

const timeLimit = 0

func TimeoutWarning(tag, detailed string, start time.Time) {
	dis := time.Now().Sub(start).Seconds()
	if dis < timeLimit {
		return
		//log.Warning(log.CENTER_COMMON_WARNING, tag, " detailed:", detailed, "TimeoutWarning using", dis, "s")
		//pubstr := fmt.Sprintf("%s count %v, using %f seconds", tag, count, dis)
		//stats.Publish(tag, pubstr)
	}
	logfile, err := os.OpenFile("/zihao/timeout.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return
	}

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(tag, " detailed:", detailed, "TimeoutWarning using", dis, "s")

}
