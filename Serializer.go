package gosoon

import (
    "reflect"
)

type JsonWriter interface {
    BeginObject()
    EndObject()
    WriteKey(string)
    WriteValue(kind reflect.Kind, value string)
}

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) {
    return func(object interface{}) {
        jsonWriter.BeginObject()
        if reflect.TypeOf(object).Elem().NumField() == 1 {
            jsonWriter.WriteKey(reflect.TypeOf(object).Elem().Field(0).Name)
            jsonWriter.WriteValue(reflect.String, "")
        }
        jsonWriter.EndObject()
    }
}
