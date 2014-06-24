package gosoon

import (
    _"strings"
    _"reflect"
)

type BlueJson struct {}

func (self BlueJson) Deserialize(json string, toFill interface{}) interface{} {
    return toFill
}
