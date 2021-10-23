package utils

import (
	"context"
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"log"
	"math/rand"
	"time"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/silenceper/wechat/v2/miniprogram/config"
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
		log.Printf("An API error has returned: %s", err)
		return false
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetVerificationStatus(phone, code string, c context.Context) bool {
	val, err := global.GRDB.Get(c, phone).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	return val == code
}

func GetSessionByCode(code string) (session auth.ResCode2Session) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     global.CFG.MiniProgram.AppId,
		AppSecret: global.CFG.MiniProgram.AppSecret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	session, err := a.Code2Session(code)
	if err != nil {
		log.Println("获取 session 错误")
		return
	}
	return
}

func GetMiniUserInfo(sessionKey, Encrypt, Iv string) (UserInfo *encryptor.PlainData) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     global.CFG.MiniProgram.AppId,
		AppSecret: global.CFG.MiniProgram.AppSecret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	e := mini.GetEncryptor()
	UserInfo, err := e.Decrypt(sessionKey, Encrypt, Iv)
	if err != nil {
		log.Println("解密数据错误")
		return
	}
	return
}
