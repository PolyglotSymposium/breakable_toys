package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

func Json(rawJson string) (json ParsedJson, err error) {
    if len(rawJson) < 2 {
        err = errors.New("Invalid JSON given, must be an object or array")
        return
    }

    stripWhitespace(&rawJson)

    if rune(rawJson[0]) != '{' {
        err = errors.New("Invalid JSON given, must begin with '{'")
    }

    return
}

func stripWhitespace(fromMe *string) {
    for isWhiteSpace(rune((*fromMe)[0])) {
        removeCharacter(fromMe)
    }
}

func removeCharacter(fromMe *string) {
    *fromMe = (*fromMe)[1:len(*fromMe)]
}

func isWhiteSpace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n'
}
