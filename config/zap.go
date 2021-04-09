package config

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
