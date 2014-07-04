package gosoon

import (
    "reflect"
    _"fmt"
)

func Deserialize(json ParsedJson, toFill interface{}) interface{} {

     mutable := reflect.ValueOf(toFill).Elem()
     mutable.FieldByName("Phrase").SetString(json.AttributeValue("Phrase"))

    return toFill
}
