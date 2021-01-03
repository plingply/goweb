/*
 * @Author: 彭林
 * @Date: 2020-12-23 18:31:34
 * @LastEditTime: 2020-12-23 19:15:12
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/upload_file/upload_file.go
 */
package upload_file

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

var endpoint = g.Cfg().GetString("OSS.endpoint")
var accessKeyId = g.Cfg().GetString("OSS.accessKeyId")
var accessKeySecret = g.Cfg().GetString("OSS.accessKeySecret")
var bucketName = g.Cfg().GetString("OSS.bucketName")
var fileEndpoint = g.Cfg().GetString("OSS.fileEndpoint")

func OSSUpLoad(file *ghttp.UploadFile) (name string, err error) {

	u1 := cast.ToString(uuid.NewV4())
	prefix := "." + strings.Split(file.Filename, ".")[1]
	dst := "./" + u1 + prefix

	if err := SaveUploadedFile(file, dst); err != nil {
		return "", err
	}

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// 读取本地文件。
	fd, err := os.Open(dst)
	if err != nil {
		return "", err
	}

	defer fd.Close()
	url := u1 + prefix
	// 上传文件流。
	err = bucket.PutObject(url, fd)
	if err != nil {
		return "", err
	}

	err = os.Remove(dst)

	if err != nil {
		os.Exit(-1)
		return "", errors.New("上传成功，临时文件删除失败")
	}
	fileURL := fileEndpoint + "/" + url
	return fileURL, nil
}

func SaveUploadedFile(file *ghttp.UploadFile, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
