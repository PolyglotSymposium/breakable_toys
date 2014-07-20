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
        It("Should not error out for '{}'", func() {
            _, err = Json("{}")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for ' {}'", func() {
            _, err = Json(" {}")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for '<tab>{}'", func() {
            _, err = Json("\t{}")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for '<newline>{}'", func() {
            _, err = Json("\n{}")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for '<ws><ws>{}'", func() {
            _, err = Json("\n {}")
            Expect(err).NotTo(HaveOccurred())
        })
        XIt("Should not error out for '{<ws>}'", func() {
            _, err = Json("{\t}")
            Expect(err).NotTo(HaveOccurred())
        })
    })
})
