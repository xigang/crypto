package md5

import (
	"testing"
)

// === RUN TestMd5
// --- PASS: TestMd5 (0.00s)
// 	md5_test.go:10: md5 data:  [220 115 82 91 107 101 164 56 86 143 129 100 250 211 12 68]
// PASS
// ok  	github.com/xigang/crypto/md5	0.004s
func TestMd5(t *testing.T) {
	str := "wangxigang2014@gmail.com"
	hashData := MD5([]byte(str))
	t.Log("md5 data: ", hashData)
}
