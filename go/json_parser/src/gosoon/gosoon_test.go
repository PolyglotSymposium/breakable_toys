package gosoon_test

import (
	. "gosoon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gosoon", func() {
	var (
		parser Parser
        subject JsonNode
	)

	Describe("the parse method", func() {
		Context("When given an empty JSON array", func() {
			BeforeEach(func() {
                subject = parser.Parse("[]")
			})

			It("Should return a JsonNode with no children", func() {
				Expect(subject.ElementCount()).To(Equal(0))
			})

            It("Should return a JsonNode whose type is an Array", func() {
				Expect(subject.Type()).To(Equal(JsonArray))
            })
		})

		Context("When given an nonempty JSON array", func() {
			BeforeEach(func() {
                subject = parser.Parse("[3]")
			})

			It("Should return a JsonNode with no children", func() {
				Expect(subject.ElementCount()).To(Equal(1))
			})
		})

        Context("When given an empty JSON object", func() {
			BeforeEach(func() {
                subject = parser.Parse("{}")
            })

            It("Should return a JsonNode whose type is an Object", func() {
				Expect(subject.Type()).To(Equal(JsonObject))
            })

            It("Should return an empty JSON object", func() {
                Expect(subject).To(Equal(NewJsonObject()))
            })
        })
	})
})
