package classify

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/grpc/nsfw"
	"github.com/akazwz/gin-api/model/response"
	"github.com/gin-gonic/gin"
)

func ImageFile(c *gin.Context) {
	// 获取 file header
	fileHeaderAny, _ := c.Get("fh")
	fileHeader := fileHeaderAny.(*multipart.FileHeader)

	// 打开文件， 获取 File
	file, err := fileHeader.Open()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "读取文件失败",
		})
		return
	}

	err, result := nsfw.ClassifyImage(b)
	if err != nil {
		log.Println(err)
		response.BadRequest(api.CodeCommonFailed, nil, "分类检测失败", c)
		return
	}
	response.Ok(api.CodeCommonSuccess, result, "success", c)
}
