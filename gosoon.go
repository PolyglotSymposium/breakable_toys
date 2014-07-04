package gosoon

import (
    "reflect"
    _"fmt"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    if reflect.TypeOf(toFill).Elem().NumField() > 0 {
        reflect.ValueOf(toFill).
            Elem().
            FieldByName("Phrase").
            SetString(json.AttributeValue("Phrase"))
    }
}
