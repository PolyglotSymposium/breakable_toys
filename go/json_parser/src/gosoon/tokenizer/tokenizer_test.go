package tokenizer_test

import (
	. "gosoon/tokenizer"

	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
)



var _ = Describe("Tokenizer", func() {
	var (
		tokenizer Tokenizer
	)

	Describe(".tok", func() {
		Context("With a colon", func() {
			BeforeEach(func() {
				tokenizer = Tokenizer{
					InitialText: ":",
				}
			})

			It("should return tokens containing only a Colon", func() {
				Fail("TODO")
			})
		})
	})
})
