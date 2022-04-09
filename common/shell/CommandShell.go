package shell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/context"
	hostDao "github.com/zihao-boy/zihao/assets/dao"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/innerNet"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/zihao-boy/zihao/entity/dto/host"
	"golang.org/x/crypto/ssh"
)

func ExecLocalShell(jobShell string) (string, error) {
	cmd := exec.Command("bash", "-c", jobShell)
	fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	paramOut, err := cmd.Output()
	fmt.Print("cmd 结果", string(paramOut))
	return string(paramOut), err
}

func ExecShell(host host.HostDto, cmd string) error {
	client, err := ssh.Dial("tcp", host.Ip, &ssh.ClientConfig{
		User:            host.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(host.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		fmt.Print("链接主机失败", err)
		return err
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()
	// 使用内存
	fmt.Print("主机执行指令", cmd)
	session.Output(cmd)
	return nil
}

func ExecCommonShell(host host.HostDto, cmd string) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	data["hostId"] = host.HostId
	data["shell"] = cmd

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/execShell", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecListFiles(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/listFiles", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecRemoveFile(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/removeFile", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecNewFile(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/newFile", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecRenameFile(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/renameFile", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecListFileContext(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/listFileContext", data, nil)
	if err != nil {
		return result.Error(err.Error()), err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecEditFile(host host.HostDto) (result.ResultDto, error) {
	data := make(map[string]interface{})
	ip := host.Ip
	var resultDto result.ResultDto

	appServiceDtoData, _ := json.Marshal(&host)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}

	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

	resp, err := httpReq.Post("http://"+ip+"/app/slave/editFile", data, nil)
	if err != nil {
		return resultDto, err
	}

	json.Unmarshal([]byte(resp), &resultDto)

	return resultDto, nil

}

func ExecUploadFile(host host.HostDto, file multipart.File, fileHeader *multipart.FileHeader) (result.ResultDto, error) {
	var resultDto result.ResultDto
	ip := host.Ip
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	bodyWriter.WriteField("curPath", host.CurPath)
	fileWriter, _ := bodyWriter.CreateFormFile("uploadFile", fileHeader.Filename)
	io.Copy(fileWriter, file) //将 客户端文件 复制给 用于传输的 fileWriter

	contentType := bodyWriter.FormDataContentType() //contentType
	bodyWriter.Close()
	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}
	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))
	resp, _ := http.Post("http://"+ip+"/app/slave/uploadFile", contentType, bodyBuffer)

	defer resp.Body.Close()
	// 4、结果读取
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(data), &resultDto)
	return resultDto, nil
}

func ExecDownloadFile(host host.HostDto, resWriter context.ResponseWriter) {
	ip := host.Ip
	appServiceDtoData, _ := json.Marshal(&host)
	body := bytes.NewReader(appServiceDtoData)
	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}
	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))
	resp, _ := http.Post("http://"+ip+"/app/slave/downloadFile", "application/json", body)

	defer func() {
		_ = resp.Body.Close()
	}()
	// 4、结果读取
	//data, _ := ioutil.ReadAll(resp.Body)
	resWriter.Header().Set("Content-Type", "application/octet-stream")
	resWriter.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(resWriter, resp.Body)

	resWriter.Flush()
	//return data, resp.Header
}

func ExecDownloadFileAndSave(host host.HostDto, path string) error {
	ip := host.Ip
	appServiceDtoData, _ := json.Marshal(&host)
	body := bytes.NewReader(appServiceDtoData)
	if strings.Contains(ip, ":") {
		ip = ip[0:strings.Index(ip, ":")]
	}
	ip += (":" + strconv.FormatInt(int64(config.Slave), 10))
	resp, _ := http.Post("http://"+ip+"/app/slave/downloadDir", "application/json", body)

	defer func() {
		_ = resp.Body.Close()
	}()

	if utils.IsFile(path) {
		os.Remove(path)
	}

	file, err := os.Create(path)
	defer func() {
		file.Close()
	}()

	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}
		file.WriteString(string(buf[:n]))
	}

	return nil
	//return data, resp.Header
}

func ExecStartWaf(waf waf.SlaveWafDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&waf)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, wafHost := range waf.Waf.WafHosts {
		tmpHostDto := host.HostDto{
			HostId: wafHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/startWaf", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}


func ExecStopWaf(waf waf.SlaveWafDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&waf)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, wafHost := range waf.Waf.WafHosts {
		tmpHostDto := host.HostDto{
			HostId: wafHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/stopWaf", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}

func ExecRefreshWafConfig(waf waf.SlaveWafDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&waf)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, wafHost := range waf.Waf.WafHosts {
		tmpHostDto := host.HostDto{
			HostId: wafHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/refreshWafConfig", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}





func ExecStartInnerNet(innerNet innerNet.SlaveInnerNetDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&innerNet)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, innerNetHost := range innerNet.InnerNet.InnerNetHosts {
		tmpHostDto := host.HostDto{
			HostId: innerNetHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/startInnerNet", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}


func ExecStopInnerNet(innerNet innerNet.SlaveInnerNetDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&innerNet)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, innerNetHost := range innerNet.InnerNet.InnerNetHosts {
		tmpHostDto := host.HostDto{
			HostId: innerNetHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/stopInnerNet", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}

func ExecRefreshInnerNetConfig(innerNet innerNet.SlaveInnerNetDataDto) (result.ResultDto, error) {
	// query hostInfo

	var (
		hostDao hostDao.HostDao
		resultDto result.ResultDto
	)
	data := make(map[string]interface{})
	appServiceDtoData, _ := json.Marshal(&innerNet)
	json.Unmarshal([]byte(appServiceDtoData), &data)

	for _, innerNetHost := range innerNet.InnerNet.InnerNetHosts {
		tmpHostDto := host.HostDto{
			HostId: innerNetHost.HostId,
		}
		hostDtos, _ := hostDao.GetHosts(tmpHostDto)
		if len(hostDtos) < 1 {
			continue
		}
		ip := hostDtos[0].Ip
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ip += (":" + strconv.FormatInt(int64(config.Slave), 10))

		resp, err := httpReq.Post("http://"+ip+"/app/slave/refreshInnerNetConfig", data, nil)
		if err != nil {
			return resultDto, err
		}
		json.Unmarshal([]byte(resp), &resultDto)

		if resultDto.Code != result.CODE_SUCCESS{
			return resultDto, nil
		}
	}
	return resultDto, nil
}
