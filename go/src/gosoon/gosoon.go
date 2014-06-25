package gosoon

import (
    _"strings"
    "reflect"
)

type BlueJson struct {
    Mappings map[string]reflect.Type
}

func (self BlueJson) Deserialize(json string, toFill interface{}) interface{} {
    return toFill
}

func (self *BlueJson) Inspect(me interface{}) {
    self.Mappings = map[string]reflect.Type{}
}
