FROM golang as Builder

ENV GOPROXY=https://goproxy.cn,direct

COPY . /app

WORKDIR /app
# 该二进制文件是使用动态链接方式编译了一个使用了 GLIBC 库的程序生成的，但是 alpne 镜像中没有 GLIBC 库而是用的 MUSL LIBC 库，这样就会导致该二进制文件无法被执行。
# 改为静态编译
# 如果要使用动态链接函数编译的话，不要依赖 GLIBC （比如编译 Go 程序的时候指定 CGO_ENABLED=0 ） 或者在 alpine 中编译一个依赖 MUSL LIBC 的版本
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine

COPY --from=Builder /app/main .

EXPOSE 8000:8000

ENTRYPOINT ["./main"]