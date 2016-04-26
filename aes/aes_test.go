package aes

import (
	"crypto/rand"
	"io"
	"testing"
)

func TestAES(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"1234567890"},
		{"GolangChina"},
		{"\n\b\t"},
	}

	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return
	}

	for _, test := range tests {
		result, err := AesEncrypt([]byte(test.param), key)
		if err != nil {
			t.Error("AES encrypt data error.", err)
		}

		origData, err := AesDecrypt(result, key)
		if err != nil {
			t.Error("AES decrypt data error.")
		}

		if string(origData) == test.param {
			t.Logf("encryption and decryption success, and origin data is: %v", string(origData))
		}
	}
}

//output:
// === RUN TestAES
// --- PASS: TestAES (0.00s)
// 	aes_test.go:36: encryption and decryption success, and origin data is: wangxigang2014@gmail.com
// 	aes_test.go:36: encryption and decryption success, and origin data is: 1234567890
// 	aes_test.go:36: encryption and decryption success, and origin data is: GolangChina
// 	aes_test.go:36: encryption and decryption success, and origin data is:

// PASS
// ok  	github.com/xigang/crypto/aes	0.007s
