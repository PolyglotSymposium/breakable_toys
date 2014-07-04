package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type NoAttributes struct {}

type OneStringAttribute struct {
    Phrase string
}

type TwoStringAttributes struct {
    Phrase string
    Name string
}

type MockEmptyObject struct {}

func (self MockEmptyObject) AttributeValue(foo string) string {
    return ""
}

type MockHasPhraseString struct {}

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
            twoStringAttributes TwoStringAttributes
        )
        BeforeEach(func() {
            oneStringAttribute = OneStringAttribute{}
            twoStringAttributes = TwoStringAttributes{}
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

        Context("When given a JSON object with a string field whose attribute matches the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &oneStringAttribute)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(oneStringAttribute.Phrase).To(Equal("Phrase's value"))
            })
        })

        Context("When given a JSON object with a string field whose attribute matches one of the databag's (who has 2)", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &twoStringAttributes)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(twoStringAttributes.Phrase).To(Equal("Phrase's value"))
            })
        })
    })
})
