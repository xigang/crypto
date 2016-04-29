package base64

import (
	"testing"
)

func TestBase64(t *testing.T) {
	var tests = []struct {
		decoded string
		encoded string
	}{
		// RFC 3548 examples
		{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
		{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
		{"\x14\xfb\x9c\x03", "FPucAw=="},

		// RFC 4648 examples
		{"", ""},
		{"f", "Zg=="},
		{"fo", "Zm8="},
		{"foo", "Zm9v"},
		{"foob", "Zm9vYg=="},
		{"fooba", "Zm9vYmE="},
		{"foobar", "Zm9vYmFy"},

		// Wikipedia examples
		{"sure.", "c3VyZS4="},
		{"sure", "c3VyZQ=="},
		{"sur", "c3Vy"},
		{"su", "c3U="},
		{"leasure.", "bGVhc3VyZS4="},
		{"easure.", "ZWFzdXJlLg=="},
		{"asure.", "YXN1cmUu"},
		{"sure.", "c3VyZS4="},
	}

	for _, test := range tests {
		encodeData := Base64Encode([]byte(test.decoded))

		if string(encodeData) == test.encoded {
			t.Logf("The encoded data is the same as the data that needs to be matched.")
			t.Logf("encoded data:%v equal test.encoded:%v ", string(encodeData), test.encoded)
		}
	}

	for _, test := range tests {
		decodeData, err := Base64Decode(test.encoded)
		if err != nil {
			t.Error(err)
		}

		if string(decodeData) == test.decoded {
			t.Logf("The same data is the same as the matching data through the base64 decoding.")
			t.Logf("decoded data:%v equal test.decode:%v", string(decodeData), test.decoded)
		}
	}
}

//Output:
// === RUN TestBase64
// --- PASS: TestBase64 (0.00s)
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:FPucA9l+ equal test.encoded:FPucA9l+
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:FPucA9k= equal test.encoded:FPucA9k=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:FPucAw== equal test.encoded:FPucAw==
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data: equal test.encoded:
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zg== equal test.encoded:Zg==
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zm8= equal test.encoded:Zm8=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zm9v equal test.encoded:Zm9v
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zm9vYg== equal test.encoded:Zm9vYg==
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zm9vYmE= equal test.encoded:Zm9vYmE=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:Zm9vYmFy equal test.encoded:Zm9vYmFy
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:c3VyZS4= equal test.encoded:c3VyZS4=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:c3VyZQ== equal test.encoded:c3VyZQ==
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:c3Vy equal test.encoded:c3Vy
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:c3U= equal test.encoded:c3U=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:bGVhc3VyZS4= equal test.encoded:bGVhc3VyZS4=
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:ZWFzdXJlLg== equal test.encoded:ZWFzdXJlLg==
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:YXN1cmUu equal test.encoded:YXN1cmUu
// 	base64_test.go:41: The encoded data is the same as the data that needs to be matched.
// 	base64_test.go:42: encoded data:c3VyZS4= equal test.encoded:c3VyZS4=
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:???~ equal test.decode:???~
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:??? equal test.decode:???
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:?? equal test.decode:??
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data: equal test.decode:
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:f equal test.decode:f
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:fo equal test.decode:fo
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:foo equal test.decode:foo
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:foob equal test.decode:foob
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:fooba equal test.decode:fooba
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:foobar equal test.decode:foobar
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:sure. equal test.decode:sure.
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:sure equal test.decode:sure
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:sur equal test.decode:sur
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:su equal test.decode:su
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:leasure. equal test.decode:leasure.
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:easure. equal test.decode:easure.
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:asure. equal test.decode:asure.
// 	base64_test.go:53: The same data is the same as the matching data through the base64 decoding.
// 	base64_test.go:54: decoded data:sure. equal test.decode:sure.
// PASS
// ok  	github.com/xigang/crypto/base64	0.004s
