package gosoon

import (
    "reflect"
    "strconv"
)

func Deserialize(json ParsedJson, toFill interface{}) {
    typeOfToFill := reflectType(toFill)

    for i := 0; i < typeOfToFill.NumField(); i += 1 {
        fieldName := typeOfToFill.Field(i).Name
        field := reflectValue(toFill).FieldByName(fieldName)

        if field.CanSet() {
            if field.Kind() == reflect.Int {
                field.SetInt(stringToInt64(json.AttributeValue(fieldName)))
            } else {
                field.SetString(json.AttributeValue(fieldName))
            }
        }
    }
}

func stringToInt64(value string) int64 {
    number, _ := strconv.Atoi(value)
    return int64(number)
}

func reflectValue(ofMe interface{}) reflect.Value {
    return reflect.ValueOf(ofMe).Elem()
}

func reflectType(ofMe interface{}) reflect.Type {
    return reflect.TypeOf(ofMe).Elem()
}
