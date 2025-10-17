package util

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordConvertMD5(password string) (encryption string) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryption = hex.EncodeToString(hasher.Sum(nil))
	return
}
