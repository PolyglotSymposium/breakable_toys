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
        self.Mappings["GilliRocks"] = reflect.Bool
    }
}
