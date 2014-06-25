package gosoon

import (
    _"strings"
    "reflect"
)

type BlueJson struct {
    Mappings map[string]reflect.Kind
}

func (self BlueJson) Deserialize(json string, toFill interface{}) interface{} {
    return toFill
}

func (self *BlueJson) Inspect(me interface{}) {
    reflector := reflect.TypeOf(me)

    self.Mappings = map[string]reflect.Kind{}

    for i := reflector.NumField()-1; i >= 0; i -= 1 {
        field := reflector.Field(i)
        self.Mappings[field.Name] = field.Type.Kind()
    }
}
