package utils

import (
	"testing"
)

func TestVerifySignature(t *testing.T) {
	secret := []byte("simon")
	body := []byte("testing with some body")
	sign := "sha1=61043448a05f65a3c9a97d7134de79eff648a2af"

	if !VerifySignature(secret, body, sign) {
		t.Fatalf("failed to verify signature")
	}
}
