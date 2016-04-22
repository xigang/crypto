package rsa

import (
	"testing"

	"github.com/xigang/crypto"
)

const (
	RSA_PRIVATE_KEY_PATH = "../rsa_private_key.pem"
	RSA_PUBLIC_KEY_PATH  = "../rsa_public_key.pem"
)

//output:
//=== RUN TestRsa
// --- PASS: TestRsa (0.15s)
// ����O��>{��)P���,X���KN\#�Z>$9����V-���:p��u*��?�~�u���qT�=|��7��®V��CKc-U�B��� ��7��f)	����,հJ�hv8C�޳��
// �}߶�Xh�b�1�b�)��;��YR���yz#p���Ѓ�                                                                       �<�Lά�~
//                                  �N\�
// 	rsa_test.go:45: wangxigang2014@gmail.com
// PASS
func TestRsa(t *testing.T) {
	str := "wangxigang2014@gmail.com"

	publicKey, err := crypto.GetFileContent(RSA_PUBLIC_KEY_PATH)
	if err != nil {
		t.Error("read public key error: ", err)
	}

	ciphertest, err := RsaEncrypt(publicKey, []byte(str))
	if err != nil {
		t.Error("rsa encrypt data error.")
	}
	t.Log(string(ciphertest))

	privateKey, err := crypto.GetFileContent(RSA_PRIVATE_KEY_PATH)
	if err != nil {
		t.Error("read private key error: ", err)
	}

	origData, err := RsaDecrypt(privateKey, ciphertest)
	if err != nil {
		t.Error("rsa decrypt data error: ", err)
	}

	t.Log(string(origData))
}
