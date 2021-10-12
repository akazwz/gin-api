module github.com/akazwz/go-gin-restful-api

go 1.16

replace (
	github.com/akazwz/go-gin-restful-api/model => ./model
	github.com/akazwz/go-gin-restful-api/pkg/utils => ../go-gin-restful-api/pkg/utils
	github.com/akazwz/go-gin-restful-api/routers => ../go-gin-restful-api/routers
	github.com/akazwz/go-gin-restful-api/service => ../go-gin-restful-api/service
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/juju/ratelimit v1.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
