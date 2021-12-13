package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/ls"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/system"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	SYSTEM_NAME string = "梓豪平台"
	VERSION     string = "v1.0"
)
const maxSize = 1000 * iris.MB // 第二种方法

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
		err     error
		hostDto host.HostDto
		//cmd      *exec.Cmd
		//outParam string
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}

	dir, err := ioutil.ReadDir(hostDto.CurPath)

	var lss = make([]ls.LsDto, 0)
	for _, fil := range dir {
		groupName := "d"
		if fil.IsDir() {
			lsDto := ls.LsDto{
				GroupName: groupName,
				Name:      fil.Name(),
				Privilege: fil.Mode().String(),
				Size:      fil.Size(),
			}
			lss = append(lss, lsDto)
		}
	}

	for _, fil := range dir {
		groupName := "d"
		if !fil.IsDir() {
			groupName = "-"
			lsDto := ls.LsDto{
				GroupName: groupName,
				Name:      fil.Name(),
				Privilege: fil.Mode().String(),
				Size:      fil.Size(),
			}
			lss = append(lss, lsDto)
		}
	}

	return result.SuccessData(lss), nil
}

func (s *SystemInfoService) RemoveFile(ctx iris.Context) (interface{}, error) {
	var (
		err     error
		hostDto host.HostDto
		//cmd      *exec.Cmd
		//outParam string
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}
	if hostDto.FileGroupName == "d" {
		err = os.RemoveAll(hostDto.FileName)
	} else {
		err = os.Remove(hostDto.FileName)
	}

	if err != nil {
		return result.Error(err.Error()), nil
	}

	return result.Success(), nil
}

func (s *SystemInfoService) NewFile(ctx iris.Context) (interface{}, error) {
	var (
		err     error
		hostDto host.HostDto
		//cmd      *exec.Cmd
		//outParam string
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}
	if hostDto.FileGroupName == "d" {
		err = os.Mkdir(hostDto.FileName, os.ModePerm)
	} else {
		file, _ := os.Create(hostDto.FileName)
		defer file.Close()
	}

	if err != nil {
		return result.Error(err.Error()), nil
	}

	return result.Success(), nil
}

func (s *SystemInfoService) RenameFile(ctx iris.Context) (interface{}, error) {
	var (
		err     error
		hostDto host.HostDto
		//cmd      *exec.Cmd
		//outParam string
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}
	//if hostDto.FileGroupName == "d"{
	//	err = os.Mkdir(hostDto.FileName,os.ModePerm)
	//}else{
	os.Rename(hostDto.FileName, hostDto.NewFileName)
	//}

	if err != nil {
		return result.Error(err.Error()), nil
	}

	return result.Success(), nil
}

func (s *SystemInfoService) ListFileContext(ctx iris.Context) (interface{}, interface{}) {
	var (
		err     error
		hostDto host.HostDto
		//cmd      *exec.Cmd
		//outParam string
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}

	file, err := os.Open(hostDto.FileName)
	if err != nil {
		return result.Error(err.Error()), nil
	}
	defer file.Close()
	fi, _ := file.Stat()
	if fi.Size() > 1024*1024 {
		return result.Error("文件超过1M不能读取内容"), nil
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return result.Error(err.Error()), nil
	}

	return result.SuccessData(string(content)), nil
}

func (s *SystemInfoService) EditFile(ctx iris.Context) (interface{}, interface{}) {
	var (
		err     error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		return result.Error("解析入参失败"), err
	}
	f, err := os.OpenFile(hostDto.FileName, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return result.Error(err.Error()), nil
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write([]byte (hostDto.FileContext))
	if err != nil {
		return result.Error(err.Error()), nil
	}
	writer.Flush()
	return result.Success(), nil
}

func (s *SystemInfoService) UploadFile(ctx iris.Context) (interface{}, error) {

	ctx.SetMaxRequestBodySize(maxSize)

	file, fileHeader, err := ctx.FormFile("uploadFile")
	defer file.Close()
	if err != nil {
		return result.Error("上传失败" + err.Error()), nil
	}

	dest := ctx.FormValue("curPath")

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest = filepath.Join(dest, fileHeader.Filename)

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败" + err.Error()), nil
	}

	return result.Success(), nil
}

func (s *SystemInfoService) DownloadFile(ctx iris.Context) {
	var (
		err     error
		hostDto host.HostDto
	)

	if err = ctx.ReadJSON(&hostDto); err != nil {
		fmt.Print(err)
		return
	}
	file, err := os.Open(hostDto.FileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	//content, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Print(err)
	//	return
	//}
	stat, _ := file.Stat()
	responseWriter := ctx.ResponseWriter()
	responseWriter.Header().Set("Content-Type", "application/octet-stream")
	responseWriter.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))
	responseWriter.Header().Set("Content-Disposition", "attachment; filename="+hostDto.FileName)
	//responseWriter.Write(content)
	io.Copy(responseWriter, file)
	responseWriter.Flush()
}
