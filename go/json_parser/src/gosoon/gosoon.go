package gosoon

type Parser struct { }

func (p *Parser) Parse(jsonText string) JsonObject {
	return JsonObject{}
}

type JsonObject struct {
}

func (self JsonObject) ElementCount() int {
    return 0;
}

func (self JsonObject) Type() string {
    return "Array"
}
