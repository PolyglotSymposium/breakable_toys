package gosoon

const (
    JsonArray = iota
    JsonObject
)

type Parser struct { }

func (p *Parser) Parse(jsonText string) JsonNode {
    node := JsonNode{}
    if jsonText == "{}" {
        node._type = JsonObject
    }
	return node
}

type JsonNode struct {
    _type int
}

func NewJsonObject() JsonNode {
    return JsonNode{ _type: JsonObject }
}

func (self JsonNode) ElementCount() int {
    return 0
}

func (self JsonNode) Type() int {
    return self._type
}
