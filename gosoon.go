package gosoon

import (
    "reflect"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    typeOfToFill := reflect.TypeOf(toFill).Elem()

    if typeOfToFill.NumField() > 0 {
        fieldName := typeOfToFill.Field(0).Name
        reflect.ValueOf(toFill).
            Elem().
            FieldByName(fieldName).
            SetString(json.AttributeValue(fieldName))
    }
}
