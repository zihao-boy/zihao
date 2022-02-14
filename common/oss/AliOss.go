package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/entity/dto/ls"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"io"
	"os"
	"strings"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// upload file
func SaveALiOss(filePath string, resourcesOssDto resources.ResourcesOssDto) error {
	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return err
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(resourcesOssDto.Path, filePath)
	if err != nil {
		return err
	}
	return nil
}

// upload file
func SaveALiOssByReader(reader io.Reader, resourcesOssDto resources.ResourcesOssDto) error {
	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return err
	}
	marker := resourcesOssDto.CurPath

	if strings.HasPrefix(marker, "/") {
		marker = marker[1:]
	}

	// 上传文件。
	err = bucket.PutObject(marker, reader)
	if err != nil {
		return err
	}
	return nil
}

// download file
func DownloadALiOssByReader(resWriter context.ResponseWriter, resourcesOssDto resources.ResourcesOssDto) error {
	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return err
	}

	marker := resourcesOssDto.CurPath

	if strings.HasPrefix(marker, "/") {
		marker = marker[1:]
	}
	// 上传文件。
	outReader, err := bucket.GetObject(marker)
	if err != nil {
		return err
	}
	defer outReader.Close()
	resWriter.Header().Set("Content-Type", "application/octet-stream")
	//resWriter.Header().Set("Content-Length", outReader.Get("Content-Length"))
	io.Copy(resWriter, outReader)

	resWriter.Flush()
	return nil
}

// download file
func DownloadALiOss(filePath string, resourcesOssDto resources.ResourcesOssDto) error {
	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return err
	}
	// 上传文件。
	err = bucket.GetObjectToFile(resourcesOssDto.Path, filePath)
	if err != nil {
		return err
	}
	return nil
}

// download file
func ListALiOss(resourcesOssDto resources.ResourcesOssDto) result.ResultDto {

	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return result.Error(err.Error())
	}
	// 列举文件。
	marker := resourcesOssDto.CurPath

	if strings.HasPrefix(marker, "/") {
		marker = marker[1:]
	}

	lsRes, err := bucket.ListObjectsV2(oss.Prefix(marker), oss.Delimiter("/"), oss.MaxKeys(100))
	if err != nil {
		return result.Error(err.Error())
	}
	var lss = make([]ls.LsDto, 0)
	for _, fil := range lsRes.CommonPrefixes {
		lsDto := ls.LsDto{
			GroupName: "d",
			Name:      strings.Replace(fil, marker, "", 1),
			Privilege: "",
			Size:      0,
		}
		lss = append(lss, lsDto)
	}

	for _, object := range lsRes.Objects {
		lsDto := ls.LsDto{
			GroupName:    "-",
			Name:         strings.Replace(object.Key, marker, "", 1),
			Privilege:    "",
			Size:         object.Size,
			LastModified: date.GetTimeString(object.LastModified),
		}
		lss = append(lss, lsDto)
	}

	return result.SuccessData(lss)
}

// delete oss
func DeleteALiOss(resourcesOssDto resources.ResourcesOssDto) error {
	// 创建OSSClient实例。
	client, err := oss.New(resourcesOssDto.Endpoint, resourcesOssDto.AccessKeyId, resourcesOssDto.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(resourcesOssDto.Bucket)
	if err != nil {
		return err
	}
	marker := resourcesOssDto.CurPath
	// 删除文件。
	if strings.HasPrefix(marker, "/") {
		marker = marker[1:]
	}
	// 删除文件。
	//if strings.HasSuffix(marker, "/") {
	//	marker = marker[:len(marker)-1]
	//}
	// 上传文件。
	if resourcesOssDto.FileGroupName == "-"{
		err = bucket.DeleteObject(marker)
	}else{
		lsRes, err := bucket.ListObjectsV2(oss.Prefix(marker), oss.Delimiter(""), oss.MaxKeys(100))
		if err != nil {
			return err
		}
		var keys []string
		for _,object:= range lsRes.Objects{
			keys = append(keys,object.Key)
		}

		if len(keys) > 0{
			bucket.DeleteObjects(keys)
		}
		if strings.HasSuffix(marker, "/") {
			marker = marker[:len(marker)-1]
		}
		err = bucket.DeleteObject(marker)
	}
	if err != nil {
		return err
	}
	return nil
}
