package v1

import (
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateVerificationCode(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		response.CommonFailed("phone is required", CodeParamsError, c)
		return
	}
	verificationCode := utils.CreateRandomVerificationCode()
	isSuccess := utils.SendVerifyMessage(phone, verificationCode)
	if !isSuccess {
		response.CommonFailed("send sms error", CodeSendSMSError, c)
		return
	}
	response.CommonSuccess(2000, nil, "success", c)
}
