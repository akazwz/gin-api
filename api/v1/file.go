package v1

import (
	"github.com/akazwz/go-gin-demo/model"
	"github.com/akazwz/go-gin-demo/model/request"
	"github.com/akazwz/go-gin-demo/model/response"
	"github.com/akazwz/go-gin-demo/pkg/utils"
	"github.com/akazwz/go-gin-demo/pkg/utils/upload"
	"github.com/akazwz/go-gin-demo/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	// get claims from middleware
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUuid := customClaims.UUID.String()
	// get file from form
	file, err := c.FormFile("file")
	if err != nil {
		response.CommonFailed("Get File Error", CodeGetFileError, c)
		return
	}
	// declare dirDate
	dirDate := time.Now().Format("2006-01-02")
	// declare file store dir
	dir := "public/file/" + dirDate + "/"
	// judge dir is existed if not existed make dir
	_, err = os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			response.CommonFailed("Create DIR Error", CodeUploadFileError, c)
			return
		}
	}
	// user file name prefix
	fileNamePrefix := time.Now().Format("15-04-05")
	// user full file name
	fileName := fileNamePrefix + "-" + file.Filename
	// user get file by this url
	url := dir + fileName
	// the local file real complete store path
	location := dir + uuid.NewV4().String() + "-" + file.Filename
	// the file oss store path
	objectKey := dirDate + "/" + uuid.NewV4().String() + "-" + file.Filename
	// get file md5
	md5File := utils.GetFileMD5(file)
	// get file real type
	fileType := utils.GetFileType(file)

	name := fileName
	size := file.Size
	// save file local
	if err := c.SaveUploadedFile(file, location); err != nil {
		response.CommonFailed("Upload File Error", CodeUploadFileError, c)
		return
	}
	// upload fil
	if err := upload.OSSUploadFile(file, objectKey); err != nil {
		response.CommonFailed("Upload OSS Error", CodeUploadFileError, c)
		return
	}
	// md5 file db data
	FileMD5Data := model.FileMD5{
		MD5:      md5File,
		Location: location,
		Size:     size,
		Type:     fileType,
	}
	if err := service.CreateMD5File(&FileMD5Data); err != nil {
		response.CommonFailed("Save MD5 File To DB Error", CodeDbErr, c)
		return
	}
	// user file db data
	userFileData := model.File{
		UserUuid: userUuid,
		URL:      url,
		MD5:      md5File,
		Name:     name,
		Size:     size,
		Type:     fileType,
	}
	if err := service.CreateUserFile(&userFileData); err != nil {
		response.CommonFailed("Save User File To DB Error", CodeDbErr, c)
		return
	}
	response.Created(userFileData, "File Upload Success", c)
}
