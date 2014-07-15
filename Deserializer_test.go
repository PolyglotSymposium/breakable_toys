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
            parsedJsonMockFactory ParsedJsonMockFactory
        )
        BeforeEach(func() {
            oneStringField = OneStringField{}
            twoStringFields = TwoStringFields{}
            parsedJsonMockFactory  = ParsedJsonMockFactory{}
        })

        Context("Given an empty JSON object", func() {
            emptyJsonObject := MockEmptyObject{}

            Context("and an object that has no fields", func() {
                var noFields NoFields

                BeforeEach(func() {
                    noFields = NoFields{}
                })

                It("Should be okay", func() {
                    Deserialize(emptyJsonObject, &noFields)
                })
            })

            Context("and an object that has a field", func() {
                BeforeEach(func() {
                    Deserialize(emptyJsonObject, &oneStringField)
                })

                It("Should set the object's field to its default value", func() {
                    Expect(oneStringField.Phrase).To(Equal(""))
                })
            })

        })

        Context("Given a JSON object with no null valued fields", func() {
            parsedJsonMockFactory.NullValuedAttributes = make([]string, 0)

            Context("and it has a field that matches databag's private field", func() {
                var onePrivateField OnePrivateField

                BeforeEach(func() {
                    Deserialize(MockHasPhraseAndNameStrings{}, &onePrivateField)
                })

                It("Should have empty string for that field", func() {
                    Expect(onePrivateField.getField()).To(Equal(""))
                })
            })

            Context("and it has a string field whose attribute matches the databag's", func() {
                BeforeEach(func() {
                    Deserialize(MockHasPhraseString{}, &oneStringField)
                })

                It("Should have the JSON value for the string field", func() {
                    Expect(oneStringField.Phrase).To(Equal("Phrase's value"))
                })
            })

            Context("and it has a float64 field whose attribute matches the databag's", func() {
                var oneFloat64Field OneFloat64Field
                BeforeEach(func() {
                    Deserialize(MockHasAnswerFloat64{}, &oneFloat64Field)
                })

                It("Should have the JSON value for the float64 field", func() {
                    Expect(oneFloat64Field.Answer).To(Equal(42.42))
                })
            })

            Context("and it has a float32 field whose attribute matches the databag's", func() {
                var oneFloat32Field OneFloat32Field
                BeforeEach(func() {
                    Deserialize(MockHasAnswerFloat32{}, &oneFloat32Field)
                })

                It("Should have the JSON value for the float32 field", func() {
                    Expect(oneFloat32Field.Answer > 42.41999 && oneFloat32Field.Answer < 42.4200001).To(BeTrue())
                })
            })

            Context("and it has a string field whose attribute matches one of the databag's two fields", func() {
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

            Context("and it has two string fields whose attributes matches both of the databag's", func() {
                BeforeEach(func() {
                    Deserialize(MockHasPhraseAndNameStrings{}, &twoStringFields)
                })

                It("Should have the JSON value for both the matching string fields", func() {
                    Expect(twoStringFields.Phrase).To(Equal("Phrase's value"))
                    Expect(twoStringFields.Name).To(Equal("Alien Bob"))
                })
            })

            Context("and it has an integer field whose attribute matches the databag's", func() {
                var oneIntField OneIntegerField
                BeforeEach(func() {
                    Deserialize(MockHasCountInt{}, &oneIntField)
                })

                It("Should have the integer value for the matching field", func() {
                    Expect(oneIntField.Count).To(Equal(42))
                })
            })
            Context("and it has a true boolean field whose attribute matches the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    Deserialize(MockHasIsCorrectBool{ truthy: "true" }, &oneBoolField)
                })

                It("Should have the bool value for the matching field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeTrue())
                })
            })
            Context("and it has a false boolean field whose attribute matches the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    Deserialize(MockHasIsCorrectBool{ truthy: "false" }, &oneBoolField)
                })

                It("Should have the bool value for the matching field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeFalse())
                })
            })
            Context("and it has fields that don't match the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    Deserialize(MockEmptyObject{}, &oneBoolField)
                })

                It("Should have the default bool value for its bool field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeFalse())
                })
            })
        })
    })
})
