package shell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/common/httpReq"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/zihao-boy/zihao/entity/dto/host"
	"golang.org/x/crypto/ssh"
)

func ExecLocalShell(jobShell string) (string,error)  {
	cmd := exec.Command("bash", "-c", jobShell)
	fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	paramOut, err := cmd.Output()
	fmt.Print("cmd 结果",string(paramOut))
	return string(paramOut),err
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
		return resultDto, err
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
