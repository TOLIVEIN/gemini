package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//EncodeMD5 ...
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

//EncodeBcrypt ...
func EncodeBcrypt(value string) string {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hashedValue)
}
