package middleware

import (
	"mime/multipart"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func FileSizeLimit(size int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileUp := request.UploadFile{}
		err := c.ShouldBind(&fileUp)
		if err != nil {
			response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
			return
		}

		fileHeader := fileUp.File

		// 判断文件是否为空
		if fileHeader.Size <= 0 {
			response.BadRequest(api.CodeCommonFailed, nil, "文件为空", c)
			return
		}

		if fileHeader.Size > size {
			response.BadRequest(api.CodeCommonFailed, nil, "文件大小超出限制", c)
			return
		}

		c.Set("fh", fileHeader)
		c.Next()
	}
}

// FileMimeTypeLimit 文件类型限制
func FileMimeTypeLimit(mimetypes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 file header
		fileHeaderAny, _ := c.Get("fh")
		fileHeader := fileHeaderAny.(*multipart.FileHeader)

		// 打开文件， 获取 File
		file, err := fileHeader.Open()
		if err != nil {
			response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
			return
		}
		// 获取 file mime type
		mtype, err := mimetype.DetectReader(file)
		if err != nil {
			response.BadRequest(api.CodeCommonFailed, nil, "获取文件 mime type失败", c)
			return
		}
		// mime types 为空不限制
		if len(mimetypes) < 1 {
			c.Next()
			return
		}
		// 判断 mime type
		includes := utils.ArrayIncludes(mimetypes, mtype.String())
		if !includes {
			response.BadRequest(api.CodeCommonFailed, nil, "文件 mime type错误", c)
			return
		}
		c.Set("fh", fileHeader)
		c.Next()
	}
}
