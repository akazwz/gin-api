package r2

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

// Upload 直传
func Upload(c *gin.Context) {
	// 获取 file header
	fileHeaderAny, _ := c.Get("fh")
	fileHeader := fileHeaderAny.(*multipart.FileHeader)

	// 获取文件 hash
	_, err := utils.HashFileByAlgo(fileHeader, "sha256")
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "hash 文件失败", c)
		return
	}

	// 打开文件， 获取 File
	file, err := fileHeader.Open()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	// 获取文件扩展名
	ext := utils.GetFileExtension(fileHeader.Filename)

	// 生成文件key
	key := utils.GenerateR2Key(ext)

	// 文件 url
	_ = fmt.Sprintf("%s/%s", os.Getenv("R2_HOST"), key)

	// 获取文件 mime-type
	mtype, err := mimetype.DetectReader(file)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	// 上传文件
	_, err = global.R2C.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:         aws.String(key),
		ContentType: aws.String(mtype.String()),
		Body:        file,
	})

	psClient := s3.NewPresignClient(global.R2C, func(options *s3.PresignOptions) {
		options.Expires = 3600
	})

	getObjectInput := &s3.GetObjectInput{
		Bucket: nil,
		Key:    nil,
	}

	u, err := psClient.PresignGetObject(context.Background(), getObjectInput)
	if err != nil {
		return
	}

	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	response.Created(api.CodeCommonSuccess, gin.H{
		"url": u.URL,
	}, "success", c)
}
