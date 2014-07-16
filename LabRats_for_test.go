package gosoon_test

import (
    "gosoon"
)

type OnePrivateField struct {
    phrase string
}

func (s OnePrivateField) getField() string {
    return s.phrase
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
