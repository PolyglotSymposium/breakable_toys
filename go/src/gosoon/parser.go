package gosoon

import "strings"

type LazilyParsedJson struct {
    unparsedJson string
}

func Json(rawJson string) LazilyParsedJson {
    return LazilyParsedJson{unparsedJson: rawJson}
}

func (self LazilyParsedJson) ParseOneLevel() map[string]LazilyParsedJson {
    parsedMap := make(map[string]LazilyParsedJson)
    keyvalue := strings.Split(self.unparsedJson, ":")
    if len(keyvalue) > 1 {
        parsedMap[keyvalue[0]] = Json(keyvalue[1])
    }
    return parsedMap
}
