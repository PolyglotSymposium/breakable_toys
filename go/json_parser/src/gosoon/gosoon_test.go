package gosoon_test

import (
	. "gosoon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gosoon", func() {
	var (
		parser Parser
        jsonText string
	)

	Describe("the parse method", func() {
		Context("When given an empty JSON array", func() {
			BeforeEach(func() {
                jsonText = "[]"
			})

			It("Should return a JsonObject with no children", func() {
				Expect(parser.Parse(jsonText).ElementCount()).To(Equal(0))
			})

            It("Should return a JsonObject whose type is an Array", func() {
				Expect(parser.Parse(jsonText).Type()).To(Equal("Array"))
            })
		})

        Context("When given an empty JSON object", func() {
            It("Should return an empty JSON object", func() {
                Expect(parser.Parse("{}")).To(Equal(JsonObject{}))
            })
        })
	})
})
