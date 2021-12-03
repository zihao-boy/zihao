package service

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/cache/factory"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/businessImages"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
)

type BusinessImagesService struct {
	businessImagesDao     dao.BusinessImagesDao
	businessDockerfileDao dao.BusinessDockerfileDao
}

/**
查询 系统信息
*/
func (businessImagesService *BusinessImagesService) GetBusinessImagesAll(businessImagesDto businessImages.BusinessImagesDto) ([]*businessImages.BusinessImagesDto, error) {
	var (
		err                error
		businessImagesDtos []*businessImages.BusinessImagesDto
	)

	businessImagesDtos, err = businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
	if err != nil {
		return nil, err
	}

	return businessImagesDtos, nil

}

/**
查询 系统信息
*/
func (businessImagesService *BusinessImagesService) GetBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err                error
		page               int64
		row                int64
		total              int64
		businessImagesDto  = businessImages.BusinessImagesDto{}
		businessImagesDtos []*businessImages.BusinessImagesDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessImagesDto.Row = row * page

	businessImagesDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesDto.TenantId = user.TenantId

	total, err = businessImagesService.businessImagesDao.GetBusinessImagesCount(businessImagesDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessImagesDtos, err = businessImagesService.businessImagesDao.GetBusinessImagess(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDtos, total, row)

}

/**
保存 系统信息
*/
func (businessImagesService *BusinessImagesService) SaveBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)
	ctx.SetMaxRequestBodySize(maxSize)

	file, fileHeader, err := ctx.FormFile("uploadFile")

	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	defer file.Close()
	dest := filepath.Join(config.G_AppConfig.DataPath, "/businessImages")

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	fileName := path.Ext(fileHeader.Filename)

	dest = filepath.Join(dest, seq.Generator()+fileName)

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		return result.Error("上传失败" + err.Error())
	}

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessImagesDto.TenantId = user.TenantId
	businessImagesDto.CreateUserId = user.UserId
	businessImagesDto.Id = seq.Generator()
	businessImagesDto.Version = "V" + date.GetNowAString()
	businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_IMPORT
	businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_CUSTOM
	businessImagesDto.TypeUrl = dest
	businessImagesDto.Name = ctx.FormValue("name")

	err = businessImagesService.businessImagesDao.SaveBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDto)

}

/**
修改 系统信息
*/
func (businessImagesService *BusinessImagesService) UpdateBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)

	if err = ctx.ReadJSON(&businessImagesDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesService.businessImagesDao.UpdateBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDto)

}

/**
删除 系统信息
*/
func (businessImagesService *BusinessImagesService) DeleteBusinessImages(ctx iris.Context) result.ResultDto {
	var (
		err               error
		businessImagesDto businessImages.BusinessImagesDto
	)

	if err = ctx.ReadJSON(&businessImagesDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessImagesService.businessImagesDao.DeleteBusinessImages(businessImagesDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessImagesDto)

}

/**
删除 系统信息
*/
func (businessImagesService *BusinessImagesService) GeneratorImages(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)

	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	tmpBusinessDockerfileDto := businessDockerfile.BusinessDockerfileDto{
		Id:       businessDockerfileDto.Id,
		TenantId: user.TenantId,
	}
	businessDockerfileDtos, err := businessImagesService.businessDockerfileDao.GetBusinessDockerfiles(tmpBusinessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	if len(businessDockerfileDtos) < 1 {
		return result.Error("dockerfile 不存在")
	}

	go doGeneratorImage(businessDockerfileDtos[0], user)

	return result.Success()

}

func doGeneratorImage(businessDockerfileDto *businessDockerfile.BusinessDockerfileDto, user *user.UserDto) {

	var (
		dockerfile        = businessDockerfileDto.Dockerfile
		tenantId          = businessDockerfileDto.TenantId
		id                = businessDockerfileDto.Id
		businessImagesDao dao.BusinessImagesDao
		f                 *os.File
		err               error
		cmd               *exec.Cmd
	)

	dest := filepath.Join(config.WorkSpace, tenantId+"/"+id)

	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}

	dest += "/Dockerfile"

	if !utils.IsFile(dest  ) {
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

	imageName := imageRepository + businessDockerfileDto.Name + ":" + businessDockerfileDto.Version

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
	businessImagesDto.TenantId = user.TenantId
	businessImagesDto.CreateUserId = user.UserId
	businessImagesDto.Id = seq.Generator()
	businessImagesDto.Version = "V" + date.GetNowAString()
	businessImagesDto.ImagesType = businessImages.IMAGES_TYPE_DOCKER
	businessImagesDto.ImagesFlag = businessImages.IMAGES_FLAG_CUSTOM
	businessImagesDto.TypeUrl = "docker pull " + imageName
	businessImagesDto.Name = businessDockerfileDto.Name

	err = businessImagesDao.SaveBusinessImages(businessImagesDto)
	if err != nil {
		fmt.Println("保存镜像失败" + err.Error())
	}

}
