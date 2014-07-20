package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
    var (
        err error
        parsed ParsedJson
    )
    Context("When given brutally invalid JSON", func() {
        BeforeEach(func() {
            parsed, err = Json(`F0oTo^Bar}{kkKkK`)
        })

        It("Should error out", func() {
            Expect(err).To(HaveOccurred())
        })
    })
    Context("When given an empty JSON object", func() {
        BeforeEach(func() {
            parsed, err = Json("{}")
        })

        It("Should not error out", func() {
            Expect(err).NotTo(HaveOccurred())
        })
    })
})
