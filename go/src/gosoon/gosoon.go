package gosoon

import (
    "strings"
    "reflect"
)

type BlueJson struct {}

func (self BlueJson) Deserialize(json string, toFill interface{}) interface{} {
    withoutCurlies := json[1:len(json)-1]

    fields := strings.FieldsFunc(withoutCurlies, func(r rune) bool {
        return r == ','
    })

    for i := range fields {
        pairs := strings.FieldsFunc(fields[i], func(r rune) bool {
            return r == ','
        })

        typ := reflect.TypeOf(toFill)

    }

    return toFill
}
