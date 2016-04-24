package des

import (
	"crypto/cipher"
	"crypto/des"

	"github.com/xigang/crypto"
)

//DesEncrypt return DES encryption
func DesEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	origData = crypto.PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//DesDecrypt return DES decryption
func DesDecrypt(crypted []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = crypto.PKCS5UnPadding(origData)
	return origData, nil
}

//TripleDesEncrypt return 3DES encryption
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = crypto.PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//TripleDesDecrypt return 3DES decryption
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = crypto.PKCS5UnPadding(origData)
	return origData, nil
}
