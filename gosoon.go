package gosoon

import (
    "reflect"
    _"fmt"
)

func Deserialize(json ParsedJson, toFill interface{}) {
     reflect.ValueOf(toFill).Elem().FieldByName("Phrase").SetString(json.AttributeValue("Phrase"))
}
