package gosoon_test

import (
    . "gosoon"

    "reflect"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type MockJsonWriter struct {
    json string
}

func (self *MockJsonWriter) BeginObject() {
    self.json += "{"
}

func (self *MockJsonWriter) WriteKey(key string) {
    self.json += key
}

func (self *MockJsonWriter) WriteValue(kind reflect.Kind, value string) {
    if kind == reflect.String {
        self.json += `""`
    }
}

func (self *MockJsonWriter) EndObject() {
    self.json += "}"
}

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
    Context("Given an object with one blank string field", func() {
        BeforeEach(func() {
            serialize(&struct{ Foo string }{})
        })
        It("Should serialize it as JSON object with one field", func() {
            Expect(mock.json).To(Equal(`{Foo""}`))
        })
    })
})
