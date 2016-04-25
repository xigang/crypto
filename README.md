## crypto
[![Coverage Status](https://coveralls.io/repos/github/crypto/badge.svg?branch=master)](https://gocover.io/github.com/xigang/crypto?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/xigang/crypto)](https://goreportcard.com/report/github.com/xigang/crypto)
[![GoDoc](https://godoc.org/github.com/xigang/crypto?status.svg)](https://godoc.org/github.com/xigang/crypto)

Go supplementary cryptography libaraies.

Installation
-----------

Use go get 

```go
go get github.com/xigang/crypto
``` 

Usage
------

##### Aes

```go
    func TestAes(t *testing.T) {
    str := "wangxgang201@gmail.com"

    key := make([]byte, 16)
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
        return
    }

    result, err := AesEncrypt([]byte(str), key)
    if err != nil {
        t.Error("aes encrypt data error.", err)
    }
    t.Log(base64.Base64Encode(result))

    origData, err := AesDecrypt(result, key)
    if err != nil {
        t.Error("aes decrypt data error.")
    }
    t.Log(string(origData))
}
```


##### Base64

```go
    func TestBase64(t *testing.T) {
        src := "wangxigang2014@gmail.com"
        base64Encode := Base64Encode([]byte(src))
        t.Log("base64Encode:", string(base64Encode))

        base64Decode, err := Base64Decode(base64Encode)
        if err != nil {
            t.Error("error: ", err)
        }
        t.Log("success: ", string(base64Decode))
    }
```


##### DES

```go
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
```

##### MD5

```go
    func TestMd5(t *testing.T) {
        str := "wangxigang2014@gmail.com"
        hashData := MD5([]byte(str))
        t.Log("md5 data: ", hashData)
    }
```


##### RSA 

```go
    func TestRsa(t *testing.T) {
        str := "wangxigang2014@gmail.com"

        publicKey, err := crypto.GetFileContent(RsaPublicKeyPath)
        if err != nil {
            t.Error("read public key error: ", err)
        }

        ciphertest, err := RsaEncrypt(publicKey, []byte(str))
        if err != nil {
            t.Error("rsa encrypt data error.")
        }
        t.Log(string(ciphertest))

        privateKey, err := crypto.GetFileContent(RsaPrivateKeyPath)
        if err != nil {
            t.Error("read private key error: ", err)
        }

        origData, err := RsaDecrypt(privateKey, ciphertest)
        if err != nil {
            t.Error("rsa decrypt data error: ", err)
        }

        t.Log(string(origData))
    }
```


##### SHA1

```go
    func TestSha1(t *testing.T) {
        str := "wangxigang2014@gmail.com"
        data := Sha1([]byte(str))
        t.Log("sha1 data: ", string(data))
    }
    ```


    ##### SHA256

    ```go
    func TestSha256(t *testing.T) {
        str := "wangxigang2014@gmail.com"
        data := Sha256([]byte(str))
        t.Log("sha256 data:", string(data))
    }
```

