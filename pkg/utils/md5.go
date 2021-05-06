package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/akazwz/go-gin-demo/global"
	"github.com/akazwz/go-gin-demo/model"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5Existed(md5 string) (err error, fileMD5 *model.File) {
	var f model.File
	err = global.GDB.Where("md5 = ?", md5).First(&f).Error
	return err, &f
}
