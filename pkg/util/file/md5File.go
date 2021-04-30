package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func Md5File(file multipart.FileHeader) string {
	openFile, err := file.Open()
	if err != nil {
	}

	defer func(openFile multipart.File) {
		err := openFile.Close()
		if err != nil {
		}
	}(openFile)

	md5h := md5.New()

	_, err = io.Copy(md5h, openFile)
	if err != nil {
	}

	sum := md5h.Sum(nil)
	toString := hex.EncodeToString(sum)
	return toString
}
