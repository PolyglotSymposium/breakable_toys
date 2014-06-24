package gosoon

type BlueJson struct {}

func (self BlueJson) Deserialize(json string, toFill interface{}) interface{} {
    return toFill
}
