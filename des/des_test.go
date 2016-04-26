package des

import (
	"testing"
)

// === RUN TestDES
// --- PASS: TestDES (0.00s)
// 	des_test.go:56: The success of using DES encryption, encryption of the data is: wangxigang2014@gmail.com
// 	des_test.go:56: The success of using DES encryption, encryption of the data is: HelloWorld
// 	des_test.go:56: The success of using DES encryption, encryption of the data is: GolangChina
// 	des_test.go:56: The success of using DES encryption, encryption of the data is: 123456789      ghljsdhkja
// PASS
// ok  	github.com/xigang/crypto/des	0.006s
func TestDES(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"HelloWorld"},
		{"GolangChina"},
		{"123456789\t\bghljsdhkja"},
	}
	key := []byte("sfe023f_")
	for _, test := range tests {
		result, err := DesEncrypt([]byte(test.param), key)
		if err != nil {
			t.Error("DES encrypt data error: ", err)
		}

		oridData, err := DesDecrypt(result, key)
		if err != nil {
			t.Error("DES decrypt data error: ", err)
		}

		if test.param == string(oridData) {
			t.Logf("The success of using DES encryption, encryption of the data is: %v", string(oridData))
		}
	}
}

// === RUN Test3DES
// --- PASS: Test3DES (0.00s)
// 	des_test.go:89: The success of using 3DES encryption, encryption of the data is: wangxigang2014@gmail.com
// 	des_test.go:89: The success of using 3DES encryption, encryption of the data is: HelloWorld
// 	des_test.go:89: The success of using 3DES encryption, encryption of the data is: GolangChina
// 	des_test.go:89: The success of using 3DES encryption, encryption of the data is: 123456789     ghljsdhkja
// PASS
// ok  	github.com/xigang/crypto/des	0.007s
func Test3DES(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"HelloWorld"},
		{"GolangChina"},
		{"123456789\t\bghljsdhkja"},
	}
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	for _, test := range tests {
		result, err := TripleDesEncrypt([]byte(test.param), key)
		if err != nil {
			t.Error("DES encrypt data error: ", err)
		}

		oridData, err := TripleDesDecrypt(result, key)
		if err != nil {
			t.Error("DES decrypt data error: ", err)
		}

		if test.param == string(oridData) {
			t.Logf("The success of using 3DES encryption, encryption of the data is: %v", string(oridData))
		}
	}
}
