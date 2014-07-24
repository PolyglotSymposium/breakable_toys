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
            if field.Type.Kind() == reflect.String {
                jsonWriter.WriteValue(field.Type.Kind(), "")
            } else if field.Type.Kind() == reflect.Bool {
                jsonWriter.WriteValue(field.Type.Kind(), fmt.Sprint(valueOfObject.Field(i).Bool()))
            } else {
                jsonWriter.WriteValue(field.Type.Kind(), fmt.Sprint(valueOfObject.Field(i).Int()))
            }
        }
        jsonWriter.EndObject()
    }
}
