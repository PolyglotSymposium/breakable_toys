package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type MockJsonWriter struct {
}

var _ = Describe("Gosoon serializer", func() {
    var serialize func(interface{})string
    BeforeEach(func() {
        serialize = MakeSerializer(MockJsonWriter{})
    })
    Context("Given an object with no fields", func() {
        var json string
        BeforeEach(func() {
            json = serialize(NoFields{})
        })
        It("Should serialize it as an empty JSON object", func() {
            Expect(json).To(Equal("{}"))
        })
    })
})
