package gosoon_test

import (
    "reflect"
)

type MockJsonWriter struct {
    json string
    pairSeparator string
}

func (self *MockJsonWriter) BeginObject() {
    self.json += "{"
}

func (self *MockJsonWriter) WriteKey(key string) {
    self.json += key
}

func (self *MockJsonWriter) WriteValue(kind reflect.Kind, value string) {
    if kind == reflect.String {
        self.json += `"` + value + `"`
    } else {
        self.json += value
    }
}

func (self *MockJsonWriter) WriteCommaExceptOnFirstPass() {
    self.json += self.pairSeparator
    self.pairSeparator = ","
}

func (self *MockJsonWriter) EndObject() {
    self.json += "}"
}

