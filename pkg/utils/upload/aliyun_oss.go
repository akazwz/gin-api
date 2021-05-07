package upload

import (
	"errors"
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

func OSSUploadFile(file *multipart.FileHeader, objectKey string) error {
	bucket, err := NewBucket()
	if err != nil {
		return errors.New("NewBucket() Failed:" + err.Error())
	}

	fileReader, err := file.Open()
	if err != nil {
		return errors.New("file.Open() Failed:" + err.Error())
	}

	err = bucket.PutObject(objectKey, fileReader)
	if err != nil {
		return errors.New("PutObject() Failed:" + err.Error())
	}

	return nil
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
