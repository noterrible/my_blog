package initialization

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"my_blog/global"
	"os"
)

func InitBucket() {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(global.Config.Image.UploadDomain, global.Config.Image.AccessKeyId, global.Config.Image.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := global.Config.Image.BucketName
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	global.Bucket = bucket
}
