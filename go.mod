module github.com/akaedison/go-gin-demo

go 1.16

replace (
	github.com/akaedison/go-gin-demo/middleware => ../go-gin-demo/middleware
	github.com/akaedison/go-gin-demo/model => ./model
	github.com/akaedison/go-gin-demo/pkg/e => ../go-gin-demo/pkg/e
	github.com/akaedison/go-gin-demo/pkg/setting => ../go-gin-demo/pkg/setting
	github.com/akaedison/go-gin-demo/pkg/util => ../go-gin-demo/pkg/util
	github.com/akaedison/go-gin-demo/routers => ../go-gin-demo/routers
)

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/ugorji/go v1.2.5 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.5 // indirect
	gorm.io/gorm v1.21.6 // indirect
)
