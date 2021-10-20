package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func CreateRandomVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vCode
}

func SendVerifyMessage(phone, verificationCode string) bool {
	credential := common.NewCredential(global.CFG.Tencent.SecretId, global.CFG.Tencent.SecretKey)

	cpf := profile.NewClientProfile()
	cpf.SignMethod = "HmacSHA1"

	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()

	request.SmsSdkAppId = common.StringPtr("1400576425")
	request.SignName = common.StringPtr("赵文卓工作学习")
	request.SenderId = common.StringPtr("")
	request.ExtendCode = common.StringPtr("")
	request.TemplateParamSet = common.StringPtrs([]string{verificationCode, "5"})
	request.TemplateId = common.StringPtr("1164712")
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	_, err := client.SendSms(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
		return false
	}
	return true
}
