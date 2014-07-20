package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
    var (
        toParse string
        err error
        parsed ParsedJson
    )
    BeforeEach(func() {
        parsed, err = Json(toParse)
    })

    Context("When given brutally invalid JSON", func() {
        toParse = `F0oTo^Bar}{kkKkK`

        It("Should error out", func() {
            Expect(err).To(HaveOccurred())
        })
    })
})
