FROM golang as Builder

ENV GOPROXY=https://goproxy.cn,direct

COPY . /app

WORKDIR /app

RUN go build -o main

FROM alpine:latest

COPY --from=Builder /app/main /

EXPOSE 8000:8000

ENTRYPOINT ["./main"]