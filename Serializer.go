package gosoon

import (
    "fmt"
    "reflect"
)

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) {
    return func(object interface{}) {
        jsonWriter.BeginObject()
        typeOfObject := reflect.TypeOf(object).Elem()
        valueOfObject := reflect.ValueOf(object).Elem()
        for i := 0; i < typeOfObject.NumField(); i += 1 {
            jsonWriter.WriteCommaExceptOnFirstPass()
            field := typeOfObject.Field(i)
            jsonWriter.WriteKey(field.Name)
            kind := field.Type.Kind()
            jsonWriter.WriteValue(kind, AsJsonString(kind, valueOfObject.Field(i)))
        }
        jsonWriter.EndObject()
    }
}

func AsJsonString(kind reflect.Kind, value reflect.Value) string {
    if kind == reflect.String {
        return value.String()
    }
    if kind == reflect.Bool {
        return fmt.Sprint(value.Bool())
    }
    return fmt.Sprint(value.Int())
}
