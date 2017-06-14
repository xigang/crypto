package base64

import (
	"encoding/base64"
)

//Base64EncodeString return Base64 encrypted data
func Base64EncodeString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//Base64DecodeString return Base64 decryption data
func Base64DecodeString(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

//Base64EncodeURL returns the base64 encoding of src.
func Base64EncodeURL(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

//Base64DecodeURL returns the bytes represented by the base64 string s.
func Base64DecodeURL(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

//TODO User-defined encryption.
