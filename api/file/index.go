package file

import (
	"fmt"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

// UploadFile 文件直传
func UploadFile(c *gin.Context) {
	fileUp := request.UploadFile{}

	err := c.ShouldBind(&fileUp)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}

	fileHeader := fileUp.File

	// 文件信息
	contentType := fileHeader.Header.Get("Content-Type")
	filename := fileHeader.Filename
	size := fileHeader.Size

	// 获取 sha256 hash
	sha256Hash, err := utils.HashFileByAlgo(fileHeader, "sha256")
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "读取文件失败", c)
		return
	}

	// 文件夹路径
	dst := fmt.Sprintf("public/file/%s", filename)
	// 保存文件
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "保存文件失败", c)
		return
	}

	response.Created(api.CodeCommonSuccess, gin.H{
		"content_type": contentType,
		"filename":     filename,
		"size":         size,
		"hash_sha256":  sha256Hash,
	}, "success", c)
}
