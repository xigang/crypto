package crypto

/*
import (
	"crypto/rand"
	"io"
	"testing"
)

const (
	BITS = 2048
)

const (
	RSA_PRIVATE_KEY_PATH = "./rsa_private_key.pem"
	RSA_PUBLIC_KEY_PATH  = "./rsa_public_key.pem"
)

//generate public key private key file.
func TestGenRsaKey(t *testing.T) {
	err := GenRsaKey(BITS)
	if err != nil {
		t.Error("generate pub and priv key error.", err)
	}
}

//read public or private key from file.
func TestGetFileContent(t *testing.T) {
	data, err := GetFileContent(RSA_PRIVATE_KEY_PATH)
	if err != nil {
		t.Error("read file content error.")
	}
	t.Log(string(data))
}

//Using RSA for encryption and decryption
func TestRsa(t *testing.T) {
	mailStr := "wangxigang2014@gmail.com"

	pubKeyData, err := GetFileContent(RSA_PUBLIC_KEY_PATH)
	if err != nil {
		t.Error("read public key content error.")
	}

	ciphertest, err := RsaEncrypt(pubKeyData, []byte(mailStr))
	if err != nil {
		t.Error("rsa encrypt data error.")
	}
	t.Log(string(ciphertest))

	privKeyData, err := GetFileContent(RSA_PRIVATE_KEY_PATH)
	if err != nil {
		t.Error("read private key error.")
	}

	origData, err := RsaDecrypt(privKeyData, ciphertest)
	if err != nil {
		t.Error("rsa decrypt data error.")
	}

	t.Log(string(origData))
}

//AES for encryption and decryption
// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
func TestAes(t *testing.T) {
	mailStr := "wangxgang201@gmail.com"

	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return
	}

	result, err := AesEncrypt([]byte(mailStr), key)
	if err != nil {
		t.Error("aes encrypt data error.", err)
	}
	t.Log(Base64Encode(result))

	origData, err := AesDecrypt(result, key)
	if err != nil {
		t.Error("aes decrypt data error.")
	}
	t.Log(string(origData))
}

//Des for encryption and decryption
func TestDes(t *testing.T) {
	mailStr := "wangxgang201@gmail.com"

	key := []byte("sfe023f_")

	result, err := DesEncrypt([]byte(mailStr), key)
	if err != nil {
		t.Error("des encrypt data error.")
	}
	t.Log(Base64Encode(result))

	origData, err := DesDecrypt(result, key)
	if err != nil {
		t.Error("des decrypt data error.")
	}
	t.Log(string(origData))
}

func Test3Des(t *testing.T) {
	mailStr := "wangxgang201@gmail.com"

	key := []byte("sfe023f_sefiel#fi32lf3e!")

	result, err := TripleDesEncrypt([]byte(mailStr), key)
	if err != nil {
		t.Error("des encrypt data error.")
	}
	t.Log(Base64Encode(result))

	origData, err := TripleDesDecrypt(result, key)
	if err != nil {
		t.Error("des decrypt data error.")
	}
	t.Log(string(origData))
}
*/
