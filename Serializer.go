package gosoon

import (
    "fmt"
    "reflect"
)

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) {
    writeOneKeyValuePair := func(field reflect.StructField, value reflect.Value) {
        jsonWriter.WriteCommaExceptOnFirstPass()
        jsonWriter.WriteKey(field.Name)
        kind := field.Type.Kind()
        jsonWriter.WriteValue(kind, stringify(kind, value))
    }
    return func(object interface{}) {
        jsonWriter.BeginObject()
        typeOfObject := reflect.TypeOf(object).Elem()
        valueOfObject := reflect.ValueOf(object).Elem()
        for i := 0; i < typeOfObject.NumField(); i += 1 {
            writeOneKeyValuePair(typeOfObject.Field(i), valueOfObject.Field(i))
        }
        jsonWriter.EndObject()
    }
}

func stringify(kind reflect.Kind, value reflect.Value) string {
    if kind == reflect.String {
        return value.String()
    }
    if kind == reflect.Bool {
        return fmt.Sprint(value.Bool())
    }
    return fmt.Sprint(value.Int())
}
