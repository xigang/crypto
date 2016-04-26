package rsa

import (
	"testing"

	"github.com/xigang/crypto"
)

const (
	RsaPrivateKeyPath = "../rsa_private_key.pem"
	RsaPublicKeyPath  = "../rsa_public_key.pem"
)

func TestRSA(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{"wangxigang2014@gmail.com"},
		{"HelloWorld"},
		{"\t\t\n\tdfghjkl34567890"},
	}

	for _, test := range tests {
		publicKey, err := crypto.GetFileContent(RsaPublicKeyPath)
		if err != nil {
			t.Error("read public key error: ", err)
		}

		ciphertest, err := RsaEncrypt(publicKey, []byte(test.param))
		if err != nil {
			t.Error("rsa encrypt data error.")
		}

		privateKey, err := crypto.GetFileContent(RsaPrivateKeyPath)
		if err != nil {
			t.Error("read private key error: ", err)
		}

		origData, err := RsaDecrypt(privateKey, ciphertest)
		if err != nil {
			t.Error("rsa decrypt data error: ", err)
		}

		if test.param == string(origData) {
			t.Logf("encryption and decryption success, decryption data is %v", string(origData))
		}
	}

}

//output:
//=== RUN TestRSA
// --- PASS: TestRSA (0.42s)
// 	rsa_test.go:45: encryption and decryption success, decryption data is wangxigang2014@gmail.com
// 	rsa_test.go:45: encryption and decryption success, decryption data is HelloWorld
// 	rsa_test.go:45: encryption and decryption success, decryption data is
// 			dfghjkl34567890
// PASS
// ok  	github.com/xigang/crypto/rsa	0.424s
