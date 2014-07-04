package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type OneStringAttribute struct {
    Phrase string
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
    Describe(".Deserialize", func() {
        Context("When given a JSON object with a string field (1 char), whose attribute matches the databag's", func() {
            var (
                subject OneStringAttribute
            )

            BeforeEach(func() {
                subject = OneStringAttribute{}
                Deserialize(MockHasPhraseString{}, &subject)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(subject.Phrase).To(Equal("Phrase's value"))
            })
        })
    })
})
