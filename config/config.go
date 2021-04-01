package config

type Conf struct {
	RunMode  RunMode  `yaml:"run_mode"`
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
	Zap      Zap      `yaml:"zap"`
}

type RunMode struct {
	Mode string `yaml:"mode"`
}

type App struct {
	PageSize  int    `yaml:"page_size"`
	JwtSecret string `yaml:"jwt_secret"`
}

type Database struct {
	Type string `yaml:"type"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Name string `yaml:"name"`
}

type Zap struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Prefix        string `yaml:"prefix"`
	Director      string `yaml:"director"`
	LinkName      string `yaml:"link-name"`
	ShowLine      bool   `yaml:"showLine"`
	EncodeLevel   string `yaml:"encode-level"`
	StacktraceKey string `yaml:"stacktrace-key"`
	LogInConsole  bool   `yaml:"log-in-console"`
}