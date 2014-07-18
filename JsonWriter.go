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

