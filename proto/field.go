package proto

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Field struct {
	noopVisitor
	types *TypeDictionary

	*Options
	fieldName   string
	fieldNumber string
	isRepeated  bool
	t           *Type
}

func (f *Field) Name() string {
	return f.fieldName
}

func (f *Field) Type() *Type {
	return f.t
}

func (f *Field) FieldNumber() string {
	return f.fieldNumber
}

func (f *Field) IsRepeated() bool {
	return f.isRepeated
}

func (f *Field) GetElementType() ElementType {
	return FieldElementType
}

func (f *Field) VisitField(field *parser.Field) (next bool) {
	f.fieldName = field.FieldName
	f.fieldNumber = field.FieldNumber
	f.isRepeated = field.IsRepeated
	f.t = f.types.findType(field.Type)

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
	f := &Field{
		noopVisitor: noopVisitor{},
		types:       types,
	}
	f.Options = NewOptions(f, types)
	return f
}
