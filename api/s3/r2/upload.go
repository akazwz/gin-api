package r2

import (
	"context"
	"fmt"
	"os"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// Upload 直传
func Upload(c *gin.Context) {
	var fileUp request.UploadFile
	// 绑定上传文件
	err := c.ShouldBind(&fileUp)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	// 判断文件是否为空
	if fileUp.File.Size <= 0 {
		response.BadRequest(api.CodeCommonFailed, nil, "file empty", c)
		return
	}

	// 获取文件 hash
	hash, err := utils.HashFileByAlgo(fileUp.File, "sha256")
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "hash file error", c)
		return
	}

	// 打开文件， 获取 File
	file, err := fileUp.File.Open()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	// 获取文件 meta data
	_, err = global.R2C.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(hash),
	})

	url := fmt.Sprintf("%s/%s", os.Getenv("R2_HOST"), hash)

	// 文件已经存在
	if err == nil {
		response.Created(api.CodeCommonSuccess, gin.H{
			"url": url,
		}, "object already exists", c)
		return
	}

	// 上传文件
	_, err = global.R2C.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(hash),
		Body:   file,
	})

	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	response.Created(api.CodeCommonSuccess, gin.H{
		"url": url,
	}, "success", c)
}
