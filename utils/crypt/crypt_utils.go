package crypt

import (
	"app-version-manager/utils/bcrypt"
	"crypto/rand"
	"fmt"
	"strings"
)

func CheckSignInName(value string) bool {
	if value == "" {
		return false
	}
	trimValue := strings.TrimSpace(value)
	if len(trimValue) < 4 {
		return false
	}

	return true
}

func GenerateToken() string {
	b2 := make([]byte, 20)
	_, err := rand.Read(b2)
	if err == nil {
		return fmt.Sprintf("%x", b2)
	}
	return ""
}

func PassWordEncrypted(value string) string {
	passWord := []byte(string(value))
	hashedPassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", hashedPassword)

}
