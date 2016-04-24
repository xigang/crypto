package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

//rsa
func RsaEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//rsa
func RsaDecrypt(privateKey []byte, ciphertest []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("pricate key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertest)
}

//sign
func Sign(priv *rsa.PrivateKey, hash crypto.Hash, data []byte) (signed []byte, err error) {
	return rsa.SignPKCS1v15(rand.Reader, priv, hash, data)
}

//unsign
func Unsign(pub *rsa.PublicKey, hash crypto.Hash, hashed, sign []byte) error {
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}
