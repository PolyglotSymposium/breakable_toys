package gosoon

import (
    "reflect"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    typeOfToFill := reflect.TypeOf(toFill).Elem()

    for i := 0; i < typeOfToFill.NumField(); i += 1 {
        fieldName := typeOfToFill.Field(i).Name
        reflect.ValueOf(toFill).
            Elem().
            FieldByName(fieldName).
            SetString(json.AttributeValue(fieldName))
    }
}
