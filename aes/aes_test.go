package aes

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/xigang/crypto/base64"
)

//output:
//=== RUN TestAes
// --- PASS: TestAes (0.00s)
// 	aes_test.go:23: jKPu0zi5EsU1NK5m+6ieosJQUj465gHmtMS52BdUYTY=
// 	aes_test.go:29: wangxgang201@gmail.com
// PASS
// ok  	github.com/xigang/crypto/aes	0.010s
func TestAes(t *testing.T) {
	str := "wangxgang201@gmail.com"

	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return
	}

	result, err := AesEncrypt([]byte(str), key)
	if err != nil {
		t.Error("aes encrypt data error.", err)
	}
	t.Log(base64.Base64Encode(result))

	origData, err := AesDecrypt(result, key)
	if err != nil {
		t.Error("aes decrypt data error.")
	}
	t.Log(string(origData))
}
