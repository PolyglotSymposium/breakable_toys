package gosoon

import (
    "reflect"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    typeOfToFill := reflect.TypeOf(toFill).Elem()

    for i := 0; i < typeOfToFill.NumField(); i += 1 {
        fieldName := typeOfToFill.Field(i).Name
        field := reflect.ValueOf(toFill).
            Elem().
            FieldByName(fieldName)
        if field.CanSet() {
            field.SetString(json.AttributeValue(fieldName))
        }
    }
}
