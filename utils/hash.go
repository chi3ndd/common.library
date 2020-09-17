package utils

import (
	"crypto"
	"fmt"
)

// Md5 : Algorithm
func Md5(str string) string {
	h := crypto.MD5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha1 : Algorithm
func Sha1(str string) string {
	h := crypto.SHA1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha256 : Algorithm
func Sha256(str string) string {
	h := crypto.SHA256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha512 : Algorithm
func Sha512(str string) string {
	h := crypto.SHA512.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
