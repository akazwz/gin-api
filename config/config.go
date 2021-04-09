package config

type Conf struct {
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
	JWT      JWT      `json:"jwt" yaml:"jwt"`
	Zap      Zap      `json:"zap" yaml:"zap"`
}
