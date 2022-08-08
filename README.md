# GIN-API

# GRPC golang 使用 

1. 安装 protoc  https://github.com/protocolbuffers/protobuf/releases/tag/v21.4
2. 安装 protoc 插件 https://grpc.io/docs/languages/go/quickstart/

````shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
````

3. 将以上安装的加入 PATH环境变量
4. 生成 代码

````shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\xxx.proto
````
