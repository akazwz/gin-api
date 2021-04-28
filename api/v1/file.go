package v1

import (
	"github.com/akazwz/go-gin-demo/model"
	"github.com/akazwz/go-gin-demo/model/response"
	"github.com/gin-gonic/gin"
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
	file, err := c.FormFile("file")
	if err != nil {
		response.CommonFailed("Get File Error", CodeGetFileError, c)
		return
	}
	filename := "public/file/" + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		response.CommonFailed("Upload File Error", CodeUploadFileError, c)
		return
	}
	name := file.Filename
	size := file.Size
	fileData := model.File{
		URL:  filename,
		MD5:  "",
		Name: name,
		Size: size,
		Type: "",
	}
	response.Created(fileData, "File Upload Success", c)
}

func GetFileUploadStatus(c *gin.Context) {

}
