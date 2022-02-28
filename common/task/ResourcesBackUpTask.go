package task

import (
	"bufio"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/business/dao/resourcesDbDao"
	"github.com/zihao-boy/zihao/business/dao/resourcesFtpDao"
	"github.com/zihao-boy/zihao/business/dao/resourcesOssDao"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/db/dbFactory"
	"github.com/zihao-boy/zihao/common/ftp"
	"github.com/zihao-boy/zihao/common/oss"
	"github.com/zihao-boy/zihao/common/shell"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/dbLink"
	"github.com/zihao-boy/zihao/entity/dto/host"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"io"
	"os"
	"path"
)

type ResourcesBackUpTask struct {
	ResourcesBackUpDto *resources.ResourcesBackUpDto
	resourcesDbDao     resourcesDbDao.ResourcesDbDao
	resourcesFtpDao    resourcesFtpDao.ResourcesFtpDao
	resourcesOssDao    resourcesOssDao.ResourcesOssDao

	hostDao hostDao.HostDao
}

func (h ResourcesBackUpTask) Run() {
	var (
		resourcesBackUpDto *resources.ResourcesBackUpDto
		//err   error
	)
	resourcesBackUpDto = h.ResourcesBackUpDto

	typeCd := resourcesBackUpDto.TypeCd

	if typeCd == resources.Back_up_Type_Cd_db {
		h.backUpDb(resourcesBackUpDto)
	} else {
		h.backUpDir(resourcesBackUpDto)
	}

}

// 备份数据库

func (h ResourcesBackUpTask) backUpDb(dto *resources.ResourcesBackUpDto) {

	resourcesDbDto := resources.ResourcesDbDto{
		DbId: dto.SrcId,
	}
	resourcesDbs, _ := h.resourcesDbDao.GetResourcesDbs(resourcesDbDto)

	if resourcesDbs == nil || len(resourcesDbs) < 1 {
		return
	}

	dblinkDto := dbLink.DbLinkDto{
		Id:       resourcesDbs[0].DbId,
		Name:     resourcesDbs[0].Name,
		Ip:       resourcesDbs[0].Ip,
		Port:     resourcesDbs[0].Port,
		Username: resourcesDbs[0].Username,
		Password: resourcesDbs[0].Password,
		DbName:   resourcesDbs[0].DbName,
		TenantId: resourcesDbs[0].TenantId,
	}

	workDir := path.Join(config.WorkSpace, dblinkDto.TenantId,"backUp")

	if !utils.IsDir(workDir){
		os.MkdirAll(workDir,os.ModePerm)
	}


	dbSqlDto := dbLink.DbSqlDto{
		FileName: path.Join(workDir,resourcesDbs[0].DbName + date.GetNowAString() + ".sql"),
	}

	// execute sql
	data := dbFactory.ExportSqlFile(dblinkDto, dbSqlDto,dto.SrcObject)

	if data.Code != result.CODE_SUCCESS {
		return
	}

	tmpFilePath :=  dbSqlDto.FileName

	if !utils.IsFile(tmpFilePath) {
		return
	}

	if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_ftp {
		h.saveToFtp(tmpFilePath, dto)
	} else if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_oss {
		h.saveToOss(tmpFilePath, dto)
	} else {
		h.saveToDb(tmpFilePath, dto)
	}

	//delete tmp sql file
	os.Remove(tmpFilePath)

}

// 备份文件夹

func (h ResourcesBackUpTask) backUpDir(dto *resources.ResourcesBackUpDto) {

	hostDto := host.HostDto{
		HostId: dto.SrcId,
	}
	hosts, _ := h.hostDao.GetHosts(hostDto)
	if hosts == nil || len(hosts) < 1 {
		return
	}
	hostDto = *hosts[0]
	hostDto.FileName = dto.SrcObject

	tmpFilePath := path.Join(config.WorkSpace, dto.TenantId, "backUp", dto.SrcObject+date.GetNowDateString()+".zip")

	shell.ExecDownloadFileAndSave(hostDto, tmpFilePath)

	if !utils.IsFile(tmpFilePath) {
		return
	}

	if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_ftp {
		h.saveToFtp(tmpFilePath, dto)
	} else if dto.TargetTypeCd == resources.Back_up_Target_Type_Cd_oss {
		h.saveToOss(tmpFilePath, dto)
	}

	//delete tmp sql file
	os.Remove(tmpFilePath)
}

func (h ResourcesBackUpTask) saveToFtp(filePath string, dto *resources.ResourcesBackUpDto) {

	resourcesFtpDto := resources.ResourcesFtpDto{
		FtpId: dto.TargetId,
	}

	resourcesFtpDtos, _ := h.resourcesFtpDao.GetResourcesFtps(resourcesFtpDto)

	if resourcesFtpDtos == nil || len(resourcesFtpDtos) < 1 {
		return
	}

	ftp.UploadFile(filePath, *resourcesFtpDtos[0])

}

func (h ResourcesBackUpTask) saveToOss(filePath string, dto *resources.ResourcesBackUpDto) {
	resourcesOssDto := resources.ResourcesOssDto{
		OssId: dto.TargetId,
	}

	resourcesOssDtos, _ := h.resourcesOssDao.GetResourcesOsss(resourcesOssDto)

	if resourcesOssDtos == nil || len(resourcesOssDtos) < 1 {
		return
	}

	oss.SaveALiOss(filePath, *resourcesOssDtos[0])
}

// save to db
func (h ResourcesBackUpTask) saveToDb(filePath string, dto *resources.ResourcesBackUpDto) {

	if dto.TypeCd != resources.Back_up_Type_Cd_db {
		return
	}

	resourcesDbDto := resources.ResourcesDbDto{
		DbId: dto.TargetId,
	}
	resourcesDbs, _ := h.resourcesDbDao.GetResourcesDbs(resourcesDbDto)

	if resourcesDbs == nil || len(resourcesDbs) < 1 {
		return
	}

	dblinkDto := dbLink.DbLinkDto{
		Id:       resourcesDbs[0].DbId,
		Name:     resourcesDbs[0].Name,
		Ip:       resourcesDbs[0].Ip,
		Port:     resourcesDbs[0].Port,
		Username: resourcesDbs[0].Username,
		Password: resourcesDbs[0].Password,
		DbName:   resourcesDbs[0].DbName,
		TenantId: resourcesDbs[0].TenantId,
	}

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}

	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString(';')
		if err != nil || io.EOF == err {
			break
		}

		if utils.IsEmpty(line) {
			continue
		}

		dbSqlDto := dbLink.DbSqlDto{
			Sql: line,
		}
		// execute sql
		dbFactory.ExecSql(dblinkDto, dbSqlDto)
	}
}
