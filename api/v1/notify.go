package v1

import (
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

// SetNotify
// @Summary 通知设置
// @Title 设置通知间隔和次数
// @Author zwz
// @Description 设置通知间隔和次数
// @Tags user
// @Accept json
// @Produce json
// @Param notify body request.SetNotify true "notify"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/notify [post]
func SetNotify(c *gin.Context) {
	var setNotify request.SetNotify
	if err := c.ShouldBindJSON(&setNotify); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUUID := customClaims.UUID

	var notify model.Notify
	err := global.GDB.Where("user_uuid = ?", userUUID).First(&notify).Error

	if err != nil {
		// 错误
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
			response.CommonFailed("Set Notify Fail", CodeDbErr, c)
			return
		}
		// 没有记录,新建
		err = global.GDB.Create(&model.Notify{
			UserUUID:         userUUID,
			NotifyGap:        setNotify.NotifyGap,
			NotifyCount:      setNotify.NotifyCount,
			NotifyLimitCount: 1000,
			LastNotify:       time.Now(),
		}).Error

		if err != nil {
			log.Println(err)
			response.CommonFailed("Set Notify Fail", CodeDbErr, c)
			return
		}
	}

	// 有记录,修改
	err = global.GDB.Where("user_uuid = ?", userUUID).Updates(&model.Notify{
		NotifyGap:        setNotify.NotifyGap,
		NotifyCount:      setNotify.NotifyCount,
		NotifyLimitCount: 1000,
	}).Error

	if err != nil {
		log.Println(err)
		response.CommonFailed("Set Notify Fail", CodeDbErr, c)
		return
	}
	response.SuccessWithMessage("Set Notify Success", c)
}
