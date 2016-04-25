package base64

import (
	"encoding/base64"
)

//Base64Encode return Base64 encrypted data
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//base64Decode return Base64 decryption data
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
