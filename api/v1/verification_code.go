package v1

import (
	"log"
	"time"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// GetVerificationCode
// @Summary 手机号获取验证码
// @Title Get Verification Code
// @Author zwz
// @Description Get Verification Code
// @Tags verification
// @Produce json
// @Param phone query string true "phone number"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /verification/sms [get]
func GetVerificationCode(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		response.CommonFailed("phone is required", CodeParamsError, c)
		return
	}
	verificationCode := utils.CreateRandomVerificationCode()

	// save phone and code in redis and set ex 5 minutes
	err := global.GRDB.Set(c.Request.Context(), phone, verificationCode, time.Minute*5).Err()
	if err != nil {
		log.Println(err)
		response.CommonFailed("redis error", CodeRedisError, c)
		return
	}
	// send sms
	isSuccess := utils.SendVerifyMessage(phone, verificationCode)
	if !isSuccess {
		response.CommonFailed("send sms error", CodeSendSMSError, c)
		return
	}
	response.CommonSuccess(2000, nil, "success", c)
}
