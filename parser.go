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
    keyvalue := strings.Split(self.unparsedJson[1:len(self.unparsedJson)-1], ":")
    if len(keyvalue) > 1 {
        unparsedKey := keyvalue[0]
        parsedMap[Json(unparsedKey).ParseAsString()] = Json(keyvalue[1])
    }
    return parsedMap
}

func (self LazilyParsedJson) ParseAsString() string {
    trimmed := strings.TrimSpace(self.unparsedJson)
    return trimmed[1:len(trimmed)-1]
}
