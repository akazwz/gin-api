package config

type Conf struct {
	RunMode  RunMode  `json:"run_mode" yaml:"run_mode"`
	App      App      `json:"app" yaml:"app"`
	Database Database `json:"database" yaml:"database"`
	JWT      JWT      `json:"jwt" yaml:"jwt"`
	Zap      Zap      `json:"zap" yaml:"zap"`
}

type RunMode struct {
	Mode string `yaml:"mode"`
}

type App struct {
	PageSize int `yaml:"page_size"`
}
