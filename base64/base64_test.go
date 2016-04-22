package base64

import (
	"testing"
)

// === RUN TestBase64
// --- PASS: TestBase64 (0.00s)
// 	base64_test.go:10: base64Encode: d2FuZ3hpZ2FuZzIwMTRAZ21haWwuY29t
// 	base64_test.go:16: success:  wangxigang2014@gmail.com
// PASS
// ok  	github.com/xigang/crypto/base64	0.004s
func TestBase64(t *testing.T) {
	src := "wangxigang2014@gmail.com"
	base64Encode := Base64Encode([]byte(src))
	t.Log("base64Encode:", string(base64Encode))

	base64Decode, err := Base64Decode(base64Encode)
	if err != nil {
		t.Error("error: ", err)
	}
	t.Log("success: ", string(base64Decode))
}
