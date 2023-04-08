package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Md5(data interface{}) (string, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	sh := md5.New()
	_, err = sh.Write(bt)
	if err != nil {
		return "", err
	}
	buf := hex.EncodeToString(sh.Sum(nil))
	return string(buf), err
}
