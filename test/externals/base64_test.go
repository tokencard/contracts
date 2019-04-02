package externals_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "fmt"

)

//Golang  tests: https://github.com/golang/go/blob/master/src/encoding/base64/base64_test.go

type testpair struct {
	decoded, encoded string
}

var pairs = []testpair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	// {"", ""}, //excluded because we demand that the return string is not the empty string in the solidity code
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

var nonCorrupt = []testpair{
	// {"", "\n"},
	// {"\x00\x00","AAA=\n"},
	// {"\x00\x00\x00","AAAA\n"},
	{"\x00", "AA=="},
	{"\x00\x00", "AAA="},
	{"\x00\x00\x00", "AAAA"},
}

// var corrupt  = []string{
// 	"!!!!",
// 	"====",
// 	"x===",
// 	"=AAA",
// 	"A=AA",
// 	"AA=A",
// 	"AA==A",
// 	"AAA=AAAA",
// 	"AAAAA",
// 	"AAAAAA",
// 	"A=",
// 	"A==",
// 	"AA=",
// 	"AAAAAA=",
// 	"YWJjZA=====",
// 	"A!\n",
// 	"A=\n",
// }


//JS: https://github.com/mathiasbynens/base64/blob/master/tests/tests.js

var jsValidPairs = []testpair{
	{"a", "YQ=="},
	{"aa", "YWE="},
	{"aaa", "YWFh"},
	// {"a", "YQ"},
	// {"a", "YR"},
	{"foo bar baz", "Zm9vIGJhciBiYXo="},
	{"foo bar", "Zm9vIGJhcg=="},
	{"foo\x00", "Zm9vAA=="},
	{"foo\x00\x00", "Zm9vAAA="},
	{"i\xB7\x1D", "abcd"},
	{"i", "ab=="},
}

var jsInvalidPairs = []string{
	// RFC 3548 examples
	// "AAECA\t\n\f\r wQFBgcICQoLDA0ODx\t\n\f\r AREhMUFRYXGBkaGxwdHh8gIS\t\n\f\r IjJCUmJygpKissLS4vMDEyMzQ1Njc4OT\t\n\f\r o7PD0+P0BBQkNERUZHSElKS0xNT\t\n\f\r k9QUVJTVFVWV1hZWltcXV5fY\t\n\f\r GFiY2RlZmdoaWprbG\t\n\f\r 1ub3BxcnN0dXZ3eH\t\n\f\r l6e3x9fn8=",
	"YQ===", //invalid char
	"abcd===",
	// "abcd ===",
	"A",
}

var _ = Describe("base64", func() {

	When("a valid base64 encoded string is passed", func() {

		It("Should decode All possible octets", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("AAECAwQFBgcICQoLDA0ODxAREhMUFRYXGBkaGxwdHh8gISIjJCUmJygpKissLS4vMDEyMzQ1Njc4OTo7PD0+P0BBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWltcXV5fYGFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6e3x9fn8="))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x20\x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f")))
		})

		for _, p := range pairs {
			p:=p
			It("Should succeed", func() {
				dec, err := Base64Exporter.Base64decode(nil, []byte(p.encoded))
				Expect(err).ToNot(HaveOccurred())
				Expect(dec).To(Equal([]byte(p.decoded)))
			})
		}

		for _, p := range jsValidPairs {
			p:=p
			It("Should succeed", func() {
				dec, err := Base64Exporter.Base64decode(nil, []byte(p.encoded))
				Expect(err).ToNot(HaveOccurred())
				Expect(dec).To(Equal([]byte(p.decoded)))
			})
		}

		for _, p := range jsInvalidPairs {
			p:=p
			It("Should succeed", func() {
				_, err := Base64Exporter.Base64decode(nil, []byte(p))
				Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
			})
		}

	})

	// When("there is corruption", func() {
	// 	for _, c := range corrupt {
	// 		c:=c
	// 		It("Should fail", func() {
	// 			_, err := Base64Exporter.Base64decode(nil, []byte(c))
	// 			if err == nil{
	// 				fmt.Println("no error: " + c)
	// 			}else{
	// 				fmt.Println("error: " + c)
	// 			}
	// 			Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
	// 		})
	// 	}
	// })

	When("there is no corruption", func() {
		for _, nc := range nonCorrupt {
			nc:=nc
			It("Should succeed", func() {
				dec, err := Base64Exporter.Base64decode(nil, []byte(nc.encoded))
				Expect(err).ToNot(HaveOccurred())
				Expect(dec).To(Equal([]byte(nc.decoded)))
			})
		}
	})

	When("an invalid base64 encoded string is passed", func() {

		It("Should fail (trigger require)", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte(""))
			Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("Z"))
			Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("zZ="))
			Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("zZz=="))
			Expect(err.Error()).To(ContainSubstring("invalid base64 encoding"))
		})

	})
})
