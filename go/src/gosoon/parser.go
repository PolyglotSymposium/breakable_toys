package gosoon

type LazilyParsedJson struct {
}

func Json(rawJson string) LazilyParsedJson {
    return LazilyParsedJson{}
}

func (self LazilyParsedJson) ParseOneLevel() map[string]LazilyParsedJson {
    return nil
}
