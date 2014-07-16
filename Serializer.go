package gosoon

type JsonWriter interface {
}

func MakeSerializer(jsonWriter JsonWriter) func(interface{}) string {
    return func(object interface{}) string {
        return "{}"
    }
}
