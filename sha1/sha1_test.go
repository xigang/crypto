package sha1

import (
	"testing"
)

// === RUN TestSHA1
// --- PASS: TestSHA1 (0.00s)
// 	sha1_test.go:30: �6`�u�@������+ԓh5�=
// 	sha1_test.go:31: The success of using SHA1 encryption, the encrypted data is wangxigang2014@gmail.com
// 	sha1_test.go:30: ۊ��Y�ԡ1�S���_1�T�
// 	sha1_test.go:31: The success of using SHA1 encryption, the encrypted data is HelloWorld
// 	sha1_test.go:30: >�ѕȞR�X��	7��~Y��
// 	sha1_test.go:31: The success of using SHA1 encryption, the encrypted data is GolangChina
// 	sha1_test.go:30: %s�}�",��������f6'
// 	sha1_test.go:31: The success of using SHA1 encryption, the encrypted data is 1234567890
// PASS
// ok  	github.com/xigang/crypto/sha1	0.004s

func TestSHA1(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"HelloWorld"},
		{"GolangChina"},
		{"1234567890\t\b"},
	}

	for _, test := range tests {
		data := Sha1([]byte(test.param))
		t.Log(string(data))
		t.Logf("The success of using SHA1 encryption, the encrypted data is %v", test.param)
	}
}
