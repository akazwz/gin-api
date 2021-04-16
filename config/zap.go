package config

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
