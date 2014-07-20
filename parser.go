package gosoon

import "errors"

type ParsedJson interface {
    AttributeValue(string) string
    AttributeIsNull(string) bool
}

func Json(rawJson string) (json ParsedJson, err error) {
    if !validBeginning(rune(rawJson[0])) && !validBeginning(rune(rawJson[1])) {
        err = errors.New("Invalid JSON given")
    }

    return
}

func validBeginning(r rune) bool {
    return r == '{' || r == ' ' || r == '\t' || r == '\n'
}
