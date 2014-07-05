package gosoon

import (
    "reflect"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    typeOfToFill := reflectType(toFill)

    for i := 0; i < typeOfToFill.NumField(); i += 1 {
        fieldName := typeOfToFill.Field(i).Name
        field := reflectValue(toFill).FieldByName(fieldName)

        if field.CanSet() {
            field.SetString(json.AttributeValue(fieldName))
        }
    }
}

func reflectValue(ofMe interface{}) reflect.Value {
    return reflect.ValueOf(ofMe).Elem()
}

func reflectType(ofMe interface{}) reflect.Type {
    return reflect.TypeOf(ofMe).Elem()
}
