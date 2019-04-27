package md5

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 return MD5 hash data
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
