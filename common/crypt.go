package common

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// CryptPass encrypts a password.
func CryptPass(pass string) string {
	passwordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(pass)))
	hashed, err := bcrypt.GenerateFromPassword([]byte(passwordMD5), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

// IsSamePass verifies two passwords (an unhashed and an hashed one) to be one and the same.
func IsSamePass(passedPass string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(passedPass))
	if err == nil {
		return true
	}
	return false
}
