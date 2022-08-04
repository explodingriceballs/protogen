package proto

import (
	"github.com/rs/zerolog/log"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"
)

type Options struct {
	noopVisitor
	types  *TypeDictionary
	parent Element

	Options []*Option
}

func (o *Options) getOptionType(name string) *Type {
	for _, option := range o.Options {
		var (
			optionPackageField = ""
			optionPackage      = ""
			optionField        = ""
		)
		if strings.Contains(option.OptionName, "(") {
			optionPackageField = option.OptionName[1:strings.LastIndex(option.OptionName, ")")]
			if !strings.Contains(optionPackageField, ".") {

			}
		}

		_ = optionPackageField
		_ = optionPackage
		_ = optionField
	}

	if strings.Contains(name, ".") {

	}
	log.Error().Msg("NAME!: " + name)
	return nil
}

func (o *Options) HasOption(name string) bool {
	if o.getOptionType(name) == nil {
		return false
	}

	return true
}

func (o *Options) GetOption(name string) map[string]interface{} {
	return nil
}

func (o *Options) VisitOption(option *parser.Option) (next bool) {
	o.AddOption(option.OptionName, option.Constant)

	// Ignore comments
	return false
}

func (o *Options) AddOption(name string, value string) {
	o.Options = append(o.Options, &Option{
		OptionName:  name,
		OptionValue: value,
	})
}

func NewOptions(parent Element, types *TypeDictionary) *Options {
	return &Options{
		parent: parent,
		types:  types,
	}
}

type Option struct {
	noopVisitor
	OptionName  string
	OptionValue string
}
