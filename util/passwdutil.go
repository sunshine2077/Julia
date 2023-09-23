package util

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 获取盐
func GetSault(size int) string {
	rand.Seed(time.Now().UnixNano())
	words := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	var sb strings.Builder
	for i := 0; i < size; i++ {
		index := rand.Intn(len(words))
		sb.WriteByte(words[index])
	}
	return sb.String()
}

// 对密码加密，避免明文
func EncodingPwd(sault string, str string) string {
	ans := fmt.Sprintf("%X", md5.Sum([]byte(string(sha512.New().Sum([]byte(sault)))+str)))
	return ans
}
