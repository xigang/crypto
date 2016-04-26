package crypto

import (
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
