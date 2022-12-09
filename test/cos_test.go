package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUplodeByFilepath(t *testing.T) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://1-1255907395.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: define.TencentSecretID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/1.2.png"

	// _, _, err := client.Object.Upload(
	// 	context.Background(), key, "./img/1.png", nil,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	f, err := os.ReadFile("./img/1.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(context.Background(), key, bytes.NewReader(f), nil)
	if err != nil {
		panic(err)
	}
}

// 分片上传初始化
func TestTesInitPartUpload(t *testing.T) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://1-1255907395.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: define.TencentSecretID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: define.TencentSecretKey,
		},
	})
	//key := "cloud-disk/fhg.jpg"
	key := "cloud-disk/pubg.mp4"
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID //1670573579a27125b5470c1e3ca8baf3c0359a13686a32b3cc8d98a425fb07965184916ea0
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {

	u, _ := url.Parse("https://1-1255907395.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "cloud-disk/pubg.mp4"
	UploadID := "1670573579a27125b5470c1e3ca8baf3c0359a13686a32b3cc8d98a425fb07965184916ea0"
	//f, err := os.ReadFile("temp/0.chunk") // md5 : 237f3f68be96267622a03d210bd096a7
	//f, err := os.ReadFile("temp/1.chunk") // md5 : 1c92ab0324b2f17622dbf8904c5aa982
	f, err := os.ReadFile("temp/2.chunk") // md5 : d18fd381c29f52a06590959f6fe6f3e8

	if err != nil {
		t.Fatal(err)
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 3, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://1-1255907395.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	//key := "cloud-disk/fhg.jpg"
	key := "cloud-disk/pubg.mp4"
	UploadID := "1670573579a27125b5470c1e3ca8baf3c0359a13686a32b3cc8d98a425fb07965184916ea0"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "237f3f68be96267622a03d210bd096a7"},
		cos.Object{
			PartNumber: 2, ETag: "1c92ab0324b2f17622dbf8904c5aa982"},
		cos.Object{
			PartNumber: 3, ETag: "d18fd381c29f52a06590959f6fe6f3e8"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
