package config

type Conf struct {
	Server      Server      `yaml:"server"`
	Database    Database    `yaml:"database"`
	RedisDB     RedisDB     `yaml:"redisDB"`
	Tencent     Tencent     `yaml:"tencent"`
	JWT         JWT         `yaml:"jwt"`
	Zap         Zap         `yaml:"zap"`
	MiniProgram MiniProgram `yaml:"miniProgram"`
}
