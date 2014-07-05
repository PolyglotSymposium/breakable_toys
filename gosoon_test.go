package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type NoFields struct {}

type OnePrivateField struct {
    phrase string
}

func (s OnePrivateField) getField() string {
    return s.phrase
}

type OneStringField struct {
    Phrase string
}

type TwoStringFields struct {
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

type MockHasPhraseAndNameStrings struct {}

func (self MockHasPhraseAndNameStrings) AttributeValue(foo string) string {
    returnMe := (MockHasPhraseString{}).AttributeValue(foo)
    if foo == "Name" {
        returnMe = "Alien Bob"
    }
    return returnMe
}

var _ = Describe("Gosoon", func() {
    Describe("Deserialize", func() {
        var (
            oneStringField OneStringField
            twoStringFields TwoStringFields
        )
        BeforeEach(func() {
            oneStringField = OneStringField{}
            twoStringFields = TwoStringFields{}
        })

        Context("When given a JSON object with no fields, and an Object that has no fields", func() {
            var noFields NoFields

            BeforeEach(func() {
                noFields = NoFields{}
            })

            It("Should be okay", func() {
                Deserialize(MockEmptyObject{}, &noFields)
            })
        })

        Context("When given a JSON object with no fields, and an Object that has a field", func() {
            BeforeEach(func() {
                Deserialize(MockEmptyObject{}, &oneStringField)
            })

            It("Should have an empty value for the string field", func() {
                Expect(oneStringField.Phrase).To(Equal(""))
            })
        })

        Context("When given a JSON object with a string field whose attribute matches the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &oneStringField)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(oneStringField.Phrase).To(Equal("Phrase's value"))
            })
        })

        Context("When given a JSON object with a string field whose attribute matches one of the databag's (who has 2)", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &twoStringFields)
            })

            It("Should have the JSON value for the matching string field", func() {
                Expect(twoStringFields.Phrase).To(Equal("Phrase's value"))
            })

            It("Should have an empty value for the non-matching string field", func() {
                Expect(twoStringFields.Name).To(Equal(""))
            })
        })

        Context("When given an empty JSON object and a databag with one, private field", func() {
            var onePrivateField OnePrivateField

            BeforeEach(func() {
                Deserialize(MockHasPhraseAndNameStrings{}, &onePrivateField)
            })

            It("Should have empty string for that field", func() {
                Expect(onePrivateField.getField()).To(Equal(""))
            })
        })

        Context("When given a JSON object with two string fields whose attributes matches both of the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseAndNameStrings{}, &twoStringFields)
            })

            It("Should have the JSON value for the matching string field", func() {
                Expect(twoStringFields.Phrase).To(Equal("Phrase's value"))
            })

            It("Should have an empty value for the non-matching string field", func() {
                Expect(twoStringFields.Name).To(Equal("Alien Bob"))
            })
        })
    })
})
