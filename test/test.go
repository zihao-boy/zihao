package main

import (
	"fmt"
	"strings"
)

func main() {
	dockerfile := "# 指定源于一个基础镜像\nFROM registry.cn-beijing.aliyuncs.com/sxd/ubuntu-java8:1.0\n" +
		"# 维护者/拥有者\nMAINTAINER xxx <xxx@xx.com>\n# 从宿主机上传文件 ，这里上传一个脚本，\n" +
		"# 点击选择业务包上传\n" +
		"# ADD bin/start_api.sh /root/\n" +
		"ADD       460b7bda-7241-459d-8a1a-b659ce9421dd/service-oa.jar    /root    \nADD ae979956-99b0-440e-9d04-05c70af1df25/start_jar.sh /root\n\n# 容器内执行相应指令\nRUN chmod u+x /root/start_jar.sh\n# 运行命令\n# CMD <command>   or CMD [<command>]\n# 整个Dockerfile 中只能有一个,多个会被覆盖的\nCMD [\"/root/start_jar.sh\", \"oa\"]"
	fmt.Println(dockerfile)
	dockerfileLines := strings.Split(dockerfile,"\n")
	for _,dockerfileLine := range dockerfileLines{
		dockerfileLine = strings.TrimLeft(dockerfileLine," ")
		dockerfileLine = strings.TrimRight(dockerfileLine," ")

		// comment
		if(strings.HasPrefix(dockerfileLine,"#")){
			continue
		}

		if(strings.HasPrefix(dockerfileLine,"ADD") || strings.HasPrefix(dockerfileLine,"COPY")){
			start := strings.Index(dockerfileLine," ")
			end := strings.LastIndex(dockerfileLine," ")
			addLine := strings.TrimLeft(dockerfileLine[start:end]," ")
			addLine = strings.TrimRight(addLine," ")
			fmt.Println(addLine)
		}
	}
}
