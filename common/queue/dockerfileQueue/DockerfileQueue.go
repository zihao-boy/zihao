package dockerfileQueue

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/softService/dao"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var lock sync.Mutex
var que chan *businessDockerfile.BusinessDockerfileDto

/**
初始化
*/
func initQueue() {
	lock.Lock()
	if que != nil {
		lock.Unlock()
		return
	}
	que = make(chan *businessDockerfile.BusinessDockerfileDto, 100)
	lock.Unlock()

	go readData(que)

}

func SendData(businessDockerfileDto *businessDockerfile.BusinessDockerfileDto) {
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
		dockerfile        = businessDockerfileDto.Dockerfile
		tenantId          = businessDockerfileDto.TenantId
		businessImagesDao dao.BusinessImagesDao
		f                 *os.File
		err               error
		cmd               *exec.Cmd
		version string = "V"+date.GetNowAString()
	)

	dest := filepath.Join(config.WorkSpace, "businessPackage/"+tenantId)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest += "/Dockerfile"

	if  utils.IsFile(dest) {
		f, err = os.OpenFile(dest, os.O_RDONLY|os.O_TRUNC, 0600)
	} else {
		f, err = os.Create(dest)
	}

	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(dockerfile))
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	imageRepository, _ := factory.GetMappingValue("IMAGES_REPOSITORY")

	imageName := imageRepository + businessDockerfileDto.Name + ":" +version

	shellScript := "docker build -f " + dest + " -t " + imageName + " ."
	//生成镜像
	cmd = exec.Command(shellScript)
	output, _ := cmd.Output()
	fmt.Print("构建镜像：" + shellScript +" 返回："+  string(output))

	dockerRepositoryUrl, _ := factory.GetMappingValue("DOCKER_REPOSITORY_URL")
	username, _ := factory.GetMappingValue("DOCKER_USERNAME")
	password, _ := factory.GetMappingValue("DOCKER_PASSWORD")
	//登录镜像仓库
	shellScript = "docker login --username=" + username + " --password=" + password + " " + dockerRepositoryUrl
	cmd = exec.Command(shellScript)

	output, _ = cmd.Output()
	fmt.Print("登录：" + shellScript +" 返回："+  string(output))

	//推镜像
	shellScript = "docker push " + imageName

	cmd = exec.Command(shellScript)

	output, _ = cmd.Output()

	fmt.Print("推镜像：" + shellScript +" 返回："+ string(output))

	businessImagesDto := businessImages.BusinessImagesDto{}
	businessImagesDto.TenantId = businessDockerfileDto.TenantId
	businessImagesDto.CreateUserId = businessDockerfileDto.CreateUserId
	businessImagesDto.Id = seq.Generator()
	businessImagesDto.Version = version
	businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_DOCKER
	businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_CUSTOM
	businessImagesDto.TypeUrl = "docker pull " + imageName
	businessImagesDto.Name = businessDockerfileDto.Name

	err = businessImagesDao.SaveBusinessImages(businessImagesDto)
	if err != nil {
		fmt.Println("保存镜像失败" + err.Error())
	}
}
