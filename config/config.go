package config

import "time"

type Conf struct {
	Server      Server      `yaml:"server"`
	Database    Database    `yaml:"database"`
	RedisDB     RedisDB     `yaml:"redisDB"`
	JWT         JWT         `yaml:"jwt"`
	Tencent     Tencent     `yaml:"tencent"`
	Zap         Zap         `yaml:"zap"`
	MiniProgram MiniProgram `yaml:"miniProgram"`
}

type Server struct {
	Mode         string        `yaml:"mode"`
	Addr         int           `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	AliOSS       AliOSS        `yaml:"aliOSS"`
}

type AliOSS struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	BucketName      string `json:"bucketName" yaml:"bucketName"`
	BucketUrl       string `json:"bucketUrl" yaml:"bucketUrl"`
}

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type RedisDB struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type JWT struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int64  `yaml:"expiresTime"`
	BufferTime  int64  `yaml:"bufferTime"`
}

type Tencent struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
}

type MiniProgram struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type Zap struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Prefix        string `yaml:"prefix"`
	Director      string `yaml:"director"`
	LinkName      string `yaml:"linkName"`
	ShowLine      bool   `yaml:"showLine"`
	EncodeLevel   string `yaml:"encodeLevel"`
	StacktraceKey string `yaml:"stacktraceKey"`
	LogInConsole  bool   `yaml:"logInConsole"`
}
