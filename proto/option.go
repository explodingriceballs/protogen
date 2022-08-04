package proto

import "github.com/yoheimuta/go-protoparser/v4/parser"

type Option struct {
	noopVisitor
	OptionName  string
	OptionValue string
}

func (o *Option) VisitOption(option *parser.Option) (next bool) {
	o.OptionName = option.OptionName
	o.OptionValue = option.Constant

	// Ignore comments
	return false
}

func NewOption() *Option {
	return &Option{}
}
