package base64

import (
	"testing"
)

func TestBase64(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"HelloWorld"},
		{"GolangChina"},
		{"1234567890"},
		{"\r\n\t\b2345678"},
	}

	for _, test := range tests {
		encodeData := Base64Encode([]byte(test.param))
		decodeData, err := Base64Decode(string(encodeData))
		if err != nil {
			t.Error(err)
		}

		if string(decodeData) == test.param {
			t.Logf("Base64 encoding data is: %v and decoding data is: %v", test.param, string(decodeData))
		}
	}
}

//output:
// === RUN TestBase64
// --- PASS: TestBase64 (0.00s)
// 	base64_test.go:44: Base64 encoding data is: wangxigang2014@gmail.com and decoding data is: wangxigang2014@gmail.com
// 	base64_test.go:44: Base64 encoding data is: HelloWorld and decoding data is: HelloWorld
// 	base64_test.go:44: Base64 encoding data is: GolangChina and decoding data is: GolangChina
// 	base64_test.go:44: Base64 encoding data is: 1234567890 and decoding data is: 1234567890
// 	base64_test.go:44: Base64 encoding data is:
// 		       2345678 and decoding data is:
// 		       2345678
// PASS
// ok  	github.com/xigang/crypto/base64	0.004s
