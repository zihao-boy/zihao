package service

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/ls"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/system"
	"os/exec"
	"strings"
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
		param         string
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

	dockerRun := "docker run -d "

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
			if tmpVars.VarSpec == "--param" {
				param += (" " + tmpVars.VarValue)
				continue
			}
			dockerRun += (" -e \"" + tmpVars.VarSpec + "=" + tmpVars.VarValue + "\"")
		}
	}

	//dockerRun += " --name=\"" + appServiceDto.AsName + "_" + seq.Generator() + "\" " + imagesUrl

	dockerRun += (" " + imagesUrl + param)

	//运行镜像
	cmd = exec.Command("bash", "-c", dockerRun)
	output, _ = cmd.CombinedOutput()
	fmt.Print("启动容器：" + dockerRun + " 返回：" + string(output))
	paramOut := map[string]interface{}{
		"ContainerId": strings.Replace(string(output), "\n", "", -1),
	}
	return result.SuccessData(paramOut), nil
}

func (s *SystemInfoService) StopContainer(ctx iris.Context) (interface{}, error) {
	var (
		err                    error
		appServiceContainerDto appService.AppServiceContainerDto
		cmd                    *exec.Cmd
	)

	if err = ctx.ReadJSON(&appServiceContainerDto); err != nil {
		return result.Error("解析入参失败"), err
	}

	dockerpull := "docker stop " + appServiceContainerDto.DockerContainerId

	//停止容器
	cmd = exec.Command("bash", "-c", dockerpull)
	output, _ := cmd.CombinedOutput()

	fmt.Print("停止容器：" + string(output))

	dockerpull = "docker rm " + appServiceContainerDto.DockerContainerId

	//停止容器
	cmd = exec.Command("bash", "-c", dockerpull)
	output, _ = cmd.CombinedOutput()

	fmt.Print("删除容器：" + string(output))

	return result.Success(), nil
}

// list files
func (s *SystemInfoService) ListFiles(ctx iris.Context) (interface{}, error) {
	var (
		err      error
		hostDto  host.HostDto
		cmd      *exec.Cmd
		outParam string
		lss      []ls.LsDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}

	shellStr := ("ls -lth " + hostDto.CurPath)

	//停止容器
	cmd = exec.Command("bash", "-c", shellStr)
	output, _ := cmd.CombinedOutput()

	outParam = string(output)
	fmt.Print("删除容器：" + outParam)

	lines := strings.Split(outParam, "\n")

	for index, line := range lines {
		newLine := strings.Split(line, " ")
		groupName := "d"
		if strings.HasSuffix(newLine[0], "-") {
			groupName = "-"
		}
		if len(newLine) < 8 {
			continue
		}
		lsDto := ls.LsDto{
			GroupName: groupName,
			Name:      newLine[8],
			Privilege: newLine[0],
		}
		lss[index] = lsDto
	}

	return result.SuccessData(lss), nil
}
