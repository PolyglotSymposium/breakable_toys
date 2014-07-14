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

type OneFloat64Field struct {
    Answer float64
}

type MockHasAnswerFloat64 struct {}

func (self MockHasAnswerFloat64) AttributeValue(foo string) string {
    if foo == "Answer" {
        return "42.42"
    }
    return ""
}

type OneFloat32Field struct {
    Answer float32
}

type MockHasAnswerFloat32 struct {}

func (self MockHasAnswerFloat32) AttributeValue(foo string) string {
    if foo == "Answer" {
        return "42.42"
    }
    return ""
}

type OneBoolField struct {
    IsCorrect bool
}

type MockHasIsCorrectBool struct {}

func (self MockHasIsCorrectBool) AttributeValue(foo string) string {
    if foo == "IsCorrect" {
        return "true"
    }
    return ""
}
