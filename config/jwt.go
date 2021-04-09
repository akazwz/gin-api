package config

type JWT struct {
	SigningKey  string `json:"signing_key" yaml:"signing_key"`
	ExpiresTime int64  `json:"expires_time" yaml:"expires_time"`
	BufferTime  int64  `json:"buffer_time" yaml:"buffer_time"`
}
