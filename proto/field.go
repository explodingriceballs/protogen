package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Field struct {
	noopVisitor
	types *TypeDictionary

	*Options
	FieldName   string
	FieldNumber string
	IsRepeated  bool
	Type        *Type
}

func (f *Field) GetName() string {
	return f.FieldName
}

func (f *Field) GetType() ElementType {
	return FieldElementType
}

func (f *Field) VisitField(field *parser.Field) (next bool) {
	f.FieldName = field.FieldName
	f.FieldNumber = field.FieldNumber
	f.IsRepeated = field.IsRepeated
	f.Type = f.types.findType(field.Type)

	if len(field.FieldOptions) > 0 {
		f.Options = NewOptions(f, f.types)

		for _, fieldOption := range field.FieldOptions {
			f.Options.AddOption(fieldOption.OptionName, fieldOption.Constant)
		}
	}

	return true
}

func (f *Field) VisitOption(option *parser.Option) (next bool) {
	return f.Options.VisitOption(option)
}

func NewField(types *TypeDictionary) *Field {
	return &Field{
		noopVisitor: noopVisitor{},
		types:       types,
	}
}
