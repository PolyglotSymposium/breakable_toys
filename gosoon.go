package gosoon

import (
    "reflect"
    "strconv"
)

type deserializer struct {
    provider ParsedJson
    receiver interface{}
}

type fieldSetter struct {
    value string
    field reflect.Value
}

type fieldNameFunc func(string)

func Deserialize(json ParsedJson, toFill interface{}) {
    deserializer{
        provider: json,
        receiver: toFill }.MapFields()
}

func (self deserializer) MapFields() {
    self.eachFieldName(func(fieldName string) {
        fieldSetter{
            field: self.receiverValue().FieldByName(fieldName),
            value: self.provider.AttributeValue(fieldName) }.set()
    })
}

func (self deserializer) receiverType() reflect.Type {
    return reflect.TypeOf(self.receiver).Elem()
}

func (self deserializer) receiverValue() reflect.Value {
    return reflect.ValueOf(self.receiver).Elem()
}

func (self deserializer) eachFieldName(fn fieldNameFunc) {
    for i := 0; i < self.receiverType().NumField(); i += 1 {
        fn(self.receiverType().Field(i).Name)
    }
}

func (self fieldSetter) set() {
    if self.field.CanSet() {
        if self.field.Kind() == reflect.Int {
            self.field.SetInt(self.valueAsInt64())
        } else {
            self.field.SetString(self.value)
        }
    }
}

func (self fieldSetter) valueAsInt64() int64 {
    number, _ := strconv.Atoi(self.value)
    return int64(number)
}
