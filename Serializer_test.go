package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Gosoon serializer", func() {
    Context("Given an object with no fields", func() {
        var json string
        BeforeEach(func() {
            json = Serialize(NoFields{})
        })
        It("Should serialize it as an empty JSON object", func() {
            Expect(json).To(Equal("{}"))
        })
    })
})
