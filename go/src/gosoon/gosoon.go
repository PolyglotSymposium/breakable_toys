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

    if reflector.NumField() > 0 {
        pair := reflector.Field(0)
        self.Mappings[pair.Name] = pair.Type.Kind()
    }

    if reflector.NumField() > 1 {
        pair := reflector.Field(1)
        self.Mappings[pair.Name] = pair.Type.Kind()
    }
}
