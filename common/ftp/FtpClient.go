package ftp

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"io"
	"os"
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

	// TLS client authentication
	config := tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
	}

	if err = ftp.AuthTLS(&config); err != nil {
		return err
	}

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.Path, "/") {
		path = resourcesFtpDto.Path
	} else {
		path = "/" + resourcesFtpDto.Path
	}

	if err = ftp.Cwd(path); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return err
	}

	defer file.Close()

	if err := ftp.Stor(path, file); err != nil {
		return err
	}
	return nil
}

// upload file

func DownloadFile(filePath string, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	// TLS client authentication
	config := tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
	}

	if err = ftp.AuthTLS(&config); err != nil {
		return err
	}

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.Path, "/") {
		path = resourcesFtpDto.Path
	} else {
		path = "/" + resourcesFtpDto.Path
	}

	if err = ftp.Cwd(path); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Open("/tmp/test.txt"); err != nil {
		return err
	}

	defer file.Close()

	err = ftp.Walk(path, func(path string, info os.FileMode, err error) error {
		_, err = ftp.Retr(path, func(r io.Reader) error {
			var hasher = sha256.New()
			if _, err = io.Copy(hasher, r); err != nil {
				return err
			}

			hash := fmt.Sprintf("%s %x", path, hex.EncodeToString(hasher.Sum(nil)))
			fmt.Println(hash)

			return err
		})

		return nil
	})
	return nil
}
