package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type TestBag struct {
    Count int
    Phrase string
}

var _ = Describe("Gosoon", func() {
    var (
        subject TestBag
    )

    Describe("BlueJson", func() {
        Describe(".Deserialize", func() {
            Context("When given an empty JSON array", func() {
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
        })
    })
})
