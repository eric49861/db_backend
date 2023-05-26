package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptWithMD5 采用MD5算法对字符串进行加密，对网站用户的密码加密就是调用该函数
func EncryptWithMD5(targetStr string) string {
	md5Byte := md5.Sum([]byte(targetStr))

	return hex.EncodeToString(md5Byte[:])
}
