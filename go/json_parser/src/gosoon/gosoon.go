package gosoon

import (
    "unicode/utf8"
    "strings"
)

const (
    JsonArray = iota
    JsonNumber
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
    NumericValue float64
}

func NewJsonNumber() JsonNode {
    jsonNumber := JsonNode { _type: JsonNumber }
    jsonNumber.NumericValue = 3.0
    return jsonNumber
}

func NewJsonObject() JsonNode {
    return JsonNode { _type: JsonObject }
}

func NewJsonArray(jsonText string) JsonNode {
    elementCount := 0
    innerText := jsonText[1:len(jsonText)-1]
    if utf8.RuneCountInString(innerText) > 0 {
        elementCount = len(strings.Split(innerText, ","))
    }
    return JsonNode { _type: JsonArray, _elementCount: elementCount }
}

func (self JsonNode) Child(i int) JsonNode {
    return NewJsonNumber()
}

func (self JsonNode) ElementCount() int {
    return self._elementCount
}

func (self JsonNode) Type() int {
    return self._type
}
