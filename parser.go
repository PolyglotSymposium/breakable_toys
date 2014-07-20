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

type condition func(string) bool

func Json(rawJson string) (json ParsedJson, err error) {
    json, err = parser{unparsedJson: rawJson}.parse()
    return
}

func (self parser) parse() (json ParsedJson, err error) {

    self.swallowWhitespaceUntil(func(json string) bool {
        return rune(json[0]) == '{'
    }, "must begin with '{'")

    self.removeCurrentRune()

    self.swallowWhitespaceUntil(func(json string) bool {
        return rune(json[0]) == '}'
    }, "must end with '}'")

    self.removeCurrentRune()

    self.swallowRemainder()

    err = self.savedError
    return
}

func (self *parser) swallowWhitespaceUntil(cond condition, errorMessage string) {
    self.swallowWhitespace()
    if self.savedError == nil && (!self.runesRemain() || !cond(self.unparsedJson)) {
        self.savedError = errors.New("Invalid JSON given, " + errorMessage)
    }
}

func (self *parser) swallowRemainder() {
    self.swallowWhitespace()

    if self.runesRemain() {
        self.savedError = errors.New("runes should come after the final '}'")
    }
}

func (self *parser) swallowWhitespace() {
    for self.savedError == nil && self.runesRemain() && self.isWhiteSpace(self.currentRune()) {
        self.removeCurrentRune()
    }
}

func (self parser) runesRemain() bool {
    return len(self.unparsedJson) != 0
}

func (self parser) currentRune() rune {
    return rune(self.unparsedJson[0])
}

func (self *parser) removeCurrentRune() {
    if self.runesRemain() {
        self.unparsedJson = (self.unparsedJson)[1:len(self.unparsedJson)]
    }
}

func (self *parser) isWhiteSpace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n'
}
