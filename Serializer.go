package gosoon

import (
    "reflect"
)

type JsonWriter interface {
    BeginObject()
    EndObject()
    WriteKey(string)
    WriteValue(kind reflect.Kind, value string)
    WriteCommaExceptOnFirstPass()
}

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) {
    return func(object interface{}) {
        jsonWriter.BeginObject()
        for i := 0; i < reflect.TypeOf(object).Elem().NumField(); i += 1 {
            jsonWriter.WriteCommaExceptOnFirstPass()
            jsonWriter.WriteKey(reflect.TypeOf(object).Elem().Field(i).Name)
            jsonWriter.WriteValue(reflect.String, "")
        }
        jsonWriter.EndObject()
    }
}
