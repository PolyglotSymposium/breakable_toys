package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Gosoon serializer", func() {
    var serialize func(interface{})
    var mock MockJsonWriter
    BeforeEach(func() {
        mock = MockJsonWriter{}
        serialize = MakeSerializer(&mock)
    })
    Context("Given an object with no fields", func() {
        BeforeEach(func() {
            serialize(&struct{}{})
        })
        It("Should serialize it as an empty JSON object", func() {
            Expect(mock.json).To(Equal("{}"))
        })
    })

    Context("Given an object with one defaulted string field", func() {
        BeforeEach(func() {
            serialize(&struct{ Foo string }{})
        })
        It("Should serialize it as JSON object with one field", func() {
            Expect(mock.json).To(Equal(`{Foo""}`))
        })
    })

    Context("Given an object with two defaulted string fields", func() {
        BeforeEach(func() {
            serialize(&struct{ Foo string; Bar string }{})
        })
        It("Should serialize it as JSON object with two fields", func() {
            Expect(mock.json).To(Equal(`{Foo"",Bar""}`))
        })
    })

    Context("Given an object with a defaulted integer field", func() {
        BeforeEach(func() {
            serialize(&struct{ Foo int }{})
        })
        It("Should serialize it as JSON object with one field, set to zero", func() {
            Expect(mock.json).To(Equal(`{Foo0}`))
        })
    })

    Context("Given an object with a defaulted boolean field", func() {
        BeforeEach(func() {
            serialize(&struct{ Foo bool }{})
        })
        It("Should serialize it as JSON object with one field, set to false", func() {
            Expect(mock.json).To(Equal(`{Foofalse}`))
        })
    })

    Context("Given an object with a boolean field set to true", func() {
        BeforeEach(func() {
            object := struct{ Foo bool }{}
            object.Foo = true
            serialize(&object)
        })
        It("Should serialize it as JSON object with one field, set to true", func() {
            Expect(mock.json).To(Equal(`{Footrue}`))
        })
    })
})
