package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"os"
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
func ListALiOss(resourcesOssDto resources.ResourcesOssDto) error {
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
	// 列举文件。
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return err
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println("Bucket: ", object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return nil
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
	// 删除文件。
	err = bucket.DeleteObject(resourcesOssDto.Path)
	return nil
}
