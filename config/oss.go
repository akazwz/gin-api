package config

type AliOSS struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	BucketName      string `json:"bucketName" yaml:"bucketName"`
	BucketUrl       string `json:"bucketUrl" yaml:"bucketUrl"`
}
