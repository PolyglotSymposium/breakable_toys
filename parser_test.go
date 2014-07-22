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

var _ = Describe("JsonString", func() {
    Context("Given an empty string with no JSON", func() {
        It("Should error out", func() {
            _, err := JsonString("")
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given a string with just a double-quote", func() {
        It("Should error out", func() {
            _, err := JsonString(`"`)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given a string with just a double-quote that does not have a closing quote", func() {
        It("Should error out", func() {
            _, err := JsonString(`"no ending double quote`)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("Given an empty JSON string", func() {
        It("Should not error out", func() {
            _, err := JsonString(`""`)
            Expect(err).ToNot(HaveOccurred())
        })
        It("Should return an empty string", func() {
            value, _ := JsonString(`""`)
            Expect(value).To(Equal(""))
        })
    })

    Context("Given a JSON string with unicode characters in it", func() {
        It("Should not error out", func() {
            _, err := JsonString(`"omega=Ω"`)
            Expect(err).ToNot(HaveOccurred())
        })
        It(`Should return the correct one character string for "<u-char>"`, func() {
            value, _ := JsonString(`"Ω"`)
            Expect(value).To(Equal("Ω"))
        })
    })

    Context("Given a JSON string with simple characters in it", func() {
        It("Should not error out", func() {
            _, err := JsonString(`" asdg asdg a"`)
            Expect(err).ToNot(HaveOccurred())
        })
        It(`Should return the correct one character string for "<char>"`, func() {
            value, _ := JsonString(`"k"`)
            Expect(value).To(Equal("k"))
        })
        It(`Should return the correct two characters string for "<char><char>"`, func() {
            value, _ := JsonString(`"ke"`)
            Expect(value).To(Equal("ke"))
        })
        It(`Should return the correct three characters string for "<char><char><char>"`, func() {
            value, _ := JsonString(`"kei"`)
            Expect(value).To(Equal("kei"))
        })
    })
})
