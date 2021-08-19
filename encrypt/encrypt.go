package encrypt

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"
)

type Encryptor interface {
	// 使用 salt 加密
	Encrypt(salt string) string
	// 创建新密码
	CreateNewPassword() (password, salt string)
}

func NewEncryptor(s string) Encryptor {
	return md5Enc{s}
}

type md5Enc struct {
	s string
}

// Encrypt 使用 salt 加密密码
func (e md5Enc) Encrypt(salt string) string {
	return e.encrypt(salt)
}

// CreateNewPassword 创建新密码
func (e md5Enc) CreateNewPassword() (password, salt string) {
	salt = e.random(10)
	password = e.encrypt(salt)
	return
}

// 生成指定位数的随机字符串
func (e md5Enc) random(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*,./;':[]{}"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func (e md5Enc) encrypt(salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(e.s+salt)))
}
