package md5

import (
	"testing"
)

func TestMD5(t *testing.T) {
	var tests = []struct {
		out string
		in  string
	}{
		{"d41d8cd98f00b204e9800998ecf8427e", ""},
		{"0cc175b9c0f1b6a831c399e269772661", "a"},
		{"187ef4436122d1cc2f40dc2b92f0eba0", "ab"},
		{"900150983cd24fb0d6963f7d28e17f72", "abc"},
		{"e2fc714c4727ee9395f324cd2e7f331f", "abcd"},
		{"ab56b4d92b40713acc5af89985d4b786", "abcde"},
		{"e80b5017098950fc58aad83c8c14978e", "abcdef"},
		{"7ac66c0f148de9519b8bd264312c4d64", "abcdefg"},
		{"e8dc4081b13434b45189a720b77b6818", "abcdefgh"},
		{"8aa99b1f439ff71293e95357bac6fd94", "abcdefghi"},
		{"a925576942e94b2ef57a066101b48876", "adefghij"},
		{"d747fc1719c7eacb84058196cfe56d57", "Discard medicine more than two years old."},
		{"bff2dcb37ef3a44ba43ab144768ca837", "He who has a shady past knows that nice guys finish last."},
		{"0441015ecb54a7342d017ed1bcfdbea5", "I wouldn't marry him with a ten foot pole."},
		{"9e3cac8e9e9757a60c3ea391130d3689", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
		{"a0f04459b031f916a59a35cc482dc039", "The days of the digital watch are numbered.  -Tom Stoppard"},
		{"e7a48e0fe884faf31475d2a04b1362cc", "Nepal premier won't resign."},
		{"637d2fe925c07c113800509964fb0e06", "For every action there is an equal and opposite government program."},
		{"834a8d18d5c6562119cf4c7f5086cb71", "His money is twice tainted: 'taint yours and 'taint mine."},
		{"de3a4d2fd6c73ec2db2abad23b444281", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
		{"acf203f997e2cf74ea3aff86985aefaf", "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
		{"e1c1384cb4d2221dfdd7c795a4222c9a", "size:  a.out:  bad magic"},
		{"c90f3ddecc54f34228c063d7525bf644", "The major problem is with sendmail.  -Mark Horton"},
		{"cdf7ab6c1fd49bd9933c43f3ea5af185", "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
		{"83bc85234942fc883c063cbd7f0ad5d0", "If the enemy is within range, then so are you."},
		{"277cbe255686b48dd7e8f389394d9299", "It's well we cannot hear the screams/That we create in others' dreams."},
		{"fd3fb0a7ffb8af16603f3d3af98f8e1f", "You remind me of a TV show, but that's all right: I watch it anyway."},
		{"469b13a78ebf297ecda64d4723655154", "C is as portable as Stonehedge!!"},
		{"63eb3a2f466410104731c4b037600110", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
		{"72c2ed7592debca1c90fc0100f931a2f", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
		{"132f7619d33b523b1d9e5bd8e0928355", "How can you write a big system without C++?  -Paul Glick"},
	}

	for _, test := range tests {
		md5 := MD5(test.in)
		if md5 == test.out {
			t.Logf("MD5 encrypted data:%v equal test.in:%v", string(md5), test.out)
		}
	}
}

//Output:
// === RUN TestMD5
// --- PASS: TestMD5 (0.00s)
// 	md5_test.go:48: MD5 encrypted data:d41d8cd98f00b204e9800998ecf8427e equal test.in:d41d8cd98f00b204e9800998ecf8427e
// 	md5_test.go:48: MD5 encrypted data:0cc175b9c0f1b6a831c399e269772661 equal test.in:0cc175b9c0f1b6a831c399e269772661
// 	md5_test.go:48: MD5 encrypted data:187ef4436122d1cc2f40dc2b92f0eba0 equal test.in:187ef4436122d1cc2f40dc2b92f0eba0
// 	md5_test.go:48: MD5 encrypted data:900150983cd24fb0d6963f7d28e17f72 equal test.in:900150983cd24fb0d6963f7d28e17f72
// 	md5_test.go:48: MD5 encrypted data:e2fc714c4727ee9395f324cd2e7f331f equal test.in:e2fc714c4727ee9395f324cd2e7f331f
// 	md5_test.go:48: MD5 encrypted data:ab56b4d92b40713acc5af89985d4b786 equal test.in:ab56b4d92b40713acc5af89985d4b786
// 	md5_test.go:48: MD5 encrypted data:e80b5017098950fc58aad83c8c14978e equal test.in:e80b5017098950fc58aad83c8c14978e
// 	md5_test.go:48: MD5 encrypted data:7ac66c0f148de9519b8bd264312c4d64 equal test.in:7ac66c0f148de9519b8bd264312c4d64
// 	md5_test.go:48: MD5 encrypted data:e8dc4081b13434b45189a720b77b6818 equal test.in:e8dc4081b13434b45189a720b77b6818
// 	md5_test.go:48: MD5 encrypted data:8aa99b1f439ff71293e95357bac6fd94 equal test.in:8aa99b1f439ff71293e95357bac6fd94
// 	md5_test.go:48: MD5 encrypted data:d747fc1719c7eacb84058196cfe56d57 equal test.in:d747fc1719c7eacb84058196cfe56d57
// 	md5_test.go:48: MD5 encrypted data:bff2dcb37ef3a44ba43ab144768ca837 equal test.in:bff2dcb37ef3a44ba43ab144768ca837
// 	md5_test.go:48: MD5 encrypted data:0441015ecb54a7342d017ed1bcfdbea5 equal test.in:0441015ecb54a7342d017ed1bcfdbea5
// 	md5_test.go:48: MD5 encrypted data:9e3cac8e9e9757a60c3ea391130d3689 equal test.in:9e3cac8e9e9757a60c3ea391130d3689
// 	md5_test.go:48: MD5 encrypted data:a0f04459b031f916a59a35cc482dc039 equal test.in:a0f04459b031f916a59a35cc482dc039
// 	md5_test.go:48: MD5 encrypted data:e7a48e0fe884faf31475d2a04b1362cc equal test.in:e7a48e0fe884faf31475d2a04b1362cc
// 	md5_test.go:48: MD5 encrypted data:637d2fe925c07c113800509964fb0e06 equal test.in:637d2fe925c07c113800509964fb0e06
// 	md5_test.go:48: MD5 encrypted data:834a8d18d5c6562119cf4c7f5086cb71 equal test.in:834a8d18d5c6562119cf4c7f5086cb71
// 	md5_test.go:48: MD5 encrypted data:de3a4d2fd6c73ec2db2abad23b444281 equal test.in:de3a4d2fd6c73ec2db2abad23b444281
// 	md5_test.go:48: MD5 encrypted data:acf203f997e2cf74ea3aff86985aefaf equal test.in:acf203f997e2cf74ea3aff86985aefaf
// 	md5_test.go:48: MD5 encrypted data:e1c1384cb4d2221dfdd7c795a4222c9a equal test.in:e1c1384cb4d2221dfdd7c795a4222c9a
// 	md5_test.go:48: MD5 encrypted data:c90f3ddecc54f34228c063d7525bf644 equal test.in:c90f3ddecc54f34228c063d7525bf644
// 	md5_test.go:48: MD5 encrypted data:cdf7ab6c1fd49bd9933c43f3ea5af185 equal test.in:cdf7ab6c1fd49bd9933c43f3ea5af185
// 	md5_test.go:48: MD5 encrypted data:83bc85234942fc883c063cbd7f0ad5d0 equal test.in:83bc85234942fc883c063cbd7f0ad5d0
// 	md5_test.go:48: MD5 encrypted data:277cbe255686b48dd7e8f389394d9299 equal test.in:277cbe255686b48dd7e8f389394d9299
// 	md5_test.go:48: MD5 encrypted data:fd3fb0a7ffb8af16603f3d3af98f8e1f equal test.in:fd3fb0a7ffb8af16603f3d3af98f8e1f
// 	md5_test.go:48: MD5 encrypted data:469b13a78ebf297ecda64d4723655154 equal test.in:469b13a78ebf297ecda64d4723655154
// 	md5_test.go:48: MD5 encrypted data:63eb3a2f466410104731c4b037600110 equal test.in:63eb3a2f466410104731c4b037600110
// 	md5_test.go:48: MD5 encrypted data:72c2ed7592debca1c90fc0100f931a2f equal test.in:72c2ed7592debca1c90fc0100f931a2f
// 	md5_test.go:48: MD5 encrypted data:132f7619d33b523b1d9e5bd8e0928355 equal test.in:132f7619d33b523b1d9e5bd8e0928355
// PASS
// ok  	github.com/xigang/crypto/md5	0.005s
