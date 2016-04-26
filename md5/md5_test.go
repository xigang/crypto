package md5

import (
	"testing"
)

//output:
// === RUN TestMD5
// --- PASS: TestMD5 (0.00s)
// 	md5_test.go:17: MD5 encrypted data is successful, the data is: [220 115 82 91 107 101 164 56 86 143 129 100 250 211 12 68]
// 	md5_test.go:17: MD5 encrypted data is successful, the data is: [26 131 61 166 58 107 126 32 9 141 174 6 208 102 2 225]
// 	md5_test.go:17: MD5 encrypted data is successful, the data is: [197 233 176 184 158 5 240 103 67 12 238 220 2 154 163 75]
// PASS
// ok  	github.com/xigang/crypto/md5	0.004s
func TestMD5(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"helloWorld"},
		{"GolangChina"},
	}

	for _, test := range tests {
		t.Logf("MD5 encrypted data is successful, the data is: %v", MD5([]byte(test.param)))
	}
}
