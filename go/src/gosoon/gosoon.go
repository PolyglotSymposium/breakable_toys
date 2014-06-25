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

    i := reflector.NumField()

    if i > 0 {
        pair := reflector.Field(0)
        self.Mappings[pair.Name] = pair.Type.Kind()
    }

    if i > 1 {
        pair := reflector.Field(1)
        self.Mappings[pair.Name] = pair.Type.Kind()
    }
}
