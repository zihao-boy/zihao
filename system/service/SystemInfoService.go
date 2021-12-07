package service

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/system"
	"os/exec"
)

const (
	SYSTEM_NAME string = "梓豪平台"
	VERSION     string = "v1.0"
)

type SystemInfoService struct {
}

/**
查询 系统信息
*/
func (*SystemInfoService) Info(context iris.Context) system.SystemDto {
	var systemDto = system.SystemDto{Id: seq.Generator(), Name: SYSTEM_NAME, Version: VERSION, Time: date.GetNowTimeString()}
	return systemDto
}

//开启容器

func (s *SystemInfoService) StartContainer(ctx iris.Context) (interface{}, error) {

	var (
		err           error
		appServiceDto appService.AppServiceDto
		cmd           *exec.Cmd
	)

	if err = ctx.ReadJSON(&appServiceDto); err != nil {
		return result.Error("解析入参失败"), err
	}

	imagesUrl := appServiceDto.ImagesUrl
	if len(imagesUrl) < 1 {
		return nil, errors.New("no include images url")
	}
	dockerpull := "docker pull " + imagesUrl

	//从镜像仓库拉去镜像
	cmd = exec.Command("bash", "-c", dockerpull)
	output, _ := cmd.CombinedOutput()

	fmt.Print("构建镜像：" + dockerpull + " 返回：" + string(output))

	dockerRun := "docker run "

	//端口
	if len(appServiceDto.AppServicePorts) > 0 {
		for _, tmpPort := range appServiceDto.AppServicePorts {
			dockerRun += (" -p " + tmpPort.SrcPort + ":" + tmpPort.TargetPort)
		}
	}

	if len(appServiceDto.AppServiceDirs) > 0 {
		for _, tmpDir := range appServiceDto.AppServiceDirs {
			dockerRun += (" -v " + tmpDir.SrcDir + ":" + tmpDir.TargetDir)
		}
	}

	if len(appServiceDto.AppServiceHosts) > 0 {
		for _, tmpHosts := range appServiceDto.AppServiceHosts {
			dockerRun += (" --add-host=" + tmpHosts.Hostname + ":" + tmpHosts.Ip)
		}
	}

	if len(appServiceDto.AppServiceVars) > 0 {
		for _, tmpVars := range appServiceDto.AppServiceVars {
			dockerRun += (" -e \"" + tmpVars.VarSpec + "=" + tmpVars.VarValue + "\"")
		}
	}

	dockerRun += "--name=\"" + appServiceDto.AsName + "\" " + imagesUrl

	//运行镜像
	cmd = exec.Command("bash", "-c", dockerRun)
	output, _ = cmd.CombinedOutput()
	fmt.Print("构建镜像：" + dockerpull + " 返回：" + string(output))
	return result.SuccessData(string(output)), nil
}
