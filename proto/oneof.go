package proto

type OneOf struct {
	OneOfName string
	Fields    []*Field
}

func (o *OneOf) AddField(f *Field) {
	o.Fields = append(o.Fields, f)
}

func NewOneOf(name string) *OneOf {
	return &OneOf{
		OneOfName: name,
		Fields:    nil,
	}
}
