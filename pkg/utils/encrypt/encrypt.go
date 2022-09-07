package encrypt

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// CompareHashAndPassword ...
func CompareHashAndPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

// GenerateFromPassword ...
func GenerateFromPassword(password string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b)
}

// Md5 ...
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
