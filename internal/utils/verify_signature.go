package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

// VerifySignature hashes body with the secret key and compares to the passed sign header
func VerifySignature(secret []byte, body []byte, sign string) bool {
	h := hmac.New(sha1.New, secret)
	h.Write(body)
	bs := h.Sum(nil)

	return sign == fmt.Sprintf("sha1=%x", bs)
}
