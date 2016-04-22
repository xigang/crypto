package sha256

import (
	"testing"
)

// === RUN TestSha256
// --- PASS: TestSha256 (0.00s)
// 	sha256_test.go:10: sha256 data: RV���tM<�W<��xF+ܨ�L65��6ܺ
// PASS
// ok  	github.com/xigang/crypto/sha256	0.004s
func TestSha256(t *testing.T) {
	str := "wangxigang2014@gmail.com"
	data := Sha256([]byte(str))
	t.Log("sha256 data:", string(data))
}
