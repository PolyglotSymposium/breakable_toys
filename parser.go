package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

func Json(rawJson string) (json ParsedJson, err error) {
    for isWhiteSpace(rune(rawJson[0])) {
        rawJson = rawJson[1:len(rawJson)]
    }

    if rune(rawJson[0]) != '{' {
        err = errors.New("Invalid JSON given")
    }

    return
}

func isWhiteSpace(r rune) bool {
    return r == ' ' || r == '\t' || r == '\n'
}
