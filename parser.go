package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

type parser struct {
    unparsedJson string
    savedError error
}

func Json(rawJson string) (json ParsedJson, err error) {
    json, err = parser{unparsedJson: rawJson}.parse()
    return
}

func (self parser) parse() (json ParsedJson, err error) {
    self.stripWhitespace()

    if len(self.unparsedJson) == 0 || rune(self.unparsedJson[0]) != '{' {
        err = errors.New("Invalid JSON given, must begin with '{'")
        return
    }

    self.removeCharacter()



    self.stripWhitespace()

    if len(self.unparsedJson) == 0 || rune(self.unparsedJson[0]) != '}' {
        err = errors.New("Invalid JSON given, must end with '}'")
    }
    return
}

func (self *parser) stripWhitespace() {
    for len(self.unparsedJson) != 0 && self.isWhiteSpace(rune((self.unparsedJson)[0])) {
        self.removeCharacter()
    }
}

func (self *parser) removeCharacter() {
    self.unparsedJson = (self.unparsedJson)[1:len(self.unparsedJson)]
}

func (self *parser) isWhiteSpace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n'
}
