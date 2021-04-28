package model

type File struct {
	URL  string `json:"url"`
	MD5  string `json:"md5"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}
