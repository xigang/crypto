package crypto

import (
	// "io/ioutil"
	// "os"
	"testing"
)

const (
	BITS = 2048
)

const (
	RSA_PRIVATE_KEY_PATH = "./rsa_private_key.pem"
	RSA_PUBLIC_KEY_PATH  = "./rsa_public_key.pem"
)

// func TestGenRsaKey(t *testing.T) {
// 	err := GenRsaKey(BITS)
// 	if err != nil {
// 		t.Error("generate pub and priv key error.", err)
// 	}
// }

// === RUN TestGetFileContent
// --- PASS: TestGetFileContent (0.00s)
// 	utils_test.go:20: success to load file:  -----BEGIN RSA PRIVATE KEY-----
// 		MIIEpQIBAAKCAQEA1+AXqT6IGuctCBfslEVrhS54CHNpbR3DG1XQAZlwM3lhPUKm
// 		c2gskuNn1HTVKKMbzsJwQuj0J/2TlO44guFxvkuinyojlHt7L282GlBUrFh37Cgn
// 		ywTnDWhQRM9Sc4tHCZ8gm+Cw5vzqXCp3O9zZX+g5gn6WhWtdJLX0zoTegBdeTcJI
// 		WVYqWzgo8OpYout9TSRCzg0tDS+zkLfBqLPah6gAtAVjN0FJq/Attt8E4tX9RVKY
// 		yXUgDFq1yQv1HD5vIvZGPKCn6wBsmsCDKpTrrYvBc+jIM1+49CsCmA35KBiNLFxE
// 		9QWM1eIuxoGsK1cltx1TuqJCMfQH4XY8PnEZWwIDAQABAoIBAQCyHRJIsPBpih4M
// 		lHi9aX2qOEoPZvIwdqL5Vdc2ypvZzKdffTOK+wPz+i88C12oltOlx6YKftycdkuj
// 		UN5NlqdbhFpcCs+qfiuIWrEACxGh1ai0r/kNODFPTacU7x4veVYWaTXlhwFOxJS1
// 		kNAqEzFDQX7Z3Yph8NC8LmkbUGOouJrKF4+CoGSGPqjmZiDqTSHhQKwO+jB0dtYe
// 		FsuGl8TTsZqo3XU1uUm1z0V2ap2YXpsrcekRG8NeMlRRa+CXckNL5vXpJ7Xspjg/
// 		M0F1c7lAJKbT4ufetcSjXGlu2EOuhqJn0Q+5Tu7EQQLiRih8XYpq9CyA76r4svVG
// 		Y1oYbUqBAoGBAN+NkdtROMXrXCJ8w4RoEnFBgTAizB/EQxxUrewC4syvQpiKjc+z
// 		MyOyghxFNgMAEVTmLL0RX6VLzLtAO9LweXDLoLDc6d342UcddnhFCDlILx6sSTo+
// 		XcSLFVBFFNWQ7EzOOsggdp/oUjSdy7mU3Dt8AXx2vQxlNaFg5Wa5UyxBAoGBAPc1
// 		P7IvcYVWbPw8moBuFSh6Y7nZgx4ZH/l7Xr8tPH5Dx0TNv4dEafWX+UBxeaUcm09m
// 		IqkBNhoBkDJVpLIi0Utn7U0SC1HzzL6ksgxVIxu3er3d/ByuKPHysV9Es1ozvtdV
// 		jze1CeeFPgU2XUd+6GDxME4mEqtzfgcW4zgBOc6bAoGAIBTT8rJkQFr3XOeks1Kw
// 		Gtq4rGzB+GkU7HLmhrEfVRptpzBzC3R/CUdrzpzMFIAk/JcPoo5kuHb5SYH9U9c4
// 		xkwiol8BnN5CaCc9QiM9++J1CKTW2Wnr0PHYvC09wSxMDPWnj+LAw0/2xCBFBTvs
// 		owi3ueHcXXQEz0T1htlfq0ECgYEAutYn+ggZiGk1nb5AF2kb3KgDz2cxzgG3RNuV
// 		VHBgFB4t8TZ/10BBxTjDTY12HvBsAapz77/Wn/kmfqR0AZ6HCLLXizRkEeBtknjv
// 		iJqgGkamIfIwpYyLrbF7lUFbPIV3a1BjGfteLPlrUo0wGuUzxHEFzUrBOYvkaH9H
// 		5AApLJkCgYEAnNmolftnW3V6tS837BFQkXcPmWuCxBvs3YvgGr0MzL4eetZk8177
// 		J3UiDn+bkERTyaiXlB7WaglgMQpxEVcOq9trLpL10VOXc5W/CmR5UsHK65ttZ0zP
// 		LM64dVXk0xM7zi1vP5fILXSs5IYHWAJ2r3ihT0DETp7WZp2r7tOHJS0=
// 		-----END RSA PRIVATE KEY-----

// PASS
// ok  	github.com/xigang/crypto	0.006s
// func TestGetFileContent(t *testing.T) {
// 	data, err := GetFileContent(RSA_PRIVATE_KEY_PATH)
// 	if err != nil {
// 		t.Error("failed to load file.", err)
// 	}

// 	t.Log("success to load file: ", string(data))
// }

func TestGetPrividateKey(t *testing.T) {
	data, err := GetFileContent(RSA_PRIVATE_KEY_PATH)
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
	data, err := GetFileContent(RSA_PUBLIC_KEY_PATH)
	if err != nil {
		t.Error("failed to load file.", err)
	}

	pubicKey, err := GetPublicKey(data)
	if err != nil {
		t.Error("get public key error: ", err)
	}

	t.Log("public key:", *pubicKey)
}
