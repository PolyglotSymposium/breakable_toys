package gosoon_test

import (
    "gosoon"
)

type NoFields struct {}

type OnePrivateField struct {
    phrase string
}

func (s OnePrivateField) getField() string {
    return s.phrase
}

type OneStringField struct {
    Phrase string
}

type TwoStringFields struct {
    Phrase string
    Name string
}

type OneIntegerField struct {
    Count int
}

type ParsedJsonMockFactory struct {
    NullValuedAttributes []string
    AttributeToValueMappings map[string]string
}

func (self ParsedJsonMockFactory) Build() gosoon.ParsedJson {
    return ParsedJsonMock{
        attributeToValueMappings: self.AttributeToValueMappings,
        nullValuedAttributes: self.NullValuedAttributes }
}

type ParsedJsonMock struct {
    nullValuedAttributes []string
    attributeToValueMappings map[string]string
}

func (self ParsedJsonMock) AttributeValue(key string) string {
    item, ok := self.attributeToValueMappings[key]

    if ok {
        return item
    }
    return ""
}

func (self ParsedJsonMock) AttributeIsNull(attributeName string) bool {
    for _, nullAttribute := range self.nullValuedAttributes {
        if attributeName == nullAttribute {
            return true
        }
    }
    return false
}

type OneFloat64Field struct {
    Answer float64
}

type OneFloat32Field struct {
    Answer float32
}

type OneBoolField struct {
    IsCorrect bool
}
