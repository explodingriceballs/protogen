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

func (e *Enum) DeclaresMessageType(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Enum) DeclaresEnumType(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Enum) GetName() string {
	return e.EnumName
}

func (e *Enum) GetType() ElementType {
	return EnumElementType
}

func (e *Enum) VisitEnum(enum *parser.Enum) (next bool) {
	e.EnumName = enum.EnumName
	e.types.startScope(e)
	e.types.RegisterEnum(e)
	e.types.endScope(e)
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
