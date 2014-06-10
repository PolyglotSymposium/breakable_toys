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

	BeforeEach(func() {
		tokenizer = Tokenizer{
			InitialText: "something",
		}
	})
})
