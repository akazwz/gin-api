package config

type Server struct {
	Mode         string `json:"mode" yaml:"mode"`
	Addr         int    `json:"addr" yaml:"addr"`
	ReadTimeout  int64  `json:"read_timeout"`
	WriteTimeout int64  `json:"write_timeout"`
}
