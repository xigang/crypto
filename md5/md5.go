package md5

import (
	"crypto/md5"
)

func MD5(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}
