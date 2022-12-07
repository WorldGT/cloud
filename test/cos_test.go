package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
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
