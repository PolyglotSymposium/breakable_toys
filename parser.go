package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

func Json(rawJson string) (json ParsedJson, err error) {
    stripWhitespace(&rawJson)

    if rune(rawJson[0]) != '{' {
        err = errors.New("Invalid JSON given")
    }

    return
}

func stripWhitespace(fromMe *string) {
    for isWhiteSpace(rune((*fromMe)[0])) {
        *fromMe = (*fromMe)[1:len(*fromMe)]
    }
}

func isWhiteSpace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n'
}
