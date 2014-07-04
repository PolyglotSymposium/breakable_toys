package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type OneStringAttribute struct {
    Phrase string
}

type NoAttributes struct {
}

type MockEmptyObject struct {
}

func (self MockEmptyObject) AttributeValue(foo string) string {
    return ""
}

type MockHasPhraseString struct {
}

func (self MockHasPhraseString) AttributeValue(foo string) string {
    if foo == "Phrase" {
        return "Phrase's value"
    }
    return ""
}

var _ = Describe("Gosoon", func() {
    Describe("Deserialize", func() {
        var (
            oneStringAttribute OneStringAttribute
        )
        BeforeEach(func() {
            oneStringAttribute = OneStringAttribute{}
        })

        Context("When given a JSON object with no fields, and an Object that has no fields", func() {
            var noAttributes NoAttributes

            BeforeEach(func() {
                noAttributes = NoAttributes{}
            })

            It("Should be okay", func() {
                Deserialize(MockEmptyObject{}, &noAttributes)
            })
        })

        Context("When given a JSON object with no fields, and an Object that has a field", func() {
            BeforeEach(func() {
                Deserialize(MockEmptyObject{}, &oneStringAttribute)
            })

            It("Should have an empty value for the string field", func() {
                Expect(oneStringAttribute.Phrase).To(Equal(""))
            })
        })

        Context("When given a JSON object with a string field (1 char), whose attribute matches the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &oneStringAttribute)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(oneStringAttribute.Phrase).To(Equal("Phrase's value"))
            })
        })
    })
})
