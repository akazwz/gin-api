module github.com/akazwz/go-gin-demo

go 1.16

replace (
	github.com/akazwz/go-gin-demo/model => ./model
	github.com/akazwz/go-gin-demo/pkg/util => ../go-gin-demo/pkg/util
	github.com/akazwz/go-gin-demo/routers => ../go-gin-demo/routers
	github.com/akazwz/go-gin-demo/service => ../go-gin-demo/service
)

require (
	github.com/akazwz/go-gin-demo v0.0.0-20210422144646-1c7eeda006a6
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	go.uber.org/zap v1.16.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
