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
    Context("Given brutally invalid JSON", func() {
        It("Should error out", func() {
            _, err = Json(`F0oTo^Bar}{kkKkK`)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given an empty string", func() {
        It("Should error out", func() {
            _, err = Json("")
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given JSON that beings correctly, but goes o' so badly", func() {
        It("Should error out for '{'", func() {
            _, err = Json("{")
            Expect(err).To(HaveOccurred())
        })
        It("Should error out for '{<ws>'", func() {
            _, err = Json("{\n")
            Expect(err).To(HaveOccurred())
        })
        It("Should error out for '{a'", func() {
            _, err = Json("{a")
            Expect(err).To(HaveOccurred())
        })
        It("Should error out for '{}a'", func() {
            _, err = Json("{}a")
            Expect(err).To(HaveOccurred())
        })
        It("Should error out for '{}<ws>a'", func() {
            _, err = Json("{} a")
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given JSON with an invalid key", func() {
        It(`Should error out for '{a:""}'`, func() {
            _, err = Json(`{a:""}`)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given an empty JSON object", func() {
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
        It("Should not error out for '{<ws>}'", func() {
            _, err = Json("{\t}")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for '{}<ws>'", func() {
            _, err = Json("{}\n")
            Expect(err).NotTo(HaveOccurred())
        })
        It("Should not error out for '{}<ws><ws>'", func() {
            _, err = Json("{}\n\t")
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Context(`Given a JSON object with a key of "", and a value of null`, func() {
        It(`Should not error out for '{"":null}'`, func() {
            _, err = Json(`{"":null}`)
            Expect(err).NotTo(HaveOccurred())
        })
        It(`Should not error out for '{<ws>"":null}'`, func() {
            _, err = Json(`{ "":null}`)
            Expect(err).NotTo(HaveOccurred())
        })
        Describe("AttributeIsNull", func() {
            It(`Should be true for ""`, func() {
                parsed, err = Json(`{"":null}`)
                Expect(parsed.AttributeIsNull("")).To(BeTrue())
            })
        })
    })
})
