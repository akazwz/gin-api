package config

type Server struct {
	Mode         string `yaml:"mode"`
	Addr         int    `yaml:"addr"`
	ReadTimeout  int64  `yaml:"readTimeout"`
	WriteTimeout int64  `yaml:"writeTimeout"`
}
