package dockerfileQueue

import (
	"bufio"
	"fmt"
	appServiceDao "github.com/zihao-boy/zihao/appService/dao"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/containerScheduling"
	"github.com/zihao-boy/zihao/common/costTime"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/notifyMessage"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/softService/dao"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex
var que chan *businessDockerfile.BusinessDockerfileDto

/**
初始化
*/
func initQueue() {

	if que != nil {
		return
	}
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()
	if que != nil {
		return
	}
	que = make(chan *businessDockerfile.BusinessDockerfileDto, 100)

	go readData(que)

}

func SendData(businessDockerfileDto *businessDockerfile.BusinessDockerfileDto) {
	defer costTime.TimeoutWarning("DockerfileQueue", "SendData", time.Now())
	initQueue()
	que <- businessDockerfileDto
}

func readData(que chan *businessDockerfile.BusinessDockerfileDto) {
	for {
		select {
		case data := <-que:
			dealData(data)
		}
	}
}

func dealData(businessDockerfileDto *businessDockerfile.BusinessDockerfileDto) {
	var (
		dockerfile           = businessDockerfileDto.Dockerfile
		tenantId             = businessDockerfileDto.TenantId
		businessImagesDao    dao.BusinessImagesDao
		businessImagesVerDao dao.BusinessImagesVerDao
		appServiceDao        appServiceDao.AppServiceDao
		hostDao              hostDao.HostDao
		f                    *os.File
		err                  error
		cmd                  *exec.Cmd
		version              string = "V" + date.GetNowAString()
	)
	defer costTime.TimeoutWarning("DockerfileQueue", "dealData", time.Now())

	dest := filepath.Join(config.WorkSpace, "businessPackage/"+tenantId)

	tenantDesc := dest
	//打开日志文件
	var logFile *os.File
	if businessDockerfileDto.LogPath == "" {
		businessDockerfileDto.LogPath = path.Join(tenantDesc, seq.Generator()+".log")
		logFile, err = os.Create(businessDockerfileDto.LogPath)
	} else {
		logFile, err = os.OpenFile(businessDockerfileDto.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	defer func() {
		logFile.Close()
	}()
	//及时关闭file句柄
	write := bufio.NewWriter(logFile)
	write.WriteString(">>>>>>>>>>>>>>>>>>>开始制作镜像" + businessDockerfileDto.Name)
	write.Flush()

	notifyMessage.SendMsg(tenantId, "开始制作镜像>"+businessDockerfileDto.Name)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest += "/Dockerfile"

	if utils.IsFile(dest) {
		//f, err = os.OpenFile(dest, os.O_RDWR, 0600)
		os.Remove(dest)
	}
	f, err = os.Create(dest)

	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(dockerfile))
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	//ignore file
	dockerignore := tenantDesc + "/.dockerignore"

	if utils.IsFile(dockerignore) {
		//f, err = os.OpenFile(dest, os.O_RDWR, 0600)
		os.Remove(dockerignore)
	}
	fIgnore, err := os.Create(dockerignore)

	defer fIgnore.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		dockerfileLines := strings.Split(dockerfile, "\n")
		ignoreContext := "*\n"
		for _, dockerfileLine := range dockerfileLines {
			dockerfileLine = strings.TrimLeft(dockerfileLine, " ")
			dockerfileLine = strings.TrimRight(dockerfileLine, " ")

			// comment
			if strings.HasPrefix(dockerfileLine, "#") {
				continue
			}

			if strings.HasPrefix(dockerfileLine, "ADD") || strings.HasPrefix(dockerfileLine, "COPY") {
				start := strings.Index(dockerfileLine, " ")
				end := strings.LastIndex(dockerfileLine, " ")
				addLine := strings.TrimLeft(dockerfileLine[start:end], " ")
				addLine = strings.TrimRight(addLine, " ")
				fmt.Println(addLine)
				ignoreContext += ("!" + addLine + "\n")
			}
		}
		_, err = fIgnore.Write([]byte(ignoreContext))

		if err != nil {
			fmt.Print(err.Error())
		}
	}

	imageRepository, _ := factory.GetMappingValue("IMAGES_REPOSITORY")

	imageName := imageRepository + businessDockerfileDto.Name + ":" + version

	shellScript := "docker build -f " + dest + " -t " + imageName + " ."
	//生成镜像
	cmd = exec.Command("bash", "-c", shellScript)
	cmd.Dir = tenantDesc
	output, _ := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	write.WriteString("构建镜像：" + shellScript + " 返回：" + string(output))
	write.Flush()

	dockerRepositoryUrl, _ := factory.GetMappingValue("DOCKER_REPOSITORY_URL")
	username, _ := factory.GetMappingValue("DOCKER_USERNAME")
	password, _ := factory.GetMappingValue("DOCKER_PASSWORD")
	//登录镜像仓库
	shellScript = "docker login --username=" + username + " --password=" + password + " " + dockerRepositoryUrl
	cmd = exec.Command("bash", "-c", shellScript)

	output, _ = cmd.CombinedOutput()
	write.WriteString("登录：" + shellScript + " 返回：" + string(output))
	write.Flush()
	//推镜像
	notifyMessage.SendMsg(tenantId, "构建镜像完成并开始推镜像>"+businessDockerfileDto.Name)
	shellScript = "docker push " + imageName

	cmd = exec.Command("bash", "-c", shellScript)

	output, _ = cmd.CombinedOutput()

	write.WriteString("推镜像：" + shellScript + " 返回：" + string(output))
	write.Flush()
	//if exits docker images
	tempBusinessImagesDto := businessImages.BusinessImagesDto{
		TenantId: businessDockerfileDto.TenantId,
		Name:     businessDockerfileDto.Name,
	}
	tempBusinessImagesDtos, _ := businessImagesDao.GetBusinessImagess(tempBusinessImagesDto)
	businessImagesDto := businessImages.BusinessImagesDto{}
	if len(tempBusinessImagesDtos) > 0 {
		businessImagesDto.TenantId = businessDockerfileDto.TenantId
		businessImagesDto.Id = tempBusinessImagesDtos[0].Id
		businessImagesDto.Version = version
		businessImagesDto.TypeUrl = imageName
		err = businessImagesDao.UpdateBusinessImages(businessImagesDto)
	} else {
		businessImagesDto.TenantId = businessDockerfileDto.TenantId
		businessImagesDto.CreateUserId = businessDockerfileDto.CreateUserId
		businessImagesDto.Id = seq.Generator()
		if businessDockerfileDto.ImagesId != "" {
			businessImagesDto.Id = businessDockerfileDto.ImagesId
		}
		businessImagesDto.Version = version
		businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_DOCKER
		businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_CUSTOM
		businessImagesDto.TypeUrl = imageName
		businessImagesDto.Name = businessDockerfileDto.Name
		err = businessImagesDao.SaveBusinessImages(businessImagesDto)
	}
	if err != nil {
		fmt.Println("保存镜像失败" + err.Error())
		write.WriteString("保存镜像失败" + err.Error())
		write.Flush()
	}

	// save docker images version
	// dj
	businessImagesVerDto := businessImages.BusinessImagesVerDto{
		Id:       seq.Generator(),
		ImagesId: businessImagesDto.Id,
		Version:  businessImagesDto.Version,
		TypeUrl:  imageName,
		TenantId: businessDockerfileDto.TenantId,
	}

	if businessDockerfileDto.VerId != "" {
		businessImagesVerDto.Id = businessDockerfileDto.VerId
	}
	businessImagesVerDao.SaveBusinessImagesVer(businessImagesVerDto)

	notifyMessage.SendMsg(tenantId, "推镜像完成>"+businessDockerfileDto.Name)
	write.WriteString(">>>>>>>>>>>>>>>>>>>制作镜像（" + businessDockerfileDto.Name + "）完成\n")
	write.Flush()

	// if ActionBuildStart
	if businessDockerfileDto.Action != businessDockerfile.ActionBuildStart {
		return
	}

	appServiceDto := appService.AppServiceDto{
		ImagesId: businessImagesDto.Id,
		//State:    appService.STATE_ONLINE,
	}
	appServiceDtos, err := appServiceDao.GetAppServices(appServiceDto)

	if err != nil || len(appServiceDtos) < 1 {
		fmt.Println("服务不存在或者未运行状态")
		return
	}
	var hosts []*host.HostDto
	for _, appServiceDto := range appServiceDtos {

		if !ifInAvgIds(appServiceDto,businessDockerfileDto.AvgIds){
			continue
		}
		tmpAppServiceDto := appService.AppServiceDto{
			AsId:  appServiceDto.AsId,
			VerId: businessImagesVerDto.Id,
		}
		appServiceDao.UpdateAppService(tmpAppServiceDto)

		//stop app service
		containerScheduling.StopContainer(appServiceDto)

		//start app service

		if appServiceDto.AsDeployType == appService.AS_DEPLOY_TYPE_HOST {
			hostDto := host.HostDto{
				HostId: appServiceDto.AsDeployId,
			}
			hosts, _ = hostDao.GetHosts(hostDto)
		} else {
			hostDto := host.HostDto{
				GroupId: appServiceDto.AsDeployId,
			}
			hosts, _ = hostDao.GetHosts(hostDto)
		}

		if len(hosts) < 1 {
			continue
		}
		//appServiceDto.VerId = businessImagesVerDto.Id

		tmpAppServiceDto = appService.AppServiceDto{
			AsId: appServiceDto.AsId,
			//State:    appService.STATE_ONLINE,
		}
		tmpAppServiceDtos, _ := appServiceDao.GetAppServices(tmpAppServiceDto)

		if err != nil || len(tmpAppServiceDtos) < 1 {
			continue
		}

		containerScheduling.ContainerScheduling(hosts, tmpAppServiceDtos[0])

	}

}

func ifInAvgIds(dto *appService.AppServiceDto, ids string) bool {
	if utils.IsEmpty(ids){
		return true
	}

	ids1 := strings.Split(ids,",")

	for _,id := range ids1{
		if dto.AsGroupId == id{
			return true
		}
	}

	return false;
}
