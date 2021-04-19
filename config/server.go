package config

import "time"

type Server struct {
	Mode         string        `yaml:"mode"`
	Addr         int           `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}
