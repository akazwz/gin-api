package config

type RedisDB struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}
