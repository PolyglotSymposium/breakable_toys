package gosoon_test

import (
	. "gosoon"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Gosoon", func() {
	var (
		parser Parser
	)

	Describe("the parse method", func() {
		Context("When given an empty JSON array", func() {
			BeforeEach(func() {
				parser = Parser {
					InitialText: "[]",
				}
			})

			It("Should return an empty slice of JSON objects", func() {
				Fail("TODO")
			})
		})
	})
})
