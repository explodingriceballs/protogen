package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Field struct {
	noopVisitor
	types *TypeDictionary

	FieldName   string
	FieldNumber string
	IsRepeated  bool
	Type        *Type
}

func (f *Field) VisitField(field *parser.Field) (next bool) {
	f.FieldName = field.FieldName
	f.FieldNumber = field.FieldNumber
	f.IsRepeated = field.IsRepeated
	f.Type = f.types.findType(field.Type)

	return true
}

func NewField(types *TypeDictionary) *Field {
	return &Field{
		noopVisitor: noopVisitor{},
		types:       types,
	}
}
