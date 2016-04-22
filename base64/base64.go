package base64

import (
	"encoding/base64"
)

//Base64Encode编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//base64解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
