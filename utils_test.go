package crypto

import (
	"crypto/des"
	"testing"
)

const (
	BITS = 2048
)

const (
	RsaPrivateKeyPath = "./rsa_private_key.pem"
	RsaPublicKeyPath  = "./rsa_public_key.pem"
)

func TestGenRsaKey(t *testing.T) {
	err := GenRsaKey(BITS)
	if err != nil {
		t.Error("generate pub and priv key error.", err)
	}
}

func TestGetFileContent(t *testing.T) {
	data, err := GetFileContent(RsaPrivateKeyPath)
	if err != nil {
		t.Error("failed to load file.", err)
	}

	t.Log("success to load file: ", string(data))
}

func TestGetPrividateKey(t *testing.T) {
	data, err := GetFileContent(RsaPrivateKeyPath)
	if err != nil {
		t.Error("failed to load file.", err)
	}
	privateKey, err := GetPrivateKey(data)
	if err != nil {
		t.Error("get providate key error: ", err)
	}

	t.Log("private key: ", *privateKey)
}

func TestGetPublicKey(t *testing.T) {
	data, err := GetFileContent(RsaPublicKeyPath)
	if err != nil {
		t.Error("failed to load file.", err)
	}

	pubicKey, err := GetPublicKey(data)
	if err != nil {
		t.Error("get public key error: ", err)
	}

	t.Log("public key:", *pubicKey)
}

func TestZeroPadding(t *testing.T) {
	var tests = []struct {
		data string
		key  string
	}{
		{"wangxigang2014@gmail.com", "sfe023f_"},
	}

	for _, test := range tests {
		t.Log("start data length", len(string(test.data)))
		block, err := des.NewCipher([]byte(test.key))
		if err != nil {
			t.Error(err)
		}
		origData := ZeroPadding([]byte(test.data), block.BlockSize())
		t.Logf("padding data:%v and length:%v", string(origData), len(string(origData)))

		data := ZeroUnPadding(origData)
		t.Logf("data is:%v and data length:%v", string(data), len(string(data)))
	}
}

func TestPKCS5Padding(t *testing.T) {
	var tests = []struct {
		data string
		key  string
	}{
		{"wangxigang2014@gmail.com", "sfe023f_"},
	}

	for _, test := range tests {
		t.Log("start data length", len(string(test.data)))
		block, err := des.NewCipher([]byte(test.key))
		if err != nil {
			t.Error(err)
		}
		origData := PKCS5Padding([]byte(test.data), block.BlockSize())
		t.Logf("padding data:%v and length:%v", string(origData), len(string(origData)))

		data := PKCS5UnPadding(origData)
		t.Logf("data is:%v and data length:%v", string(data), len(string(data)))

	}
}
