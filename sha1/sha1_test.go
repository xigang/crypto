package sha1

import (
	"testing"
)

// === RUN TestSha1
// --- PASS: TestSha1 (0.00s)
// 	sha1_test.go:10: sha1 data:  ?6`?u?@??????+ԓh5?=
// PASS
// ok  	github.com/xigang/crypto/sha1	0.008s
func TestSha1(t *testing.T) {
	str := "wangxigang2014@gmail.com"
	data := Sha1([]byte(str))
	t.Log("sha1 data: ", string(data))
}
