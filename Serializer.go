package gosoon

type JsonWriter interface {
    BeginObject()
    EndObject()
}

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) {
    return func(object interface{}) {
        jsonWriter.BeginObject()
        jsonWriter.EndObject()
    }
}
