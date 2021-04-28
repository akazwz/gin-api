package upload

import (
	"github.com/akazwz/go-gin-demo/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type AliOSS struct{}

func (*AliOSS) UploadFile(file *multipart.File) (string, string, error) {
	return "", "", nil
}

func NewBucket() (*oss.Bucket, error) {
	endpoint := global.CFG.Server.AliOSS.Endpoint
	AccessKeyId := global.CFG.Server.AliOSS.AccessKeyId
	AccessKeySecret := global.CFG.Server.AliOSS.AccessKeySecret
	BucketName := global.CFG.Server.AliOSS.BucketName
	client, err := oss.New(endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
