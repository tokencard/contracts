package externals_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("base64", func() {

	Context("When a valid base64 encoded string is passed", func() {

		It("Should decode 'Zg==' to 'f'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zg=="))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("f")))
		})

		It("Should decode 'Zm8=' to 'fo'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zm8="))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("fo")))
		})

		It("Should decode 'Zm9v' to 'foo'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zm9v"))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("foo")))
		})

		It("Should decode 'Zm9vYg==' to 'foob'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zm9vYg=="))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("foob")))
		})

		It("Should decode 'Zm9vYmE=' to 'fooba'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zm9vYmE="))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("fooba")))
		})

		It("Should decode 'Zm9vYmFy' to 'foobar'", func() {
			dec, err := Base64Exporter.Base64decode(nil, []byte("Zm9vYmFy"))
			Expect(err).ToNot(HaveOccurred())
			Expect(dec).To(Equal([]byte("foobar")))
		})

	})

	Context("When an invalid base64 encoded string is passed", func() {

		It("Should fail (trigger require)", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte(""))
			Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("Z"))
			Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("zZ="))
			Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
		})

		It("Should fail (trigger require)'", func() {
			_, err := Base64Exporter.Base64decode(nil, []byte("zZz=="))
			Expect(err).To(MatchError(errors.New("abi: improperly formatted output")))
		})

	})
})
