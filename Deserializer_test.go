package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

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

        Context("Given a JSON object with no fields and an object that has no fields", func() {
            var noFields NoFields

            BeforeEach(func() {
                noFields = NoFields{}
            })

            It("Should be okay", func() {
                Deserialize(MockEmptyObject{}, &noFields)
            })
        })

        Context("Given a JSON object with no fields and an object that has a field", func() {
            BeforeEach(func() {
                Deserialize(MockEmptyObject{}, &oneStringField)
            })

            It("Should set the object's field to its default value", func() {
                Expect(oneStringField.Phrase).To(Equal(""))
            })
        })

        Context("Given a JSON object with a string field whose attribute matches the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseString{}, &oneStringField)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(oneStringField.Phrase).To(Equal("Phrase's value"))
            })
        })

        Context("Given a JSON object with a string field whose attribute matches one of the databag's two fields", func() {
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

        Context("Given an empty JSON object and a databag with one private field", func() {
            var onePrivateField OnePrivateField

            BeforeEach(func() {
                Deserialize(MockHasPhraseAndNameStrings{}, &onePrivateField)
            })

            It("Should have empty string for that field", func() {
                Expect(onePrivateField.getField()).To(Equal(""))
            })
        })

        Context("Given a JSON object with two string fields whose attributes matches both of the databag's", func() {
            BeforeEach(func() {
                Deserialize(MockHasPhraseAndNameStrings{}, &twoStringFields)
            })

            It("Should have the JSON value for both the matching string fields", func() {
                Expect(twoStringFields.Phrase).To(Equal("Phrase's value"))
                Expect(twoStringFields.Name).To(Equal("Alien Bob"))
            })
        })

        Context("Given a JSON object with an integer field whose attribute matches the databag's", func() {
            var oneIntField OneIntegerField
            BeforeEach(func() {
                Deserialize(MockHasCountInt{}, &oneIntField)
            })

            It("Should have the integer value for the matching field", func() {
                Expect(oneIntField.Count).To(Equal(42))
            })
        })
    })
})
