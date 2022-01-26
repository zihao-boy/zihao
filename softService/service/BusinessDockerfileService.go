package service

import (
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/softService/dao"
)

type BusinessDockerfileService struct {
	businessDockerfileDao dao.BusinessDockerfileDao
	businessPackageDao dao.BusinessPackageDao
}

/**
查询 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) GetBusinessDockerfileAll(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) ([]*businessDockerfile.BusinessDockerfileDto, error) {
	var (
		err                    error
		businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	)

	businessDockerfileDtos, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfiles(businessDockerfileDto)
	if err != nil {
		return nil, err
	}

	return businessDockerfileDtos, nil

}

/**
查询 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) GetBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                    error
		page                   int64
		row                    int64
		total                  int64
		businessDockerfileDto  = businessDockerfile.BusinessDockerfileDto{}
		businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	businessDockerfileDto.Row = row * page

	businessDockerfileDto.Page = (page - 1) * row
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.Name = ctx.URLParam("name")
	businessDockerfileDto.Version = ctx.URLParam("version")
	businessDockerfileDto.Id = ctx.URLParam("id")

	total, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfileCount(businessDockerfileDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	businessDockerfileDtos, err = businessDockerfileService.businessDockerfileDao.GetBusinessDockerfiles(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDtos, total, row)

}

/**
保存 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) SaveBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)
	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	businessDockerfileDto.Id = seq.Generator()
	businessDockerfileDto.Version = "V" + date.GetNowAString()

	//save log

	logPath := filepath.Join(config.WorkSpace, "businessPackage/"+user.TenantId, businessDockerfileDto.Id+".log")

	businessDockerfileDto.LogPath = logPath

	err = businessDockerfileService.businessDockerfileDao.SaveBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}


func (businessDockerfileService *BusinessDockerfileService) SaveBusinessDockerfileCommon(ctx iris.Context) interface{} {
	var (
		err                   error
		businessDockerfileCommonDto businessDockerfile.BusinessDockerfileCommonDto
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
		businessPackageDto businessPackage.BusinessPackageDto
	)
	if err = ctx.ReadJSON(&businessDockerfileCommonDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	//save shell file
	businessPackageDto.Id = seq.Generator()

	curDest := filepath.Join("businessPackage", user.TenantId, businessPackageDto.Id)
	dest := filepath.Join(config.WorkSpace, curDest)
	if !utils.IsDir(dest) {
		utils.CreateDir(dest)
	}
	dest = filepath.Join(dest, "start_"+businessDockerfileCommonDto.Name+".sh")
	if utils.IsFile(dest) {
		os.Remove(dest)
	}
	f, _ := os.Create(dest)
	defer func() {
		f.Close()
	}()
	f.WriteString(businessDockerfileCommonDto.ShellContext)
	businessPackageDto.TenantId = user.TenantId
	businessPackageDto.CreateUserId = user.UserId
	//businessPackageDto.Path = filepath.Join(curDest, fileHeader.Filename)
	businessPackageDto.Path = filepath.Join(businessPackageDto.Id, "start_"+businessDockerfileCommonDto.Name+".sh")
	businessPackageDto.Varsion = "V" + date.GetNowAString()
	businessPackageDto.Name = "start_"+businessDockerfileCommonDto.Name+".sh"

	err = businessDockerfileService.businessPackageDao.SaveBusinessPackage(businessPackageDto)
	if err != nil {
		return result.Error(err.Error())
	}
	var dockerfile = "# 指定源于一个基础镜像\n"
	if businessDockerfileCommonDto.DeployType == appService.DeployTypeJava {
		dockerfile += "FROM registry.cn-beijing.aliyuncs.com/sxd/ubuntu-java8:1.0\n"
	} else {
		dockerfile += "FROM centos:centos7\n"
	}
	dockerfile = dockerfile +
		"# 维护者/拥有者\nMAINTAINER xxx <xxx@xx.com>\n" +
		"# 从宿主机上传文件 ，这里上传一个脚本，\n" +
		"# 具体脚本可以去业务包上传后复制路径\n" +
		"ADD "+businessDockerfileCommonDto.Path+" /root/\n" +
		"# 从宿主机上传文件 ，这里上传一个业务文件，\n" +
		"# 具体脚本可以去业务包上传后复制路径\n" +
		"ADD "+businessPackageDto.Path+" /root/\n" +
		"# 容器内执行相应指令\n" +
		"RUN chmod u+x /root/start_"+businessDockerfileCommonDto.Name+".sh\n" +
		"# 运行命令\n" +
		"# CMD <command>   or CMD [<command>]\n" +
		"# 整个Dockerfile 中只能有一个,多个会被覆盖的\n" +
		"CMD [\"/root/start_"+businessDockerfileCommonDto.Name+".sh\"]\n"

	businessDockerfileDto.Name = businessDockerfileCommonDto.Name
	businessDockerfileDto.Dockerfile = dockerfile
	businessDockerfileDto.TenantId = user.TenantId
	businessDockerfileDto.CreateUserId = user.UserId
	businessDockerfileDto.Id = seq.Generator()
	businessDockerfileDto.Version = "V" + date.GetNowAString()

	//save log

	logPath := filepath.Join(config.WorkSpace, "businessPackage/"+user.TenantId, businessDockerfileDto.Id+".log")

	businessDockerfileDto.LogPath = logPath

	err = businessDockerfileService.businessDockerfileDao.SaveBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)
}
/**
修改 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) UpdateBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessDockerfileService.businessDockerfileDao.UpdateBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}

/**
删除 系统信息
*/
func (businessDockerfileService *BusinessDockerfileService) DeleteBusinessDockerfiles(ctx iris.Context) result.ResultDto {
	var (
		err                   error
		businessDockerfileDto businessDockerfile.BusinessDockerfileDto
	)

	if err = ctx.ReadJSON(&businessDockerfileDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = businessDockerfileService.businessDockerfileDao.DeleteBusinessDockerfile(businessDockerfileDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(businessDockerfileDto)

}


