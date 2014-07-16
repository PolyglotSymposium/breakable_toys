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
            parsedJsonMockFactory = ParsedJsonMockFactory{}
        })

        Context("Given an empty JSON object", func() {
            emptyJsonObject := parsedJsonMockFactory.Build()

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
            BeforeEach(func() {
                parsedJsonMockFactory.NullValuedAttributes = make([]string, 0)
            })

            Context("and it has a field that matches databag's private field", func() {
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "phrase": "foobar" }
                })
                var onePrivateField OnePrivateField

                BeforeEach(func() {
                    Deserialize(parsedJsonMockFactory.Build(), &onePrivateField)
                })

                It("Should have empty string for that field", func() {
                    Expect(onePrivateField.getField()).To(Equal(""))
                })
            })

            Context("and it has a string field whose attribute matches the databag's", func() {
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Phrase": "Phrase's value" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneStringField)
                })

                It("Should have the JSON value for the string field", func() {
                    Expect(oneStringField.Phrase).To(Equal("Phrase's value"))
                })
            })

            Context("and it has a float64 field whose attribute matches the databag's", func() {
                var oneFloat64Field OneFloat64Field
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Answer": "42.42" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneFloat64Field)
                })

                It("Should have the JSON value for the float64 field", func() {
                    Expect(oneFloat64Field.Answer).To(Equal(42.42))
                })
            })

            Context("and it has a float32 field whose attribute matches the databag's", func() {
                var oneFloat32Field OneFloat32Field
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Answer": "42.42" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneFloat32Field)
                })

                It("Should have the JSON value for the float32 field", func() {
                    Expect(oneFloat32Field.Answer > 42.41999 && oneFloat32Field.Answer < 42.4200001).To(BeTrue())
                })
            })

            Context("and it has a string field whose attribute matches one of the databag's two fields", func() {
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Phrase": "Phrase's value" }
                    Deserialize(parsedJsonMockFactory.Build(), &twoStringFields)
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
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Phrase": "Phrase's value",
                        "Name": "Alien Bob" }
                    Deserialize(parsedJsonMockFactory.Build(), &twoStringFields)
                })

                It("Should have the JSON value for both the matching string fields", func() {
                    Expect(twoStringFields.Phrase).To(Equal("Phrase's value"))
                    Expect(twoStringFields.Name).To(Equal("Alien Bob"))
                })
            })

            Context("and it has an integer field whose attribute matches the databag's", func() {
                var oneIntField OneIntegerField
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Count": "42" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneIntField)
                })

                It("Should have the integer value for the matching field", func() {
                    Expect(oneIntField.Count).To(Equal(42))
                })
            })

            Context("and it has a true boolean field whose attribute matches the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "IsCorrect": "true" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneBoolField)
                })

                It("Should have the bool value for the matching field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeTrue())
                })
            })

            Context("and it has a false boolean field whose attribute matches the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "IsCorrect": "false" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneBoolField)
                })

                It("Should have the bool value for the matching field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeFalse())
                })
            })

            Context("and it has fields that don't match the databag's", func() {
                var oneBoolField OneBoolField
                BeforeEach(func() {
                    Deserialize(parsedJsonMockFactory.Build(), &oneBoolField)
                })

                It("Should have the default bool value for its bool field", func() {
                    Expect(oneBoolField.IsCorrect).To(BeFalse())
                })
            })
        })

        Context("Given a JSON object with a null string field", func() {
            BeforeEach(func() {
                parsedJsonMockFactory.NullValuedAttributes = []string{ "Phrase" }
            })

            Context("and that null field is conventionally the same as a databag's field", func() {
                BeforeEach(func() {
                    parsedJsonMockFactory.AttributeToValueMappings = map[string]string{
                        "Phrase": "null" }
                    Deserialize(parsedJsonMockFactory.Build(), &oneStringField)
                })

                It("Should have the default string value for that field", func() {
                    Expect(oneStringField.Phrase).To(Equal(""))
                })
            })
        })
    })
})
