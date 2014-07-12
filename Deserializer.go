package gosoon

import (
    "reflect"
    "strconv"
)

type deserializer struct {
    provider ParsedJson
    receiver interface{}
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

type fieldSetter struct {
    value string
    field reflect.Value
}

func (self fieldSetter) set() {
    if self.field.CanSet() {
        switch self.field.Kind() {
        case reflect.Int:
            self.field.SetInt(self.valueAsInt64())
        case reflect.Float64, reflect.Float32:
            self.field.SetFloat(self.valueAsFloat64())
        default:
            self.field.SetString(self.value)
        }
    }
}

func (self fieldSetter) valueAsFloat64() float64 {
    float, _ := strconv.ParseFloat(self.value, 64)
    return float64(float)
}

func (self fieldSetter) valueAsInt64() int64 {
    number, _ := strconv.Atoi(self.value)
    return int64(number)
}
