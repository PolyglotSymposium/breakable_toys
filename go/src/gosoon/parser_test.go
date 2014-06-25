package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Lazy JSON parser", func() {
    var parsedMap map[string]LazilyParsedJson

    Context("when passed an empty JSON object", func() {
        BeforeEach(func() {
            parsedMap = Json("{}").ParseOneLevel()
        })
        It("Should return an empty map", func() {
            Expect(len(parsedMap)).To(Equal(0))
        })
    })

    Context("when passed an JSON object with one key-value pair", func() {
        BeforeEach(func() {
            parsedMap = Json("{\"a\":\"A\"}").ParseOneLevel()
        })
        It("Should return a map with one lazily parsed hunk of JSON", func() {
            Expect(len(parsedMap)).To(Equal(1))
        })
    })
})
