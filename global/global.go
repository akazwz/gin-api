package global

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

var (
	GDB *gorm.DB
	R2C *s3.Client
)
