package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Lazy JSON parser", func() {
    Context("when passed an empty JSON object", func() {
        var parsedMap map[string]LazilyParsedJson
        BeforeEach(func() {
            parsedMap = Json("{}").ParseOneLevel()
        })
        It("Should return an empty map of strings to lazily parsed hunks of JSON", func() {
            Expect(len(parsedMap)).To(Equal(0))
        })
    })
})
