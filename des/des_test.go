package des

import (
	"testing"

	"github.com/xigang/crypto/base64"
)

//output:
//=== RUN TestDes
// --- PASS: TestDes (0.00s)
// 	des_test.go:17: VQxt7qbs8wNU1G42hS1qQqyGFWwUDR98z5m/XjlDPiI=
// 	des_test.go:24: wangxigang2014@gmail.com
// PASS
// ok  	github.com/xigang/crypto/des	0.004s
func TestDes(t *testing.T) {
	str := "wangxigang2014@gmail.com"
	key := []byte("sfe023f_")

	result, err := DesEncrypt([]byte(str), key)
	if err != nil {
		t.Error("des encrypt data error: ", err)
	}
	t.Log(base64.Base64Encode(result))

	oridData, err := DesDecrypt(result, key)
	if err != nil {
		t.Error("des decrypt data error: ", err)
	}

	t.Log(string(oridData))
}

//output:
//=== RUN Test3Des
// --- PASS: Test3Des (0.00s)
// 	des_test.go:43: zJVj3lLyZJF5O/tR2LX/trr8bXZXCFpm
// 	des_test.go:49: wangxgang201@gmail.com
// PASS
// ok  	github.com/xigang/crypto/des	0.004s
func Test3Des(t *testing.T) {
	str := "wangxgang201@gmail.com"

	key := []byte("sfe023f_sefiel#fi32lf3e!")

	result, err := TripleDesEncrypt([]byte(str), key)
	if err != nil {
		t.Error("des encrypt data error.")
	}
	t.Log(base64.Base64Encode(result))

	origData, err := TripleDesDecrypt(result, key)
	if err != nil {
		t.Error("des decrypt data error.")
	}
	t.Log(string(origData))
}
