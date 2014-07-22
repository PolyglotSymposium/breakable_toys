package gosoon

import (
    "errors"
)

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

type parser struct {
    unparsedJson string
    savedError error
    result ParsedJson
}

type parserResult struct {
}

type condition func(string) bool

func Json(rawJson string) (json ParsedJson, err error) {
    json, err = parser{unparsedJson: rawJson}.parse()
    return
}

func JsonString(rawJson string) (stringsValue string, err error) {
    if len(rawJson) < 2 {
        err = errors.New("Invalid JSON string")
    }

    for i := 1; i < len(rawJson); i += 1 {
        if []rune(rawJson)[i] == '"' {
            return
        }
        if []rune(rawJson)[i] == '\\' {
            if []rune(rawJson)[i+1] == '"' {
                i += 1
            } else {
                stringsValue += map[rune]string{
                    'n': "\n",
                    'r': "\r",
                    'f': "\f",
                    'b': "\b",
                    '∕': "∕",
                    't': "\t" }[[]rune(rawJson)[i+1]]
                i += 2
            }
        }
        stringsValue += string([]rune(rawJson)[i])
    }

    return "", errors.New("Invald JSON string, must end in double quote")
}

func (self parser) parse() (json ParsedJson, err error) {
    self.result = parserResult{}

    self.swallowWhitespaceUntil(func(json string) bool {
        return rune(json[0]) == '{'
    }, "must begin with '{'")

    self.removeCurrentRune()

    self.swallowWhitespace()

    if self.runesRemain() && self.currentRune() == '"' {
        for self.currentRune() != '}' {
            self.removeCurrentRune()
        }
    }

    self.swallowWhitespaceUntil(func(json string) bool {
        return rune(json[0]) == '}'
    }, "must end with '}'")

    self.removeCurrentRune()

    self.swallowRemainder()

    err = self.savedError
    json = self.result
    return
}

func (self parserResult) AttributeIsNull(attribute string) bool {
    return true
}

func (self parserResult) AttributeValue(attribute string) string {
    return ""
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
        self.savedError = errors.New("runes should not come after the final '}'")
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
