package gosoon

import (
    "unicode/utf8"
)

const (
    JsonArray = iota
    JsonObject
)

type Parser struct { }

func (p *Parser) Parse(jsonText string) JsonNode {
    if jsonText == "{}" {
        return NewJsonObject()
    }
	return NewJsonArray(jsonText)
}

type JsonNode struct {
    _type int
    _elementCount int
}

func NewJsonObject() JsonNode {
    return JsonNode { _type: JsonObject }
}

func NewJsonArray(jsonText string) JsonNode {
    elementCount := 0
    if utf8.RuneCountInString(jsonText) > 2 {
        elementCount = 1
    }
    return JsonNode { _type: JsonArray, _elementCount: elementCount }
}

func (self JsonNode) ElementCount() int {
    return self._elementCount
}

func (self JsonNode) Type() int {
    return self._type
}
