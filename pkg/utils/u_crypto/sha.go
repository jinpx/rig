package u_crypto

import (
	"crypto/sha1"
	"crypto/sha256"
)

func Sha1(text string) []byte {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}

func Sha256(text string) []byte {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}
