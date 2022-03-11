package ftp

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/entity/dto/ls"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/dutchcoders/goftp.v1"
)

// upload file

func UploadFile(filePath string, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var pathTmp string

	if strings.HasPrefix(resourcesFtpDto.Path, "/") {
		pathTmp = resourcesFtpDto.Path
	} else {
		pathTmp = "/" + resourcesFtpDto.Path
	}

	if err = ftp.Cwd(pathTmp); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return err
	}

	fileName :=  path.Base(filePath)

	defer file.Close()

	if err := ftp.Stor(path.Join(pathTmp,fileName), file); err != nil {
		return err
	}
	return nil
}

// upload file

func DownloadFile(resWriter context.ResponseWriter, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}
	resWriter.Header().Set("Content-Type", "application/octet-stream")

	//err = ftp.Walk(path, func(path string, info os.FileMode, err error) error {
	//	_, err = ftp.Retr(path, func(r io.Reader) error {
	//		io.Copy(resWriter, r)
	//		return err
	//	})
	//	return nil
	//})

	s, err := ftp.Retr(path, func(r io.Reader) error {
		io.Copy(resWriter, r)
		return err
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(s)
	resWriter.Flush()
	return nil
}

func ListFile(resourcesFtpDto resources.ResourcesFtpDto) result.ResultDto {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return result.Error(err.Error())
	}

	defer ftp.Close()

	// TLS client authentication
	//config := tls.Config{
	//	InsecureSkipVerify: true,
	//	ClientAuth:         tls.RequestClientCert,
	//}

	//if err = ftp.AuthTLS(&config); err != nil {
	//	return result.Error(err.Error())
	//}

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return result.Error(err.Error())
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}

	//if err = ftp.Cwd(path); err != nil {
	//	return result.Error(err.Error())
	//}

	dirs, err := ftp.List(path)

	if err != nil {
		return result.Error(err.Error())
	}
	var lss = make([]ls.LsDto, 0)
	fmt.Println(dirs)
	for _, fil := range dirs {
		if strings.Contains(fil, ";") {
			lsrs := strings.Split(fil, ";")
			if len(lsrs) == 4 {
				name := strings.Trim(lsrs[3], " ")
				name = strings.ReplaceAll(name, "\r", "")
				name = strings.ReplaceAll(name, "\n", "")
				lsDto := ls.LsDto{
					GroupName:    "d",
					Name:         name,
					Privilege:    strings.Split(lsrs[2], "=")[1],
					Size:         0,
					LastModified: strings.Split(lsrs[1], "=")[1],
				}
				lss = append(lss, lsDto)
			}
		}else{
			lsrs := getLsLines(fil)
			if len(lsrs) == 9 {
				name := strings.Trim(lsrs[len(lsrs)-1], " ")
				name = strings.ReplaceAll(name, "\r", "")
				name = strings.ReplaceAll(name, "\n", "")
				if strings.HasPrefix(lsrs[0], "d") {
					lsDto := ls.LsDto{
						GroupName:    "d",
						Name:         name,
						Privilege:    lsrs[0],
						Size:         0,
						LastModified: lsrs[5]+" "+lsrs[6]+" "+lsrs[7],
					}
					lss = append(lss, lsDto)
				}
			}

		}
	}

	for _, fil := range dirs {
		if strings.Contains(fil, ";") {
			lsrs := strings.Split(fil, ";")
			if len(lsrs) == 5 {
				size, _ := strconv.ParseInt(strings.Split(lsrs[1], "=")[1], 10, 64)
				name := strings.Trim(lsrs[4], " ")
				name = strings.ReplaceAll(name, "\r", "")
				name = strings.ReplaceAll(name, "\n", "")
				lsDto := ls.LsDto{
					GroupName:    "-",
					Name:         name,
					Privilege:    strings.Split(lsrs[3], "=")[1],
					Size:         size,
					LastModified: strings.Split(lsrs[2], "=")[1],
				}
				lss = append(lss, lsDto)
			}
		}else{
			lsrs := getLsLines(fil)
			if len(lsrs) == 9 {
				size, _ := strconv.ParseInt(lsrs[4], 10, 64)
				name := strings.Trim(lsrs[len(lsrs)-1], " ")
				name = strings.ReplaceAll(name, "\r", "")
				name = strings.ReplaceAll(name, "\n", "")
				if !strings.HasPrefix(lsrs[0], "d") {
					lsDto := ls.LsDto{
						GroupName:    "-",
						Name:         name,
						Privilege:    lsrs[0],
						Size:         size,
						LastModified: lsrs[5] + " " + lsrs[6] + " " + lsrs[7],
					}
					lss = append(lss, lsDto)
				}
			}
		}

	}
	return result.SuccessData(lss)
}

func getLsLines(line string) []string {
	lines := strings.Split(line," ")
	var newLines []string
	for _,l := range lines{
		if utils.IsEmpty(l){
			continue
		}
		newLines = append(newLines,l)
	}
	return newLines
}

func NewFileOrDir(resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}
	if resourcesFtpDto.FileGroupName == "-" {
		//err = ftp.(path)
	} else {
		err = ftp.Mkd(path)
	}

	if err != nil {
		return err
	}

	return nil
}

// rename file or dir

func RenameFileOrDir(resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	//var path string
	//
	//if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
	//	path = resourcesFtpDto.CurPath
	//} else {
	//	path = "/" + resourcesFtpDto.CurPath
	//}
	if resourcesFtpDto.FileGroupName == "-" {
		//err = ftp.(path)
		err = ftp.Rename(resourcesFtpDto.Name, resourcesFtpDto.NewName)
	} else {
		err = ftp.Rename(resourcesFtpDto.Name, resourcesFtpDto.NewName)
	}

	if err != nil {
		return err
	}

	return nil
}

func UploadFileByReader(f multipart.File, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}

	if strings.Contains(path, "/") {
		pos := strings.LastIndex(path, "/")
		dirStr := path[0:pos]
		dirs := strings.Split(dirStr, "/")
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i]
			if utils.IsEmpty(dir) || "/" == dir {
				continue
			}

			rs, _ := ftp.Stat(dir)
			if len(rs) > 0 {
				ftp.Cwd(dir)
				continue
			}
			ftp.Mkd(dir)
			ftp.Cwd(dir)
		}
	}

	if err := ftp.Stor(path, f); err != nil {
		return err
	}
	return nil
}

func DeleteFile(resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}
	if resourcesFtpDto.FileGroupName == "-" {
		err = ftp.Dele(path)
	} else {
		err = deleteDirAndFile(path, ftp)
	}

	if err != nil {
		return err
	}

	return nil
}

// delete dir and file

func deleteDirAndFile(dirPath string, ftp *goftp.FTP) error {

	dirs, err := ftp.List(dirPath)

	if err != nil {
		return err
	}
	for _, fil := range dirs {
		lsrs := strings.Split(fil, ";")
		if len(lsrs) == 4 {
			name := strings.Trim(lsrs[3], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			err = deleteDirAndFile(path.Join(dirPath, name), ftp)
			if err != nil {
				return err
			}
		}

		if len(lsrs) == 5 {
			name := strings.Trim(lsrs[4], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			err = ftp.Dele(path.Join(dirPath, name))
			if err != nil {
				return err
			}
		}
	}

	err = ftp.Rmd(dirPath)

	if err != nil {
		return err
	}
	return nil
}
