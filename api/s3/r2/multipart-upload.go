package r2

import (
	"context"
	"github.com/akazwz/gin-api/global"
	"os"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// CreateMultipartUpload 创建多块上传
func CreateMultipartUpload(c *gin.Context) {
	key := c.Param("key")

	// 创建 multipart upload
	multiUpload, err := global.R2C.CreateMultipartUpload(context.TODO(), &s3.CreateMultipartUploadInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(key),
	})

	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	response.Created(api.CodeCommonSuccess, gin.H{
		"upload_id": multiUpload.UploadId,
		"key":       key,
	}, "success", c)
}

// UploadPart 上传部分
func UploadPart(c *gin.Context) {
	var partUploadParams request.MultipartUpload
	err := c.ShouldBind(&partUploadParams)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	uploadId := partUploadParams.UploadId
	if len(uploadId) < 1 {
		response.BadRequest(api.CodeCommonFailed, nil, "upload id error", c)
		return
	}

	fh := partUploadParams.File

	// 判断文件是否为空
	if fh.Size <= 0 {
		response.BadRequest(api.CodeCommonFailed, nil, "file empty", c)
		return
	}

	// 打开文件， 获取 File
	file, err := fh.Open()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	_, err = global.R2C.UploadPart(context.TODO(), &s3.UploadPartInput{
		Bucket:     aws.String(os.Getenv("R2_BUCKET_NAME")),
		UploadId:   aws.String(uploadId),
		Key:        aws.String(partUploadParams.Key),
		PartNumber: partUploadParams.PartNumber,
		Body:       file,
	})
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Created(api.CodeCommonSuccess, nil, "success", c)
}

// CompleteMultipartUpload 完成多块上传
func CompleteMultipartUpload(c *gin.Context) {
	_, err := global.R2C.CompleteMultipartUpload(context.TODO(), &s3.CompleteMultipartUploadInput{
		Bucket:   nil,
		Key:      nil,
		UploadId: nil,
	})
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}

// AbortMultipartUpload 放弃多块上传
func AbortMultipartUpload(c *gin.Context) {
	_, err := global.R2C.AbortMultipartUpload(context.TODO(), &s3.AbortMultipartUploadInput{
		Bucket:   nil,
		Key:      nil,
		UploadId: nil,
	})
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}

// ListParts 列出分块
func ListParts(c *gin.Context) {
	_, err := global.R2C.ListParts(context.TODO(), &s3.ListPartsInput{
		Bucket:   nil,
		Key:      nil,
		UploadId: nil,
	})
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}

func ListMultipartUploads(c *gin.Context) {
	_, err := global.R2C.ListMultipartUploads(context.TODO(), &s3.ListMultipartUploadsInput{
		Bucket: nil,
	})
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}
