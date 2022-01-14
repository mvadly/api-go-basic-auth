package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func SaltMD5(username, password string) string {
	salt := os.Getenv("SALT_KEY")
	result := md5.Sum([]byte(username + password + salt))
	final := hex.EncodeToString(result[:])
	return final
}
