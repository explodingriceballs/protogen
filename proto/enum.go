package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type EnumValue struct {
	identifier string
	number     string
}

func (e *EnumValue) Identifier() string {
	return e.identifier
}

func (e *EnumValue) Number() string {
	return e.number
}

type Enum struct {
	noopVisitor
	types *TypeDictionary

	*Options
	enumName   string
	enumValues []*EnumValue
}

func (e *Enum) Name() string {
	return e.enumName
}

func (e *Enum) GetElementType() ElementType {
	return EnumElementType
}

func (e *Enum) Values() []*EnumValue {
	return e.enumValues
}

func (e *Enum) VisitEnum(enum *parser.Enum) (next bool) {
	e.enumName = enum.EnumName
	e.types.startScope(e)
	e.types.RegisterEnum(e)
	e.types.endScope(e)
	return true
}

func (e *Enum) VisitEnumField(field *parser.EnumField) (next bool) {
	e.enumValues = append(e.enumValues, &EnumValue{
		identifier: field.Ident,
		number:     field.Number,
	})
	// Skip comments
	return false
}

func (e *Enum) VisitOption(option *parser.Option) (next bool) {
	return e.Options.VisitOption(option)
}

func NewEnum(types *TypeDictionary) *Enum {
	e := &Enum{
		types: types,
	}
	e.Options = NewOptions(e, types)
	return e
}
