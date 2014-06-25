package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "reflect"
)

type Blank struct {}

type OneAttribute struct {
    GilliRocks bool
}

type TestBag struct {
    Count int
    Phrase string
}

var _ = Describe("Gosoon", func() {
    Describe("BlueJson", func() {
        Describe(".Inspect", func() {
            Context("when given an object with no attbitutes", func() {
                blueJson := BlueJson{}

                BeforeEach(func() {
                    blueJson.Inspect(Blank{})
                })

                Describe(".Mappings", func() {
                    It("Should return an empty set of mappings", func() {
                        Expect(blueJson.Mappings).To(Equal(map[string]reflect.Type{}))
                    })
                })
            })
        })
        Describe(".Deserialize", func() {
            Context("When given an empty JSON array", func() {
                var (
                    subject TestBag
                )

                BeforeEach(func() {
                    subject = (BlueJson{}).Deserialize("{}", TestBag{}).(TestBag)
                })

                It("Should have the default value for its integer field", func() {
                    Expect(subject.Count).To(Equal(0))
                })

                It("Should have the default value for its string field", func() {
                    Expect(subject.Phrase).To(Equal(""))
                })
            })

            Context("When given a JSON object, but serializing to an object with no attributes", func() {
                var (
                    subject Blank
                )

                BeforeEach(func() {
                    subject = (BlueJson{}).Deserialize("{ \"Phrase\": \"a\" }", Blank{}).(Blank)
                })

                It("Should return a blank object", func() {
                    Expect(subject).To(Equal(Blank{}))
                })
            })

            Context("When none of the properties on the JSON match the object's properties", func() {
                var (
                    subject OneAttribute
                )

                BeforeEach(func() {
                    subject = (BlueJson{}).Deserialize("{ \"Phrase\": \"a\" }", OneAttribute{}).(OneAttribute)
                })

                It("Should return a blank object", func() {
                    Expect(subject).To(Equal(OneAttribute{}))
                })
            })


            Context("When given a JSON object with a string field (1 char), whose attribute matches the databag's", func() {
                var (
                    subject TestBag
                )

                BeforeEach(func() {
                    subject = (BlueJson{}).Deserialize("{ \"Phrase\": \"a\" }", TestBag{}).(TestBag)
                })

                XIt("Should have the JSON value for the string field", func() {
                    Expect(subject.Phrase).To(Equal("a"))
                })
            })
        })
    })
})
