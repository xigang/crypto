package sha1

import (
	"crypto/sha1"
)

//Sha1 return SHA1 encrypted data
func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
