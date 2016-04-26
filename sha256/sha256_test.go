package sha256

import (
	"testing"
)

// === RUN TestSHA256
// --- PASS: TestSHA256 (0.00s)
// 	sha256_test.go:18: The success of using SHA256 encryption, encryption of the data is [82 86 26 182 171 202 116 22 77 17 60 182 87 60 212 199 27 75 120 70 43 220 168 196 76 54 53 236 242 54 220 186]
// 	sha256_test.go:18: The success of using SHA256 encryption, encryption of the data is [118 218 169 183 78 119 76 134 114 132 25 68 193 23 5 95 18 192 13 79 126 174 96 226 89 110 205 83 211 15 28 161]
// PASS
// ok  	github.com/xigang/crypto/sha256	0.004s
func TestSHA256(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		// {12124},
		{"\t\r\b"},
	}

	for _, test := range tests {
		data := Sha256([]byte(test.param))
		t.Logf("The success of using SHA256 encryption, encryption of the data is %v", data)
	}
}
