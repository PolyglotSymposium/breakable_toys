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

func Deserialize(json ParsedJson, toFill interface{}) {
    deserializer{
        provider: json,
        receiver: toFill }.MapFields()
}

func (self deserializer) MapFields() {
    for i := 0; i < self.receiverType().NumField(); i += 1 {
        self.mapFieldByIndex(i)
    }
}

func (self deserializer) receiverType() reflect.Type {
    return reflect.TypeOf(self.receiver).Elem()
}

func (self deserializer) receiverValue() reflect.Value {
    return reflect.ValueOf(self.receiver).Elem()
}

func (self deserializer) mapFieldByIndex(index int) {
    fieldName := self.receiverType().Field(index).Name
    fieldSetter{
        field: self.receiverValue().FieldByName(fieldName),
        value: self.provider.AttributeValue(fieldName) }.set()
}

func (self fieldSetter) set() {
    if self.field.CanSet() {
        if self.field.Kind() == reflect.Int {
            number, _ := strconv.Atoi(self.value)
            self.field.SetInt(int64(number))
        } else {
            self.field.SetString(self.value)
        }
    }
}
