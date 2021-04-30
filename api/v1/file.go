package v1

import (
	"github.com/akazwz/go-gin-demo/model"
	"github.com/akazwz/go-gin-demo/model/request"
	"github.com/akazwz/go-gin-demo/model/response"
	"github.com/akazwz/go-gin-demo/pkg/util/upload"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// CreateFile
// @Summary Create File
// @Title Create File
// @Author zwz
// @Description create file
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Param token header string true "token"
// @Success 201 {object} response.Response
// @Failure 400,401 {object} response.Response
// @Router /file [post]
func CreateFile(c *gin.Context) {
	claims, _ := c.Get("claims")
	customClaims := claims.(*request.CustomClaims)
	_ = customClaims.UUID.String()
	file, err := c.FormFile("file")
	if err != nil {
		response.CommonFailed("Get File Error", CodeGetFileError, c)
		return
	}
	dirDate := time.Now().Format("2006-01-02")

	dir := "public/file/" + dirDate + "/"

	// 判断目录是否存在,不存在创建目录
	_, err = os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			response.CommonFailed("Create DIR Error", CodeUploadFileError, c)
			return
		}
	}

	fileNamePrefix := time.Now().Format("15-04-05")

	fileName := fileNamePrefix + "-" + file.Filename

	localFile := dir + fileName

	if err := c.SaveUploadedFile(file, localFile); err != nil {
		response.CommonFailed("Upload File Error", CodeUploadFileError, c)
		return
	}

	if err := upload.OSSUploadFile(file); err != nil {
		response.CommonFailed("Upload OSS Error", CodeUploadFileError, c)
		return
	}

	name := fileName
	size := file.Size
	fileData := model.File{
		URL:  localFile,
		MD5:  "",
		Name: name,
		Size: size,
		Type: "",
	}
	response.Created(fileData, "File Upload Success", c)
}
