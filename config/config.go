package config

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
	Zap      Zap      `yaml:"zap"`
}
