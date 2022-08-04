package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type EnumValue struct {
	Identifier string
	Number     string
}

type Enum struct {
	noopVisitor
	types *TypeDictionary

	EnumName   string
	EnumValues []*EnumValue
}

func (e *Enum) VisitEnum(enum *parser.Enum) (next bool) {
	e.EnumName = enum.EnumName
	e.types.RegisterEnum(e)
	return true
}

func (e *Enum) VisitEnumField(field *parser.EnumField) (next bool) {
	e.EnumValues = append(e.EnumValues, &EnumValue{
		Identifier: field.Ident,
		Number:     field.Number,
	})
	// Skip comments
	return false
}

func NewEnum(types *TypeDictionary) *Enum {
	return &Enum{
		types: types,
	}
}
