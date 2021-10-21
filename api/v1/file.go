package v1

import (
	"os"
	"time"

	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/pkg/utils"
	"github.com/akazwz/go-gin-restful-api/pkg/utils/upload"
	"github.com/akazwz/go-gin-restful-api/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
// @Param md5 formData string true "md5"
// @Param token header string true "token"
// @Success 201 {object} response.Response
// @Failure 400,401 {object} response.Response
// @Router /file [post]
func CreateFile(c *gin.Context) {
	fullPath := c.FullPath()
	userUuid := ""
	if fullPath != "/v1/avatar" {
		// get claims from middleware
		claims, _ := c.Get("claims")
		// convent claims to type *request.CustomClaims
		customClaims := claims.(*request.CustomClaims)
		// get user uuid to store who upload this file
		userUuid = customClaims.UUID.String()
	}
	// md5 a pass
	md5 := c.PostForm("md5")
	err, fileExisted := utils.MD5Existed(md5)
	if err == nil {
		userFileData := model.File{
			UserUuid: userUuid,
			URL:      fileExisted.URL,
			MD5:      fileExisted.MD5,
			Name:     fileExisted.Name,
			Size:     fileExisted.Size,
			Type:     fileExisted.Type,
		}
		if err := service.CreateUserFile(&userFileData); err != nil {
			response.CommonFailed("Save User File To DB Error", CodeDbErr, c)
			return
		}
		response.Created(userFileData, "A Pass Success", c)
		return
	}
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
	fileNamePrefix := uuid.NewV4().String()
	// user full file name
	fileName := fileNamePrefix + "-" + file.Filename
	// user get file by this url
	url := dir + fileName
	// the local file real complete store path
	location := dir + fileName
	// the file oss store path
	objectKey := dirDate + "/" + fileName
	// get file md5
	md5File := utils.GetFileMD5(file)
	if md5File != md5 {
		response.CommonFailed("Upload File Damaged", CodeUploadFileError, c)
		return
	}
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
	fileMD5Data := model.FileMD5{
		MD5:      md5File,
		UserUuid: userUuid,
		Location: location,
		Size:     size,
		Type:     fileType,
	}
	if err := service.CreateMD5File(&fileMD5Data); err != nil {
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
