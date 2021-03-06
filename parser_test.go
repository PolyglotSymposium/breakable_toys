package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
    var (
        err error
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

    Context("Given a JSON object with at string key and a string field and some whitespace between tokens", func() {
        parsedJson, _ := Json("{   \"Golang's\"  : \t  \"pretty fun\"\n }")

        Describe("AttributeValue", func() {
            It("Should be that attribute's value", func() {
                Expect(parsedJson.AttributeValue("Golang's")).To(Equal("pretty fun"))
            })
        })
    })

    Context("Given a JSON object with a string key, and a string field", func() {
        var parsedJson ParsedJson
        var err error

        BeforeEach(func() {
            parsedJson, err = Json(`{"Kazark":"The Man"}`)
        })

        It("Should not error out", func() {
            Expect(err).NotTo(HaveOccurred())
        })

        Describe("AttributeValue", func() {
            Context("Given a key other than that of the object", func() {
                It("Should return an empty string", func() {
                    Expect(parsedJson.AttributeValue("")).To(Equal(""))
                })
            })
            Context("Given the key that is in the object", func() {
                It("Should return the string's value", func() {
                    Expect(parsedJson.AttributeValue("Kazark")).To(Equal("The Man"))
                })
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
        It(`Should return the correct two character string for "<char><u-char>"`, func() {
            value, _ := JsonString(`"kΩ"`)
            Expect(value).To(Equal("kΩ"))
        })
    })

    Context("Given a JSON string with an escaped character in it", func() {
        It("Should not error out", func() {
            _, err := JsonString(`" asdg \" asdg a"`)
            Expect(err).ToNot(HaveOccurred())
        })
        It("Should return a string containing the quotation mark", func() {
            value, _ := JsonString(`" This is a \" test"`)
            Expect(value).To(Equal(` This is a " test`))
        })
        It("Should return a string containing the newline", func() {
            value, _ := JsonString(`" This is a \n test"`)
            Expect(value).To(Equal(" This is a \n test"))
        })
        It("Should return a string containing the tab", func() {
            value, _ := JsonString(`" This is a \t test"`)
            Expect(value).To(Equal(" This is a \t test"))
        })
        It("Should return a string containing the return", func() {
            value, _ := JsonString(`" This is a \r test"`)
            Expect(value).To(Equal(" This is a \r test"))
        })
        It("Should return a string containing the formfeed", func() {
            value, _ := JsonString(`" This is a \f test"`)
            Expect(value).To(Equal(" This is a \f test"))
        })
        It("Should return a string containing the backspace", func() {
            value, _ := JsonString(`" This is a \b test"`)
            Expect(value).To(Equal(" This is a \b test"))
        })
        It("Should return a string containing the solidus", func() {
            value, _ := JsonString(`" This is a \∕ test"`)
            Expect(value).To(Equal(" This is a ∕ test"))
        })
        It("Should return a string containing the reverse solidus", func() {
            value, _ := JsonString(`" This is a \\ test"`)
            Expect(value).To(Equal(" This is a \\ test"))
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
