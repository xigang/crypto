package sha256

import (
	"crypto/sha256"
)

//Sha256 return SHA256 encrypted data
func Sha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
