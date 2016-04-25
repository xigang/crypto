## crypto
[![Coverage Status](https://coveralls.io/repos/github.com/xigang/crypto/badge.svg?branch=master)](https://gocover.io/github.com/xigang/crypto?branch=master)
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


Base64
```

```
