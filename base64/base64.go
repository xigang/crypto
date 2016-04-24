package base64

import (
	"encoding/base64"
)

//Base64Encode
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//base64Decode
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
