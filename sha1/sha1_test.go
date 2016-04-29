package sha1

import (
	"fmt"
	"testing"
)

func TestSHA1(t *testing.T) {
	var tests = []struct {
		out string
		in  string
	}{
		{"da39a3ee5e6b4b0d3255bfef95601890afd80709", ""},
		{"86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", "a"},
		{"da23614e02469a0d7c7bd1bdab5c9c474b1904dc", "ab"},
		{"a9993e364706816aba3e25717850c26c9cd0d89d", "abc"},
		{"81fe8bfe87576c3ecb22426f8e57847382917acf", "abcd"},
		{"03de6c570bfe24bfc328ccd7ca46b76eadaf4334", "abcde"},
		{"1f8ac10f23c5b5bc1167bda84b833e5c057a77d2", "abcdef"},
		{"2fb5e13419fc89246865e7a324f476ec624e8740", "abcdefg"},
		{"425af12a0743502b322e93a015bcf868e324d56a", "abcdefgh"},
		{"c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1", "abcdefghi"},
		{"d68c19a0a345b7eab78d5e11e991c026ec60db63", "abcdefghij"},
		{"ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7", "Discard medicine more than two years old."},
		{"e5dea09392dd886ca63531aaa00571dc07554bb6", "He who has a shady past knows that nice guys finish last."},
		{"45988f7234467b94e3e9494434c96ee3609d8f8f", "I wouldn't marry him with a ten foot pole."},
		{"55dee037eb7460d5a692d1ce11330b260e40c988", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
		{"b7bc5fb91080c7de6b582ea281f8a396d7c0aee8", "The days of the digital watch are numbered.  -Tom Stoppard"},
		{"c3aed9358f7c77f523afe86135f06b95b3999797", "Nepal premier won't resign."},
		{"6e29d302bf6e3a5e4305ff318d983197d6906bb9", "For every action there is an equal and opposite government program."},
		{"597f6a540010f94c15d71806a99a2c8710e747bd", "His money is twice tainted: 'taint yours and 'taint mine."},
		{"6859733b2590a8a091cecf50086febc5ceef1e80", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
		{"514b2630ec089b8aee18795fc0cf1f4860cdacad", "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
		{"c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69", "size:  a.out:  bad magic"},
		{"74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a", "The major problem is with sendmail.  -Mark Horton"},
		{"0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66", "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
		{"3ae7937dd790315beb0f48330e8642237c61550a", "If the enemy is within range, then so are you."},
		{"410a2b296df92b9a47412b13281df8f830a9f44b", "It's well we cannot hear the screams/That we create in others' dreams."},
		{"841e7c85ca1adcddbdd0187f1289acb5c642f7f5", "You remind me of a TV show, but that's all right: I watch it anyway."},
		{"163173b825d03b952601376b25212df66763e1db", "C is as portable as Stonehedge!!"},
		{"32b0377f2687eb88e22106f133c586ab314d5279", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
		{"0885aaf99b569542fd165fa44e322718f4a984e0", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
		{"6627d6904d71420b0bf3886ab629623538689f45", "How can you write a big system without C++?  -Paul Glick"},
	}

	for _, test := range tests {
		s := fmt.Sprintf("%x", Sha1([]byte(test.in)))
		if s == test.out {
			t.Logf("sha1 encrypt data:%v equal test.out:%v", s, test.out)
		}
	}
}

//Output:
// === RUN TestSHA1
// --- PASS: TestSHA1 (0.00s)
// 	sha1_test.go:50: sha1 encrypt data:da39a3ee5e6b4b0d3255bfef95601890afd80709 equal test.out:da39a3ee5e6b4b0d3255bfef95601890afd80709
// 	sha1_test.go:50: sha1 encrypt data:86f7e437faa5a7fce15d1ddcb9eaeaea377667b8 equal test.out:86f7e437faa5a7fce15d1ddcb9eaeaea377667b8
// 	sha1_test.go:50: sha1 encrypt data:da23614e02469a0d7c7bd1bdab5c9c474b1904dc equal test.out:da23614e02469a0d7c7bd1bdab5c9c474b1904dc
// 	sha1_test.go:50: sha1 encrypt data:a9993e364706816aba3e25717850c26c9cd0d89d equal test.out:a9993e364706816aba3e25717850c26c9cd0d89d
// 	sha1_test.go:50: sha1 encrypt data:81fe8bfe87576c3ecb22426f8e57847382917acf equal test.out:81fe8bfe87576c3ecb22426f8e57847382917acf
// 	sha1_test.go:50: sha1 encrypt data:03de6c570bfe24bfc328ccd7ca46b76eadaf4334 equal test.out:03de6c570bfe24bfc328ccd7ca46b76eadaf4334
// 	sha1_test.go:50: sha1 encrypt data:1f8ac10f23c5b5bc1167bda84b833e5c057a77d2 equal test.out:1f8ac10f23c5b5bc1167bda84b833e5c057a77d2
// 	sha1_test.go:50: sha1 encrypt data:2fb5e13419fc89246865e7a324f476ec624e8740 equal test.out:2fb5e13419fc89246865e7a324f476ec624e8740
// 	sha1_test.go:50: sha1 encrypt data:425af12a0743502b322e93a015bcf868e324d56a equal test.out:425af12a0743502b322e93a015bcf868e324d56a
// 	sha1_test.go:50: sha1 encrypt data:c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1 equal test.out:c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1
// 	sha1_test.go:50: sha1 encrypt data:d68c19a0a345b7eab78d5e11e991c026ec60db63 equal test.out:d68c19a0a345b7eab78d5e11e991c026ec60db63
// 	sha1_test.go:50: sha1 encrypt data:ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7 equal test.out:ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7
// 	sha1_test.go:50: sha1 encrypt data:e5dea09392dd886ca63531aaa00571dc07554bb6 equal test.out:e5dea09392dd886ca63531aaa00571dc07554bb6
// 	sha1_test.go:50: sha1 encrypt data:45988f7234467b94e3e9494434c96ee3609d8f8f equal test.out:45988f7234467b94e3e9494434c96ee3609d8f8f
// 	sha1_test.go:50: sha1 encrypt data:55dee037eb7460d5a692d1ce11330b260e40c988 equal test.out:55dee037eb7460d5a692d1ce11330b260e40c988
// 	sha1_test.go:50: sha1 encrypt data:b7bc5fb91080c7de6b582ea281f8a396d7c0aee8 equal test.out:b7bc5fb91080c7de6b582ea281f8a396d7c0aee8
// 	sha1_test.go:50: sha1 encrypt data:c3aed9358f7c77f523afe86135f06b95b3999797 equal test.out:c3aed9358f7c77f523afe86135f06b95b3999797
// 	sha1_test.go:50: sha1 encrypt data:6e29d302bf6e3a5e4305ff318d983197d6906bb9 equal test.out:6e29d302bf6e3a5e4305ff318d983197d6906bb9
// 	sha1_test.go:50: sha1 encrypt data:597f6a540010f94c15d71806a99a2c8710e747bd equal test.out:597f6a540010f94c15d71806a99a2c8710e747bd
// 	sha1_test.go:50: sha1 encrypt data:6859733b2590a8a091cecf50086febc5ceef1e80 equal test.out:6859733b2590a8a091cecf50086febc5ceef1e80
// 	sha1_test.go:50: sha1 encrypt data:514b2630ec089b8aee18795fc0cf1f4860cdacad equal test.out:514b2630ec089b8aee18795fc0cf1f4860cdacad
// 	sha1_test.go:50: sha1 encrypt data:c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69 equal test.out:c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69
// 	sha1_test.go:50: sha1 encrypt data:74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a equal test.out:74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a
// 	sha1_test.go:50: sha1 encrypt data:0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66 equal test.out:0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66
// 	sha1_test.go:50: sha1 encrypt data:3ae7937dd790315beb0f48330e8642237c61550a equal test.out:3ae7937dd790315beb0f48330e8642237c61550a
// 	sha1_test.go:50: sha1 encrypt data:410a2b296df92b9a47412b13281df8f830a9f44b equal test.out:410a2b296df92b9a47412b13281df8f830a9f44b
// 	sha1_test.go:50: sha1 encrypt data:841e7c85ca1adcddbdd0187f1289acb5c642f7f5 equal test.out:841e7c85ca1adcddbdd0187f1289acb5c642f7f5
// 	sha1_test.go:50: sha1 encrypt data:163173b825d03b952601376b25212df66763e1db equal test.out:163173b825d03b952601376b25212df66763e1db
// 	sha1_test.go:50: sha1 encrypt data:32b0377f2687eb88e22106f133c586ab314d5279 equal test.out:32b0377f2687eb88e22106f133c586ab314d5279
// 	sha1_test.go:50: sha1 encrypt data:0885aaf99b569542fd165fa44e322718f4a984e0 equal test.out:0885aaf99b569542fd165fa44e322718f4a984e0
// 	sha1_test.go:50: sha1 encrypt data:6627d6904d71420b0bf3886ab629623538689f45 equal test.out:6627d6904d71420b0bf3886ab629623538689f45
// PASS
// ok  	github.com/xigang/crypto/sha1	0.009s
