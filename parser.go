package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

func Json(rawJson string) (json ParsedJson, err error) {
    if !validBeginningRune(rune(rawJson[0])) {
        err = errors.New("Invalid JSON given")
    }

    return
}

func validBeginningRune(r rune) bool {
    return r == '{' || r == ' ' || r == '\t'
}
