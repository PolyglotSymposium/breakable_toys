package gosoon_test

type NoFields struct {}

type OnePrivateField struct {
    phrase string
}

func (s OnePrivateField) getField() string {
    return s.phrase
}

type OneStringField struct {
    Phrase string
}

type TwoStringFields struct {
    Phrase string
    Name string
}

type OneIntegerField struct {
    Count int
}

type MockEmptyObject struct {}

func (self MockEmptyObject) AttributeValue(foo string) string {
    return ""
}

type MockHasCountInt struct {}

func (self MockHasCountInt) AttributeValue(foo string) string {
    if foo == "Count" {
        return "42"
    }
    return ""
}

type MockHasPhraseString struct {}

func (self MockHasPhraseString) AttributeValue(foo string) string {
    if foo == "Phrase" {
        return "Phrase's value"
    }
    return ""
}

type MockHasPhraseAndNameStrings struct {}

func (self MockHasPhraseAndNameStrings) AttributeValue(foo string) string {
    returnMe := (MockHasPhraseString{}).AttributeValue(foo)
    if foo == "Name" {
        returnMe = "Alien Bob"
    }
    return returnMe
}

