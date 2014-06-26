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
        It("Should return a map with the key and value set appropriately", func() {
            Expect(parsedMap["a"]).To(Equal(Json("\"A\"")))
        })
    })

    Context("When parsing JSON to a string", func() {
        var parsedString string
        Context("When presented with a string surrounding by double-quotes and no extra spaces", func() {
            BeforeEach(func() {
                parsedString = Json("\"6![]i|@nd\"").ParseAsString()
            })
            It("Should parse everything but the outer quotes as the string", func() {
                Expect(parsedString).To(Equal("6![]i|@nd"))
            })
        })
    })
})
