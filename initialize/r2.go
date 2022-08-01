package initialize

import (
	"log"
	"os"

	"github.com/akazwz/gin-api/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitR2Client() *s3.Client {
	accountId := os.Getenv("R2_ACCOUNT_ID")
	accessKeyId := os.Getenv("R2_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("R2_ACCESS_KEY_SECRET")

	// 生成 s3 client
	client, err := utils.GenerateR2Client(accountId, accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatalln("初始化 R2 Client 失败")
	}
	return client
}
